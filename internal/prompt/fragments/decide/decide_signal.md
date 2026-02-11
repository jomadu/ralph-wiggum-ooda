# Decide Signal

Your task is to decide which signal to emit based on blockers, completion status, and remaining work.

## Check for Blockers

If any of these exist, decide to emit FAILURE:
- **Missing information** - Unclear requirements, ambiguous specifications, missing context
- **Missing tools/commands** - Build commands, test commands, or tools referenced in AGENTS.md not available
- **Missing dependencies** - Code libraries, external services, or data sources not accessible
- **Work tracking issues** - Cannot query tasks, cannot update status, work tracking system unavailable
- **Conflicting requirements** - Specifications contradict each other or implementation
- **Insufficient permissions** - Cannot read/write required files or execute required commands

## Check for Completion

If the procedure's goal is fully accomplished, decide to emit SUCCESS:
- **Work tracking procedures (build):** No ready work remains
- **Planning/auditing procedures:** Output is complete and accurate
- **Single-task procedures (sync):** The task is complete

## Check for Continuation

If no blockers exist and work remains, decide to continue iterating (no signal).

## Examples

**Blocker detected → Decide FAILURE:**
"Cannot proceed: Missing authentication module specification. The OAuth2 integration requires a detailed spec defining token refresh behavior. Decision: Emit FAILURE signal."

**Goal achieved → Decide SUCCESS:**
"All ready work completed. No tasks remain in work tracking. Decision: Emit SUCCESS signal."

**Work remains → Decide to continue:**
"Completed task #42. Three more ready tasks available. Decision: Continue iterating (no signal)."
