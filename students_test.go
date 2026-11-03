package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCourseRequiresEnrollment(t *testing.T) {
	t.Setenv("STUDENT_DATA_DIR", t.TempDir())
	h := academy(handler())
	for _, path := range []string{"/app.js", "/curriculum.js", "/student/state"} {
		w := httptest.NewRecorder()
		h.ServeHTTP(w, httptest.NewRequest("GET", path, nil))
		if w.Code != 401 {
			t.Fatalf("%s returned %d", path, w.Code)
		}
	}
	w := httptest.NewRecorder()
	h.ServeHTTP(w, httptest.NewRequest("GET", "/", nil))
	if !strings.Contains(w.Body.String(), "Sign in with GitHub") {
		t.Fatal("missing public sign-in")
	}
}
func TestEligibilityRequiresIndependentEvidence(t *testing.T) {
	s := school{units: []unit{{ID: "hello", Answer: 0}, {ID: "project-first", Answer: -1}}}
	r := record{Work: map[string]submission{"hello": {Status: "approved"}}, Progress: json.RawMessage(`{"completed":["hello"],"quiz":{"hello":0}}`)}
	if s.eligible(r) {
		t.Fatal("missing project approval granted credit")
	}
	r.Work["project-first"] = submission{Status: "approved"}
	if !s.eligible(r) {
		t.Fatal("complete record rejected")
	}
	r.Progress = json.RawMessage(`{"completed":["hello"],"quiz":{}}`)
	if s.eligible(r) {
		t.Fatal("missing quiz answer counted as answer zero")
	}
}
func TestReviewerCannotApproveOwnWork(t *testing.T) {
	t.Setenv("REVIEWER_GITHUB_IDS", "17")
	s := school{}
	req := httptest.NewRequest(http.MethodPost, "/student/review", strings.NewReader("id=17"))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	s.actions(w, req, session{ID: 17}, record{})
	if w.Code != 403 {
		t.Fatalf("self-approval returned %d", w.Code)
	}
}
func TestRecordSurvivesRestart(t *testing.T) {
	dir := t.TempDir()
	s := school{dir: dir}
	r := record{Repo: 99, Version: 4, Work: map[string]submission{"hello": {SHA: "a", Status: "pending"}}}
	if e := s.store(12, r); e != nil {
		t.Fatal(e)
	}
	restarted := school{dir: dir}
	got, e := restarted.load(12)
	if e != nil || got.Repo != 99 || got.Work["hello"].Status != "pending" {
		t.Fatal("stored evidence lost", e)
	}
}
