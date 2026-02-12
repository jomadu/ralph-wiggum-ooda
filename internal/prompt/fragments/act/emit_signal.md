# Emit Signal

Execute the signal decision from the Decide phase.

## Signal Types

**SUCCESS**: Emit when procedure goal is fully accomplished.

**FAILURE**: Emit when blockers prevent continuation.

**Continue**: Provide summary without signal to continue iterating.

## Format

```
<promise>SUCCESS</promise>
```

```
<promise>FAILURE</promise>
```

## Examples

**SUCCESS** (no work remaining):
```
<promise>SUCCESS</promise>

All ready work completed.
```

**FAILURE** (blocked):
```
<promise>FAILURE</promise>

Cannot proceed: Missing auth spec.
```

**Continue** (more work):
```
Completed task #42. 3 more tasks ready.
```
