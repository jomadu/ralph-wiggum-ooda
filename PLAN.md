# Plan: Add Makefile for Build/Test/Lint

## Overview

Add a Makefile to standardize build, test, and lint operations across bash (v0.1.0) and Go (v2) implementations, providing a consistent developer interface.

## Specification Changes

### 1. Update distribution.md spec

**Location:** `specs/distribution.md`

**Changes needed:**
- Add Makefile to build process section
- Document `make build` as primary build interface
- Update build examples to show both `make build` and direct `go build`
- Add Makefile to "Build-time" dependencies section

**Rationale:** Distribution spec covers build process; Makefile is a build tool

### 2. Update operational-knowledge.md spec (optional)

**Location:** `specs/operational-knowledge.md`

**Changes needed:**
- Add Makefile detection to bootstrap workflow
- Document that AGENTS.md should reference `make` commands when Makefile exists

**Rationale:** Bootstrap should detect and document Makefile if present

## Implementation Tasks

### Task 1: Create Makefile with core targets

**Priority:** P1 (high)

**Description:** Create `Makefile` at project root with targets: test, build, lint, all, clean

**Acceptance criteria:**
- `make test` runs bash verification (`./rooda.sh --version`, `./rooda.sh --list-procedures`) and `go test ./...`
- `make build` runs `go build -o bin/rooda ./cmd/rooda`
- `make lint` runs `shellcheck rooda.sh` (warn if missing) and `go vet ./...`
- `make all` runs lint, test, build in sequence
- `make clean` removes `bin/` directory
- All targets have `.PHONY` declarations

**Implementation notes:**
```makefile
.PHONY: test build lint all clean

test:
	@echo "Running bash verification..."
	./rooda.sh --version
	./rooda.sh --list-procedures
	@echo "Running Go tests..."
	go test ./...

build:
	@echo "Building Go binary..."
	go build -o bin/rooda ./cmd/rooda

lint:
	@echo "Running shellcheck..."
	@which shellcheck > /dev/null && shellcheck rooda.sh || echo "WARN: shellcheck not installed, skipping"
	@echo "Running go vet..."
	go vet ./...

all: lint test build

clean:
	rm -rf bin/
```

**Estimated effort:** 30 minutes

---

### Task 2: Add help target to Makefile

**Priority:** P2 (medium)

**Description:** Add `make help` target that lists available commands with descriptions

**Acceptance criteria:**
- `make help` displays all available targets with one-line descriptions
- `make` (no target) defaults to `help`

**Implementation notes:**
```makefile
.DEFAULT_GOAL := help

.PHONY: help
help:
	@echo "Available targets:"
	@echo "  make test    - Run tests for bash and Go implementations"
	@echo "  make build   - Build Go binary (bin/rooda)"
	@echo "  make lint    - Run shellcheck and go vet"
	@echo "  make all     - Run lint, test, and build"
	@echo "  make clean   - Remove build artifacts"
	@echo "  make help    - Show this help message"
```

**Estimated effort:** 15 minutes

---

### Task 3: Update AGENTS.md to reference make commands

**Priority:** P1 (high)

**Description:** Update AGENTS.md Build/Test/Lint Commands section to show `make` as primary interface

**Acceptance criteria:**
- Build/Test/Lint Commands section shows `make` commands first
- Direct commands (go test, go build, etc.) shown as alternatives
- Note added explaining Makefile provides unified interface

**Implementation notes:**

Replace current Build/Test/Lint Commands section with:

```markdown
## Build/Test/Lint Commands

**Dependencies:**
- yq >= 4.0.0 (required) — YAML parsing for rooda-config.yml
- AI CLI tool (configurable) — default: kiro-cli, can substitute with claude, aider, cursor-agent
- bd (beads CLI) — issue tracking
- Go >= 1.24.5 (required for v2 Go implementation)
- make (optional but recommended) — unified build interface

**Unified Interface (via Makefile):**
```bash
make test    # Run all tests (bash verification + Go tests)
make build   # Build Go binary
make lint    # Run all linters (shellcheck + go vet)
make all     # Run lint, test, and build
make clean   # Remove build artifacts
```

**Direct Commands (alternative):**
```bash
# Test
./rooda.sh --version              # v0.1.0 bash verification
./rooda.sh --list-procedures      # v0.1.0 bash verification
go test ./...                     # v2 Go tests
go test -v ./internal/...         # v2 Go tests (verbose)

