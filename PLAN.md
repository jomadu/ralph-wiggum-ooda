# Draft Plan: Spec-to-Implementation Gap Closure

**Status:** Draft (not yet published to work tracking)

**Context:** v2 specifications describe a Go rewrite with advanced features (promise signals, failure tracking, timeouts, dry-run, context injection, signal handling, iteration statistics, audit procedures). Current implementation is v0.1.0 bash with basic OODA loop functionality. Gap analysis identified 13 critical missing features and 4 lower-priority improvements.

**Strategy:** Implement v2 Go rewrite as the foundation, then layer features in dependency order.

---

## Priority 1: Foundation (Go Rewrite Core)

### Task 1.1: Go Project Bootstrap
- Initialize Go module structure (cmd/, internal/, pkg/)
- Set up build system (Makefile, go.mod, go.sum)
- Embed default prompt files using go:embed
- Create main.go entry point with version flag
- **Acceptance:** `go build` produces working binary, `./rooda --version` displays version

### Task 1.2: Configuration System
- Implement 3-tier config loading (workspace, global, env vars)
- Implement YAML parsing for rooda-config.yml
- Implement config validation (required fields, prompt file existence)
- Implement provenance tracking (record where each value came from)
- Implement AI command resolution (direct, alias, env var, default)
- **Acceptance:** Config loads from all tiers with correct precedence, validation catches errors, provenance displayed in dry-run

### Task 1.3: CLI Argument Parsing
- Implement flag parsing with cobra or stdlib flag
- Support long flags (--max-iterations, --dry-run, --verbose, --quiet, --context, --context-file, --ai-cmd, --ai-cmd-alias, --observe, --orient, --decide, --act, --config, --log-level, --unlimited)
- Support short flags (-v, -q, -n, -u, -d, -c)
- Implement mutual exclusivity checks (--verbose/--quiet, --max-iterations/--unlimited)
- Implement help text generation (global help, procedure-specific help, --list-procedures)
- **Acceptance:** All flags parse correctly, help text displays, mutual exclusivity enforced

### Task 1.4: Prompt Composition
- Implement prompt assembly from 4 OODA phase files
- Support embedded default prompts (go:embed)
- Support filesystem prompt overrides
- Implement context injection (--context and --context-file)
- Implement path resolution (relative to config file directory)
- **Acceptance:** Prompts assemble correctly, context injected in correct location, overrides work

---

## Priority 2: Core Loop Features

### Task 2.1: Basic Iteration Loop
- Implement iteration loop with counter and max iterations
- Implement AI CLI process spawning (exec.Command)
- Implement output capture (stdout/stderr)
- Implement iteration progress display
- Implement loop termination (max iterations, Ctrl+C)
- **Acceptance:** Loop executes N iterations, displays progress, terminates correctly

### Task 2.2: Promise Signal Detection
- Implement output scanning for `<promise>SUCCESS</promise>` and `<promise>FAILURE</promise>`
- Implement outcome determination (promise signals override exit code)
- Implement loop termination on SUCCESS signal
- Implement FAILURE precedence when both signals present
- **Acceptance:** Loop terminates on SUCCESS signal, FAILURE signal overrides exit code 0, both signals → FAILURE wins

### Task 2.3: Failure Threshold Tracking
- Implement ConsecutiveFailures counter
- Implement failure threshold check (default: 3)
- Implement counter reset on success
- Implement loop abort on threshold exceeded
- Implement configurable threshold (loop.failure_threshold, procedure.failure_threshold)
- **Acceptance:** Loop aborts after 3 consecutive failures, counter resets on success, threshold configurable

### Task 2.4: Iteration Timeout
- Implement per-iteration timeout (loop.iteration_timeout, procedure.iteration_timeout)
- Implement process killing (SIGTERM → wait 5s → SIGKILL)
- Implement timeout as failure (increments ConsecutiveFailures)
- Implement partial output capture from timed-out processes
- **Acceptance:** Timed-out iterations killed and counted as failures, partial output captured

### Task 2.5: Signal Handling
- Implement SIGINT/SIGTERM handler
- Implement AI CLI process cleanup on interrupt
- Implement graceful termination (kill AI CLI, wait, exit with code 130)
- Implement zombie process prevention
- **Acceptance:** Ctrl+C kills AI CLI cleanly, no zombie processes, exit code 130

---

## Priority 3: Observability & Usability

### Task 3.1: Dry-Run Mode
- Implement --dry-run flag
- Implement config validation without execution
- Implement prompt assembly display
- Implement AI command validation (binary exists and executable)
- Implement exit codes (0=valid, 1=user error, 2=config error)
- **Acceptance:** Dry-run validates config, displays assembled prompt, checks AI command, exits without executing

### Task 3.2: Iteration Statistics
- Implement timing tracking (start/end per iteration)
- Implement running statistics (min, max, mean, stddev)
- Implement constant-memory statistics (O(1) space)
- Implement statistics display at loop completion
- Implement conditional stddev display (omit when count < 2)
- **Acceptance:** Statistics displayed at completion, correct values, constant memory usage

### Task 3.3: Output Buffering
- Implement configurable output buffer (loop.max_output_buffer, default: 10MB)
- Implement buffer truncation from beginning when exceeded
- Implement warning log on truncation
- Implement per-procedure buffer override (procedure.max_output_buffer)
- **Acceptance:** Output buffered up to limit, truncated from beginning when exceeded, warning logged

### Task 3.4: Logging System
- Implement structured logging (iteration, command, exit code, duration)
- Implement log levels (debug, info, warn, error)
- Implement configurable log level (loop.log_level, --log-level, ROODA_LOOP_LOG_LEVEL)
- Implement --verbose and --quiet overrides
- **Acceptance:** Logs structured, log level configurable, --verbose/--quiet work

