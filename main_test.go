package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestServer(t *testing.T) {
	for _, tc := range []struct {
		path     string
		status   int
		contains string
	}{
		{"/", 200, "Go Path"},
		{"/app.js", 200, "gopath.v1"},
		{"/curriculum.js", 200, "CURRICULUM"},
		{"/sources.json", 200, "go.dev"},
		{"/not-present", 404, "404"},
	} {
		t.Run(tc.path, func(t *testing.T) {
			w := httptest.NewRecorder()
			handler().ServeHTTP(w, httptest.NewRequest("GET", tc.path, nil))
			if w.Code != tc.status || !strings.Contains(w.Body.String(), tc.contains) {
				t.Fatalf("%s: status %d, expected %d and %q", tc.path, w.Code, tc.status, tc.contains)
			}
			if w.Header().Get("X-Content-Type-Options") != "nosniff" {
				t.Fatal("missing security header")
			}
		})
	}
}
func TestHealthAndMethods(t *testing.T) {
	h := handler()
	w := httptest.NewRecorder()
	h.ServeHTTP(w, httptest.NewRequest("GET", "/healthz", nil))
	var result map[string]string
	if err := json.Unmarshal(w.Body.Bytes(), &result); err != nil || result["status"] != "ok" {
		t.Fatal("invalid health response")
	}
	w = httptest.NewRecorder()
	h.ServeHTTP(w, httptest.NewRequest(http.MethodPost, "/", nil))
	if w.Code != http.StatusMethodNotAllowed {
		t.Fatalf("unexpected write method status %d", w.Code)
	}
}
