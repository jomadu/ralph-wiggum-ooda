# Draft Plan: Spec to Implementation Gap Analysis

## Priority 1: Critical Acceptance Criteria Gaps

### Task 1: Implement CLI procedure-based invocation validation
**Gap:** cli-interface.md specifies procedure-based invocation should load OODA files from config (AC line 1), but implementation doesn't validate this works correctly.

**Current State:** rooda.sh lines 70-90 implement config lookup, but no test coverage exists.

**Acceptance Criteria:**
- Procedure name resolves to four OODA phase files from config
- Missing procedure produces clear error message
- Invalid config structure produces clear error message

**Implementation:** Add validation test cases or manual verification steps.

---

### Task 2: Implement explicit flag override behavior
**Gap:** cli-interface.md specifies explicit flags should override config-based procedure settings (AC line 2), but this behavior is not explicitly tested or documented in implementation.

**Current State:** rooda.sh argument parsing supports both modes, but precedence is implicit.

**Acceptance Criteria:**
- Explicit --observe/--orient/--decide/--act flags override procedure config
- Documented in rooda.sh comments or help text
- Test case verifies override behavior

**Implementation:** Add test case or documentation clarifying precedence.

---

### Task 3: Implement config file resolution relative to script location
**Gap:** cli-interface.md specifies config file should resolve relative to script location (AC line 3), implemented at line 28 but not validated.

**Current State:** `SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"` and `CONFIG_FILE="${SCRIPT_DIR}/rooda-config.yml"` implement this.

**Acceptance Criteria:**
- Config resolves correctly when script invoked from different directories
- Test case verifies resolution behavior

**Implementation:** Add test case for different invocation paths.

---

### Task 4: Implement missing file error messages
**Gap:** cli-interface.md specifies missing files should produce clear error messages (AC line 4), implemented at lines 107-111 but not all edge cases covered.

**Current State:** File validation exists but doesn't distinguish between missing OODA files vs missing config.

**Acceptance Criteria:**
- Error message identifies which file is missing
- Error message includes full path to missing file
- Different error messages for config vs OODA phase files

**Implementation:** Enhance error messages in validation blocks.

---

### Task 5: Implement invalid argument error handling
**Gap:** cli-interface.md specifies invalid arguments should produce usage help (AC line 5), implemented at lines 68-72 but usage message is inconsistent.

**Current State:** Unknown option triggers error, but usage message shows `<task-id>` which is not actually used.

**Acceptance Criteria:**
- Usage message matches actual argument structure
- Remove `<task-id>` from usage (not implemented)
- Consistent error messages across all validation points

**Implementation:** Fix usage messages at lines 68-72, 95-103, 117-125.

---

### Task 6: Implement max iterations default behavior
**Gap:** cli-interface.md specifies max iterations can be specified or defaults to procedure config (AC line 6), implemented at lines 88-91 but not fully tested.

**Current State:** Default iteration loading works, but fallback to 0 (unlimited) is implicit.

**Acceptance Criteria:**
- Command-line --max-iterations takes precedence
- Config default_iterations used if CLI not specified
- Defaults to 0 (unlimited) if neither specified
- Documented in help text

**Implementation:** Add documentation and test cases for three-tier default system.

---

## Priority 2: Iteration Loop Acceptance Criteria Gaps

### Task 7: Implement max iterations termination condition
**Gap:** iteration-loop.md specifies loop should execute until max iterations reached or Ctrl+C (AC line 1), implemented at lines 163-166 but not validated.

**Current State:** Termination logic exists but no test coverage.

**Acceptance Criteria:**
- Loop terminates when ITERATION >= MAX_ITERATIONS
- Loop runs indefinitely when MAX_ITERATIONS = 0
- Ctrl+C terminates immediately

**Implementation:** Add test cases for termination conditions.

---

### Task 8: Implement context clearing between iterations
**Gap:** iteration-loop.md specifies each iteration should exit completely, clearing AI context (AC line 2), but this is architectural - script doesn't actually exit between iterations.

**Current State:** Script runs in single bash process, AI CLI is invoked per iteration but bash context persists.

