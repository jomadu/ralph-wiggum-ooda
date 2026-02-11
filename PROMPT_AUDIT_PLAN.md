# Prompt Fragment Audit Plan

**Branch:** `audit-prompts`  
**Goal:** Review and improve all prompt fragments in `internal/prompt/fragments/`  
**Date Started:** 2026-02-11

## Audit Status Legend

- [ ] Not Started
- [~] In Progress
- [x] Complete

## Fragment Inventory (Sorted by Usage)

### High-Impact Fragments (10+ uses)

- [x] `act/emit_signal.md` - **16 procedures** (all) - REPLACED emit_signal.md + emit_failure.md
- [x] `observe/study_agents_md.md` - **16 procedures** (all) - RENAMED from read_agents_md.md
- [x] `observe/study_specs.md` - **10 procedures** - RENAMED from read_specs.md
- [x] `observe/study_task_input.md` - **10 procedures** - RENAMED from read_task_input.md
- [x] `decide/decide_signal.md` - **10 procedures** - RENAMED from check_if_blocked.md, REFACTORED to decide SUCCESS/FAILURE/continue
- [x] `decide/break_down_into_tasks.md` - **10 procedures**
- [x] `decide/prioritize_tasks.md` - **10 procedures**
- [x] `act/write_draft_plan.md` - **10 procedures**

### Medium-Impact Fragments (3-9 uses)

- [x] `observe/study_impl.md` - **9 procedures** - RENAMED from read_impl.md
- [ ] `act/write_audit_report.md` - **3 procedures**

### Low-Impact Fragments (2 uses)

- [ ] `observe/scan_repo_structure.md` - **2 procedures**
- [ ] `observe/detect_build_system.md` - **2 procedures**
- [ ] `observe/query_work_tracking.md` - **2 procedures**
- [ ] `orient/identify_drift.md` - **2 procedures**
- [ ] `orient/evaluate_against_quality_criteria.md` - **2 procedures**
- [ ] `orient/understand_feature_requirements.md` - **2 procedures**
- [ ] `orient/understand_bug_root_cause.md` - **2 procedures**
- [ ] `orient/identify_maintenance_needs.md` - **2 procedures**
- [ ] `orient/identify_affected_code.md` - **2 procedures**
- [ ] `decide/identify_issues.md` - **2 procedures**
- [ ] `decide/prioritize_findings.md` - **2 procedures**
- [ ] `decide/prioritize_gaps_by_impact.md` - **2 procedures**
- [ ] `act/commit_changes.md` - **2 procedures**
- [ ] `act/write_gap_report.md` - **2 procedures**

### Single-Use Fragments (1 use)

- [ ] `observe/detect_work_tracking.md` - **1 procedure**
- [x] `observe/study_task_details.md` - **1 procedure** - RENAMED from read_task_details.md
- [x] `observe/study_draft_plan.md` - **1 procedure** - RENAMED from read_draft_plan.md
- [ ] `observe/run_tests.md` - **1 procedure**
- [ ] `observe/run_lints.md` - **1 procedure**
- [ ] `observe/verify_commands.md` - **1 procedure**
- [ ] `orient/compare_detected_vs_documented.md` - **1 procedure**
- [ ] `orient/understand_task_requirements.md` - **1 procedure**
- [ ] `orient/search_codebase.md` - **1 procedure**
- [ ] `orient/identify_affected_files.md` - **1 procedure**
- [ ] `orient/parse_plan_tasks.md` - **1 procedure**
- [ ] `orient/map_to_work_tracking_format.md` - **1 procedure**
- [ ] `orient/compare_documented_vs_actual.md` - **1 procedure**
- [ ] `orient/identify_specified_but_not_implemented.md` - **1 procedure**
- [ ] `orient/identify_implemented_but_not_specified.md` - **1 procedure**
- [ ] `orient/identify_affected_specs.md` - **1 procedure**
- [ ] `orient/identify_spec_deficiencies.md` - **1 procedure**
- [ ] `orient/identify_structural_issues.md` - **1 procedure**
- [ ] `orient/identify_duplication.md` - **1 procedure**
- [ ] `orient/identify_code_smells.md` - **1 procedure**
- [ ] `orient/identify_complexity_issues.md` - **1 procedure**
- [ ] `decide/determine_sections_to_update.md` - **1 procedure**
- [ ] `decide/pick_task.md` - **1 procedure**
- [ ] `decide/plan_implementation_approach.md` - **1 procedure**
- [ ] `decide/determine_import_strategy.md` - **1 procedure**
- [ ] `decide/categorize_drift_severity.md` - **1 procedure**
- [ ] `act/write_agents_md.md` - **1 procedure**
- [ ] `act/modify_files.md` - **1 procedure**
- [ ] `act/run_tests.md` - **1 procedure**
- [ ] `act/update_work_tracking.md` - **1 procedure**
- [ ] `act/create_work_items.md` - **1 procedure**
- [ ] `act/update_draft_plan_status.md` - **1 procedure**

