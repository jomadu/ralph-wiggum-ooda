# Gap Analysis Plan: v2 Specs vs v0.1.0 Implementation

**Generated:** 2026-02-08T22:23:26-08:00  
**Source:** Gap analysis comparing v2 specifications to v0.1.0 bash implementation  
**Status:** Draft

## Executive Summary

v2 specs define 16 procedures and 13 critical features for a Go rewrite. v0.1.0 bash implements 9 of 16 procedures (56%) and 0 of 13 critical features (0%). The bash implementation is a working proof-of-concept that validates the core OODA loop pattern but lacks the robustness, observability, and extensibility specified in v2.

**Root cause:** v2 specs describe a Go rewrite with production-grade features; v0.1.0 is a bash prototype.

**Resolution path:** Implement missing procedures in bash (short-term) OR proceed with Go rewrite per PLAN.md (long-term, already planned).

## Gap Summary

### Missing Procedures (7 of 16)

**Audit procedures (5):**
- audit-spec — Quality assessment of specs
- audit-impl — Quality assessment of implementation  
- audit-agents — Accuracy assessment of AGENTS.md vs repo state
- (audit-spec-to-impl and audit-impl-to-spec exist but named draft-plan-*)

**Planning procedures (2):**
- draft-plan-spec-chore — Plan spec maintenance
- draft-plan-impl-feat — Plan new capability implementation
- draft-plan-impl-fix — Plan implementation correction
- draft-plan-impl-chore — Plan implementation maintenance

### Missing Critical Features (13 of 13)

**Loop control:**
1. Promise signals — AI can't signal completion (`<promise>SUCCESS</promise>`)
2. Failure tracking — No consecutive failure counter or abort threshold
3. Iteration modes — Only supports max-iterations, not unlimited mode
4. Exit code semantics — No standardized exit codes (0/1/2/130)

**Observability:**
5. Iteration statistics — No min/max/mean/stddev timing
6. Dry-run mode — Can't validate config without executing
7. Provenance tracking — Can't see where config values came from

**Robustness:**
8. Timeouts — No iteration timeout, processes can run forever
9. Signal handling — SIGINT/SIGTERM not handled cleanly
10. Output buffering — No max buffer size, no truncation handling

**Configuration:**
11. Context injection — No `--context` or `--context-file` flags
12. Config tiers — Only workspace config, no global/env vars/built-in defaults
13. Go implementation — Bash is platform-specific, not cross-platform binary

### Undocumented Features (4)

Bash implementation has features not in specs:
1. Git push automation with fallback (`git push` → `git push -u`)
2. Platform detection (macOS vs Linux install instructions)
3. Fuzzy procedure name matching (suggests closest match on typo)
4. AI tool preset resolution (hardcoded + custom from config)

## Priority 0: Critical Gaps (Blocking)

### Task 1: Implement promise signal scanning

**Description:** Add `<promise>SUCCESS</promise>` and `<promise>FAILURE</promise>` detection to loop.

**Acceptance criteria:**
- Loop scans AI CLI output for promise signals after each iteration
- `<promise>SUCCESS</promise>` terminates loop with exit code 0
- `<promise>FAILURE</promise>` increments consecutive failure counter
- If both signals present, FAILURE takes precedence

**Rationale:** Without promise signals, loop can't detect when AI has completed work. Currently relies only on max iterations.

**Estimated effort:** 2 hours (add grep to output, update termination logic)

---

### Task 2: Implement failure tracking

**Description:** Add consecutive failure counter and abort threshold.

**Acceptance criteria:**
- Track consecutive AI CLI failures (non-zero exit code OR `<promise>FAILURE</promise>`)
- Reset counter to 0 on any successful iteration
- Abort loop when consecutive failures >= threshold (default: 3)
- Exit with code 1 when aborted

**Rationale:** Without failure tracking, loop continues indefinitely on repeated failures, wasting API calls.

**Estimated effort:** 3 hours (add counter, threshold config, abort logic)

---

### Task 3: Add missing audit procedures

**Description:** Implement 5 missing audit procedures.

**Acceptance criteria:**
- audit-spec procedure exists in rooda-config.yml
- audit-impl procedure exists in rooda-config.yml
- audit-agents procedure exists in rooda-config.yml
- All use existing prompt files (observe_specs.md, observe_impl.md, observe_bootstrap.md)
- All have default_iterations: 1

**Rationale:** Audits are read-only assessments that feed into planning. Missing audits means no quality gates before planning.

**Estimated effort:** 1 hour (config entries only, prompts already exist)

---

### Task 4: Add missing planning procedures

**Description:** Implement 4 missing planning procedures.

**Acceptance criteria:**
- draft-plan-spec-chore procedure exists
- draft-plan-impl-feat procedure exists
- draft-plan-impl-fix procedure exists
- draft-plan-impl-chore procedure exists
- All use appropriate prompt files (may need to create new prompts)

**Rationale:** Missing planning procedures means incomplete coverage of conventional commit types (feat/fix/refactor/chore).