**Acceptance Criteria:**
- Document that AI context clears (kiro-cli exits after each invocation)
- Bash script context persists (by design)
- File-based state provides continuity

**Implementation:** Add documentation clarifying what "exit completely" means (AI CLI exits, not bash script).

---

### Task 9: Implement iteration counter display
**Gap:** iteration-loop.md specifies iteration counter should increment correctly (AC line 3), implemented at line 177 but display is confusing.

**Current State:** Counter increments after iteration completes, displays next iteration number.

**Acceptance Criteria:**
- Counter starts at 0
- Increments after each iteration
- Display message clearly indicates next iteration number

**Implementation:** Fix display message or document current behavior.

---

### Task 10: Implement git push per iteration
**Gap:** iteration-loop.md specifies git push should happen after each iteration (AC line 6), implemented at lines 171-174 but error handling is minimal.

**Current State:** Push attempts, creates branch if needed, but other failures are silent.

**Acceptance Criteria:**
- Push succeeds after each iteration
- Creates remote branch if missing
- Handles other push failures gracefully (auth, network, conflicts)

**Implementation:** Enhance error handling for git push failures.

---

## Priority 3: Configuration Schema Acceptance Criteria Gaps

### Task 11: Implement YAML structure validation
**Gap:** configuration-schema.md specifies YAML structure should support nested procedure definitions (AC line 1), implemented in rooda-config.yml but not validated at runtime.

**Current State:** yq queries assume valid structure, no validation before queries.

**Acceptance Criteria:**
- Config file structure validated before use
- Invalid YAML produces clear error message
- Missing required fields detected early

**Implementation:** Add config validation function called at startup.

---

### Task 12: Implement required fields validation
**Gap:** configuration-schema.md specifies required fields should be validated at runtime (AC line 2), partially implemented at lines 82-86 but only checks for null.

**Current State:** Null check exists but doesn't validate field types or values.

**Acceptance Criteria:**
- Validate observe/orient/decide/act are non-empty strings
- Validate file paths exist (optional, could be deferred)
- Clear error messages for missing required fields

**Implementation:** Enhance validation in config lookup section.

---

### Task 13: Implement optional fields support
**Gap:** configuration-schema.md specifies optional fields should be supported (AC line 3), implemented for default_iterations but not for display/summary/description.

**Current State:** display/summary/description defined in config but not used by script.

**Acceptance Criteria:**
- Optional fields don't cause errors if missing
- display/summary/description available for future help text generation
- default_iterations defaults to 0 if not specified

**Implementation:** Document that optional fields are for future use.

---

### Task 14: Implement yq query error handling
**Gap:** configuration-schema.md specifies yq queries should successfully extract procedure configuration (AC line 4), implemented but no error handling for yq failures.

**Current State:** yq errors propagate to user but aren't caught or explained.

**Acceptance Criteria:**
- yq parse errors produce clear error messages
- Invalid config structure detected and explained
- Suggest fixes for common config errors

**Implementation:** Wrap yq queries in error handling.

---

### Task 15: Implement missing procedure error messages
**Gap:** configuration-schema.md specifies missing procedures should return clear error messages (AC line 5), implemented at lines 82-86 but message could be clearer.

**Current State:** Error says "Procedure 'X' not found in config" but doesn't list available procedures.

**Acceptance Criteria:**
- Error message lists available procedures
- Suggests closest match (fuzzy matching)
- Includes config file path in error

**Implementation:** Enhance error message with available procedures list.

---

### Task 16: Implement file path resolution
**Gap:** configuration-schema.md specifies file paths should be resolved relative to script directory (AC line 6), implemented but not documented in config file.

**Current State:** Paths in config are relative to script location, but this isn't obvious from config file alone.

**Acceptance Criteria:**
- Config file includes comment explaining path resolution
- Documentation clarifies relative path behavior
- Test case verifies resolution from different working directories

**Implementation:** Add comments to rooda-config.yml explaining path resolution.

---

## Priority 4: AI CLI Integration Acceptance Criteria Gaps

### Task 17: Implement kiro-cli stdin prompt piping
**Gap:** ai-cli-integration.md specifies prompt should be piped to kiro-cli via stdin (AC line 1), implemented at line 169 but not validated.