### Task 3.5: Exit Code Semantics
- Implement exit code constants (0=success, 1=user error, 2=config error, 3=execution error, 130=interrupted)
- Implement exit code selection based on loop status
- Implement status tracking (success, max-iters, aborted, interrupted)
- **Acceptance:** Exit codes match spec, status tracked correctly

---

## Priority 4: Procedures

### Task 4.1: Procedure Metadata
- Implement procedure categories (direct-action, audit, planning)
- Implement procedure validation (name, description, OODA files, iteration limits)
- Implement --list-procedures with categories
- Implement procedure-specific help text
- **Acceptance:** Procedures categorized, validation works, help text displays

### Task 4.2: Audit Procedures (5 procedures)
- Create prompts for audit-spec (observe_specs.md, orient_quality.md, decide_audit_report.md, act_audit_report.md)
- Create prompts for audit-impl (observe_impl.md, orient_quality.md, decide_audit_report.md, act_audit_report.md)
- Create prompts for audit-agents (observe_agents.md, orient_agents_accuracy.md, decide_audit_report.md, act_audit_report.md)
- Create prompts for audit-spec-to-impl (observe_plan_specs_impl.md, orient_gap.md, decide_audit_report.md, act_audit_report.md)
- Create prompts for audit-impl-to-spec (observe_plan_specs_impl.md, orient_gap.md, decide_audit_report.md, act_audit_report.md)
- Add procedure definitions to rooda-config.yml
- **Acceptance:** All 5 audit procedures defined, prompts exist, procedures execute

### Task 4.3: Iteration Mode System
- Implement IterationMode enum (max-iterations, unlimited, promise-driven)
- Implement mode resolution (CLI > procedure > loop > default)
- Implement promise-driven mode (loop until SUCCESS signal)
- Implement --unlimited flag
- **Acceptance:** All 3 modes work, resolution precedence correct, --unlimited overrides max iterations

---

## Priority 5: Distribution & Polish

### Task 5.1: Single Binary Distribution
- Implement go:embed for all default prompts
- Implement build script for cross-platform binaries (macOS, Linux)
- Implement version embedding (ldflags)
- Create release artifacts (tar.gz, checksums)
- **Acceptance:** Single binary runs without external dependencies, version embedded, cross-platform builds work

### Task 5.2: Homebrew Distribution
- Create Homebrew formula (docs/HOMEBREW_SETUP.md already exists)
- Test installation via brew
- Document installation process
- **Acceptance:** `brew install rooda` works, binary in PATH

### Task 5.3: Documentation
- Update README.md for v2
- Update AGENTS.md with v2 commands
- Create CHANGELOG.md
- Document migration from v0.1.0 bash to v2 Go
- **Acceptance:** Documentation complete, migration guide clear

---

## Priority 6: Testing & Quality

### Task 6.1: Unit Tests
- Write tests for config loading and validation
- Write tests for prompt composition
- Write tests for promise signal detection
- Write tests for failure tracking
- Write tests for iteration statistics
- **Acceptance:** Core logic covered by unit tests, tests pass

### Task 6.2: Integration Tests
- Write tests for full loop execution
- Write tests for timeout handling
- Write tests for signal handling
- Write tests for dry-run mode
- **Acceptance:** End-to-end scenarios covered, tests pass

### Task 6.3: Linting & Formatting
- Set up golangci-lint
- Configure linters (gofmt, govet, staticcheck, errcheck)
- Add lint command to Makefile
- Fix all lint errors
- **Acceptance:** `make lint` passes with no errors

---

## Task Count Summary

- **Foundation:** 4 tasks
- **Core Loop:** 5 tasks
- **Observability:** 5 tasks
- **Procedures:** 3 tasks
- **Distribution:** 3 tasks
- **Testing:** 3 tasks

**Total:** 23 tasks

---

## Dependencies

```
1.1 (Go Bootstrap) → 1.2 (Config) → 1.3 (CLI) → 1.4 (Prompts)
                   ↓
                  2.1 (Basic Loop) → 2.2 (Promise Signals) → 2.3 (Failure Tracking)
                                  ↓
                                 2.4 (Timeout) → 2.5 (Signal Handling)
                                  ↓
                                 3.1 (Dry-Run)
                                  ↓
                                 3.2 (Statistics) → 3.3 (Buffering) → 3.4 (Logging) → 3.5 (Exit Codes)
                                  ↓
                                 4.1 (Metadata) → 4.2 (Audit Procedures) → 4.3 (Iteration Modes)
                                  ↓
                                 5.1 (Binary) → 5.2 (Homebrew) → 5.3 (Docs)
                                  ↓
                                 6.1 (Unit Tests) → 6.2 (Integration Tests) → 6.3 (Linting)
```

---

## Notes

- **Bash implementation preservation:** v0.1.0 bash implementation remains functional during Go rewrite. No breaking changes to existing workflows until v2 is ready.
- **Incremental delivery:** Each task produces a working artifact. Foundation tasks produce a minimal working binary. Core loop tasks add features incrementally.
- **Prompt reuse:** Many audit procedures reuse existing prompts (observe_specs.md, observe_impl.md, orient_gap.md, orient_quality.md). Only need to create audit-specific decide/act prompts.
- **Testing strategy:** Unit tests written alongside implementation (not deferred to end). Integration tests added after core loop complete.
- **Migration path:** v0.1.0 bash and v2 Go can coexist. Users migrate when v2 reaches feature parity (after Task 4.3).
