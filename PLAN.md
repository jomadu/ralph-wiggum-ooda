# Plan: Rename components/ to prompts/

## Goal
Eliminate naming discrepancy between framework (`src/components/`) and consumer projects (`prompts/`) by using `prompts/` consistently everywhere.

## Tasks

### 1. Rename directory
- Rename `src/components/` to `src/prompts/`
- Priority: High (blocks other tasks)
- Dependencies: None

### 2. Update src/rooda-config.yml (38 references)
- Replace all `src/components/*.md` paths with `src/prompts/*.md`
- Verify all procedure definitions reference correct paths
- Priority: High (blocks testing)
- Dependencies: Task 1

### 3. Update specs/component-authoring.md (26 references)
- Replace `src/components/` with `src/prompts/`
- Update terminology from "components" to "prompts" where referring to the files
- Keep "component" when referring to OODA phase components conceptually
- Priority: Medium
- Dependencies: Task 1

### 4. Update specs/cli-interface.md (16 references)
- Replace example paths from `src/components/` to `src/prompts/`
- Priority: Medium
- Dependencies: Task 1

### 5. Update src/README.md (12 references)
- Title: "Procedures and Components Specification" → "Procedures and Prompts Specification"
- Replace `src/components/` references with `src/prompts/`
- Update section headers if needed
- Priority: Medium
- Dependencies: Task 1

### 6. Update specs/configuration-schema.md (5 references)
- Replace example paths from `src/components/` to `src/prompts/`
- Priority: Medium
- Dependencies: Task 1

### 7. Update README.md (5 references)
- Change installation: `cp -r ralph-wiggum-ooda/src/components ./prompts` → `cp -r ralph-wiggum-ooda/src/prompts ./prompts`
- Update explanation text about component composition
- Update sample repository structure diagram
- Priority: High (user-facing)
- Dependencies: Task 1

### 8. Update specs/agents-md-format.md (4 references)
- Remove "Path Conventions" section entirely
- Update any remaining references to `src/components/`
- Simplify documentation now that naming is consistent
- Priority: Medium
- Dependencies: Task 1

### 9. Update AGENTS.md (2 references)
- Implementation Definition: `src/components/*.md` → `src/prompts/*.md`
- Priority: Medium
- Dependencies: Task 1

### 10. Verify functionality
- Run `shellcheck src/rooda.sh` to verify no errors
- Run `./rooda.sh bootstrap --max-iterations 1` to test
- Confirm script loads prompts correctly
- Check for any errors referencing old paths
- Priority: High (validation)
- Dependencies: Tasks 1-9

### 11. Final grep verification
- Run `grep -ri "components" .` excluding .git and .beads
- Verify only conceptual uses remain (not path references)
- Update any missed references
- Priority: Low (cleanup)
- Dependencies: Task 10

## Success Criteria
- [ ] Directory renamed from `src/components/` to `src/prompts/`
- [ ] All 128 references updated across 14 files
- [ ] shellcheck passes
- [ ] Bootstrap procedure runs successfully
- [ ] No path references to old `components/` name remain
- [ ] Installation instructions are simpler and clearer
- [ ] "Path Conventions" section removed from agents-md-format.md

## Notes
- .beads/issues.jsonl and git logs contain historical references - leave unchanged
- Keep "component" terminology when referring to OODA phase composition conceptually
- Only change "components" to "prompts" when referring to the directory or files
