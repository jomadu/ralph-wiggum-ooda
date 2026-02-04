# Implementation Plan

## Task 1: Create user-documentation.md spec

**Priority:** High
**Effort:** Low
**Dependencies:** None

Define specification for user-facing documentation in `docs/`.

**Acceptance Criteria:**
- Spec follows TEMPLATE.md structure (Job to be Done, Activities, Acceptance Criteria, Examples)
- Defines what constitutes user-facing documentation vs specifications
- Specifies data flow: specs → implementation → docs
- Establishes acceptance criteria for documentation quality
- Defines structure: what goes in README vs docs/ vs specs/
- Includes examples of good vs bad documentation

**Files to create:**
- `specs/user-documentation.md`

## Task 2: Update specs/README.md to include user-documentation.md

**Priority:** Medium
**Effort:** Low
**Dependencies:** Task 1

Add new spec to the index.

**Acceptance Criteria:**
- specs/README.md lists user-documentation.md with extracted JTBD
- Follows existing README structure

**Files to modify:**
- `specs/README.md`

## Task 3: Add documentation quality criteria to AGENTS.md

**Priority:** Medium
**Effort:** Low
**Dependencies:** Task 1

Incorporate documentation quality criteria into Quality Criteria section.

**Acceptance Criteria:**
- Quality Criteria section includes boolean checks for documentation
- Examples: "All docs/ files have clear purpose (PASS/FAIL)", "All examples in docs are verified working (PASS/FAIL)"
- Criteria reference user-documentation.md spec

**Files to modify:**
- `AGENTS.md`

## Task 4: Verify existing docs/ files against new spec

**Priority:** Low
**Effort:** Medium
**Dependencies:** Task 1, Task 3

Audit current documentation for compliance.

**Acceptance Criteria:**
- Each doc in docs/ evaluated against acceptance criteria
- Gaps documented (missing purpose, broken examples, contradictions)
- Results inform whether refactoring plan needed

**Files to audit:**
- `docs/ooda-loop.md`
- `docs/ralph-loop.md`
- `docs/beads.md`
- `docs/README.md`
- `README.md` (root)
