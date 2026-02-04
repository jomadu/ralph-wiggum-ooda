# Bug: VALIDATION-*.md Files Created in Project Root

## Issue Description

The `./rooda.sh build` procedure creates `VALIDATION-<issue-id>.md` files in the project root. These files accumulate and clutter the repository structure.

## Root Cause

**Missing specification for OODA components**: The `src/README.md` file describes component structure and common steps, but it's not formatted as a proper specification in `specs/`. This means:

1. Components lack clear guidance on what constitutes "operational learning" that belongs in AGENTS.md
2. The A6 common step ("Update AGENTS.md if learned something new") is too vague
3. Without a spec defining component behavior, `act_build.md` evolved to include validation file creation
4. No clear boundaries between "operational" (belongs in AGENTS.md) vs "historical" (doesn't belong)

The validation file pattern emerged because components don't have a specification that defines:
- What should be added to AGENTS.md (operational commands, file paths, quality criteria)
- What should NOT be added to AGENTS.md (test artifacts, validation patterns, historical notes)
- How to distinguish between the two

## Current State

- 18 VALIDATION-*.md files exist in project root
- Files document manual test cases for acceptance criteria validation
- No guidance in AGENTS.md about where to create these files or when to delete them
- `src/README.md` exists but isn't a proper spec following `specs/TEMPLATE.md`

## Expected Behavior

Components should have a specification that:
- Defines component structure and common steps formally
- Provides clear guidance on A6 ("Update AGENTS.md if learned something new")
- References `specs/agents-md-format.md` to clarify what belongs in AGENTS.md
- Prevents components from adding test artifacts or validation patterns to AGENTS.md

## Solution

1. Create `specs/component-system.md` following TEMPLATE.md format
2. Move content from `src/README.md` into proper spec structure
3. Add explicit guidance about AGENTS.md updates (what qualifies as "operational learning")
4. Update `act_build.md` to reference the component spec for AGENTS.md update guidance
5. Delete all 18 existing VALIDATION-*.md files from project root

## Acceptance Criteria

- [ ] `specs/component-system.md` created following TEMPLATE.md format
- [ ] Spec defines component structure, common steps, and principles
- [ ] Spec clarifies what should/shouldn't be added to AGENTS.md
- [ ] `src/README.md` updated to reference the spec (or removed if redundant)
- [ ] All 18 existing VALIDATION-*.md files deleted from project root
- [ ] Future build procedures don't create validation files in project root
