# Gap Analysis Plan: Spec to Implementation

**Generated:** 2026-02-09
**Source:** Gap analysis comparing v2 Go rewrite specifications against current bash implementation

## Executive Summary

The project is in a **specification-complete, implementation-missing** state. All 11 v2 Go rewrite specifications are complete with JTBD structure, acceptance criteria, and examples. The current implementation is the v0.1.0 bash script (`rooda.sh`) which implements a subset of v2 features. **No Go implementation exists yet** — no `go.mod`, no `*.go` files, no Go project structure.

The 10 ready issues in beads correctly identify the highest-priority implementation gaps. This plan organizes them by dependency order and adds missing foundational tasks.

## Critical Gaps (P0) - Foundation Required for Any Implementation

### 1. Initialize Go project structure
**Priority:** P0 (blocks all other work)
**Description:** Create Go module, directory structure (cmd/, internal/, pkg/), and main.go entry point. Establish build toolchain.
**Acceptance:**
- `go.mod` exists with module path `github.com/jomadu/rooda`
- `cmd/rooda/main.go` exists with version flag working
- `go build ./cmd/rooda` produces binary
- Directory structure: `cmd/rooda/`, `internal/config/`, `internal/loop/`, `internal/prompt/`, `internal/ai/`, `internal/agents/`
**Dependencies:** None
**Beads issue:** Not filed yet

### 2. Embed default prompts and procedures (ralph-wiggum-ooda-hrz0)
**Priority:** P0 (required for zero-config startup)
**Description:** Use `//go:embed` to embed 25 prompt files from `prompts/` directory. Define 16 built-in procedures in `internal/config/defaults.go`.
**Acceptance:** Binary contains embedded prompts, `rooda --list-procedures` shows all 16.
**Dependencies:** Task 1 (Go project structure)
**Beads issue:** ralph-wiggum-ooda-hrz0

### 3. Implement configuration system (configuration.md)
**Priority:** P0 (required for all procedures)
**Description:** Three-tier config system (built-in defaults, global config, workspace config, env vars, CLI flags). YAML parsing, provenance tracking, validation.
**Acceptance:**
- Config loads from all tiers with correct precedence
- `--dry-run` displays resolved config with provenance
- Invalid config produces clear error messages
- Zero-config startup works (uses built-in defaults)
**Dependencies:** Task 2 (embedded defaults)
**Beads issue:** Not filed yet

### 4. Implement CLI interface (cli-interface.md)
**Priority:** P0 (required for all procedures)
**Description:** Argument parsing, procedure invocation, help text, flag handling. Support for `--help`, `--version`, `--list-procedures`, `--dry-run`, `--verbose`, `--context`, OODA phase overrides.
**Acceptance:**
- `rooda <procedure>` invokes named procedure
- `rooda --help` displays usage
- `rooda --list-procedures` lists all procedures
- `rooda --version` displays version
- All flags from cli-interface.md acceptance criteria work
**Dependencies:** Task 3 (configuration system)
**Beads issue:** Not filed yet

## High Priority (P1) - Core Loop Functionality

### 5. Implement basic iteration loop (ralph-wiggum-ooda-gm29)
**Priority:** P1
**Description:** Define IterationState, LoopStatus types. Implement loop termination logic (max iterations, Ctrl+C). Implement iteration counter and timing.
**Acceptance:** `rooda build --max-iterations 3` runs 3 iterations and exits.
**Dependencies:** Task 4 (CLI interface)
**Beads issue:** ralph-wiggum-ooda-gm29

### 6. Implement prompt composition (prompt-composition.md)
**Priority:** P1 (required for loop to assemble prompts)
**Description:** Load fragments from embedded resources or filesystem. Process Go templates. Concatenate fragments by phase. Inject user context.
**Acceptance:**
- Prompts assembled from fragment arrays
- `builtin:` prefix loads embedded fragments
- Relative paths load from filesystem
- Template processing works with parameters
- `--context` flag injects user context
**Dependencies:** Task 2 (embedded prompts), Task 3 (configuration)
**Beads issue:** Not filed yet

### 7. Implement AI CLI integration (ai-cli-integration.md)
**Priority:** P1 (required for loop to execute AI)
**Description:** Resolve AI command from precedence chain. Spawn AI CLI process. Pipe prompt to stdin. Capture stdout/stderr. Stream output when `--verbose`.
**Acceptance:**
- AI command resolved from config/flags
- Built-in aliases work (kiro-cli, claude, copilot, cursor-agent)
- Output captured and returned
- `--verbose` streams output to terminal
**Dependencies:** Task 3 (configuration)
**Beads issue:** Not filed yet

### 8. Implement promise signal scanning (ralph-wiggum-ooda-sont)
**Priority:** P1
**Description:** Scan AI CLI output for `<promise>SUCCESS</promise>` and `<promise>FAILURE</promise>`. Implement outcome matrix (exit code + output signals → iteration outcome). Terminate loop on SUCCESS signal.
**Acceptance:** Loop terminates when AI outputs `<promise>SUCCESS</promise>`.
**Dependencies:** Task 5 (basic loop), Task 7 (AI CLI integration)
**Beads issue:** ralph-wiggum-ooda-sont

