# Gap Analysis Plan: v0.1.0 vs v2 Specifications

## Context

Gap analysis was run comparing the v2 Go rewrite specifications (11 complete specs in `specs/`) against the current v0.1.0 bash implementation (`rooda.sh`). This analysis correctly identified that the v2 Go rewrite has not been implemented.

**Critical Finding:** This is not a gap to fix—it's a planned v2 rewrite that is a separate project initiative. The specifications were written first (spec-driven development), and implementation has not yet started.

## Current State

**v0.1.0 bash implementation (working):**
- `rooda.sh` - Main OODA loop script (bash, v0.1.0)
- `rooda-config.yml` - Procedure definitions and AI tool presets
- `prompts/*.md` - 25 OODA prompt component files
- 9 procedures: bootstrap, build, draft-plan-story-to-spec, draft-plan-bug-to-spec, draft-plan-spec-to-impl, draft-plan-impl-to-spec, draft-plan-spec-refactor, draft-plan-impl-refactor, publish-plan
- Single-tier config system (workspace only)
- Manual verification (no automated tests)
- Requires yq >= 4.0.0 for YAML parsing

**v2 Go rewrite specifications (complete, not implemented):**
- 11 complete specs with JTBD structure, acceptance criteria, examples
- 16 procedures with fragment-based composition
- 55 prompt fragments organized by OODA phase
- Three-tier config system (workspace/global/built-in)
- Promise signals for iteration control
- Iteration statistics with Welford's algorithm
- Dry-run mode, provenance tracking, template system
- Single binary distribution with embedded resources
- No external dependencies (no yq required)

## Recommended Actions

### 1. Update AGENTS.md to clarify v2 rewrite status

**Priority:** P2 (documentation clarity)

**Description:** Add a section to AGENTS.md that explicitly documents:
- v2 specifications are complete (11 specs in `specs/`)
- v2 implementation has not started (no Go code exists)
- v0.1.0 bash is the current working implementation
- v2 is a planned rewrite, not a gap in v0.1.0
- The "gap" identified by this analysis is intentional (spec-driven development)

**Acceptance Criteria:**
- AGENTS.md has a "v2 Rewrite Status" section
- Section clarifies that v2 specs are complete but implementation hasn't started
- Section notes that v0.1.0 bash is the current working version
- Section explains that v2 is a separate planned initiative

**Rationale:** Prevents future confusion when gap analysis is run again. Makes it clear that the v2 rewrite is a known, planned project.

### 2. (Optional) Create v2 rewrite tracking

**Priority:** P3 (optional, depends on team decision)

**Description:** If the team wants to track v2 rewrite progress, create a separate plan or epic in the work tracking system. This would be a new initiative with its own tasks, not a gap-filling exercise.

**Acceptance Criteria:**
- Decision made on whether to track v2 rewrite
- If yes: separate plan/epic created with v2 implementation tasks
- If no: document that v2 rewrite is not currently prioritized

**Rationale:** Provides visibility into v2 rewrite progress if the team decides to pursue it. Keeps v2 work separate from v0.1.0 maintenance.

## Non-Actions

**Do NOT:**
- Create tasks to "implement missing features" from v2 specs into v0.1.0 bash
- Treat the v2 rewrite as a gap to be filled incrementally
- Attempt to backport v2 features (fragment composition, promise signals, etc.) into v0.1.0
- Create issues for "missing" v2 features—they're not missing, they're unimplemented by design

**Rationale:** v0.1.0 bash and v2 Go are fundamentally different architectures. Incremental migration is not feasible. v2 is a clean rewrite, not an evolution of v0.1.0.

## Learnings

**What worked:**
- Gap analysis correctly identified that v2 is not implemented
- Specifications are complete and internally consistent
- v0.1.0 bash implementation is working and documented

**What to improve:**
- AGENTS.md should explicitly document v2 rewrite status to prevent confusion
- Gap analysis procedure should handle "planned rewrite" scenarios differently than "missing features"

**Operational note:** When running gap analysis on a codebase with complete specs but no implementation, the result should acknowledge this is a planned rewrite, not a gap to fix.
