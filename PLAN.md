# Specification Refactoring Plan

## Priority 1: Verify Command Examples in Specifications

**Impact:** HIGH - Incorrect command examples mislead users and cause runtime failures

**Tasks:**
- Verify all bash command examples in external-dependencies.md work as documented
- Verify all bash command examples in cli-interface.md work as documented
- Verify all bash command examples in iteration-loop.md work as documented
- Verify all bash command examples in configuration-schema.md work as documented
- Verify all bash command examples in ai-cli-integration.md work as documented
- Verify all bash command examples in component-authoring.md work as documented
- Verify all bash command examples in agents-md-format.md work as documented
- Verify all bash command examples in component-system.md work as documented (deprecated but retained)
- Verify all bash command examples in prompt-composition.md work as documented (deprecated but retained)

**Acceptance Criteria:**
- All command examples execute successfully or are marked as pseudocode/illustrative
- Commands that reference actual implementation (./rooda.sh, yq, bd, kiro-cli) are verified working
- Pseudocode examples are clearly marked as non-executable
- Any broken commands are corrected or removed

**Effort:** MEDIUM - Requires running each command and validating output

**Risk:** LOW - Verification only, no structural changes to specs

## Priority 2: Document Command Verification Process

**Impact:** MEDIUM - Prevents future drift between specs and implementation

**Tasks:**
- Add "Command Verification" section to AGENTS.md quality criteria
- Document how to verify command examples (manual execution, automated testing)
- Define what constitutes "verified working" (exit code 0, expected output, no errors)
- Clarify distinction between executable commands and pseudocode examples

**Acceptance Criteria:**
- AGENTS.md contains clear guidance on command verification
- Future spec authors know how to verify command examples
- Quality criteria checking includes command verification step

**Effort:** LOW - Documentation update only

**Risk:** LOW - Clarifies existing process

## Priority 3: Add Verification Markers to Specs

**Impact:** LOW - Improves transparency but doesn't affect functionality

**Tasks:**
- Consider adding verification markers to command examples (e.g., "✓ Verified 2026-02-03")
- Evaluate if verification markers add value or create maintenance burden
- If valuable, add markers to all verified commands

**Acceptance Criteria:**
- Decision made on whether to use verification markers
- If using markers, all verified commands are marked
- If not using markers, rationale documented

**Effort:** LOW - Simple annotation if adopted

**Risk:** LOW - Optional enhancement

## Notes

**Why Command Verification Failed:**
The quality criterion "All command examples in specs are verified working" is boolean (PASS/FAIL), but no verification process was established when specs were created. Command examples were written based on implementation understanding but not empirically tested.

**Scope:**
This plan focuses on specification quality only. Implementation quality criteria (shellcheck, procedure execution, cross-platform compatibility) are separate and not assessed here.

**Deprecated Specs:**
component-system.md and prompt-composition.md are deprecated but retained for historical reference. Their command examples should still be verified to ensure historical accuracy.