**Estimated effort:** 4 hours (2 hours config, 2 hours new prompts if needed)

**Dependencies:** May require new prompt files if existing prompts don't cover these use cases.

## Priority 1: High-Impact Gaps

### Task 5: Implement dry-run mode

**Description:** Add `--dry-run` flag that validates config and displays assembled prompt without executing.

**Acceptance criteria:**
- `--dry-run` flag accepted by CLI
- Validates all OODA phase files exist
- Validates AI command binary exists and is executable
- Displays assembled prompt with section markers
- Displays resolved configuration
- Exits with code 0 if validation passes, 1 if fails
- Does not execute AI CLI

**Rationale:** Dry-run is essential for debugging config issues and verifying prompt assembly before execution.

**Estimated effort:** 4 hours (add flag parsing, validation logic, display formatting)

---

### Task 6: Implement context injection

**Description:** Add `--context` and `--context-file` flags for user-provided context.

**Acceptance criteria:**
- `--context <text>` flag injects text into prompt
- `--context-file <path>` flag reads file and injects content
- Multiple `--context` flags accumulate
- Multiple `--context-file` flags accumulate
- Context appears as dedicated section before OODA phases
- Context passed through verbatim (not interpreted)

**Rationale:** Context injection allows steering the agent without modifying prompt files.

**Estimated effort:** 3 hours (add flag parsing, prompt assembly modification)

---

### Task 7: Implement iteration statistics

**Description:** Add timing statistics (min/max/mean/stddev) displayed at loop completion.

**Acceptance criteria:**
- Track start/end time for each iteration
- Calculate min, max, mean, stddev using Welford's online algorithm
- Display statistics at loop completion (info level)
- Format: "Iteration timing: count=N min=Xs max=Xs mean=Xs stddev=Xs"
- Omit stddev when count < 2
- Use constant memory (O(1)) regardless of iteration count

**Rationale:** Statistics help identify slow iterations and tune iteration limits.

**Estimated effort:** 4 hours (add timing capture, statistics calculation, display formatting)

## Priority 2: Medium-Impact Gaps

### Task 8: Implement iteration timeout

**Description:** Add per-iteration timeout that kills AI CLI process if exceeded.

**Acceptance criteria:**
- `loop.iteration_timeout` config field (seconds, nil = no timeout)
- `procedure.iteration_timeout` overrides loop-level
- `ROODA_LOOP_ITERATION_TIMEOUT` env var sets loop-level
- If AI CLI exceeds timeout, process killed
- Iteration counts as failure (increments consecutive failure counter)
- Warning logged when timeout occurs

**Rationale:** Timeouts prevent runaway AI processes from blocking the loop indefinitely.

**Estimated effort:** 5 hours (add timeout logic, process kill, config fields)

**Dependencies:** Requires bash timeout command or custom implementation.

---

### Task 9: Implement signal handling

**Description:** Handle SIGINT/SIGTERM cleanly.

