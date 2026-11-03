package main

import (
	"encoding/base64"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

type ghTransport func(*http.Request) (*http.Response, error)

func (f ghTransport) RoundTrip(r *http.Request) (*http.Response, error) { return f(r) }
func ghReply(code int, v any) *http.Response {
	b, _ := json.Marshal(v)
	return &http.Response{StatusCode: code, Body: io.NopCloser(strings.NewReader(string(b)))}
}
func form(values url.Values) *http.Request {
	q := httptest.NewRequest("POST", "/", strings.NewReader(values.Encode()))
	q.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	return q
}

func TestEnrollmentSubmissionReviewAndPublication(t *testing.T) {
	t.Setenv("REVIEWER_GITHUB_IDS", "88")
	sha := strings.Repeat("a", 40)
	student := session{ID: 7, Login: "alice", Token: "example"}
	fork := true
	published := ""
	s := school{dir: t.TempDir(), units: []unit{{ID: "hello", Answer: 1}}}
	s.client = &http.Client{Transport: ghTransport(func(q *http.Request) (*http.Response, error) {
		if q.Header.Get("Authorization") != "Bearer example" {
			t.Fatal("missing user authorization")
		}
		switch q.URL.Path {
		case "/repos/alice/Go-Path-Student":
			return ghReply(200, map[string]any{"id": 100, "name": "Go-Path-Student", "owner": map[string]int{"id": 7}, "fork": fork}), nil
		case "/repos/alice/Go-Path-Student/commits/" + sha:
			return ghReply(200, map[string]any{"sha": sha, "author": map[string]int{"id": 7}, "files": []map[string]string{{"filename": "work/hello/main.go", "status": "added"}}}), nil
		case "/repos/alice/Go-Path-Student/contents/go-path/progress.json":
			if q.Method == "GET" {
				return ghReply(404, nil), nil
			}
			var data map[string]string
			json.NewDecoder(q.Body).Decode(&data)
			b, _ := base64.StdEncoding.DecodeString(data["content"])
			published = string(b)
			return ghReply(201, nil), nil
		}
		t.Fatalf("unexpected GitHub request %s", q.URL.Path)
		return nil, nil
	})}
	act := func(user session, path string, q *http.Request, want int) {
		t.Helper()
		r, e := s.load(user.ID)
		if e != nil {
			t.Fatal(e)
		}
		q.URL.Path = path
		w := httptest.NewRecorder()
		s.actions(w, q, user, r)
		if w.Code != want {
			t.Fatalf("%s: got %d (%s), want %d", path, w.Code, w.Body.String(), want)
		}
	}
	enroll := url.Values{"agree": {"yes"}}
	act(student, "/student/enroll", form(enroll), 400)
	fork = false
	act(student, "/student/enroll", form(enroll), 303)
	act(student, "/student/submit", form(url.Values{"unit": {"hello"}, "sha": {sha}}), 303)
	progress := `{"version":0,"progress":{"completed":["hello"],"quiz":{"hello":1},"notes":{"hello":"PRIVATE-NOTE"}}}`
	act(student, "/student/progress", httptest.NewRequest("POST", "/", strings.NewReader(progress)), 200)
	if strings.Contains(published, "PRIVATE-NOTE") || !strings.Contains(published, "hello") {
		t.Fatal("incorrect public snapshot")
	}
	act(student, "/student/progress", httptest.NewRequest("POST", "/", strings.NewReader(progress)), 409)
	r, _ := s.load(7)
	if r.Certificate != "" {
		t.Fatal("self-reported progress issued a credential")
	}
	review := url.Values{"id": {"7"}, "unit": {"hello"}, "sha": {sha}, "decision": {"approved"}, "comment": {"Tests passed and walkthrough reviewed."}}
	act(student, "/student/review", form(review), 403)
	act(session{ID: 88, Login: "instructor"}, "/student/review", form(review), 303)
	r, _ = s.load(7)
	if r.Certificate == "" || r.Work["hello"].Reviewer != "instructor" {
		t.Fatal("missing independently verified credential")
	}
}
