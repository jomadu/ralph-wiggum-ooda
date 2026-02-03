# Draft Plan: Spec to Implementation Gap Analysis

## Priority 1: Critical Gaps (Missing Core Functionality)

### 1. Implement Version Validation for yq Dependency
**Gap:** external-dependencies.md specifies yq v4.0.0 minimum version requirement, but rooda.sh only checks if yq exists (line 15), not version compatibility.

**Impact:** Users with yq v3 will get cryptic YAML parsing errors instead of clear version mismatch message.

**Acceptance Criteria:**
- Check yq version at startup
- Warn if version < v4.0.0
- Provide upgrade instructions

**Implementation:**
- Modify rooda.sh dependency check section (lines 15-19)
- Extract version from `yq --version` output
- Compare against minimum version 4.0.0

---

### 2. Implement Early Dependency Checks for kiro-cli and bd
**Gap:** external-dependencies.md specifies kiro-cli and bd as required dependencies, but rooda.sh doesn't check for them at startup. Users discover missing tools only when procedures execute (late failure).

**Impact:** Poor user experience - script starts, parses config, then fails mid-execution.

**Acceptance Criteria:**
- Check for kiro-cli at startup
- Check for bd at startup (optional - only if procedure uses work tracking)
- Provide installation instructions if missing
- Exit with clear error before entering iteration loop

**Implementation:**
- Add dependency checks after yq validation
- Use `command -v kiro-cli` and `command -v bd`
- Display installation instructions per platform

---

### 3. Implement Error Handling for kiro-cli Failures
**Gap:** ai-cli-integration.md and iteration-loop.md note that kiro-cli exit status is not checked. Script continues to git push even if AI CLI fails.

**Impact:** Repeated failures without termination, pushing incomplete/invalid changes.

**Acceptance Criteria:**
- Check kiro-cli exit status after each iteration
- Break loop after N consecutive failures (e.g., 3)
- Display clear error message
- Don't push changes if kiro-cli failed

**Implementation:**
- Capture exit status: `create_prompt | kiro-cli chat --no-interactive --trust-all-tools; EXIT_CODE=$?`
- Track consecutive failures counter
- Conditional git push based on success

---

## Priority 2: High-Impact Gaps (Quality & Usability)

### 4. Add Help Flag Support
**Gap:** cli-interface.md notes "No `--help` or `-h` flag support. Users must trigger an error to see usage."

**Impact:** Poor discoverability - users can't easily learn how to use the tool.

**Acceptance Criteria:**
- `./rooda.sh --help` displays usage information
- `./rooda.sh -h` displays usage information
- Help text includes examples for both invocation modes
- Help text lists available procedures from config

**Implementation:**
- Add --help/-h case in argument parsing
- Display usage, examples, and available procedures
- Exit with status 0

---

### 5. Add Version Flag Support
**Gap:** cli-interface.md notes "No `--version` flag to show script version."

**Impact:** Users can't determine which version they're running for troubleshooting.

**Acceptance Criteria:**
- `./rooda.sh --version` displays version number
- Version matches git tag or commit hash
- Exit with status 0

**Implementation:**
- Add VERSION variable at top of script
- Add --version case in argument parsing
- Display version and exit

---

### 6. Implement Timeout Mechanism for AI CLI
**Gap:** ai-cli-integration.md notes "No timeout: If kiro-cli hangs, the script waits indefinitely."

**Impact:** Script can hang forever if AI CLI has issues.

**Acceptance Criteria:**
- Set timeout for kiro-cli invocation (e.g., 10 minutes)
- Kill process if timeout exceeded
- Display timeout error message
- Continue to next iteration or exit based on policy

**Implementation:**
- Use `timeout` command wrapper: `timeout 600 kiro-cli chat ...`
- Check exit code 124 (timeout)
- Handle timeout as failure

---

### 7. Implement Prompt Size Validation
**Gap:** ai-cli-integration.md notes "No prompt size validation: Large OODA phase files could exceed token limits."

**Impact:** Silent truncation or failures when prompts are too large.

**Acceptance Criteria:**
- Calculate approximate token count before piping to AI
- Warn if approaching token limits (e.g., 80% of 200K)
- Display prompt size in characters and estimated tokens
- Allow override with --force flag

**Implementation:**
- Count characters in assembled prompt
- Estimate tokens (rough: chars / 4)
- Display warning if > 160K tokens
- Proceed with confirmation or --force

---

## Priority 3: Code Quality Improvements

### 8. Remove Duplicate Validation Blocks
**Gap:** cli-interface.md notes "Lines 95-103 and 117-125 contain identical validation logic."

**Impact:** Code duplication makes maintenance harder.

**Acceptance Criteria:**
- Single validation function
- Called once after argument parsing
- No duplicate code

**Implementation:**
- Extract validation to function
- Call after all argument parsing complete
- Remove duplicate blocks at lines 95-103 and 117-125

---

### 9. Fix Iteration Display Off-by-One
**Gap:** iteration-loop.md notes "Separator shows 'LOOP $ITERATION' after incrementing, displays next iteration number not the one that just completed."

**Impact:** Confusing progress messages.

**Acceptance Criteria:**
- Display "Completed iteration N" after iteration finishes
- Display "Starting iteration N" before iteration starts
- Clear distinction between completed and upcoming

**Implementation:**
- Display message before incrementing counter
- Or adjust display to show completed iteration number

---

### 10. Add Config File Structure Validation
**Gap:** cli-interface.md notes "Script doesn't validate config file structure before querying with yq."

**Impact:** Cryptic yq errors if config is malformed.

**Acceptance Criteria:**
- Validate YAML syntax before querying
- Check for required top-level keys (procedures)
- Provide clear error messages for malformed config

