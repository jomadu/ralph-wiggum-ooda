# Plan: Clarify Promise Signal Format

## Problem

Contradiction between specs and implementation fragments regarding promise signal format:

**Specs say**: Exact string match only - `<promise>SUCCESS</promise>` and `<promise>FAILURE</promise>`

**Fragments imply**: Reasons might be included in signals (ambiguous guidance)

This could lead AI agents to emit invalid signals like `<promise>FAILURE: reason</promise>` which won't be detected by the parser.

## Goal

Ensure all specs and prompt fragments consistently communicate:
1. Promise signals use exact format with no variations
2. Explanatory text comes AFTER the signal, not within it
3. Examples show correct separation of signal and explanation
4. SUCCESS signal means procedure goal achieved (for build: no ready work remains; for planning/auditing: output validated as minimal, complete, and accurate), not just iteration success

## Files to Update

### Prompt Fragments

1. `internal/prompt/fragments/act/emit_success.md`
   - Add explicit format: `<promise>SUCCESS</promise>`
   - Show example with explanation after signal
   - Clarify signal must be exact, explanation separate

2. `internal/prompt/fragments/act/emit_failure.md`
   - Add explicit format: `<promise>FAILURE</promise>`
   - Show example with explanation after signal
   - Clarify signal must be exact, explanation separate

3. `internal/prompt/fragments/decide/check_if_blocked.md`
   - Change "emit FAILURE with explanation" to "emit FAILURE signal, then explain"
   - Clarify signal format is exact, explanation comes after

4. `internal/prompt/fragments/act/write_audit_report.md`
   - Add validation step: review report for minimal, complete, and accurate
   - Only proceed to emit_success after validation

5. `internal/prompt/fragments/act/write_draft_plan.md`
   - Add validation step: review plan for minimal, complete, and accurate
   - Only proceed to emit_success after validation

6. `internal/prompt/fragments/act/write_gap_report.md`
   - Add validation step: review report for minimal, complete, and accurate
   - Only proceed to emit_success after validation

### Spec Files

7. `specs/error-handling.md`
   - Review all examples showing FAILURE signals
   - Ensure no examples show reasons embedded in signal tags
   - Add explicit guidance that explanations come after signal
   - Verify SUCCESS signal semantics align with procedure goals (not just iteration success)
   - Correct any inconsistencies with signal format or semantics

8. `specs/iteration-loop.md`
   - Review all examples showing SUCCESS/FAILURE signals
   - Ensure examples show proper separation
   - Add note about signal placement (end of output, after explanations)
   - Clarify SUCCESS means procedure goal achieved (build: no ready work; planning/auditing: validated output)
   - Correct any inconsistencies with signal format or semantics

9. `specs/ai-cli-integration.md`
   - Review signal scanning examples
   - Ensure examples show correct format
   - Verify no ambiguous examples
   - Correct any inconsistencies with signal format or semantics

10. `specs/procedures.md`
    - Review procedure descriptions and examples
    - Ensure SUCCESS criteria align with procedure goals
    - Verify examples show correct signal format
    - Correct any inconsistencies with signal format or semantics

## Changes Required

### emit_success.md
```markdown
# Emit Success

You must output a SUCCESS promise when the procedure's goal is fully achieved.

For work tracking procedures (build): Signal SUCCESS only when no ready work remains.
For planning/auditing procedures: Signal SUCCESS only when output is minimal, complete, and accurate.
For single-task procedures (sync): Signal SUCCESS when the task completes.

Actions:
- Output the exact signal: <promise>SUCCESS</promise>
- After the signal, include summary of what was accomplished
- List files modified or created
- Note any follow-up actions needed

Example (build procedure, no work remaining):
<promise>SUCCESS</promise>

All ready work completed:
- Implemented task #42: Add user authentication
- All tests passing
- No ready tasks remaining in work tracking

Example (audit procedure, validated output):
<promise>SUCCESS</promise>

Audit completed and validated:
- Reviewed 12 specification files
- Generated audit report at docs/audit-2024-01-15.md
- Report is minimal yet complete and accurate
- Found 3 issues requiring attention

Example (draft plan procedure, validated output):
<promise>SUCCESS</promise>

Draft plan completed and validated:
- Created plan at docs/draft-plan-auth-feature.md
- Plan is minimal yet complete and accurate
- Broken down into 8 actionable tasks
- Ready for import to work tracking
```

### emit_failure.md
```markdown
# Emit Failure

You must output a FAILURE promise to indicate the procedure cannot proceed.

Actions:
- Output the exact signal: <promise>FAILURE</promise>
- After the signal, explain why work is blocked
- List missing prerequisites or dependencies
- Suggest what needs to happen to unblock
- Provide actionable next steps

Example:
<promise>FAILURE</promise>

Work is blocked because:
- No work tracking system detected (.beads/ or .github/issues not found)
- AGENTS.md does not specify work tracking configuration

To unblock:
1. Run `rooda bootstrap` to configure work tracking
2. Or manually create .beads/ directory and initialize beads
```

