# Draft Plan: Spec to Implementation Gap Analysis

## High Priority Gaps (Missing Core Features)

### 1. AI CLI Integration - Missing ai_tools Configuration Section
**Gap:** Specs define ai_tools section in rooda-config.yml for custom presets, but implementation doesn't have this section in the actual config file.

**Evidence:**
- `specs/configuration-schema.md` documents ai_tools section with examples (fast, thorough, custom presets)
- `specs/ai-cli-integration.md` extensively documents custom presets in config
- `src/rooda-config.yml` has commented-out example but no actual ai_tools section
- `src/rooda.sh` lines 102-143 implement resolve_ai_tool_preset() that queries .ai_tools.$preset

**Impact:** Users cannot define custom AI tool presets in config as documented. The --ai-cmd-preset flag only works with hardcoded presets (kiro-cli, claude, aider).

**Acceptance Criteria:**
- Add ai_tools section to src/rooda-config.yml with at least one example preset
- Verify --ai-cmd-preset resolves custom presets from config
- Update AGENTS.md if needed to document the feature

**Dependencies:** None

---

### 2. External Dependencies - Missing yq Version Validation
**Gap:** Specs require yq v4.0.0+ validation at startup, but implementation doesn't validate version.

**Evidence:**
- `specs/external-dependencies.md` specifies yq minimum version 4.0.0 with validation algorithm
- `specs/external-dependencies.md` includes Example 3 showing version check error message
- `src/rooda.sh` checks for yq existence but doesn't validate version
- AGENTS.md notes "yq version validation (requires v4.0.0+)" but implementation missing

**Impact:** Script may fail with cryptic errors if user has yq v3 installed (incompatible syntax).

**Acceptance Criteria:**
- Add yq version check at startup (lines 15-19 area)
- Extract version from `yq --version` output
- Compare against minimum 4.0.0
- Exit with clear error message if version too old
- Include upgrade instructions in error message

**Dependencies:** None

---

### 3. CLI Interface - Missing --list-procedures Implementation
**Gap:** Specs document --list-procedures flag, implementation has function but may not be fully integrated.

**Evidence:**
- `specs/cli-interface.md` acceptance criteria includes --list-procedures flag
- `src/rooda.sh` lines 70-100 implement list_procedures() function
- Help text shows --list-procedures option
- Need to verify it works correctly with all procedures

**Impact:** Users cannot easily discover available procedures without reading config file.

**Acceptance Criteria:**
- Verify --list-procedures works with current config
- Test output format matches expectations
- Ensure display and summary fields are shown correctly

**Dependencies:** None

---

## Medium Priority Gaps (Documentation & Validation)

### 4. Documentation - Code Examples Not Verified
**Gap:** Quality criteria require all code examples in docs/ to be verified working, but no verification process exists.

**Evidence:**
- AGENTS.md quality criteria: "All code examples in docs/ are verified working (PASS/FAIL)"
- `specs/user-documentation.md` acceptance criteria: "All code examples in documentation are verified working"
- No test script or verification process for docs/ examples
- docs/beads.md, docs/ooda-loop.md, docs/ralph-loop.md contain code examples

**Impact:** Documentation examples may be outdated or incorrect, misleading users.