### 9. Implement failure tracking (ralph-wiggum-ooda-pn2x)
**Priority:** P1
**Description:** Track ConsecutiveFailures counter. Reset counter on success. Abort loop when threshold exceeded.
**Acceptance:** Loop aborts after 3 consecutive failures (default threshold).
**Dependencies:** Task 8 (promise signal scanning)
**Beads issue:** ralph-wiggum-ooda-pn2x

### 10. Implement iteration timeouts (ralph-wiggum-ooda-ac1c)
**Priority:** P1
**Description:** Add `iteration_timeout` config field. Kill AI CLI process if timeout exceeded. Count timeout as failure.
**Acceptance:** Loop kills AI CLI after configured timeout, increments failure counter.
**Dependencies:** Task 7 (AI CLI integration), Task 9 (failure tracking)
**Beads issue:** ralph-wiggum-ooda-ac1c

### 11. Implement signal handling (ralph-wiggum-ooda-07xa)
**Priority:** P1
**Description:** Register SIGINT/SIGTERM handlers. Kill AI CLI process on interrupt. Wait for cleanup (5s timeout). Exit with code 130.
**Acceptance:** Ctrl+C kills AI CLI cleanly, exits with code 130.
**Dependencies:** Task 7 (AI CLI integration)
**Beads issue:** ralph-wiggum-ooda-07xa

## Medium Priority (P2) - Enhanced Functionality

### 12. Implement dry-run mode (ralph-wiggum-ooda-xdf7)
**Priority:** P2
**Description:** Add `--dry-run` flag. Validate config, prompt files, AI command. Display assembled prompt and resolved config with provenance. Exit with code 0 (success) or 1 (validation failed).
**Acceptance:** `rooda build --dry-run` validates and displays prompt without executing.
**Dependencies:** Task 6 (prompt composition), Task 7 (AI CLI integration)
**Beads issue:** ralph-wiggum-ooda-xdf7

### 13. Implement iteration statistics (ralph-wiggum-ooda-jobd)
**Priority:** P2
**Description:** Track min/max/mean/stddev using Welford's online algorithm. Display statistics at loop completion. Use constant memory (O(1)).
**Acceptance:** Loop displays 'Iteration timing: count=N min=Xs max=Xs mean=Xs stddev=Xs'.
**Dependencies:** Task 5 (basic loop)
**Beads issue:** ralph-wiggum-ooda-jobd

### 14. Implement context injection (ralph-wiggum-ooda-hur3)
**Priority:** P2
**Description:** Add `--context <text>` and `--context-file <path>` flags. Accumulate multiple contexts. Inject as dedicated section before OODA phases.
**Acceptance:** `rooda build --context 'focus on auth'` injects context into prompt.
**Dependencies:** Task 6 (prompt composition)
**Beads issue:** ralph-wiggum-ooda-hur3

### 15. Implement provenance display (ralph-wiggum-ooda-rayx)
**Priority:** P2
**Description:** Track where each config value came from. Display provenance in dry-run mode. Display provenance with `--verbose`.
**Acceptance:** Dry-run shows 'max_iterations: 10 (from: workspace config)'.
**Dependencies:** Task 3 (configuration system), Task 12 (dry-run mode)
**Beads issue:** ralph-wiggum-ooda-rayx

### 16. Implement structured logging (observability.md)
**Priority:** P2
**Description:** Log events at four levels (debug, info, warn, error). Configurable log level and timestamp format. Structured fields (logfmt). Progress display.
**Acceptance:**
- Log level configurable via config/env/flags
- Timestamp format configurable
- Progress messages at iteration start/complete
- Loop completion displays status and timing
**Dependencies:** Task 5 (basic loop)
**Beads issue:** Not filed yet

## Lower Priority (P3) - Advanced Features

### 17. Implement AGENTS.md parser (operational-knowledge.md, agents-md-format.md)
**Priority:** P3 (required for procedures to read project conventions)
**Description:** Parse AGENTS.md into structured data. Extract build commands, test commands, spec paths, impl paths, work tracking config, quality criteria.
**Acceptance:**
- AGENTS.md parsed into AgentsMD struct
- All 10 required sections extracted
- Validation detects missing sections
**Dependencies:** Task 3 (configuration system)
**Beads issue:** Not filed yet

### 18. Implement AGENTS.md bootstrap (operational-knowledge.md)
**Priority:** P3
**Description:** Detect build system, test system, spec paths, impl paths, work tracking system. Generate AGENTS.md from template with detected values.
**Acceptance:**
- Bootstrap detects Go project structure
- Bootstrap creates AGENTS.md with detected values
- Bootstrap commits AGENTS.md
**Dependencies:** Task 17 (AGENTS.md parser)
**Beads issue:** Not filed yet