### check_if_blocked.md
```markdown
# Check If Blocked

Your task is to determine if work can proceed or if blockers exist.

Check for:
- Missing dependencies or prerequisites
- Unresolved questions or ambiguities
- Required resources not available
- Conflicting requirements

If blocked, output the exact signal: <promise>FAILURE</promise>
Then explain the blockers.

Otherwise, proceed to the ACT phase.
```

### write_audit_report.md
```markdown
# Write Audit Report

You must create an audit report documenting the findings. Use the file writing tool.

Include:
- Executive summary of findings
- Detailed list of issues with evidence
- Quality criteria results (PASS/FAIL)
- Prioritized recommendations
- Metrics and statistics

After writing, validate the report:
- Is it minimal? (No unnecessary content)
- Is it complete? (All findings documented)
- Is it accurate? (Evidence supports conclusions)

Only proceed to emit_success if validation passes.
```

### write_draft_plan.md
```markdown
# Write Draft Plan

You must create a draft plan file (PLAN.md) with the prioritized tasks. Use the file writing tool.

Include:
- Overview and objectives
- Task list with descriptions
- Priorities and dependencies
- Acceptance criteria for each task
- Implementation notes

After writing, validate the plan:
- Is it minimal? (No unnecessary tasks)
- Is it complete? (All requirements covered)
- Is it accurate? (Tasks are actionable and correctly scoped)

Only proceed to emit_success if validation passes.
```

### write_gap_report.md
```markdown
# Write Gap Report

You must create a gap analysis report documenting the findings. Use the file writing tool.

Include:
- Summary of gaps found
- Spec-to-impl gaps (specified but not implemented)
- Impl-to-spec gaps (implemented but not specified)
- Impact assessment for each gap
- Recommended actions

After writing, validate the report:
- Is it minimal? (No unnecessary content)
- Is it complete? (All gaps documented)
- Is it accurate? (Gap analysis is correct)

Only proceed to emit_success if validation passes.
```

### Spec Updates

All spec files must be reviewed for inconsistencies and corrected:

**Signal Format Consistency:**
- All examples showing signals must follow exact format: `<promise>SUCCESS</promise>` or `<promise>FAILURE</promise>`
- No examples with reasons embedded: `<promise>FAILURE: reason</promise>` ❌
- Explanations always come after signal, not within tags

**Signal Semantics Consistency:**
- SUCCESS means procedure goal achieved, not iteration success
- For build procedure: SUCCESS = no ready work remains
- For planning/auditing procedures: SUCCESS = output validated as minimal, complete, and accurate
- For single-task procedures: SUCCESS = task completes

**Examples to investigate and correct:**

Pattern to find:
```
[AI output and work...]

<promise>SUCCESS</promise>

[Optional summary after signal]
```

NOT:
```
<promise>SUCCESS: completed task</promise>
<promise>FAILURE: missing API key</promise>
```

**Specific areas to review in each spec:**
- error-handling.md: Failure detection examples, outcome matrix examples
- iteration-loop.md: Loop termination examples, iteration outcome examples
- ai-cli-integration.md: Signal scanning examples, execution result examples
- procedures.md: Procedure descriptions, built-in procedure examples

## Validation

After updates:
- [ ] All prompt fragments show exact signal format
- [ ] All prompt fragments show examples with proper separation
- [ ] emit_success.md clarifies SUCCESS means procedure goal achieved, not iteration success
- [ ] emit_success.md specifies for build procedure: SUCCESS only when no ready work remains
- [ ] emit_success.md specifies for planning/auditing procedures: SUCCESS only when output validated as minimal, complete, and accurate
- [ ] write_audit_report.md includes validation step before proceeding to emit_success
- [ ] write_draft_plan.md includes validation step before proceeding to emit_success
- [ ] write_gap_report.md includes validation step before proceeding to emit_success
- [ ] All spec examples use correct format
- [ ] No spec examples show reasons embedded in signal tags
- [ ] Spec files consistently describe SUCCESS as procedure goal achieved, not iteration success
- [ ] Spec files consistently describe signal format as exact match with explanations after
- [ ] All inconsistencies between specs and fragments resolved
- [ ] Signal scanning logic remains simple (exact string match)

## Rationale

**Why exact format?**
- Simple string matching (no regex, no parsing)
- No ambiguity in detection
- Forces AI to follow specification precisely
- Prevents false negatives from format variations

**Why explanation after signal?**
- Signal detection happens first (fast path)
- Explanations are for human debugging, not parsing
- Keeps signal scanning O(1) complexity
- Buffer truncation preserves signals at end
