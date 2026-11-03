package main

import (
	"bytes"
	"crypto/rand"
	"crypto/sha256"
	_ "embed"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

//go:embed units.json
var catalog []byte

type unit struct {
	ID, Title string
	Answer    int
}
type session struct {
	ID           int64
	Login, Token string
	Until        time.Time
}
type submission struct{ SHA, Status, Reviewer, Comment string }
type record struct {
	Repo        int64
	Progress    json.RawMessage
	Version     int
	Work        map[string]submission
	Certificate string
	Published   string
	Login       string
}
type school struct {
	mu          sync.Mutex
	recordLocks [64]sync.Mutex
	sessions    map[string]session
	units       []unit
	dir, origin string
	client      *http.Client
}

func nonce() string {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		panic(err)
	}
	return base64.RawURLEncoding.EncodeToString(b)
}
func (s *school) path(id int64) string {
	return filepath.Join(s.dir, strconv.FormatInt(id, 10)+".json")
}
func (s *school) load(id int64) (record, error) {
	r := record{Work: map[string]submission{}}
	b, e := os.ReadFile(s.path(id))
	if errors.Is(e, os.ErrNotExist) {
		return r, nil
	}
	if e == nil {
		e = json.Unmarshal(b, &r)
	}
	return r, e
}
func (s *school) store(id int64, r record) error {
	b, e := json.Marshal(r)
	if e != nil {
		return e
	}
	p := s.path(id)
	if e = os.WriteFile(p+".tmp", b, 0600); e != nil {
		return e
	}
	return os.Rename(p+".tmp", p)
}
func die(w http.ResponseWriter, code int, msg string) { http.Error(w, msg, code) }
func send(w http.ResponseWriter, v any) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(v)
}
func cookie(w http.ResponseWriter, name, value string, age int) {
	http.SetCookie(w, &http.Cookie{Name: name, Value: value, Path: "/", HttpOnly: true, Secure: true, SameSite: http.SameSiteLaxMode, MaxAge: age})
}
func (s *school) api(sess session, method, path string, data, out any) error {
	b, e := json.Marshal(data)
	if e != nil {
		return e
	}
	req, e := http.NewRequest(method, "https://api.github.com"+path, bytes.NewReader(b))
	if e != nil {
		return e
	}
	req.Header.Set("Authorization", "Bearer "+sess.Token)
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("Content-Type", "application/json")
	resp, e := s.client.Do(req)
	if e != nil {
		return e
	}
	defer resp.Body.Close()
	if resp.StatusCode/100 != 2 {
		return fmt.Errorf("GitHub returned %d", resp.StatusCode)
	}
	if out != nil {
		return json.NewDecoder(resp.Body).Decode(out)
	}
	return nil
}
func (s *school) signin(w http.ResponseWriter, r *http.Request) {
	if s.origin == "" || os.Getenv("GITHUB_CLIENT_ID") == "" || os.Getenv("GITHUB_CLIENT_SECRET") == "" {
		die(w, 503, "GitHub sign-in has not been configured by the course operator.")
		return
	}
	state, key := nonce(), nonce()
	cookie(w, "__Host-go-oauth", state+"."+key, 600)
	hash := sha256.Sum256([]byte(key))
	q := url.Values{"client_id": {os.Getenv("GITHUB_CLIENT_ID")}, "redirect_uri": {s.origin + "/student/callback"}, "scope": {"public_repo"}, "state": {state}, "code_challenge": {base64.RawURLEncoding.EncodeToString(hash[:])}, "code_challenge_method": {"S256"}}
	http.Redirect(w, r, "https://github.com/login/oauth/authorize?"+q.Encode(), 302)
}
func (s *school) callback(w http.ResponseWriter, r *http.Request) {
	c, e := r.Cookie("__Host-go-oauth")
	cookie(w, "__Host-go-oauth", "", -1)
	if e != nil {
		die(w, 403, "Login request expired.")
		return
	}
	p := strings.Split(c.Value, ".")
	if len(p) != 2 || r.URL.Query().Get("state") != p[0] {
		die(w, 403, "Login request mismatch.")
		return
	}
	q := url.Values{"client_id": {os.Getenv("GITHUB_CLIENT_ID")}, "client_secret": {os.Getenv("GITHUB_CLIENT_SECRET")}, "code": {r.URL.Query().Get("code")}, "code_verifier": {p[1]}, "redirect_uri": {s.origin + "/student/callback"}}
	req, _ := http.NewRequest("POST", "https://github.com/login/oauth/access_token", strings.NewReader(q.Encode()))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, e := s.client.Do(req)
	if e != nil {
		die(w, 502, "GitHub unavailable.")
		return
	}
	defer resp.Body.Close()
	var t struct {
		Token string `json:"access_token"`
	}
	if json.NewDecoder(resp.Body).Decode(&t) != nil || t.Token == "" {
		die(w, 403, "GitHub did not authorize sign-in.")
		return
	}
	ss := session{Token: t.Token, Until: time.Now().Add(8 * time.Hour)}
	var user struct {
		ID    int64
		Login string
	}
	if s.api(ss, "GET", "/user", nil, &user) != nil || user.ID == 0 {
		die(w, 403, "Could not verify GitHub identity.")
		return
	}
	ss.ID = user.ID
	ss.Login = user.Login
	key := nonce()
	s.mu.Lock()
	for k, v := range s.sessions {
		if time.Now().After(v.Until) {
			delete(s.sessions, k)
		}
	}
	s.sessions[key] = ss
	s.mu.Unlock()
	cookie(w, "__Host-go-session", key, 28800)
	http.Redirect(w, r, "/", 303)
}
func page(w http.ResponseWriter, title, body string) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, "<!doctype html><html lang=en><meta name=viewport content='width=device-width,initial-scale=1'><title>%s · Go Path</title><link rel=stylesheet href=/style.css><main><article class='lesson-body'><h1>%s</h1>%s</article></main></html>", title, title, body)
}
func (s *school) repo(ss session) (int64, error) {
	var repo struct {
		ID                 int64
		Name               string
		Private, Fork      bool
		Owner              struct{ ID int64 }
		TemplateRepository any `json:"template_repository"`
	}
	e := s.api(ss, "GET", "/repos/"+ss.Login+"/Go-Path-Student", nil, &repo)
	if e != nil {
		return 0, e
	}
	if repo.Owner.ID != ss.ID || repo.Private || repo.Fork || repo.TemplateRepository != nil || repo.Name != "Go-Path-Student" {
		return 0, errors.New("Use a new public Go-Path-Student repository owned by your account.")
	}
	return repo.ID, nil
}
func (s *school) reviewer(id int64) bool {
	for _, x := range strings.Split(os.Getenv("REVIEWER_GITHUB_IDS"), ",") {
		if strings.TrimSpace(x) == strconv.FormatInt(id, 10) {
			return true
		}
	}
	return false
}
func (s *school) portal(w http.ResponseWriter, ss session, r record) {
	var b strings.Builder
	b.WriteString("<p>Signed in as " + template.HTMLEscapeString(ss.Login) + ". Your notes and drafts stay private. Course milestones, evidence links and certificate status are public when you sync to GitHub.</p>")
	if r.Repo == 0 {
		b.WriteString("<p>Create a new public <b>Go-Path-Student</b> repository with a README. Submit your own work. Forked, copied or cloned solutions are ineligible. Use AI only when the assignment explicitly permits it, and document that use. Instructors may ask for a walkthrough.</p><p><a href='https://github.com/new?name=Go-Path-Student' target='_blank' rel='noopener'>Create your student repository ↗</a></p><form method=post action=/student/enroll><label><input required type=checkbox name=agree value=yes> I accept the originality policy and public progress tracking.</label><p><button>Verify repository and enroll</button></p></form>")
	} else {
		b.WriteString("<p><a href='/'>Open learning path →</a></p><p>For each item, commit original work under <code>work/&lt;item ID&gt;/</code> in Go-Path-Student. Include code, tests and an explanation. Submit its full commit SHA. An instructor reviews the pinned evidence.</p><form method=post action=/student/submit><select name=unit>")
		for _, u := range s.units {
			fmt.Fprintf(&b, "<option value='%s'>%s · %s</option>", u.ID, u.ID, template.HTMLEscapeString(u.Title))
		}
		b.WriteString("</select><input name=sha required pattern='[a-f0-9]{40}' placeholder='40-character commit SHA'><button>Submit for review</button></form><ul>")
		for _, u := range s.units {
			x := r.Work[u.ID]
			fmt.Fprintf(&b, "<li>%s · %s %s</li>", template.HTMLEscapeString(u.Title), template.HTMLEscapeString(x.Status), template.HTMLEscapeString(x.Comment))
		}
		b.WriteString("</ul><form method=post action=/student/sync><button>Sync milestones to public GitHub repository</button></form>")
		if r.Certificate != "" {
			fmt.Fprintf(&b, "<p><a href='/certificate/%d/%s'>View your completion certificate</a></p>", ss.ID, r.Certificate)
		}
	}
	if s.reviewer(ss.ID) {
		b.WriteString("<h2>Pending submissions</h2><ul>")
		entries, _ := os.ReadDir(s.dir)
		for _, entry := range entries {
			id, err := strconv.ParseInt(strings.TrimSuffix(entry.Name(), ".json"), 10, 64)
			if err != nil {
				continue
			}
			student, err := s.load(id)
			if err != nil {
				continue
			}
			for item, x := range student.Work {
				if x.Status == "pending" {
					fmt.Fprintf(&b, "<li>Student %d / %s / <a href='https://github.com/%s/Go-Path-Student/commit/%s'>%s</a></li>", id, template.HTMLEscapeString(item), url.PathEscape(student.Login), x.SHA, x.SHA)
				}
			}
		}
		b.WriteString("</ul>")
		b.WriteString("<h2>Instructor review</h2><p>Review the pinned diff, run its tests in an isolated environment, and verify originality and permitted AI use with the learner. Enter the student's numeric GitHub ID, item ID and exact reviewed SHA.</p><form method=post action=/student/review><input required name=id placeholder='Student GitHub ID'><input required name=unit placeholder='Item ID'><input required name=sha placeholder='Reviewed SHA'><input required name=comment placeholder='Review feedback'><select name=decision><option>approved</option><option>revise</option></select><button>Record review</button></form>")
	}
	fmt.Fprintf(&b, "<p>Your GitHub student ID: %d</p><form method=post action=/student/logout><button>Sign out</button></form>", ss.ID)
	page(w, "Student record", b.String())
}