### Unused Fragments (0 uses)

**None** - All fragments now in use after emit_signal unification

**Total Fragments:** 55 files (54 unique) - reduced from 57 after consolidation

**Fragments Completed:** 11 of 55 (20%)
- 2 replaced (emit_success, emit_failure → emit_signal)
- 7 renamed and enhanced (read_* → study_*, check_if_blocked → decide_signal)
- 2 improved (emit_signal decision logic, decide_signal comprehensive rewrite)

---

## Procedure Breakdown

*Note: Extract complete composed prompts with `./bin/rooda run <procedure> --dry-run > prompts-baseline/<procedure>.txt`*

### 1. agents-sync
**Purpose:** Synchronize AGENTS.md with actual repository state

**Fragments:**
- Observe: `read_agents_md`, `scan_repo_structure`, `detect_build_system`, `detect_work_tracking`
- Orient: `compare_detected_vs_documented`, `identify_drift`
- Decide: `determine_sections_to_update`, `check_if_blocked`
- Act: `write_agents_md`, `commit_changes`, `emit_success`

**Composed Prompt:** `prompts-baseline/agents-sync.txt`  
**Audit Status:** [ ]

---

### 2. build
**Purpose:** Implement a task from work tracking

**Fragments:**
- Observe: `read_agents_md`, `query_work_tracking`, `read_specs`, `read_impl`, `read_task_details`
- Orient: `understand_task_requirements`, `search_codebase`, `identify_affected_files`
- Decide: `pick_task`, `plan_implementation_approach`, `check_if_blocked`
- Act: `modify_files`, `run_tests`, `update_work_tracking`, `commit_changes`, `emit_success`

**Composed Prompt:** `prompts-baseline/build.txt`  
**Audit Status:** [ ]

---

### 3. publish-plan
**Purpose:** Import draft plan into work tracking system

**Fragments:**
- Observe: `read_agents_md`, `read_draft_plan`, `query_work_tracking`
- Orient: `parse_plan_tasks`, `map_to_work_tracking_format`
- Decide: `determine_import_strategy`, `check_if_blocked`
- Act: `create_work_items`, `update_draft_plan_status`, `emit_success`

**Composed Prompt:** `prompts-baseline/publish-plan.txt`  
**Audit Status:** [ ]

---

### 4. audit-spec
**Purpose:** Audit specification files for quality issues

**Fragments:**
- Observe: `read_agents_md`, `read_specs`
- Orient: `evaluate_against_quality_criteria`
- Decide: `identify_issues`, `prioritize_findings`
- Act: `write_audit_report`, `emit_success`

**Composed Prompt:** `prompts-baseline/audit-spec.txt`  
**Audit Status:** [ ]

---

### 5. audit-impl
**Purpose:** Audit implementation files for quality issues

**Fragments:**
- Observe: `read_agents_md`, `read_impl`, `run_tests`, `run_lints`
- Orient: `evaluate_against_quality_criteria`
- Decide: `identify_issues`, `prioritize_findings`
- Act: `write_audit_report`, `emit_success`

