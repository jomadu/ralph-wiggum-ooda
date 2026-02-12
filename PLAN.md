# Plan: Distill Prompt Fragments to <100 Words

## Objective

Reduce all 56 prompt fragments in `internal/prompt/fragments/` to under 100 words each while preserving essential instructions and clarity.

## Current State

- Total fragments: 56
- Average words per fragment: 70 (already close to target)
- Fragments over 100 words: 3 critical ones need reduction

**Impact Analysis (usage × words = total impact):**
- Top 3 fragments account for 12,893 total words across all procedures (33% of total impact)
- Optimizing these 3 fragments yields maximum ROI

## Tasks

### P0: Reduce High-Impact Fragments (impact >2000)

**Task 1: Distill emit_signal.md (415→<100 words, impact: 6640)**
- Priority: P0
- Usage: 16 procedures (100% - used by ALL procedures)
- Current: 415 words with extensive examples
- Impact: Highest - reducing this saves 5,040+ words across all procedure prompts
- Strategy: Remove verbose examples, keep signal format and one example per type
- Acceptance: <100 words, preserves SUCCESS/FAILURE/continue distinction

**Task 2: Distill study_agents_md.md (232→<100 words, impact: 3712)**
- Priority: P0
- Usage: 16 procedures (100% - used by ALL procedures)
- Current: 232 words listing 10 required topics
- Impact: Second highest - reducing this saves 2,112+ words across all procedure prompts
- Strategy: Compress topic list to bullet format, remove parsing instructions
- Acceptance: <100 words, preserves all 10 required topics

**Task 3: Distill decide_signal.md (231→<100 words, impact: 2541)**
- Priority: P0
- Usage: 11 procedures (69% - used by most procedures)
- Current: 231 words with blocker categories and examples
- Impact: Third highest - reducing this saves 1,441+ words across procedure prompts
- Strategy: Condense blocker types, remove examples
- Acceptance: <100 words, preserves blocker/completion/continuation logic

### P1: Reduce Medium-Impact Fragments (impact 600-800)

**Task 4: Distill high-usage study fragments**
- Priority: P1
- Targets (all used 8-10 times):
  - `observe/study_specs.md` (78 words, 10 uses, impact: 780)
  - `observe/study_impl.md` (69 words, 10 uses, impact: 690)
  - `act/write_draft_plan.md` (86 words, 8 uses, impact: 688)
  - `observe/study_task_input.md` (79 words, 8 uses, impact: 632)
- Strategy: Tighten instructions, remove redundancy
- Acceptance: All <75 words, preserve core instructions

### P2: Reduce Low-Impact Fragments (impact <300)

**Task 5: Distill remaining fragments over 80 words**
- Priority: P2
- Targets (used 1-3 times):
  - `act/write_audit_report.md` (85 words, 3 uses, impact: 255)
  - `act/write_gap_report.md` (88 words, 2 uses, impact: 176)
  - `observe/study_task_details.md` (86 words, 1 use, impact: 86)
- Strategy: Minimal edits to stay under 80 words
- Acceptance: All <80 words, preserve core instructions

### P3: Verify All Fragments

**Task 6: Audit all 56 fragments for <100 word compliance**
- Priority: P3
- Dependencies: Tasks 1-5 complete
- Strategy: Run word count check, verify no regressions
- Acceptance: All fragments <100 words, average drops from 70 to ~55 words

### P4: Update Tests

**Task 7: Verify fragment composition tests still pass**
- Priority: P4
- Dependencies: Task 6 complete
- Strategy: Run `go test ./internal/prompt/...`
- Acceptance: All tests pass, no broken references

## Expected Impact

**Before optimization:**
- Total words across all procedures: ~15,680 words (16 procedures × ~980 words avg)
- Longest fragments: 415, 232, 231 words

**After P0 completion (Tasks 1-3):**
- Savings: ~8,593 words removed from procedure prompts
- Total words: ~7,087 words (55% reduction)
- All critical fragments <100 words

**After full completion (Tasks 1-7):**
- All 56 fragments <100 words
- Average fragment size: ~55 words (22% reduction from current 70)
- Total procedure prompt size: ~6,160 words (61% reduction)

## Implementation Notes

- Preserve imperative tone ("Your task is to...")
- Keep fragment structure (# Title, task description, bullets)
- Maintain AGENTS.md references (critical for runtime behavior)
- Don't remove essential validation criteria
- Test each fragment in context (compose full procedure prompt)

## Success Criteria

- [ ] All 56 fragments <100 words
- [ ] No loss of essential instructions
- [ ] All tests pass (`go test ./...`)
- [ ] Procedures still compose correctly
- [ ] Signal system still functions (SUCCESS/FAILURE detection)
