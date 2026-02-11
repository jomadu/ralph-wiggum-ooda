# Emit Signal

Your task is to emit the signal you decided on in the Decide phase. If you decided not to emit a signal, the loop will continue automatically to the next iteration.

## Signal Types

### SUCCESS - Goal Achieved
Emit when you decided the procedure's goal is fully accomplished and no further iteration is needed.

### FAILURE - Cannot Proceed
Emit when you decided the procedure cannot continue due to blockers identified in the Decide phase.

### Default Behavior - Continue Iterating
If you decided not to emit a signal, simply provide a summary of what was accomplished and the loop will continue to the next iteration with fresh context.

## Output Format

Output the exact signal format, then provide context:

```
<promise>SUCCESS</promise>
```

```
<promise>FAILURE</promise>
```

Or omit the signal to continue iterating:

```
Completed task #42: Add user authentication
- Implemented OAuth2 flow
- All tests passing
- 3 more ready tasks available in work tracking

(No signal - loop continues automatically)
```

## Examples

### SUCCESS Example (build procedure, no work remaining)
```
<promise>SUCCESS</promise>

All ready work completed:
- Implemented task #42: Add user authentication
- All tests passing
- No ready tasks remaining in work tracking
```

### SUCCESS Example (audit procedure, validated output)
```
<promise>SUCCESS</promise>

Audit completed and validated:
- Reviewed 12 specification files
- Generated audit report at docs/audit-2024-01-15.md
- Report is minimal yet complete and accurate
- Found 3 issues requiring attention
```

### Continue Example (build procedure, more work available)
```
Completed task #42: Add user authentication
- Implemented OAuth2 flow
- All tests passing
- 3 more ready tasks available in work tracking
```

### Continue Example (audit procedure, needs refinement)
```
Draft audit report generated but needs validation:
- Reviewed 12 specification files
- Report at docs/audit-2024-01-15.md may be too verbose
- Will verify findings against quality criteria in next iteration
```

### FAILURE Example (missing prerequisites)
```
<promise>FAILURE</promise>

Cannot proceed: Missing authentication module specification. The OAuth2 integration requires a detailed spec defining token refresh behavior and error handling patterns.

Next steps:
1. Create specs/auth-oauth2.md with token lifecycle specification
2. Define error handling patterns for expired tokens
3. Document refresh token rotation policy
```

### FAILURE Example (unresolvable error)
```
<promise>FAILURE</promise>

Cannot proceed: Build system failure. Tests fail with segmentation fault in auth module.

Next steps:
1. Debug segfault in src/auth/oauth2.go:142
2. Check for null pointer dereference in token validation
3. Run under debugger to capture stack trace
```

## Notes

- Only two signals exist: `<promise>SUCCESS</promise>` and `<promise>FAILURE</promise>`
- The loop orchestrator scans for these exact formats to determine iteration outcome
- If no signal is found, the loop continues automatically to the next iteration
- Always provide context explaining the current state
- For SUCCESS: summarize accomplishments and list modified files
- For FAILURE: explain why blocked and provide actionable next steps
- For continuing: summarize progress and what remains to be done
