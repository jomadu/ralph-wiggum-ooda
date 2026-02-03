# Draft Plan: Spec to Implementation Gap Analysis

## Priority 1: Critical Missing Features

### 1. Dependency Checking for kiro-cli and bd
**Gap:** external-dependencies.md specifies that kiro-cli and bd should be checked at startup, but rooda.sh only checks for yq (lines 15-19). Users discover missing tools only when procedures execute.

**Spec Reference:** external-dependencies.md - "Late failure for kiro-cli/bd: Script doesn't check for kiro-cli or bd at startup"

**Acceptance Criteria:**
- Add startup checks for kiro-cli and bd before entering iteration loop
- Display clear error messages with installation instructions if missing
- Exit with status 1 if required tools are missing

**Implementation:**
- Add check_dependency() function that validates command existence
- Call for yq, kiro-cli, and bd before argument parsing
- Provide platform-specific installation instructions in error messages

---

### 2. Version Validation for yq
**Gap:** external-dependencies.md specifies minimum yq v4.0.0 required, but rooda.sh only checks if yq exists, not version. Users with yq v3 get cryptic YAML parsing errors.

**Spec Reference:** external-dependencies.md - "No version validation: Script checks if yq exists but not if it's v4+"

**Acceptance Criteria:**
- Extract yq version from `yq --version` output
- Compare against minimum version 4.0.0
- Display clear error if version is too old
- Provide upgrade instructions

**Implementation:**
- Parse `yq --version` output to extract version number
- Use version comparison logic (major.minor.patch)
- Error message: "yq v4.0.0+ required, found vX.Y.Z. Upgrade with: brew upgrade yq"

---

### 3. Error Handling for kiro-cli Failures
**Gap:** iteration-loop.md and ai-cli-integration.md specify that kiro-cli exit status should be checked, but rooda.sh line 167 pipes to kiro-cli without checking exit status. Failed iterations continue to git push.

**Spec Reference:** 
- iteration-loop.md - "No kiro-cli error handling: If kiro-cli exits with non-zero status, the loop continues anyway"
- ai-cli-integration.md - "No error handling: Script continues to git push even if kiro-cli fails"

**Acceptance Criteria:**
- Capture kiro-cli exit status after each iteration
- Break loop after N consecutive failures (N=3)
- Display clear error message when breaking due to failures
- Skip git push if kiro-cli failed

**Implementation:**
- Store kiro-cli exit status in variable
- Track consecutive failure count
- Add conditional git push based on success
- Reset failure count on successful iteration

---

## Priority 2: Code Quality Issues

### 4. Remove Duplicate Validation Blocks
**Gap:** cli-interface.md identifies duplicate validation at lines 95-103 and 117-125 in rooda.sh. This violates DRY principle and creates maintenance burden.

**Spec Reference:** cli-interface.md - "Duplicate validation blocks: Lines 95-103 and 117-125 contain identical validation logic"

**Acceptance Criteria:**
- Single validation block for OODA phase files
- Single validation block for file existence
- No change in validation behavior
- shellcheck passes

**Implementation:**
- Keep validation at lines 95-103 (after config loading)
- Remove duplicate validation at lines 117-125
- Verify all code paths reach validation before iteration loop

---

### 5. Fix Iteration Display Off-by-One
**Gap:** iteration-loop.md notes that separator shows "LOOP $ITERATION" after incrementing, displaying next iteration number instead of completed iteration number. This is confusing.

**Spec Reference:** iteration-loop.md - "Iteration display off-by-one: The separator shows 'LOOP $ITERATION' after incrementing"

**Acceptance Criteria:**
- Separator displays completed iteration number
- First iteration shows "LOOP 1" after completion
- User-facing messages are 1-indexed (not 0-indexed)

**Implementation:**
- Display separator before incrementing counter
- Or adjust display to show ITERATION+1 before increment

---

## Priority 3: Missing Features (Nice-to-Have)

### 6. Add --help Flag Support
**Gap:** cli-interface.md identifies missing --help flag. Users must trigger errors to see usage.

**Spec Reference:** cli-interface.md - "Help flag: No --help or -h flag support"

**Acceptance Criteria:**
- `./rooda.sh --help` displays usage information
- `./rooda.sh -h` displays usage information
- Help text includes both invocation modes (procedure and explicit flags)
- Help text includes examples
- Exit with status 0 after displaying help

**Implementation:**
- Add --help|-h case in argument parsing
- Display usage, examples, and available procedures
- Query config file to list available procedures with summaries

---

### 7. Add --version Flag Support
**Gap:** cli-interface.md identifies missing --version flag.

**Spec Reference:** cli-interface.md - "Version flag: No --version flag to show script version"

**Acceptance Criteria:**
- `./rooda.sh --version` displays version number
- Version follows semantic versioning
- Exit with status 0 after displaying version

**Implementation:**
- Add VERSION variable at top of script
- Add --version case in argument parsing
- Display: "rooda.sh version X.Y.Z"

---

### 8. Improve Git Push Error Handling
**Gap:** iteration-loop.md notes that git push failures (other than missing remote branch) are silent.

**Spec Reference:** iteration-loop.md - "Git push failures: If git push fails for reasons other than missing remote branch, the error is silent"

**Acceptance Criteria:**
- Distinguish between "no remote branch" and other git push failures
- Display clear error messages for push failures
- Continue loop after push failure (non-fatal)
- Log push failures for debugging

**Implementation:**
- Check git push exit status
- Parse error output to identify failure type
- Display appropriate error message
- Don't break loop (push failures are non-fatal)

---

## Priority 4: Documentation Gaps

### 9. Document Prompt Component Structure
**Gap:** prompt-composition.md exists as spec but there's no corresponding documentation for prompt component file structure. Users don't know how to write custom OODA phase files.

**Spec Reference:** prompt-composition.md defines assembly but not component structure

**Acceptance Criteria:**
- Document markdown structure for OODA phase files
- Provide examples of observe/orient/decide/act components
- Explain how components are composed
- Document best practices for writing custom components

**Implementation:**
- Create docs/prompt-components.md
- Include examples from src/components/
- Document naming conventions
- Link from README.md

---

## Out of Scope

The following gaps were identified but are intentionally deferred:

- **Config validation:** Validating config file structure before querying (low priority, yq errors are clear enough)
- **Short flags:** Adding -o, -m, etc. (low priority, explicit flags are clearer)
- **Verbose/quiet modes:** Output verbosity control (low priority, current output is reasonable)
- **Dry-run mode:** Preview without execution (nice-to-have, not critical)
- **Resume capability:** Save/restore iteration state (complex, low ROI)
- **Timeout mechanism:** Timeout for kiro-cli invocation (edge case, Ctrl+C works)
- **Prompt size validation:** Check token limits (AI CLI responsibility)
- **Alternative CLI support:** Document other AI CLIs (kiro-cli is the target)

---

## Implementation Notes

**Testing Strategy:**
- Manual verification of each change (no automated tests per AGENTS.md)
- Test both procedure-based and explicit flag invocations
- Test error conditions (missing tools, bad versions, invalid args)
- Verify shellcheck passes after changes

**Rollout:**
- Implement Priority 1 tasks first (critical for user experience)
- Priority 2 improves code quality (low risk)
- Priority 3 adds convenience features (optional)
- Priority 4 improves documentation (can be done in parallel)

**Dependencies:**
- Tasks 1-3 are independent (can be parallelized)
- Task 4 should be done after tasks 1-3 (avoid merge conflicts)
- Tasks 6-7 are independent
- Task 9 is independent (documentation only)