**Current State:** `create_prompt | kiro-cli chat --no-interactive --trust-all-tools` implements this.

**Acceptance Criteria:**
- Prompt successfully piped to kiro-cli
- kiro-cli reads full prompt from stdin
- No truncation or corruption

**Implementation:** Add test case or manual verification.

---

### Task 18: Implement --no-interactive flag
**Gap:** ai-cli-integration.md specifies --no-interactive flag should disable interactive prompts (AC line 2), implemented but not validated.

**Current State:** Flag passed to kiro-cli, assumes kiro-cli honors it.

**Acceptance Criteria:**
- kiro-cli runs without user input
- No interactive prompts during execution
- Script continues automatically

**Implementation:** Validate kiro-cli supports flag and honors it.

---

### Task 19: Implement --trust-all-tools flag
**Gap:** ai-cli-integration.md specifies --trust-all-tools flag should bypass permission prompts (AC line 3), implemented but not validated.

**Current State:** Flag passed to kiro-cli, assumes kiro-cli honors it.

**Acceptance Criteria:**
- kiro-cli executes tools without permission prompts
- File read/write operations proceed automatically
- Command execution proceeds automatically

**Implementation:** Validate kiro-cli supports flag and honors it.

---

### Task 20: Implement AI file read capability validation
**Gap:** ai-cli-integration.md specifies AI should be able to read files from repository (AC line 4), not validated in implementation.

**Current State:** Assumes kiro-cli provides file read tools.

**Acceptance Criteria:**
- AI can read any file in repository
- File read operations succeed
- No permission errors

**Implementation:** Add test case or documentation.

---

### Task 21: Implement AI file write capability validation
**Gap:** ai-cli-integration.md specifies AI should be able to write/modify files (AC line 5), not validated in implementation.

**Current State:** Assumes kiro-cli provides file write tools.

**Acceptance Criteria:**
- AI can create new files
- AI can modify existing files
- Changes persist to disk

**Implementation:** Add test case or documentation.

---

### Task 22: Implement AI bash command execution validation
**Gap:** ai-cli-integration.md specifies AI should be able to execute bash commands (AC line 6), not validated in implementation.

**Current State:** Assumes kiro-cli provides command execution tools.

**Acceptance Criteria:**
- AI can execute bash commands
- Command output captured
- Command exit status available

**Implementation:** Add test case or documentation.

---

### Task 23: Implement AI git commit capability validation
**Gap:** ai-cli-integration.md specifies AI should be able to commit changes to git (AC line 7), not validated in implementation.

**Current State:** Assumes kiro-cli provides git tools or can execute git commands.

**Acceptance Criteria:**
- AI can create git commits
- Commit messages are descriptive
- Commits include all relevant changes

**Implementation:** Add test case or documentation.

---

### Task 24: Implement AI CLI exit status handling
**Gap:** ai-cli-integration.md specifies script should continue regardless of AI CLI exit status (AC line 8), implemented but not explicitly documented.

**Current State:** No exit status check after kiro-cli invocation at line 169.

**Acceptance Criteria:**
- Script continues to git push even if kiro-cli fails
- No error handling for kiro-cli failures
- Document this as intentional design decision

**Implementation:** Add comment explaining why exit status is ignored.

---

## Priority 5: External Dependencies Acceptance Criteria Gaps

### Task 25: Implement yq dependency check
**Gap:** external-dependencies.md specifies all external dependencies should be documented (AC line 1), yq is checked at lines 15-19 but other dependencies are not.

**Current State:** Only yq is checked at startup.

**Acceptance Criteria:**
- yq checked and error message provided if missing
- kiro-cli checked at startup (not currently implemented)
- bd checked at startup (not currently implemented)
- git checked at startup (optional)

**Implementation:** Add dependency checks for kiro-cli and bd.

---

### Task 26: Implement version requirements validation
**Gap:** external-dependencies.md specifies version requirements should be specified (AC line 2), documented in spec but not validated in implementation.

**Current State:** No version checking for any dependencies.

**Acceptance Criteria:**
- yq version >= 4.0.0 validated
- kiro-cli version >= 1.0.0 validated
- bd version >= 0.1.0 validated
- Clear error messages for incompatible versions

