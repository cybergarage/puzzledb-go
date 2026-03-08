# Repository Guidelines

## Project Structure & Module Organization
- `puzzledb/`: core library and server internals (auth, cluster, coordinator, plugins, API bindings).
- `cmd/`: entrypoints for binaries (`puzzledb-server`, `puzzledb-cli`).
- `puzzledbtest/`: integration-style test helpers and suites, including plugin-specific tests.
- `doc/`: architecture, specs, protocol docs, and generated CLI/server/API docs.
- `scripts/`: utility scripts (for example FoundationDB setup helpers).
- Root configs: `Makefile`, `.golangci.yaml`, `go.mod`.

## Build, Test, and Development Commands
- `make build`: build server and CLI binaries.
- `make install`: install binaries into `$(go env GOPATH)/bin`.
- `make run`: install and start `puzzledb-server` locally.
- `make certs`: generate test certificates under `puzzledbtest/certs`.
- `make unittest`: run Go tests and generate `puzzledb-cover.out` + HTML coverage.
- `make test`: full local gate (`format` + `lint` + tests).
- `make lint`: run `golangci-lint` on `puzzledb/`, `puzzledbtest/`, and `cmd/`.
- `make format`: refresh version file and apply `gofmt -s`.

## Coding Style & Naming Conventions
- Language: Go (CI uses Go `1.25.x`).
- Formatting: always run `make format` (includes `gofmt -s`).
- Linting: follow `.golangci.yaml`; resolve warnings before opening a PR.
- Naming: use idiomatic Go naming (`CamelCase` exports, `mixedCaps` locals, short receiver names).
- Keep package boundaries aligned with existing domains (`plugins/*`, `store/*`, `query/*`).

## Testing Guidelines
- Primary framework: Go `testing` with package/unit tests and `puzzledbtest` integration suites.
- Test file naming: `*_test.go`; prefer table-driven tests for protocol/storage behavior.
- Before PR: run `make certs && make unittest` (matches CI behavior) and attach relevant results if failures are environment-specific.

## Commit & Pull Request Guidelines
- Prefer concise, imperative commit subjects.
- Common patterns in history: `refactor(scope): ...`, `style: ...`, `Update ...`, `Update doc: <path>`.
- Keep commits scoped; separate refactors, style-only changes, and docs updates.
- PRs should target `main`, describe behavior changes, list validation commands run, and include doc/spec updates when interfaces, proto files, or commands change.
