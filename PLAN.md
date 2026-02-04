# Specification Refactoring Plan

## Quality Assessment Results

**Criterion 1: All specs have "Job to be Done" section** - **FAIL**
- agents-md-format.md uses "## Purpose" instead of "## Job to be Done"
- All other 8 specs (excluding TEMPLATE.md, README.md, specification-system.md) have "## Job to be Done"

**Criterion 2: All specs have "Acceptance Criteria" section** - **FAIL**
- agents-md-format.md missing "## Acceptance Criteria" section
- All other 8 specs have "## Acceptance Criteria"

**Criterion 3: All specs have "Examples" section** - **FAIL**
- agents-md-format.md missing "## Examples" section
- All other 8 specs have "## Examples"

**Criterion 4: All command examples in specs are verified working** - **FAIL**
- 44 bash code blocks identified across 8 specs
- No verification process defined
- No empirical testing performed
- No distinction between executable commands and pseudocode/illustrative examples

**Criterion 5: No specs marked as DEPRECATED without replacement** - **PASS**
- component-system.md marked DEPRECATED, replacement is component-authoring.md
- prompt-composition.md marked DEPRECATED, replacement is component-authoring.md
- Both deprecated specs correctly reference replacement

## Refactoring Tasks (Priority Order)

### Task 1: Fix agents-md-format.md Structure (HIGH PRIORITY)
**Impact:** Criteria 1, 2, 3 all fail due to this single spec
**Effort:** Low
**Risk:** Low - structural change only

**Actions:**
- Rename "## Purpose" section to "## Job to be Done"
- Add "## Acceptance Criteria" section with checkboxes for required AGENTS.md sections
- Add "## Examples" section with sample AGENTS.md files for different project types

**Acceptance Criteria:**
- agents-md-format.md has "## Job to be Done" section
- agents-md-format.md has "## Acceptance Criteria" section
- agents-md-format.md has "## Examples" section
- Content from "## Purpose" preserved in "## Job to be Done"

### Task 2: Define Command Example Verification Process (HIGH PRIORITY)
**Impact:** Criterion 4 fails - no verification process exists
**Effort:** Medium
**Risk:** Low - documentation only

**Actions:**
- Create verification process definition in AGENTS.md or separate doc
- Distinguish executable commands (./rooda.sh, yq, bd, kiro-cli) from pseudocode
- Define verification steps: execute command, validate output, document result
- Specify which commands must work vs which are illustrative

**Acceptance Criteria:**
- Verification process documented
- Clear distinction between executable and pseudocode examples
- Verification steps defined (execute, validate, document)
- Criteria for "verified working" specified

### Task 3: Execute Verification Pass on All Specs (MEDIUM PRIORITY)
**Impact:** Criterion 4 fails - examples not empirically tested
**Effort:** High - 44 bash code blocks to verify
**Risk:** Medium - may discover broken examples

**Actions:**
- Systematically execute all bash code blocks marked as executable
- Validate output matches expected behavior
- Document verification results per spec
- Fix broken examples or mark as pseudocode

**Acceptance Criteria:**
- All executable bash code blocks tested
- Verification results documented
- Broken examples fixed or marked as pseudocode
- Pseudocode examples clearly labeled

### Task 4: Mark Non-Executable Examples Clearly (MEDIUM PRIORITY)
**Impact:** Criterion 4 - prevents confusion between executable and illustrative
**Effort:** Low
**Risk:** Low - documentation clarification

**Actions:**
- Add comments to pseudocode blocks: "# Pseudocode - not executable"
- Add comments to illustrative examples: "# Example - adjust for your project"
- Ensure executable examples have no such markers

**Acceptance Criteria:**
- All pseudocode blocks marked clearly
- All illustrative examples marked clearly
- Executable examples unmarked (default assumption)

### Task 5: Automate Verification Where Possible (LOW PRIORITY)
**Impact:** Criterion 4 - enables continuous verification
**Effort:** High - requires tooling
**Risk:** Medium - automation complexity

**Actions:**
- Create script to extract bash code blocks from specs
- Identify which blocks are executable vs pseudocode
- Run executable blocks and capture results
- Generate verification report

**Acceptance Criteria:**
- Verification script exists
- Script distinguishes executable from pseudocode
- Script executes commands safely (dry-run mode)
- Verification report generated

## Critical Path

1. Task 1 (fix agents-md-format.md) → Achieves criteria 1, 2, 3 compliance
2. Task 2 (define verification process) → Prerequisite for Task 3
3. Task 3 (execute verification) → Achieves criterion 4 compliance
4. Task 4 (mark non-executable) → Enhances criterion 4 compliance
5. Task 5 (automate) → Maintains criterion 4 compliance over time

## Summary

- **Criteria failing:** 4 of 5 (criteria 1, 2, 3, 4)
- **Criteria passing:** 1 of 5 (criterion 5)
- **Root cause:** agents-md-format.md doesn't follow template (affects 3 criteria), no verification process (affects 1 criterion)
- **Recommended first action:** Task 1 (fix agents-md-format.md structure) - highest impact, lowest effort