var shaPattern = regexp.MustCompile("^[0-9a-f]{40}$")

func (s *school) actions(w http.ResponseWriter, q *http.Request, ss session, r record) {
	fail := func(msg string) { die(w, 400, msg) }
	path := q.URL.Path
	if path == "/student/enroll" {
		id, e := s.repo(ss)
		if e != nil || q.FormValue("agree") != "yes" {
			fail("Create your original public Go-Path-Student repository, initialize it with a README and accept the policy.")
			return
		}
		if r.Repo != 0 && r.Repo != id {
			fail("Repository identity changed.")
			return
		}
		r.Repo = id
		r.Login = ss.Login
	} else if path == "/student/review" {
		id, e := strconv.ParseInt(q.FormValue("id"), 10, 64)
		if !s.reviewer(ss.ID) || id == ss.ID || e != nil {
			die(w, 403, "Independent instructor review required.")
			return
		}
		target, e := s.load(id)
		if e != nil {
			die(w, 500, "Record unreadable.")
			return
		}
		item := q.FormValue("unit")
		x := target.Work[item]
		decision := q.FormValue("decision")
		if x.Status != "pending" || x.SHA != q.FormValue("sha") || (decision != "approved" && decision != "revise") || strings.TrimSpace(q.FormValue("comment")) == "" {
			fail("Check student, item, pending SHA and feedback.")
			return
		}
		x.Status = decision
		x.Comment = q.FormValue("comment")
		x.Reviewer = ss.Login
		target.Work[item] = x
		if s.eligible(target) && target.Certificate == "" {
			target.Certificate = nonce()
		}
		if s.store(id, target) != nil {
			die(w, 500, "Could not save review.")
			return
		}
		http.Redirect(w, q, "/student", 303)
		return
	} else {
		if r.Repo == 0 {
			die(w, 403, "Enroll before recording work.")
			return
		}
		switch path {
		case "/student/progress":
			var update struct {
				Version  int
				Progress json.RawMessage
			}
			if json.NewDecoder(q.Body).Decode(&update) != nil || len(update.Progress) == 0 || update.Progress[0] != '{' {
				fail("Invalid progress.")
				return
			}
			if update.Version != r.Version {
				die(w, 409, "Progress changed elsewhere. Reload before saving.")
				return
			}
			r.Version++
			r.Progress = update.Progress
		case "/student/submit", "/student/sync":
			id, e := s.repo(ss)
			if e != nil || id != r.Repo {
				fail("Student repository ownership, identity or visibility changed.")
				return
			}
			if path == "/student/sync" {
				if e = s.syncGit(ss, &r); e != nil {
					die(w, 502, "GitHub sync failed. Confirm write authorization, initialize the repository and retry.")
					return
				}
				http.Redirect(w, q, "/student", 303)
				return
			}
			item, sha := q.FormValue("unit"), q.FormValue("sha")
			valid := false
			for _, u := range s.units {
				if u.ID == item {
					valid = true
				}
			}
			if !valid || !shaPattern.MatchString(sha) || r.Work[item].Status == "approved" {
				fail("Choose an unapproved course item and a full commit SHA.")
				return
			}
			var c struct {
				SHA    string
				Author struct{ ID int64 }
				Files  []struct{ Filename, Status string }
			}
			base := "/repos/" + ss.Login + "/Go-Path-Student"
			if s.api(ss, "GET", base+"/commits/"+sha, nil, &c) != nil || c.SHA != sha || c.Author.ID != ss.ID {
				fail("Commit must resolve in your repository and be attributed to your account.")
				return
			}
			found := false
			for _, f := range c.Files {
				if strings.HasPrefix(f.Filename, "work/"+item+"/") && f.Status != "removed" {
					found = true
				}
			}
			if !found {
				fail("Commit must change evidence inside work/" + item + "/.")
				return
			}
			r.Work[item] = submission{SHA: sha, Status: "pending"}
		default:
			die(w, 404, "Unknown action.")
			return
		}
	}
	if s.eligible(r) && r.Certificate == "" {
		r.Certificate = nonce()
	}
	if s.store(ss.ID, r) != nil {
		die(w, 500, "Could not save student record.")
		return
	}
	if path == "/student/progress" {
		warning := ""
		if err := s.syncGit(ss, &r); err != nil {
			warning = "Saved privately; GitHub sync pending. Retry in Student record."
		} else if err := s.store(ss.ID, r); err != nil {
			warning = "Saved privately; sync receipt pending."
		}
		send(w, map[string]any{"version": r.Version, "warning": warning})
		return
	}
	http.Redirect(w, q, "/student", 303)
}
func (s *school) eligible(r record) bool {
	var p struct {
		Completed []string
		Quiz      map[string]int
		Projects  map[string]map[string]bool
	}
	if json.Unmarshal(r.Progress, &p) != nil {
		return false
	}
	complete := map[string]bool{}
	for _, id := range p.Completed {
		complete[id] = true
	}
	for _, u := range s.units {
		if r.Work[u.ID].Status != "approved" {
			return false
		}
		if u.Answer >= 0 {
			a, ok := p.Quiz[u.ID]
			if !ok || a != u.Answer || !complete[u.ID] {
				return false
			}
		}
	}
	return len(s.units) > 0
}
func (s *school) syncGit(ss session, r *record) error {
	path := "/repos/" + ss.Login + "/Go-Path-Student/contents/go-path/progress.json"
	var old struct{ SHA string } // A missing file is created without a SHA. GitHub rejects conflicting writes.

	milestones := map[string]any{}
	for k, v := range r.Work {
		milestones[k] = map[string]string{"sha": v.SHA, "review": v.Status}
	}
	var study struct {
		Completed []string
		Last      *string
		Projects  map[string]map[string]bool
	}
	_ = json.Unmarshal(r.Progress, &study)
	b, _ := json.MarshalIndent(map[string]any{"learner": ss.Login, "completedLessons": study.Completed, "projectChecklists": study.Projects, "lastLesson": study.Last, "reviews": milestones, "certified": r.Certificate != ""}, "", "  ")
	digest := fmt.Sprintf("%x", sha256.Sum256(b))
	if digest == r.Published {
		return nil
	}
	id, err := s.repo(ss)
	if err != nil || id != r.Repo {
		return errors.New("Enrollment repository changed")
	}
	_ = s.api(ss, "GET", path, nil, &old)
	data := map[string]string{"message": "Record Go Path learning progress", "content": base64.StdEncoding.EncodeToString(b)}
	if old.SHA != "" {
		data["sha"] = old.SHA
	}
	if err := s.api(ss, "PUT", path, data, nil); err != nil {
		return err
	}
	r.Published = digest
	return nil
}
func academy(next http.Handler) http.Handler {
	s := &school{sessions: map[string]session{}, dir: os.Getenv("STUDENT_DATA_DIR"), origin: strings.TrimRight(os.Getenv("PUBLIC_ORIGIN"), "/"), client: &http.Client{Timeout: 12 * time.Second}}
	if s.dir == "" {
		s.dir = ".student-data"
	}
	if err := os.MkdirAll(s.dir, 0700); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(catalog, &s.units); err != nil {
		panic(err)
	}
	return http.HandlerFunc(func(w http.ResponseWriter, q *http.Request) {
		p := q.URL.Path
		if p == "/healthz" || p == "/style.css" {
			next.ServeHTTP(w, q)
			return
		}
		w.Header().Set("Cache-Control", "no-store")
		w.Header().Set("Referrer-Policy", "no-referrer")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		if q.Method == "GET" && p == "/student/login" {
			s.signin(w, q)
			return
		}
		if q.Method == "GET" && p == "/student/callback" {
			s.callback(w, q)
			return
		}

		if q.Method == "GET" && strings.HasPrefix(p, "/certificate/") {
			v := strings.Split(strings.TrimPrefix(p, "/certificate/"), "/")
			if len(v) == 2 {
				id, e := strconv.ParseInt(v[0], 10, 64)
				r, err := s.load(id)
				if e == nil && err == nil && r.Certificate != "" && r.Certificate == v[1] {
					page(w, "Certificate of completion", fmt.Sprintf("<p>Go Path · Beginner to Full Stack</p><p>GitHub account ID %d completed all %d instructor-verified coursework items.</p><p>Credential ID: %s</p>", id, len(s.units), r.Certificate))
					return
				}
			}
			die(w, 404, "Unknown credential.")
			return
		}
		c, _ := q.Cookie("__Host-go-session")
		var ss session
		if c != nil {
			s.mu.Lock()
			ss = s.sessions[c.Value]
			s.mu.Unlock()
		}
		if ss.ID == 0 || time.Now().After(ss.Until) {
			if p == "/" && q.Method == "GET" {
				if _, e := q.Cookie("__Host-go-visit"); e != nil {
					cookie(w, "__Host-go-visit", "seen", 86400)
					f, err := os.OpenFile(filepath.Join(s.dir, "visits.log"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
					if err == nil {
						fmt.Fprintln(f, time.Now().UTC().Format("2006-01-02"))
						f.Close()
					}
				}
				page(w, "Learn Go with GitHub", "<p>Sign in to enroll. We count anonymous visits by day. Enrollment requires a public Go-Path-Student repository. GitHub requests public repository write access so you can publish milestones. Your private notes stay on the course server.</p><a href='/student/login'>Sign in with GitHub →</a>")
				return
			}
			die(w, 401, "Sign in with GitHub.")
			return
		}
		if q.Method == "POST" {
			q.Body = http.MaxBytesReader(w, q.Body, 256<<10)
		}
		lockID := ss.ID
		if p == "/student/review" && s.reviewer(ss.ID) {
			if id, err := strconv.ParseInt(q.FormValue("id"), 10, 64); err == nil && id > 0 {
				lockID = id
			}
		}
		lock := &s.recordLocks[uint64(lockID)%64]
		lock.Lock()
		defer lock.Unlock()
		r, e := s.load(ss.ID)
		if e != nil {
			die(w, 500, "Record unreadable.")
			return
		}
		if q.Method == "POST" {
			if s.origin == "" || q.Header.Get("Origin") != s.origin {
				die(w, 403, "Origin mismatch.")
				return
			}
			q.Body = http.MaxBytesReader(w, q.Body, 256<<10)
			if p == "/student/logout" {
				s.mu.Lock()
				delete(s.sessions, c.Value)
				s.mu.Unlock()
				cookie(w, "__Host-go-session", "", -1)
				http.Redirect(w, q, "/", 303)
				return
			}
			s.actions(w, q, ss, r)
			return
		}
		if p == "/student" || (p == "/" && r.Repo == 0) {
			s.portal(w, ss, r)
			return
		}
		if r.Repo == 0 {
			die(w, 403, "Enrollment required.")
			return
		}
		if p == "/student/state" {
			send(w, map[string]any{"id": ss.ID, "version": r.Version, "progress": r.Progress})
			return
		}
		next.ServeHTTP(w, q)
	})
}
