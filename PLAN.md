# Gap Analysis Plan: v2 Go Implementation

**Status:** All 11 v2 specifications complete. Zero Go implementation exists.

**Gap:** Complete rewrite from bash (v0.1.0) to Go required.

## Priority 1: Foundation (Blocks Everything)

- [ ] Initialize Go module (go.mod)
- [ ] Create directory structure (cmd/rooda/, internal/)
- [ ] Implement iteration loop core (internal/loop/)
- [ ] Implement configuration loading (internal/config/)
- [ ] Implement prompt composition (internal/prompt/)
- [ ] Implement AI CLI execution (internal/ai/)

## Priority 2: User Interface

- [ ] Implement CLI argument parsing (cmd/rooda/main.go)
- [ ] Implement --help, --version, --list-procedures
- [ ] Implement --dry-run mode
- [ ] Implement --verbose and --quiet flags
- [ ] Implement --max-iterations and --unlimited

## Priority 3: Procedures System

- [ ] Create fragments/ directory with 55 embedded fragments
- [ ] Implement fragment loading (builtin: prefix + filesystem)
- [ ] Implement template processing (Go text/template)
- [ ] Define 16 built-in procedures in code
- [ ] Implement procedure validation

## Priority 4: Error Handling & Observability

- [ ] Implement failure detection (exit codes + promise signals)
- [ ] Implement consecutive failure tracking
- [ ] Implement timeout handling
- [ ] Implement structured logging
- [ ] Implement iteration statistics (Welford's algorithm)
- [ ] Implement signal handling (SIGINT/SIGTERM)

## Priority 5: AGENTS.md Integration

- [ ] Implement AGENTS.md parser (internal/agents/)
- [ ] Implement empirical verification
- [ ] Implement drift detection
- [ ] Implement in-place updates
- [ ] Implement bootstrap workflow

## Priority 6: Distribution

- [ ] Implement version embedding (-ldflags)
- [ ] Create build script (scripts/build.sh)
- [ ] Create install script (scripts/install.sh)
- [ ] Embed default prompts (go:embed)
- [ ] Cross-compile for platforms (darwin, linux, windows)
- [ ] Generate checksums

## Priority 7: Quality Gates

- [ ] Write unit tests for core modules
- [ ] Write integration tests for procedures
- [ ] Set up CI/CD pipeline (.github/workflows/)
- [ ] Configure golangci-lint
- [ ] Verify all acceptance criteria from specs

## Dependencies

- Go 1.21+ (for go:embed)
- yq removed (Go YAML library replaces it)
- All other dependencies remain (AI CLI tools, bd, etc.)

## Notes

- v0.1.0 bash implementation remains functional during v2 development
- Bash implementation archived after v2 reaches feature parity
- Fragment-based prompt system is new in v2 (v1 used monolithic prompts)
- Built-in procedures expand from 9 to 16 in v2