**Composed Prompt:** `prompts-baseline/audit-impl.txt`  
**Audit Status:** [ ]

---

### 6. audit-agents
**Purpose:** Audit AGENTS.md for accuracy and completeness

**Fragments:**
- Observe: `read_agents_md`, `scan_repo_structure`, `detect_build_system`, `verify_commands`
- Orient: `compare_documented_vs_actual`, `identify_drift`
- Decide: `categorize_drift_severity`
- Act: `write_audit_report`, `emit_success`

**Composed Prompt:** `prompts-baseline/audit-agents.txt`  
**Audit Status:** [ ]

---

### 7. audit-spec-to-impl
**Purpose:** Find specifications not implemented in code

**Fragments:**
- Observe: `read_agents_md`, `read_specs`, `read_impl`
- Orient: `identify_specified_but_not_implemented`
- Decide: `prioritize_gaps_by_impact`
- Act: `write_gap_report`, `emit_success`

**Composed Prompt:** `prompts-baseline/audit-spec-to-impl.txt`  
**Audit Status:** [ ]

---

### 8. audit-impl-to-spec
**Purpose:** Find implementation not covered by specifications

**Fragments:**
- Observe: `read_agents_md`, `read_impl`, `read_specs`
- Orient: `identify_implemented_but_not_specified`
- Decide: `prioritize_gaps_by_impact`
- Act: `write_gap_report`, `emit_success`

**Composed Prompt:** `prompts-baseline/audit-impl-to-spec.txt`  
**Audit Status:** [ ]

---

### 9. draft-plan-spec-feat
**Purpose:** Create plan for new specification feature

**Fragments:**
- Observe: `read_agents_md`, `read_task_input`, `read_specs`, `read_impl`
- Orient: `understand_feature_requirements`, `identify_affected_specs`
- Decide: `break_down_into_tasks`, `prioritize_tasks`, `check_if_blocked`
- Act: `write_draft_plan`, `emit_success`

**Composed Prompt:** `prompts-baseline/draft-plan-spec-feat.txt`  
**Audit Status:** [ ]

---

### 10. draft-plan-spec-fix
**Purpose:** Create plan for specification bug fix

**Fragments:**
- Observe: `read_agents_md`, `read_task_input`, `read_specs`, `read_impl`
- Orient: `understand_bug_root_cause`, `identify_spec_deficiencies`
- Decide: `break_down_into_tasks`, `prioritize_tasks`, `check_if_blocked`
- Act: `write_draft_plan`, `emit_success`

**Composed Prompt:** `prompts-baseline/draft-plan-spec-fix.txt`  
**Audit Status:** [ ]

---

### 11. draft-plan-spec-refactor
**Purpose:** Create plan for specification refactoring

**Fragments:**
- Observe: `read_agents_md`, `read_task_input`, `read_specs`
- Orient: `identify_structural_issues`, `identify_duplication`
- Decide: `break_down_into_tasks`, `prioritize_tasks`, `check_if_blocked`
- Act: `write_draft_plan`, `emit_success`

**Composed Prompt:** `prompts-baseline/draft-plan-spec-refactor.txt`  
**Audit Status:** [ ]

---

### 12. draft-plan-spec-chore
**Purpose:** Create plan for specification maintenance tasks

**Fragments:**
- Observe: `read_agents_md`, `read_task_input`, `read_specs`
- Orient: `identify_maintenance_needs`
- Decide: `break_down_into_tasks`, `prioritize_tasks`, `check_if_blocked`
- Act: `write_draft_plan`, `emit_success`

**Composed Prompt:** `prompts-baseline/draft-plan-spec-chore.txt`  
**Audit Status:** [ ]

---

### 13. draft-plan-impl-feat
**Purpose:** Create plan for new implementation feature

**Fragments:**
- Observe: `read_agents_md`, `read_task_input`, `read_specs`, `read_impl`
- Orient: `understand_feature_requirements`, `identify_affected_code`
- Decide: `break_down_into_tasks`, `prioritize_tasks`, `check_if_blocked`
- Act: `write_draft_plan`, `emit_success`