**Acceptance Criteria:**
- Create script to extract and verify code examples from docs/*.md
- Add to quality criteria verification process
- Document verification approach in AGENTS.md

**Dependencies:** None

---

### 5. Specification System - Missing Spec Index Regeneration
**Gap:** Specs describe automatic spec index regeneration in specs/README.md, but implementation doesn't maintain this.

**Evidence:**
- `specs/component-authoring.md` mentions "A3.6: Regenerate Spec Index (If Specs Modified)"
- `src/prompts/act_build.md` includes step A3.6 for regenerating spec index
- specs/README.md exists but may not be automatically maintained
- No clear algorithm for what should be in the index

**Impact:** Spec index may become stale as specs are added/modified/deleted.

**Acceptance Criteria:**
- Define spec index structure requirements
- Implement automatic regeneration when specs change
- Verify index stays current with spec changes

**Dependencies:** None

---

### 6. Iteration Loop - Missing Error Handling for AI CLI Failures
**Gap:** Specs note "No kiro-cli error handling" as known issue, but this is a significant gap.

**Evidence:**
- `specs/iteration-loop.md` Known Issues: "No kiro-cli error handling: If kiro-cli exits with non-zero status, the loop continues anyway"
- `specs/ai-cli-integration.md` Known Issues: "No error handling: Script continues to git push even if AI CLI fails"
- Implementation doesn't check AI CLI exit status

**Impact:** Loop continues after AI CLI failures, potentially pushing incomplete or invalid changes.

**Acceptance Criteria:**
- Check AI CLI exit status after each iteration
- Implement failure handling (retry, skip push, abort loop)
- Add --max-failures threshold option
- Document behavior in AGENTS.md

**Dependencies:** None

---

## Low Priority Gaps (Nice-to-Have Features)

### 7. CLI Interface - Missing Dry-Run Mode
**Gap:** Specs suggest --dry-run mode as area for improvement, not implemented.

**Evidence:**
- `specs/iteration-loop.md` Areas for Improvement: "Dry-run mode: Support --dry-run flag to show what would execute without actually running the AI CLI"
- No --dry-run flag in implementation
- Would be useful for testing procedures without executing

**Impact:** Users cannot preview what a procedure would do without actually running it.

**Acceptance Criteria:**
- Add --dry-run flag to CLI
- Show assembled prompt without executing AI CLI
- Display what files would be read/modified
- Skip git push in dry-run mode

**Dependencies:** None

---

### 8. External Dependencies - Missing Automated Dependency Checker
**Gap:** Specs describe dependency checking algorithm, but implementation only checks yq.

**Evidence:**
- `specs/external-dependencies.md` documents comprehensive dependency checking algorithm
- Implementation only checks for yq at startup
- No checks for kiro-cli, bd, shellcheck, git
- No version validation for any tools

**Impact:** Users may encounter runtime errors due to missing dependencies.

**Acceptance Criteria:**
- Implement check_dependencies() function per spec algorithm
- Check all required dependencies at startup
- Provide installation instructions for missing tools
- Validate versions where specified

**Dependencies:** Task #2 (yq version validation)

---

### 9. Configuration Schema - Missing Config Validation
**Gap:** Specs note "Config validation: Script doesn't validate config file structure before querying with yq" as known issue.

**Evidence:**
- `specs/cli-interface.md` Known Issues: "Config validation: Script doesn't validate config file structure before querying with yq"
- `src/rooda.sh` lines 195-298 implement validate_config() function
- Need to verify this function is comprehensive

**Impact:** Invalid YAML or missing required fields cause cryptic yq errors.

**Acceptance Criteria:**
- Verify validate_config() checks all required fields
- Add validation for ai_tools section structure
- Provide clear error messages for common config mistakes
- Test with intentionally malformed configs

**Dependencies:** Task #1 (ai_tools section)

---

### 10. Iteration Loop - Missing Progress Indicators
**Gap:** Specs suggest progress indicators as area for improvement.

**Evidence:**
- `specs/iteration-loop.md` Areas for Improvement: "Progress indicators: Show which OODA phase is executing during long-running iterations"
- `specs/iteration-loop.md` Areas for Improvement: "Iteration timing: Display elapsed time per iteration"
- No progress indicators in current implementation

**Impact:** Users have no visibility into what's happening during long-running iterations.

**Acceptance Criteria:**
- Show current OODA phase during execution
- Display elapsed time per iteration
- Add progress bar or spinner for long operations
- Make optional via --verbose flag

**Dependencies:** None

---

## Completeness Assessment

**Specifications Coverage:** ~85%
- Core functionality (OODA loop, procedures, CLI, config) fully implemented
- AI CLI integration mostly implemented (missing ai_tools config section)
- Dependency checking partially implemented (yq only)
- Error handling minimal
- Documentation validation missing

**Implementation Coverage:** ~90%
- All 9 procedures defined and functional
- All 25 prompt files valid and structured correctly
- CLI interface complete with all documented flags
- Config schema supports all required fields
- Quality validation scripts (validate-prompts.sh, audit-links.sh) working

**Critical Gaps:** 3 (tasks #1, #2, #3)
**Important Gaps:** 6 (tasks #4-#9)
**Enhancement Gaps:** 1 (task #10)

**Highest Priority:** Task #1 (ai_tools config section) - documented feature that doesn't work
**Blocking Issues:** None - all gaps are independent and can be addressed in parallel
