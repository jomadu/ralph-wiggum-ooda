# Task: Add Makefile for Build/Test/Lint

## Intent

Standardize build, test, and lint operations using `make` to provide a consistent interface across the project lifecycle.

## Motivation

Currently, build/test/lint commands are documented in AGENTS.md but require developers to remember different commands for bash (v0.1.0) vs Go (v2) implementations. A Makefile provides:

- Single command interface: `make test`, `make build`, `make lint`
- Automatic detection of which implementation to run
- Consistent developer experience
- Easy CI/CD integration

## Desired Outcome

A `Makefile` at project root that provides:

```bash
make test    # Run tests for both bash and Go implementations
make build   # Build Go binary (v2)
make lint    # Run shellcheck (v0.1.0) and go vet (v2)
make all     # Run lint, test, build
make clean   # Remove build artifacts
```

## Acceptance Criteria

- [ ] `Makefile` exists at project root
- [ ] `make test` runs both bash verification and `go test ./...`
- [ ] `make build` produces `bin/rooda` binary
- [ ] `make lint` runs shellcheck on rooda.sh and go vet on Go code
- [ ] `make all` runs lint, test, and build in sequence
- [ ] `make clean` removes `bin/` directory
- [ ] AGENTS.md updated to reference `make` commands as primary interface
- [ ] All make targets have `.PHONY` declarations where appropriate

## Notes

- shellcheck may not be installed on all machines; lint target should warn but not fail if missing
- Consider adding `make install` target for installing to `~/.local/bin/rooda`
- Consider adding `make help` target to list available commands
