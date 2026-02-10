# Test Cases for Signal Examples in iteration-loop.md

## Test 1: Signal Placement Note Exists
**Expected:** A note in the "Promise Signal Format" or "Design Rationale" section explaining that signals should be emitted at the END of output

**Verification:** Search for text mentioning "end of output" or "after all work complete" in context of promise signals

## Test 2: SUCCESS Semantics Clarified
**Expected:** Documentation clarifies that SUCCESS means the procedure's specific goal is achieved:
- build procedure: no more ready work items
- audit procedures: validated audit report produced
- planning procedures: validated draft plan produced
- agents-sync: AGENTS.md synchronized with repository state

**Verification:** Check if the spec explains what SUCCESS means for different procedure types

## Test 3: Signal Separation in Examples
**Expected:** All examples show signals on their own line, not embedded in explanatory text

**Examples to check:**
- Example 3: Verbose Mode (line ~525)
- Example 4: Dry-Run Mode (preamble examples)
- Example 5: Dry-Run Mode with User Context (preamble examples)

**Verification:** Signals appear as standalone lines like:
```
<promise>SUCCESS</promise>
```

NOT like:
```
Task complete <promise>SUCCESS</promise> - all tests passing
```

## Test 4: Consistency Across Examples
**Expected:** All examples follow the same pattern:
1. Signal appears at the end of AI output
2. Signal is on its own line
3. Explanatory text (if any) comes AFTER the signal

**Verification:** Review all examples in the Examples section

## Test 5: Preamble Instructions Match Best Practices
**Expected:** The preamble examples include the instruction:
"Explanations should come AFTER the signal, not embedded in the tag"

**Verification:** Check all preamble examples (Examples 4, 5, 6, 7, 8)