**Implementation:**
- Use `yq eval '.' config.yml` to validate syntax
- Check for `.procedures` key existence
- Display helpful error if validation fails

---

## Priority 4: Documentation Gaps

### 11. Document Tested Version Combinations
**Gap:** external-dependencies.md notes "No dependency matrix: Document tested version combinations."

**Impact:** Users don't know which versions are compatible.

**Acceptance Criteria:**
- Document tested versions in external-dependencies.md
- Include OS (macOS, Linux), yq version, kiro-cli version, bd version
- Update after testing new combinations

**Implementation:**
- Add "Tested Combinations" section to spec
- Create table with OS, tool versions, test date, status

---

### 12. Add Platform-Specific Installation Instructions
**Gap:** external-dependencies.md notes "Installation commands assume macOS (brew) or Linux package managers. Windows/WSL users need different instructions."

**Impact:** Windows users can't easily install dependencies.

**Acceptance Criteria:**
- Document Windows/WSL installation for all dependencies
- Include alternative installation methods (binary downloads, cargo, npm)
- Test instructions on Windows/WSL

**Implementation:**
- Expand installation sections in external-dependencies.md
- Add Windows-specific commands
- Add alternative installation methods

---

## Priority 5: Future Enhancements (Low Priority)

### 13. Add Short Flag Alternatives
**Gap:** cli-interface.md notes "No short flag alternatives (e.g., `-o` for `--observe`, `-m` for `--max-iterations`)."

**Impact:** More typing required for common operations.

**Acceptance Criteria:**
- Support `-o` for `--observe`
- Support `-r` for `--orient`
- Support `-d` for `--decide`
- Support `-a` for `--act`
- Support `-m` for `--max-iterations`
- Support `-c` for `--config`

**Implementation:**
- Add short flag cases in argument parsing
- Update usage/help text

---

### 14. Add Verbose/Quiet Modes
**Gap:** cli-interface.md notes "No control over output verbosity."

**Impact:** Can't suppress or enhance output for different use cases.

**Acceptance Criteria:**
- `--verbose` flag shows detailed execution info
- `--quiet` flag suppresses non-error output
- Default mode shows standard progress

**Implementation:**
- Add VERBOSE and QUIET variables
- Conditional echo statements based on mode
- Add flags to argument parsing

---

### 15. Implement Resume Capability
**Gap:** iteration-loop.md notes "No resume capability: Save iteration state to file, allowing resume after Ctrl+C or failure."

**Impact:** Can't resume interrupted loops.

**Acceptance Criteria:**
- Save iteration state to .rooda-state file
- `--resume` flag continues from saved state
- Clear state on successful completion

**Implementation:**
- Write ITERATION, PROCEDURE, MAX_ITERATIONS to state file
- Read state file on --resume
- Clean up state file on completion

---

### 16. Add Dry-Run Mode
**Gap:** iteration-loop.md notes "No dry-run mode: Support --dry-run flag to show what would execute without actually running the AI CLI."

**Impact:** Can't preview execution without running.

**Acceptance Criteria:**
- `--dry-run` flag shows assembled prompt
- Displays what would be executed
- Exits without running kiro-cli

**Implementation:**
- Add --dry-run flag
- Display prompt instead of piping to kiro-cli
- Exit after display

---

### 17. Add Iteration Timing Display
**Gap:** iteration-loop.md notes "No iteration timing: Display elapsed time per iteration to help users understand performance."

**Impact:** Can't measure iteration performance.

**Acceptance Criteria:**
- Display start time before iteration
- Display end time and duration after iteration
- Show cumulative time for all iterations

**Implementation:**
- Capture start time with `date +%s`
- Calculate duration after iteration
- Display in progress messages

---

### 18. Add Progress Indicators During Execution
**Gap:** iteration-loop.md notes "No progress indicators: Show which OODA phase is executing during long-running iterations."

**Impact:** No visibility into what's happening during long iterations.

**Acceptance Criteria:**
- Display "Executing OBSERVE phase..."
- Display "Executing ORIENT phase..."
- Display "Executing DECIDE phase..."
- Display "Executing ACT phase..."

**Implementation:**
- Requires kiro-cli to support phase callbacks or streaming output
- Or add markers in prompt that AI echoes back
- Parse output to detect phase transitions

---

## Non-Gaps (Specs Accurately Describe Implementation)

The following specs accurately describe the current implementation with no gaps:

- **prompt-composition.md** - create_prompt() function matches spec exactly (lines 143-159)
- **configuration-schema.md** - rooda-config.yml structure matches spec, all 9 procedures defined correctly
- **agents-md-format.md** - AGENTS.md follows specified structure with all required sections
- **specification-system.md** - specs/ directory structure matches spec (README, TEMPLATE, individual specs)

---

## Implementation Notes

**Prioritization Rationale:**
- P1: Critical gaps that cause user-facing failures or data integrity issues
- P2: High-impact gaps that significantly improve usability and reliability
- P3: Code quality improvements that reduce technical debt
- P4: Documentation gaps that help users but don't affect functionality
- P5: Future enhancements that add convenience but aren't essential

**Dependencies:**
- Task 3 (error handling) should be implemented before task 6 (timeout) for consistent failure handling
- Task 8 (remove duplication) should be done early to simplify subsequent changes
- Task 11-12 (documentation) can be done in parallel with implementation tasks

**Testing Strategy:**
- Manual testing for each task (no automated test suite exists per AGENTS.md)
- Verify with `shellcheck src/rooda.sh` after each change
- Test both invocation modes (procedure-based and explicit flags)
- Test error conditions (missing files, invalid config, etc.)
