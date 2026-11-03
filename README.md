# Go Path

A working, self-paced Go learning platform, from the first program to a full-stack project.

## Run

Install a supported Go release (1.22 or newer is the syntax baseline), open this folder in a terminal, then run:

```
go run .
```

Open http://127.0.0.1:8080. Stop the server with Ctrl+C. No dependencies or database are required for the learning platform. You can alternatively serve `dist/` with any static web server. Use the same address each time to retain browser-local progress. Device storage is scoped by browser and site address.

```
go test ./...
go vet ./...
```

## What works

- Ten progressive modules and thirty detailed lessons with original explanations, code examples, exercises, hints, worked solutions and sources.
- Search and filter lessons; navigate modules, checkpoints and project briefs.
- Interactive quizzes with explanatory feedback and retry support.
- Lesson completion requires a correct answer and a learner-reviewed exercise checklist.
- Ten projects with requirements and saved self-assessment rubrics.
- Saved code drafts, personal notes, completion and quiz answers on the current device.
- JSON export/import for backups and moving progress between devices. Import validates its schema and asks before replacement.
- Go examples and drafts can be downloaded, then run locally or pasted into the official Go Playground.
- Responsive layout, keyboard controls, labeled form fields, live feedback and a skip link.
- A dependency-free Go server embedding the complete learning app.

## Scope and architecture

The learning UI is HTML/CSS/JavaScript. The local server is Go `net/http`; the hosted edition serves the same static assets. Learning records use `localStorage`, with no login, server database or cross-device sync. The curriculum teaches Go `net/http`, PostgreSQL through `database/sql`, HTML templates and browser fetch, testing, authentication design, security and deployment. Those integrations are course material and project work, not hidden services operated by this app.

The editor does not execute untrusted code, and review checklists do not automatically grade submissions. Checkpoint questions are graded in the browser. A passed checkpoint is evidence of a knowledge check, not a professional certification. No credentials or private data are sent to a code runner.

Fonts load from Google Fonts when online; the interface falls back to system fonts. Official references and the Go Playground require an internet connection. Course content is bundled locally.

## Source grounding

See `dist/sources.json` for the source and coverage audit, `dist/devdocs-topics.json` for the retrieved DevDocs topic inventory, and `CURRICULUM.md` for a readable export. DevDocs indexes references; a teaching sequence must also draw on the official Go tutorials, specification, package documentation, database guides and security guidance. The course adds PostgreSQL and MDN references for the full-stack portions.

The scope is a practical core path, not a claim to have manually reviewed every standard-library symbol. Specialist topics are mapped to further reading. Version-sensitive APIs should always be checked against the toolchain used by the learner.

## Files

`dist/curriculum.js` contains curriculum data. `dist/app.js` implements routing, quizzes, progress, notes and backup validation. `dist/style.css` contains responsive styling. `main.go` embeds and serves assets with read/write timeouts and response headers. Static hosting should configure equivalent headers at the hosting layer when available.

To edit the course, preserve stable module and lesson identifiers so existing progress remains usable. Incompatible curriculum changes should migrate the stored progress schema rather than silently discarding records.
