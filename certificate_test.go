package main

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestCurrentCredentialRequirements(t *testing.T) {
	var s school
	if err := json.Unmarshal(catalog, &s.units); err != nil {
		t.Fatal(err)
	}
	r := record{Work: map[string]submission{}}
	lessons := []string{}
	quiz := map[string]int{}
	projects := 0
	for _, u := range s.units {
		r.Work[u.ID] = submission{Status: "approved"}
		if u.Answer < 0 {
			projects++
		} else {
			lessons = append(lessons, u.ID)
			quiz[u.ID] = u.Answer
		}
	}
	if len(lessons) != 40 || projects != 12 {
		t.Fatal("catalog and certificate do not match")
	}
	r.Progress, _ = json.Marshal(map[string]any{"completed": lessons, "quiz": quiz})
	last := s.units[len(s.units)-1]
	delete(r.Work, last.ID)
	s.issue(&r)
	if r.Certificate != "" {
		t.Fatal("missing approval issued certificate")
	}
	r.Work[last.ID] = submission{Status: "approved"}
	s.issue(&r)
	if r.Certificate == "" || r.Issued == "" || r.LessonCount != 40 || r.ProjectCount != 12 {
		t.Fatal("invalid issued credential")
	}
	var out strings.Builder
	if err := sampleCredential(&out); err != nil {
		t.Fatal(err)
	}
	for _, expected := range []string{"data:image/png;base64,", "GO via PATH LEARNING CENTER of excellence", "Alex Morgan", "@alex-example", "40 lesson checkpoints and 12 projects", "SAMPLE-GP-0042", "November 3, 2026", "Sample only."} {
		if !strings.Contains(out.String(), expected) {
			t.Fatalf("missing %s", expected)
		}
	}
}