**Composed Prompt:** `prompts-baseline/draft-plan-impl-feat.txt`  
**Audit Status:** [ ]

---

### 14. draft-plan-impl-fix
**Purpose:** Create plan for implementation bug fix

**Fragments:**
- Observe: `read_agents_md`, `read_task_input`, `read_specs`, `read_impl`
- Orient: `understand_bug_root_cause`, `identify_affected_code`
- Decide: `break_down_into_tasks`, `prioritize_tasks`, `check_if_blocked`
- Act: `write_draft_plan`, `emit_success`

**Composed Prompt:** `prompts-baseline/draft-plan-impl-fix.txt`  
**Audit Status:** [ ]

---

### 15. draft-plan-impl-refactor
**Purpose:** Create plan for code refactoring

**Fragments:**
- Observe: `read_agents_md`, `read_task_input`, `read_impl`
- Orient: `identify_code_smells`, `identify_complexity_issues`
- Decide: `break_down_into_tasks`, `prioritize_tasks`, `check_if_blocked`
- Act: `write_draft_plan`, `emit_success`

**Composed Prompt:** `prompts-baseline/draft-plan-impl-refactor.txt`  
**Audit Status:** [ ]

---

### 16. draft-plan-impl-chore
**Purpose:** Create plan for code maintenance tasks

**Fragments:**
- Observe: `read_agents_md`, `read_task_input`, `read_impl`
- Orient: `identify_maintenance_needs`
- Decide: `break_down_into_tasks`, `prioritize_tasks`, `check_if_blocked`
- Act: `write_draft_plan`, `emit_success`

**Composed Prompt:** `prompts-baseline/draft-plan-impl-chore.txt`  
**Audit Status:** [ ]

---

## Audit Criteria

For each fragment, evaluate:

1. **Clarity** - Is the instruction clear and unambiguous?
2. **Completeness** - Does it cover all necessary aspects?
3. **Consistency** - Does it align with other fragments and OODA principles?
4. **Actionability** - Can an AI agent execute this effectively?
5. **Specificity** - Is it specific enough without being overly prescriptive?
6. **Error Handling** - Does it address failure cases appropriately?

## Audit Process

1. Review fragment content
2. Check usage across procedures
3. Identify improvement opportunities
4. Document findings
5. Implement changes
6. Update status checkboxes

## Prompt Extraction Procedure

To evaluate prompt quality over time, extract the complete composed prompts for each procedure using dry-run mode.

### Extract Single Procedure

```bash
./bin/rooda run <procedure-name> --dry-run > prompts-baseline/<procedure-name>.txt
```

### Extract All Procedures (Baseline)

```bash
# Create baseline directory
mkdir -p prompts-baseline

# Extract all 16 procedures
./bin/rooda run agents-sync --dry-run > prompts-baseline/agents-sync.txt
./bin/rooda run build --dry-run > prompts-baseline/build.txt
./bin/rooda run publish-plan --dry-run > prompts-baseline/publish-plan.txt
./bin/rooda run audit-spec --dry-run > prompts-baseline/audit-spec.txt
./bin/rooda run audit-impl --dry-run > prompts-baseline/audit-impl.txt
./bin/rooda run audit-agents --dry-run > prompts-baseline/audit-agents.txt
./bin/rooda run audit-spec-to-impl --dry-run > prompts-baseline/audit-spec-to-impl.txt
./bin/rooda run audit-impl-to-spec --dry-run > prompts-baseline/audit-impl-to-spec.txt
./bin/rooda run draft-plan-spec-feat --dry-run > prompts-baseline/draft-plan-spec-feat.txt
./bin/rooda run draft-plan-spec-fix --dry-run > prompts-baseline/draft-plan-spec-fix.txt
./bin/rooda run draft-plan-spec-refactor --dry-run > prompts-baseline/draft-plan-spec-refactor.txt
./bin/rooda run draft-plan-spec-chore --dry-run > prompts-baseline/draft-plan-spec-chore.txt
./bin/rooda run draft-plan-impl-feat --dry-run > prompts-baseline/draft-plan-impl-feat.txt
./bin/rooda run draft-plan-impl-fix --dry-run > prompts-baseline/draft-plan-impl-fix.txt
./bin/rooda run draft-plan-impl-refactor --dry-run > prompts-baseline/draft-plan-impl-refactor.txt
./bin/rooda run draft-plan-impl-chore --dry-run > prompts-baseline/draft-plan-impl-chore.txt
```

