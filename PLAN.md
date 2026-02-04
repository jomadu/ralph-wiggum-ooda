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

## Refactoring Tasks

### Task 1: Fix agents-md-format.md Structure
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

## Summary

- **Criteria failing:** 4 of 5 (criteria 1, 2, 3, 4)
- **Criteria passing:** 1 of 5 (criterion 5)
- **Root cause:** agents-md-format.md doesn't follow template (affects 3 criteria)
- **Action:** Task 1 (fix agents-md-format.md structure) - highest impact, lowest effort
