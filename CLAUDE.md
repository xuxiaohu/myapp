# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Commands

```bash
make build        # compile to bin/myapp (VERSION injected via git describe)
make test         # run all tests with -race
make lint         # run golangci-lint
make fmt          # gofmt + goimports
make tidy         # go mod tidy && go mod verify
make test-cover   # tests + coverage.html report
make all          # lint → test → build
```

Run a single test:
```bash
go test -run TestLoad_defaults ./internal/config/
```

## Architecture

The startup flow is: `cmd/myapp/main.go` calls `internal/config.Load()` to read env vars, then uses `pkg/version.Version` (injected at build time via `-ldflags`) to print version info.

**Package boundaries:**
- `cmd/myapp/` — binary entry point only; no business logic
- `internal/` — packages that must not be imported outside this module
- `pkg/` — packages safe for external import (currently only `version`)


## Version injection

`pkg/version.Version` defaults to `"dev"`. The Makefile sets it at build time:
```
-X github.com/xuxiaohu/myapp/pkg/version.Version=$(git describe --tags --always --dirty)
```
Always use `make build` (not `go build` directly) to get a meaningful version string.

## Linter configuration

`.golangci.yml` uses v2 format with 13 linters enabled. `errcheck` and `gocritic` are suppressed in `_test.go` files. `goimports` expects local imports (`github.com/user/myapp/...`) to be grouped separately from stdlib and third-party.

## CI

`.github/workflows/ci.yml` runs `lint` and `test` jobs in parallel; `build` runs only after both pass. The `build` job uses `fetch-depth: 0` so `git describe` works for version tagging.