# Build
go build -o bin/rooda ./cmd/rooda  # v2 Go binary

# Lint
shellcheck rooda.sh               # v0.1.0 bash (requires shellcheck)
go vet ./...                      # v2 Go linter
```

**Note:** The Makefile provides a unified interface across both implementations. Use `make` commands for consistency.
```

**Estimated effort:** 10 minutes

---

### Task 4: Update distribution.md spec

**Priority:** P2 (medium)

**Description:** Update distribution.md spec to document Makefile in build process

**Acceptance criteria:**
- Build Process section mentions Makefile
- Build-time dependencies include make
- Examples show both `make build` and direct `go build`

**Implementation notes:**

In `specs/distribution.md`, update "Build Process" section:

```markdown
### Build Process
```
1. Embed version metadata:
   go build -ldflags "-X main.Version=$(git describe --tags) \
                      -X main.CommitSHA=$(git rev-parse HEAD) \
                      -X main.BuildDate=$(date -u +%Y-%m-%dT%H:%M:%SZ)"

2. Embed default prompts:
   // In main package
   //go:embed prompts/*.md
   var defaultPrompts embed.FS

3. Build binary:
   # Using Makefile (recommended)
   make build
   
   # Or directly
   go build -o bin/rooda ./cmd/rooda

4. Cross-compile for each platform:
   for platform in darwin/arm64 darwin/amd64 linux/amd64 linux/arm64 windows/amd64; do
     GOOS=${platform%/*} GOARCH=${platform#*/} go build -o rooda-$platform
   done
```

Update "Build-time" dependencies:
```markdown
### Build-time
- Go 1.21+ (for `go:embed` and modern stdlib)
- git (for version metadata)
- make (optional but recommended for unified build interface)
- Cross-compilation toolchain (built into Go)
```

**Estimated effort:** 15 minutes

---

### Task 5: Test Makefile on clean checkout

**Priority:** P1 (high)

**Description:** Verify Makefile works correctly on a clean repository checkout

**Acceptance criteria:**
- Clone repository to new directory
- Run `make all` successfully
- Verify `bin/rooda` binary created
- Verify all tests pass
- Verify lint runs (with warning if shellcheck missing)

**Implementation notes:**
```bash
cd /tmp
git clone <repo-url> rooda-test
cd rooda-test
make all
./bin/rooda --version
make clean
```

**Estimated effort:** 15 minutes

---

## Task Summary

| Task | Priority | Effort | Dependencies |
|------|----------|--------|--------------|
| 1. Create Makefile | P1 | 30 min | None |
| 2. Add help target | P2 | 15 min | Task 1 |
| 3. Update AGENTS.md | P1 | 10 min | Task 1 |
| 4. Update distribution.md | P2 | 15 min | Task 1 |
| 5. Test on clean checkout | P1 | 15 min | Tasks 1, 3 |

**Total estimated effort:** 1 hour 25 minutes

**Critical path:** Tasks 1 → 3 → 5 (55 minutes)

## Implementation Order

1. **Task 1** - Create Makefile (foundation)
2. **Task 3** - Update AGENTS.md (documents new interface)
3. **Task 5** - Test on clean checkout (validation)
4. **Task 2** - Add help target (nice-to-have)
5. **Task 4** - Update distribution.md spec (documentation)

## Notes

- shellcheck is optional; Makefile should warn but not fail if missing
- Consider adding `make install` target in future (out of scope for this plan)
- Makefile should work on both macOS and Linux (use POSIX-compatible syntax)
- Go commands assume Go 1.24.5+ is installed (per AGENTS.md dependencies)
