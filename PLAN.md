# Specification Refactoring Plan

## Quality Assessment Results

### Criterion 1: All specs have "Job to be Done" section
**Status:** PASS
- external-dependencies.md: ✓
- cli-interface.md: ✓
- iteration-loop.md: ✓
- component-system.md: ✓ (deprecated)
- configuration-schema.md: ✓
- ai-cli-integration.md: ✓
- agents-md-format.md: N/A (format spec, not JTBD-based)
- component-authoring.md: ✓
- prompt-composition.md: ✓ (deprecated)

### Criterion 2: All specs have "Acceptance Criteria" section
**Status:** PASS
- All specs contain "Acceptance Criteria" sections with checkboxes

### Criterion 3: All specs have "Examples" section
**Status:** PASS
- All specs contain "Examples" sections with numbered examples

### Criterion 4: All command examples in specs are verified working
**Status:** FAIL
- Per AGENTS.md operational learning (2026-02-03), this criterion requires verification process definition
- Command examples exist but have not been empirically tested
- No verification process currently defined or executed
- Examples include: `./rooda.sh bootstrap`, `bd ready --json`, `yq eval`, `kiro-cli chat`, etc.

### Criterion 5: No specs marked as DEPRECATED without replacement
**Status:** PASS
- component-system.md: DEPRECATED with replacement (component-authoring.md)
- prompt-composition.md: DEPRECATED with replacement (component-authoring.md)
- Both deprecated specs explicitly reference their replacement

### Human Markers Found
- No TODO, FIXME, or HACK markers found in spec files (only in R19 step definition text)

## Overall Assessment
**Refactoring Required:** YES (Criterion 4 fails)

## Refactoring Tasks

### 1. Define Command Verification Process
**Priority:** Critical
**Impact:** High - ensures specs remain accurate as implementation evolves
**Effort:** Medium

Create verification process that:
- Identifies executable commands vs pseudocode/illustrative examples
- Distinguishes commands that must work (./rooda.sh, yq, bd, kiro-cli) from examples
- Defines how to mark non-executable examples clearly
- Establishes verification workflow (manual or automated)
- Documents verification results

**Acceptance Criteria:**
- Verification process documented in AGENTS.md or separate verification spec
- Clear distinction between executable and illustrative examples
- Process can be executed to validate all command examples

### 2. Execute Command Verification on All Specs
**Priority:** Critical
**Impact:** High - validates current spec accuracy
**Effort:** High

Systematically verify all command examples in:
- external-dependencies.md (yq, kiro-cli, bd, shellcheck, git commands)
- cli-interface.md (./rooda.sh invocations with various flags)
- iteration-loop.md (loop control examples)
- configuration-schema.md (yq queries)
- ai-cli-integration.md (kiro-cli invocation)
- component-authoring.md (create_prompt function)

**Acceptance Criteria:**
- All executable commands tested empirically
- Non-working commands corrected or marked as illustrative
- Verification results documented
- Specs updated with corrections

### 3. Mark Non-Executable Examples Clearly
**Priority:** High
**Impact:** Medium - prevents confusion about what should work
**Effort:** Low

Add clear markers to pseudocode and illustrative examples:
- Use "Pseudocode:" prefix for algorithm descriptions
- Use "Illustrative:" prefix for conceptual examples
- Use "Example (not executable):" for hypothetical scenarios
- Ensure executable examples have no such markers

**Acceptance Criteria:**
- All non-executable examples clearly marked
- Executable examples have no confusion markers
- Consistent marking pattern across all specs

### 4. Create Verification Tracking System
**Priority:** Medium
**Impact:** Medium - enables ongoing verification
**Effort:** Medium

Establish system to track verification status:
- Add verification metadata to specs (last verified date, status)
- Create verification checklist or tracking file
- Define re-verification triggers (spec updates, implementation changes)
- Document verification ownership

**Acceptance Criteria:**
- Verification status visible for each spec
- Re-verification process defined
- Ownership assigned

### 5. Automate Command Verification Where Possible
**Priority:** Low
**Impact:** High - reduces manual verification burden
**Effort:** High

Create automated verification tooling:
- Extract executable commands from specs
- Execute commands in test environment
- Validate expected outputs
- Report verification failures
- Integrate with CI/CD if applicable

**Acceptance Criteria:**
- Automated verification script exists
- Can be run on-demand or in CI
- Reports clear pass/fail results
- Covers majority of executable commands