**Implementation:** Add version validation for all dependencies.

---

### Task 27: Implement installation instructions
**Gap:** external-dependencies.md specifies installation instructions should be provided per platform (AC line 3), documented in spec but not in script error messages.

**Current State:** yq error message includes brew install command, but not Linux alternatives.

**Acceptance Criteria:**
- Error messages include platform-specific install commands
- Detect OS and provide appropriate instructions
- Link to detailed installation docs

**Implementation:** Enhance error messages with platform detection.

---

### Task 28: Implement dependency checking for critical tools
**Gap:** external-dependencies.md specifies dependency checking should be implemented for critical tools (AC line 4), only yq is checked.

**Current State:** kiro-cli and bd failures happen at runtime, not startup.

**Acceptance Criteria:**
- Check for kiro-cli at startup
- Check for bd at startup (if work tracking uses beads)
- Provide clear error messages with installation instructions

**Implementation:** Add startup checks for kiro-cli and bd.

---

## Priority 6: Prompt Composition Acceptance Criteria Gaps

### Task 29: Implement prompt assembly algorithm documentation
**Gap:** prompt-composition.md specifies algorithm should be documented (AC line 1), implemented at lines 143-159 but not documented inline.

**Current State:** create_prompt() function exists but lacks comments.

**Acceptance Criteria:**
- Function includes comments explaining assembly process
- Heredoc structure documented
- Command substitution explained

**Implementation:** Add inline comments to create_prompt() function.

---

### Task 30: Implement file validation behavior
**Gap:** prompt-composition.md specifies file validation behavior should be specified (AC line 2), implemented at lines 107-111 and 117-125 but duplicated.

**Current State:** Duplicate validation blocks exist.

**Acceptance Criteria:**
- Single validation block (remove duplication)
- Clear error messages for missing files
- Validation happens before create_prompt() is called

**Implementation:** Refactor to remove duplicate validation blocks.

---

### Task 31: Implement error handling for missing files
**Gap:** prompt-composition.md specifies error handling for missing files should be defined (AC line 3), implemented but could be clearer.

**Current State:** File existence check at lines 107-111 and 117-125.

**Acceptance Criteria:**
- Error message includes file path
- Error message explains which OODA phase is missing
- Script exits with non-zero status

**Implementation:** Enhance error messages with OODA phase context.

---

### Task 32: Implement valid markdown output
**Gap:** prompt-composition.md specifies output format should produce valid markdown (AC line 4), implemented but not validated.

**Current State:** create_prompt() produces markdown structure.

**Acceptance Criteria:**
- Output is valid markdown
- Section headers are properly formatted
- File contents don't break markdown structure

**Implementation:** Add markdown validation or linting.

---

### Task 33: Implement OODA phase section headers
**Gap:** prompt-composition.md specifies section headers should clearly delineate OODA phases (AC line 5), implemented at lines 145-158 but not validated.

**Current State:** Headers are "## OBSERVE", "## ORIENT", "## DECIDE", "## ACT".

**Acceptance Criteria:**
- Headers are consistent across all prompts
- Headers are recognizable by AI
- Headers match expected format

**Implementation:** Add test case or documentation.

---

## Summary

**Total Tasks:** 33

**By Priority:**
- Priority 1 (Critical CLI): 6 tasks
- Priority 2 (Iteration Loop): 4 tasks
- Priority 3 (Configuration): 6 tasks
- Priority 4 (AI Integration): 8 tasks
- Priority 5 (Dependencies): 4 tasks
- Priority 6 (Prompt Composition): 5 tasks

**Implementation Status:**
- Most acceptance criteria are partially implemented
- Primary gaps are in validation, error handling, and documentation
- No missing core functionality, mostly quality and robustness improvements
- Test coverage is the biggest gap (no automated tests exist)

**Recommended Approach:**
1. Start with Priority 1 (CLI validation and error messages)
2. Add dependency checks (Priority 5, Tasks 25-28)
3. Enhance error handling (Priority 2, Priority 3)
4. Add documentation and comments (Priority 4, Priority 6)
5. Consider adding automated tests (not in current specs)
