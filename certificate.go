package main

import (
	"bytes"
	_ "embed"
	"encoding/base64"
	"html/template"
	"io"
	"strings"
	"time"
)

//go:embed certificate.html
var certificateMarkup string
var certificateView = template.Must(template.New("certificate").Parse(certificateMarkup))

type credentialView struct {
	Name, ShortName, Login, Issued, Credential, Link string
	Lessons, Projects                                int
	Sample                                           bool
	Logo                                             template.URL
}

func renderCredential(w io.Writer, v credentialView) error {
	logo, err := assets.ReadFile("dist/path-emblem.png")
	if err != nil {
		return err
	}
	v.Logo = template.URL("data:image/png;base64," + base64.StdEncoding.EncodeToString(logo))
	return certificateView.Execute(w, v)
}
func sampleCredential(w io.Writer) error {
	return renderCredential(w, credentialView{Name: "Alex Morgan", ShortName: "Alex", Login: "alex-example", Issued: "November 3, 2026", Credential: "SAMPLE-GP-0042", Lessons: 40, Projects: 12, Sample: true})
}
func (s *school) issue(r *record) {
	if r.Certificate != "" || !s.eligible(*r) {
		return
	}
	r.Certificate = nonce()
	r.Issued = time.Now().UTC().Format("January 2, 2006")
	for _, u := range s.units {
		if strings.HasPrefix(u.ID, "project-") {
			r.ProjectCount++
		} else {
			r.LessonCount++
		}
	}
}
func certificateBytes(v credentialView) ([]byte, error) {
	var b bytes.Buffer
	err := renderCredential(&b, v)
	return b.Bytes(), err
}