**Acceptance criteria:**
- SIGINT (Ctrl+C) kills AI CLI process and exits with code 130
- SIGTERM kills AI CLI process and exits with code 130
- Wait up to 5s for AI CLI to terminate after kill signal
- Log warning if AI CLI doesn't terminate within timeout
- Exit anyway after timeout (don't hang forever)

**Rationale:** Clean shutdown prevents orphaned processes and corrupted state.

**Estimated effort:** 4 hours (add trap handlers, process kill, timeout logic)

---

### Task 10: Implement provenance tracking

**Description:** Track where each config value came from (CLI flag, env var, workspace config, built-in default).

**Acceptance criteria:**
- Provenance tracked for all resolved settings
- `--verbose` displays provenance at startup
- Format: "setting: value (source: tier)"
- Helps debug "where did this value come from?" questions

**Rationale:** Provenance is essential for debugging config issues in multi-tier systems.

**Estimated effort:** 6 hours (add tracking data structure, display logic)

**Dependencies:** Requires config tier implementation (Task 12).

## Priority 3: Low-Impact Gaps

### Task 11: Standardize exit codes

**Description:** Use consistent exit codes across all termination paths.

**Acceptance criteria:**
- Exit code 0: Success (AI signaled completion)
- Exit code 1: Aborted (failure threshold exceeded)
- Exit code 2: Max iterations reached
- Exit code 130: Interrupted (SIGINT/SIGTERM)
- All exit paths use correct codes

**Rationale:** Standardized exit codes enable scripting and CI/CD integration.

**Estimated effort:** 2 hours (update all exit statements)

---

### Task 12: Implement config tiers

**Description:** Add global config, environment variables, and built-in defaults.

**Acceptance criteria:**
- Global config at `~/.config/rooda/rooda-config.yml` (or `$XDG_CONFIG_HOME/rooda/`)
- `ROODA_CONFIG_HOME` env var overrides global config directory
- Environment variables with `ROODA_` prefix override config files
- Built-in defaults for all settings
- Precedence: CLI flags > env vars > workspace > global > built-in
- Workspace config overrides global for overlapping fields

**Rationale:** Config tiers enable personal preferences (global) and project-specific settings (workspace) without conflict.

**Estimated effort:** 8 hours (add global config loading, env var parsing, merge logic)

**Dependencies:** Bash implementation is complex; may be better to wait for Go rewrite.

---

### Task 13: Implement output buffering

**Description:** Add max output buffer size with truncation from beginning.

**Acceptance criteria:**
- `loop.max_output_buffer` config field (bytes, default: 10485760 = 10MB)
- `procedure.max_output_buffer` overrides loop-level
- If AI CLI output exceeds buffer, truncate from beginning
- Keep most recent output (for signal scanning)
- Warning logged when truncation occurs
- `Truncated=true` flag in result

**Rationale:** Output buffering prevents memory exhaustion on verbose AI output.

**Estimated effort:** 5 hours (add buffer size tracking, truncation logic)

**Dependencies:** Bash string manipulation may be inefficient; consider Go rewrite.

---

### Task 14: Implement unlimited iteration mode

**Description:** Add `--unlimited` flag and `iteration_mode` config field.

**Acceptance criteria:**
- `--unlimited` CLI flag sets iteration mode to unlimited
- `loop.iteration_mode` config field ("max-iterations" or "unlimited")
- `procedure.iteration_mode` overrides loop-level
- Unlimited mode runs until SUCCESS signal, failure threshold, or Ctrl+C
- `default_max_iterations` ignored when mode is unlimited

**Rationale:** Unlimited mode is useful for long-running procedures that rely on AI signals for termination.

**Estimated effort:** 3 hours (add mode field, update termination logic)

## Priority 4: Documentation Gaps

### Task 15: Document undocumented bash features

**Description:** Add undocumented bash features to AGENTS.md Operational Learnings.

**Acceptance criteria:**
- Git push automation with fallback documented
- Platform detection documented
- Fuzzy procedure name matching documented
- AI tool preset resolution documented
- Each feature has rationale explaining why it exists

**Rationale:** Undocumented features are invisible to agents and users, leading to confusion.

**Estimated effort:** 1 hour (update AGENTS.md)

## Dependencies

```
Task 2 (failure tracking) → Task 1 (promise signals)
Task 8 (timeout) → Task 2 (failure tracking)
Task 9 (signal handling) → Task 8 (timeout)
Task 10 (provenance) → Task 12 (config tiers)
Task 11 (exit codes) → Task 2 (failure tracking)
```

## Estimated Total Effort

- Priority 0 (critical): 10 hours
- Priority 1 (high-impact): 11 hours
- Priority 2 (medium-impact): 15 hours
- Priority 3 (low-impact): 18 hours
- Priority 4 (documentation): 1 hour

**Total: 55 hours** (approximately 7 working days)

## Recommendation

**Option A: Implement missing features in bash (short-term)**
- Pros: Incremental progress, validates features before Go rewrite
- Cons: Bash is fragile, platform-specific, hard to test
- Timeline: 7 working days for full feature parity

**Option B: Proceed with Go rewrite per existing PLAN.md (long-term)**
- Pros: Production-grade implementation, cross-platform, testable
- Cons: Longer timeline, higher upfront cost
- Timeline: 23 tasks organized in 6 priorities (per existing PLAN.md)

**Recommendation:** Option B (Go rewrite). The bash implementation has served its purpose as a proof-of-concept. Implementing 13 missing features in bash would take 55 hours and still leave a fragile, platform-specific implementation. The Go rewrite plan is already defined and addresses all gaps systematically.

**Next steps:**
1. Review existing PLAN.md (Go rewrite plan)
2. Prioritize Go rewrite tasks
3. Begin with Priority 0 (foundation) tasks from Go plan
4. Preserve bash implementation in archive/ for reference

## Notes

**Why this gap exists:**
- v2 specs were written for a Go rewrite with production-grade features
- v0.1.0 bash was a proof-of-concept to validate the OODA loop pattern
- Bash implementation successfully validated: fresh context per iteration, file-based state continuity, composable OODA prompts
- Gap is expected and intentional — bash was never meant to be the final implementation

**What bash implementation proved:**
- OODA loop pattern works for autonomous AI coding
- Fresh context prevents LLM degradation
- File-based state provides continuity without conversation history
- Composable prompts enable procedure reuse
- Git push automation provides backpressure and state persistence

**What bash implementation lacks:**
- Robustness (no timeouts, no signal handling, no failure tracking)
- Observability (no statistics, no dry-run, no provenance)
- Extensibility (no config tiers, no context injection)
- Portability (bash is platform-specific, requires yq dependency)

**Operational learnings from gap analysis:**
- Bash implementation has 4 undocumented features that should be preserved in Go rewrite
- Promise signals are critical for loop termination — must be in v1.0.0
- Failure tracking prevents wasted API calls — must be in v1.0.0
- Dry-run mode is essential for debugging — must be in v1.0.0
- Config tiers are complex in bash but natural in Go — wait for rewrite
