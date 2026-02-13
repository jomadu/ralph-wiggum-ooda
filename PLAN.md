# Plan: Fix Specification and Documentation Bugs

## Overview

Remove references to non-existent Homebrew distribution channel and fix incorrect procedure names in AGENTS.md. The Homebrew installation method was documented but never implemented, creating confusion for users and agents.

## Objectives

1. Remove all Homebrew references from user-facing documentation
2. Fix AGENTS.md to reference correct procedure names
3. Update distribution spec to reflect actual implementation status
4. Ensure documentation matches reality

## Tasks

### Task 1: Fix AGENTS.md Procedure References ✓ Published

**Issue ID:** ralph-wiggum-ooda-29um  
**Published:** 2026-02-13T06:58:27-08:00

**Description:** Remove references to non-existent procedures `draft-plan-story-to-spec` and `draft-plan-bug-to-spec` from AGENTS.md.

**Changes:**
- Update "Story/Bug Input" section
- Replace with correct procedure names: `draft-plan-spec-feat`, `draft-plan-spec-fix`, `draft-plan-impl-feat`, `draft-plan-impl-fix`
- Clarify that TASK.md is input for planning procedures

**Acceptance Criteria:**
- AGENTS.md references only procedures that exist in `rooda list` output
- "Story/Bug Input" section accurately describes planning workflow
- No references to `draft-plan-story-to-spec` or `draft-plan-bug-to-spec`

**Priority:** P0 (blocks agent operations)

**Dependencies:** None

### Task 2: Remove Homebrew from README.md ✓ Published

**Issue ID:** ralph-wiggum-ooda-6vsh  
**Published:** 2026-02-13T06:58:27-08:00

**Description:** Remove Homebrew installation method from README.md since it was never implemented.

**Changes:**
- Remove "Homebrew" reference from Installation section
- Keep reference to docs/installation.md for detailed instructions
- Ensure Quick Start uses curl-based installation

**Acceptance Criteria:**
- No Homebrew references in README.md
- Installation section points to docs/installation.md
- Quick Start installation command works

**Priority:** P0 (user-facing documentation error)

**Dependencies:** None

### Task 3: Remove Homebrew from docs/installation.md ✓ Published

**Issue ID:** ralph-wiggum-ooda-4e9c  
**Published:** 2026-02-13T06:58:27-08:00

**Description:** Remove Homebrew section from installation documentation.

**Changes:**
- Remove "Homebrew (macOS/Linux)" section entirely
- Keep "Quick install" (curl-based) as primary method
- Keep "Direct Download" and "Build from Source" sections
- Update section ordering if needed

**Acceptance Criteria:**
- No Homebrew section in docs/installation.md
- All remaining installation methods are verified working
- Documentation flows logically without Homebrew

**Priority:** P0 (user-facing documentation error)

**Dependencies:** None

### Task 4: Update distribution.md Spec ✓ Published

**Issue ID:** ralph-wiggum-ooda-5yml  
**Published:** 2026-02-13T06:58:27-08:00

**Description:** Update distribution spec to reflect that Homebrew is not implemented and add guidance for future distribution method changes.

**Changes:**
- Mark Homebrew acceptance criteria as "NOT IMPLEMENTED"
- Add "Implementation Status" section listing which distribution methods are actually working
- Add "Removing Distribution Methods" section with guidance
- Update examples to remove Homebrew references or mark as "planned"

**Acceptance Criteria:**
- Spec clearly indicates Homebrew is not implemented
- Implementation status section lists: curl install (working), direct download (working), go install (working), Homebrew (not implemented)
- Guidance exists for removing distribution methods
- No misleading claims about Homebrew functionality

**Priority:** P1 (prevents future confusion)

**Dependencies:** None

### Task 5: Verify Installation Script ✓ Published

**Issue ID:** ralph-wiggum-ooda-8yvk  
**Published:** 2026-02-13T06:58:27-08:00

**Description:** Verify that scripts/install.sh does not depend on Homebrew and works correctly.

**Changes:**
- Read scripts/install.sh
- Verify it uses direct download method
- Test execution (if possible)
- Document any issues found

**Acceptance Criteria:**
- scripts/install.sh does not reference Homebrew
- Script uses direct download from GitHub Releases
- Script includes platform detection and checksum verification

**Priority:** P1 (verification task)

**Dependencies:** None

## Implementation Notes

**Order of execution:**
1. Tasks 1, 2, 3, 5 can be done in parallel (no dependencies)
2. Task 4 should be done after understanding current state from Task 5

**Testing:**
- Run `rooda list` to verify procedure names
- Execute installation commands from updated docs
- Run `go test ./...` to ensure no code changes needed
- Verify all cross-references in documentation

**Quality gates:**
- All documentation examples must execute successfully
- No broken cross-references
- No AI writing patterns in updated docs
- `make test` and `make build` still pass

## Out of Scope

- Actually implementing Homebrew support (future enhancement)
- Creating CI/CD pipeline for releases
- Adding other distribution methods (Scoop, apt, etc.)
- Migration documentation for users (Homebrew never actually worked, so no users to migrate)
