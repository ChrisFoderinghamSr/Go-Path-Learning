# Validation — November 3, 2026

Passed with Go 1.26.5 on Windows and Node.js 26.5.1:

- `node --check dist/app.js` and `node --check dist/curriculum.js`.
- `node check.mjs`: 81 core routes, 40 lesson schemas, invalid-backup rejection, and imported-note HTML escaping.
- `go test ./...`: embedded assets, not-found behavior, health response, method rejection and security header checks.
- `go vet ./...`.
- 23 complete Go example/solution checks with `go test`. The multi-file module example was assembled into its documented files; the clamp test used the function from the earlier exercise solution. Contextual SQL, frontend, shell and integration snippets require the surrounding application described in their lessons and were not all executed as complete applications.

Browser verification on the local preview:

- Dashboard layout and lesson navigation.
- Exercise review, correct quiz feedback, lesson completion and the updated dashboard.
- Quiz choice, review checks, completion and notes persist after reload.
- Curriculum search returns database and capstone matches for PostgreSQL.
- Full reference explorer loads the 6,590-entry index and returns two expected matches for http.NewRequest.

The frontend backup validator is checked programmatically. Export/import and code-download controls are implemented; no claim is made that every browser-specific file-dialog path was exercised. The responsive CSS has desktop, tablet and mobile layouts; the browser visual inspection used the desktop viewport. Database, authentication, deployment and security integration exercises are learning material requiring learner-provisioned services, not operated course infrastructure.

Certificate checks verify the seal, sample fields and all 52 required approvals (40 lessons and 12 projects).
