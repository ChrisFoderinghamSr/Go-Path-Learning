# Go Path

A working, self-paced Go learning platform, from the first program to a full-stack project.

## Run

Install a supported Go release (1.22 or newer is the syntax baseline), open this folder in a terminal, then run:

```
go run .
```

Open http://127.0.0.1:8080 to view the enrollment screen, or visit /certificate/sample for the printable example. Student sign-in requires HTTPS, a registered GitHub OAuth app and persistent backend storage. Follow [the student deployment guide](STUDENT-DEPLOYMENT.md) before enrolling real students.

```
go test ./...
go vet ./...
```

## What works

- Twelve progressive modules and forty detailed lessons with original explanations, code examples, exercises, hints, worked solutions and sources.
- Search and filter lessons; navigate modules, checkpoints and project briefs.
- Interactive quizzes with explanatory feedback and retry support.
- Lesson completion requires a correct answer and a learner-reviewed exercise checklist.
- Twelve projects with requirements and saved self-assessment rubrics.
- Saved code drafts, personal notes, completion and quiz answers in the signed-in student’s private backend record.
- JSON export/import for backups and moving progress between devices. Import validates its schema and asks before replacement.
- Go examples and drafts can be downloaded, then run locally or pasted into the official Go Playground.
- Responsive layout, keyboard controls, labeled form fields, live feedback and a skip link.
- A dependency-free Go server embedding the complete learning app.

## Scope and architecture

The Go server serves the course and student API from one HTTPS origin. Students sign in with GitHub and enroll a public Go-Path-Student repository they own. Private records hold drafts, notes and quiz answers. GitHub receives a separate summary of milestones and reviewed commit SHAs. A single running service writes to persistent storage.

Coursework runs locally or in Go Playground. An instructor reviews each pinned submission. An issued certificate requires 40 completed lesson checkpoints and independent approval of all 40 lessons and 12 projects. The public verification page displays the award details recorded at issuance.

Fonts load from Google Fonts when online; the interface falls back to system fonts. Official references and the Go Playground require an internet connection. Course content is bundled locally.

## Source grounding

See `dist/sources.json` for the source and coverage audit, `dist/devdocs-topics.json` for the retrieved DevDocs topic inventory, and `CURRICULUM.md` for a readable export. DevDocs indexes references; a teaching sequence must also draw on the official Go tutorials, specification, package documentation, database guides and security guidance. The course adds PostgreSQL and MDN references for the full-stack portions.

The scope is a practical core path, not a claim to have manually reviewed every standard-library symbol. Specialist topics are mapped to further reading. Version-sensitive APIs should always be checked against the toolchain used by the learner.

## Files

`dist/curriculum.js` contains curriculum data. `dist/app.js` implements routing, quizzes, progress, notes and backup validation. `dist/style.css` contains responsive styling. `main.go` embeds and serves assets with read/write timeouts and response headers. Static hosting should configure equivalent headers at the hosting layer when available.

To edit the course, preserve stable module and lesson identifiers so existing progress remains usable. Incompatible curriculum changes should migrate the stored progress schema rather than silently discarding records.