### 19. Implement AGENTS.md empirical verification (operational-knowledge.md)
**Priority:** P3
**Description:** Verify build/test commands execute. Verify file paths exist. Detect drift between AGENTS.md and reality. Update AGENTS.md when drift detected.
**Acceptance:**
- Commands from AGENTS.md are executed and verified
- Path patterns validated against filesystem
- Drift detection updates AGENTS.md with rationale
**Dependencies:** Task 17 (AGENTS.md parser)
**Beads issue:** Not filed yet

### 20. Implement distribution (distribution.md)
**Priority:** P3 (required for installation)
**Description:** Cross-compile for macOS/Linux/Windows. Generate checksums. Create Homebrew formula. Create install script with checksum verification.
**Acceptance:**
- Binaries built for all platforms
- SHA256 checksums generated
- Install script verifies checksums
- `rooda --version` reports correct version
**Dependencies:** Task 1 (Go project structure)
**Beads issue:** Not filed yet

## Specification Gaps (None Identified)

All 11 specifications are complete:
- ✅ cli-interface.md
- ✅ configuration.md
- ✅ prompt-composition.md
- ✅ procedures.md
- ✅ iteration-loop.md
- ✅ error-handling.md
- ✅ ai-cli-integration.md
- ✅ observability.md
- ✅ operational-knowledge.md
- ✅ agents-md-format.md
- ✅ distribution.md

All specs have:
- ✅ Job to be Done section
- ✅ Acceptance Criteria section
- ✅ Examples section
- ✅ Data Structures section
- ✅ Algorithm section (where applicable)

## Implementation Status Summary

**Current implementation (v0.1.0 bash):**
- ✅ Basic OODA loop with iteration counting
- ✅ Prompt composition from 4 phase files
- ✅ AI CLI integration (pipes to kiro-cli/claude/aider)
- ✅ YAML config parsing (via yq)
- ✅ 9 procedures defined in rooda-config.yml
- ✅ 25 prompt files in prompts/
- ❌ No promise signal scanning
- ❌ No failure tracking
- ❌ No timeouts
- ❌ No dry-run mode
- ❌ No verbose mode
- ❌ No context injection
- ❌ No provenance tracking
- ❌ No AGENTS.md lifecycle
- ❌ No embedded prompts (requires external files)

**v2 Go implementation:**
- ❌ No Go project structure
- ❌ No Go code exists
- ❌ All features unimplemented

## Recommended Execution Order

**Phase 1: Foundation (P0 tasks 1-4)**
Execute sequentially — each blocks the next:
1. Initialize Go project structure
2. Embed default prompts and procedures
3. Implement configuration system
4. Implement CLI interface

**Phase 2: Core Loop (P1 tasks 5-11)**
Execute in dependency order:
1. Basic iteration loop (task 5)
2. Prompt composition (task 6) — parallel with AI CLI integration (task 7)
3. Promise signal scanning (task 8)
4. Failure tracking (task 9)
5. Iteration timeouts (task 10) — parallel with signal handling (task 11)

**Phase 3: Enhanced Features (P2 tasks 12-16)**
Can execute in parallel after Phase 2 complete:
- Dry-run mode (task 12)
- Iteration statistics (task 13)
- Context injection (task 14)
- Provenance display (task 15)
- Structured logging (task 16)

**Phase 4: Advanced Features (P3 tasks 17-20)**
Execute after Phase 3:
- AGENTS.md parser (task 17)
- AGENTS.md bootstrap (task 18) — parallel with verification (task 19)
- Distribution (task 20) — can execute anytime after task 1

## Notes

**Why no Go code exists:**
- Project is on `goify` branch which restructured files but didn't implement Go rewrite
- Bash v0.1.0 implementation moved to root level
- v1 artifacts archived in `archive/`
- Specifications written but implementation not started

**Why beads issues are correct:**
- All 10 ready issues map to P1-P2 tasks in this plan
- Issues correctly identify iteration loop, promise signals, failure tracking, timeouts, signal handling, dry-run, statistics, context injection, provenance as priorities
- Missing issues: Go project init (P0), config system (P0), CLI interface (P0), prompt composition (P1), AI CLI integration (P1), logging (P2), AGENTS.md lifecycle (P3), distribution (P3)

**Why this order:**
- P0 tasks are foundation — nothing works without them
- P1 tasks are core loop — minimum viable OODA iteration
- P2 tasks are quality-of-life — observability and control
- P3 tasks are advanced — AGENTS.md lifecycle and distribution

**Estimated effort:**
- Phase 1 (P0): 3-5 iterations (foundation is critical, must be solid)
- Phase 2 (P1): 5-8 iterations (core loop complexity)
- Phase 3 (P2): 3-5 iterations (enhancements are simpler)
- Phase 4 (P3): 5-8 iterations (AGENTS.md lifecycle is complex)
- **Total: 16-26 iterations** for complete v2 implementation

**Risk areas:**
- Go template processing in prompt composition (task 6)
- Process management for AI CLI (task 7) — stdin piping, output capture, signal handling
- Promise signal scanning edge cases (task 8) — partial output, truncation
- AGENTS.md drift detection (task 19) — distinguishing command failure from drift