### Compare Before/After Changes

After making fragment improvements:

```bash
# Extract updated prompts
mkdir -p prompts-updated
./bin/rooda run <procedure-name> --dry-run > prompts-updated/<procedure-name>.txt

# Compare changes
diff prompts-baseline/<procedure-name>.txt prompts-updated/<procedure-name>.txt
```

### Workflow

1. **Baseline extraction** - Run before making any changes to establish baseline
2. **Make fragment improvements** - Edit fragments in `internal/prompt/fragments/`
3. **Rebuild binary** - `make build` to embed updated fragments
4. **Extract updated prompts** - Run dry-run again to new directory
5. **Compare** - Use diff to see exactly what changed in composed prompts
6. **Document** - Note improvements in audit status for affected fragments

## Notes

- `emit_failure.md` exists but is not used by any procedure - investigate if needed
- `emit_signal.md` is used by all 16 procedures - critical fragment
- `check_if_blocked.md` is used by 10 procedures - high-impact fragment
- `read_agents_md.md` is used by all 16 procedures - foundational fragment

---

## Change Log

### 2026-02-11 - Read AGENTS.md Improvements

**Change:** Enhanced `observe/read_agents_md.md` with comprehensive parsing and validation

**Improvements:**
- Listed all 10 required topics explicitly (was 5 vague items)
- Added flexible parsing approach (any header level, bold text, plain text)
- Added validation rules (check all topics covered, commands extractable, paths present)
- Added error handling (missing file triggers FAILURE signal)
- Clarified storage instruction ("Remember extracted information" vs "Store parsed data")
- Aligned with `specs/agents-md-format.md` specification

**Impact:**
- All 16 procedures use this fragment
- Better guidance for agents parsing AGENTS.md files with varying formats
- Clear validation prevents silent failures from missing information

**Files Modified:**
- Updated: `internal/prompt/fragments/observe/read_agents_md.md`

**Status:** ✅ Complete

---

### 2026-02-11 - Unified Signal Emission

**Change:** Replaced `emit_signal.md` and `emit_failure.md` with unified `emit_signal.md`

**Rationale:** 
- Original design had implicit third state ("keep iterating") with no clear guidance
- `emit_failure.md` was never used (0 references)
- Unified fragment makes state machine explicit: SUCCESS stops, FAILURE stops, no signal continues
- Removed CONTINUE signal as unnecessary - default behavior is to continue iterating
- Simplified to two-signal design: only emit when you need to stop

**Impact:**
- All 16 procedures updated to use `emit_signal.md`
- Tests updated and passing
- Binary rebuilt with embedded changes

**Files Modified:**
- Created: `internal/prompt/fragments/act/emit_signal.md`
- Deleted: `internal/prompt/fragments/act/emit_success.md`
- Deleted: `internal/prompt/fragments/act/emit_failure.md`
- Deleted: `internal/prompt/fragments_signal_test.go` (obsolete test file)
- Updated: `internal/procedures/builtin.go` (16 procedures)
- Updated: `internal/procedures/builtin_test.go` (test assertions)
- Updated: `internal/prompt/fragments_test.go` (signal validation)

**Status:** ✅ Complete - All tests passing

