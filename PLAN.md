# Specification Refactoring Plan

## Quality Assessment Results

**Criteria Scores:**
- Criterion 1 (Job to be Done sections): FAIL - agents-md-format.md missing section
- Criterion 2 (Acceptance Criteria sections): FAIL - agents-md-format.md missing section
- Criterion 3 (Examples sections): FAIL - agents-md-format.md missing section
- Criterion 4 (Command examples verified working): FAIL - No verification process defined, no empirical testing
- Criterion 5 (DEPRECATED specs have replacements): PASS

**Overall: 4 of 5 criteria failed - refactoring required**

## Refactoring Tasks (Priority Order)

### Task 1: Fix agents-md-format.md Structure
**Priority:** High (correctness)
**Effort:** Low
**Impact:** Brings spec into compliance with template

Restructure agents-md-format.md to follow JTBD template:
- Add "## Job to be Done" section (convert existing "## Purpose" content)
- Add "## Acceptance Criteria" section with boolean checkboxes
- Add "## Examples" section with concrete AGENTS.md examples
- Maintain existing content in appropriate sections
- Ensure consistency with other specs

**Acceptance Criteria:**
- [ ] agents-md-format.md has "Job to be Done" section
- [ ] agents-md-format.md has "Acceptance Criteria" section
- [ ] agents-md-format.md has "Examples" section
- [ ] All existing content preserved in appropriate sections
- [ ] Structure matches TEMPLATE.md pattern

### Task 2: Define Command Example Verification Process
**Priority:** High (enables criterion 4 evaluation)
**Effort:** Medium
**Impact:** Provides clear process for validating command examples

Create verification process specification:
- Define what constitutes "verified working"
- Distinguish executable commands from pseudocode/illustrative examples
- Specify verification methodology (manual execution, automated testing)
- Define tracking mechanism for verification status
- Document verification frequency (per-commit, per-release, on-demand)

**Acceptance Criteria:**
- [ ] Verification process documented in AGENTS.md or separate spec
- [ ] Clear distinction between executable vs illustrative examples
- [ ] Verification methodology specified
- [ ] Tracking mechanism defined
- [ ] Process is actionable and repeatable

### Task 3: Execute Initial Verification Pass
**Priority:** High (validates current state)
**Effort:** High
**Impact:** Identifies broken examples, validates working examples

Execute verification on all specs:
- Identify all bash code blocks in specs/*.md
- Classify each as executable command or pseudocode
- Execute all executable commands
- Document results (pass/fail/not-applicable)
- Fix broken examples or mark as illustrative
- Update specs with verification status

**Acceptance Criteria:**
- [ ] All bash code blocks classified
- [ ] All executable commands tested
- [ ] Broken examples fixed or reclassified
- [ ] Verification results documented
- [ ] Criterion 4 can be evaluated as PASS/FAIL

### Task 4: Mark Non-Executable Examples Clearly
**Priority:** Medium (clarity)
**Effort:** Low
**Impact:** Prevents confusion about what should work

Update spec examples:
- Add markers for pseudocode (e.g., "# Pseudocode" comment)
- Add markers for illustrative examples (e.g., "# Example only - not executable")
- Ensure executable examples have no markers
- Update TEMPLATE.md with marking conventions

**Acceptance Criteria:**
- [ ] All pseudocode marked clearly
- [ ] All illustrative examples marked clearly
- [ ] Executable examples unmarked
- [ ] TEMPLATE.md documents marking conventions

### Task 5: Automate Verification Where Possible
**Priority:** Low (efficiency)
**Effort:** High
**Impact:** Reduces manual verification burden

Create automation tooling:
- Script to extract bash code blocks from specs
- Script to classify blocks (executable vs illustrative)
- Script to execute commands in safe environment
- Script to report verification status
- Integrate into CI/CD if applicable

**Acceptance Criteria:**
- [ ] Extraction script working
- [ ] Classification logic implemented
- [ ] Execution script working with safety checks
- [ ] Reporting script generates clear output
- [ ] Documentation for running verification automation