**Fragments Audited:**
- [x] `act/emit_signal.md` - **REPLACED** with `emit_signal.md`
- [x] `act/emit_failure.md` - **REPLACED** with `emit_signal.md`
- [x] `observe/read_agents_md.md` - **IMPROVED** - Added all 10 required topics, flexible parsing, validation, error handling
- [x] `observe/read_specs.md` - **IMPROVED** - Changed "load" to "study", added JTBD structure guidance, examples, implementation status
- [x] `observe/read_*` → `observe/study_*` - **REFACTORED** - All 6 read fragments renamed to study, enhanced with analytical guidance
- [x] `decide/check_if_blocked.md` → `decide/decide_signal.md` - **REFACTORED** - Now decides SUCCESS/FAILURE/continue instead of just checking blockers; aligned with OODA principle (Decide makes decisions, Act executes them)
- [x] `act/emit_signal.md` - **IMPROVED** - Removed decision-making logic, now purely executes decisions made in Decide phase

---

### 2026-02-11 - Refactor read_* to study_* Fragments

**Change:** Renamed all `observe/read_*.md` fragments to `observe/study_*.md` and enhanced with analytical guidance

**Rationale:**
- "Load" and "read" are passive, mechanical terms
- "Study" better captures the analytical work agents must perform
- Agents need to understand intent and context, not just load text
- Consistent with OODA observe phase goal: gather and comprehend information

**Improvements:**
- **study_agents_md.md** - Added "Use the file reading tool" for tool usage clarity
- **study_specs.md** - Added JTBD, acceptance criteria, examples, implementation status extraction
- **study_impl.md** - Added design patterns, architectural decisions, completeness analysis
- **study_task_input.md** - Added scope boundaries, intent understanding; removed hardcoded TASK.md path; made implementation-agnostic (supports files, APIs, commands)
- **study_task_details.md** - Added context understanding, historical decisions
- **study_draft_plan.md** - Added rationale for task breakdown and ordering; removed hardcoded PLAN.md path; made implementation-agnostic (supports files, APIs, commands)

**Impact:**
- All 16 procedures updated
- 6 fragments renamed and enhanced
- Tests updated and passing
- Binary rebuilt with embedded changes

**Files Modified:**
- Renamed: `observe/read_agents_md.md` → `observe/study_agents_md.md`
- Renamed: `observe/read_specs.md` → `observe/study_specs.md`
- Renamed: `observe/read_impl.md` → `observe/study_impl.md`
- Renamed: `observe/read_task_input.md` → `observe/study_task_input.md`
- Renamed: `observe/read_task_details.md` → `observe/study_task_details.md`
- Renamed: `observe/read_draft_plan.md` → `observe/study_draft_plan.md`
- Updated: `internal/procedures/builtin.go` (all 16 procedures)
- Updated: `internal/procedures/builtin_test.go`
- Updated: `internal/prompt/composer_test.go`
- Updated: `internal/prompt/fragments_imperative_test.go`

**Status:** ✅ Complete - All tests passing

---

### 2026-02-11 - Refactored check_if_blocked to decide_signal

**Change:** Renamed `decide/check_if_blocked.md` to `decide/decide_signal.md` and refactored to make signal decisions

**Rationale:**
- Original fragment only checked for blockers, didn't decide what to do about them
- Fragment is last Decide step before emit_signal in Act phase
- OODA principle: Decide phase makes decisions, Act phase executes them
- Fragment should decide SUCCESS/FAILURE/continue, not just identify blockers

**Improvements:**
- **Renamed** from `check_if_blocked.md` to `decide_signal.md` (clearer purpose)
- **Added SUCCESS criteria** - When to stop because goal achieved
- **Added continuation criteria** - When to keep iterating
- **Removed signal emission** - That's emit_signal's job in Act phase
- **Added concrete examples** - All three decision types with context
- **Specific blocker categories** - Context-aware (AGENTS.md, work tracking, permissions)
- **Aligned with emit_signal** - Decide makes decision, Act executes it

**Impact:**
- All 10 procedures using this fragment updated
- emit_signal.md also updated to remove decision-making logic
- Tests updated and passing
- Binary rebuilt with embedded changes

**Files Modified:**
- Renamed: `decide/check_if_blocked.md` → `decide/decide_signal.md`
- Updated: `internal/procedures/builtin.go` (10 procedures)
- Updated: `act/emit_signal.md` (removed decision logic, now purely executes)

**Status:** ✅ Complete - All tests passing

---
