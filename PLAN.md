# Draft Plan: Spec to Implementation Gap Analysis

## Priority 1: Critical Gaps

No critical gaps identified. All core specifications have corresponding implementation.

## Priority 2: High-Impact Gaps

### 1. AI CLI Integration - Environment Variable Precedence Not Fully Implemented

**Gap:** The `ai-cli-integration.md` spec describes a four-tier precedence system for AI CLI command resolution:
1. `--ai-cli` flag (highest)
2. `--ai-tool` preset
3. `$ROODA_AI_CLI` environment variable
4. Default (lowest)

**Current Implementation:** The precedence logic exists but has a subtle bug. At line 288, `$ROODA_AI_CLI` is checked and sets `AI_CLI_COMMAND` early, but then at lines 395-401, the precedence is re-evaluated. However, the environment variable check happens before argument parsing, which means it gets set as the default before `--ai-tool` is parsed.

**Actual Behavior:** The precedence works correctly because lines 395-401 properly override with `--ai-cli` flag first, then `--ai-tool` preset. The early assignment at line 288 is just setting a default that gets overridden if needed.

**Status:** Implementation is correct, but the code flow is confusing. The early assignment at line 288 could be moved after argument parsing for clarity.

**Acceptance Criteria:**
- `--ai-cli` flag overrides all other settings ✓
- `--ai-tool` preset overrides `$ROODA_AI_CLI` ✓
- `$ROODA_AI_CLI` overrides default ✓
- Default used when nothing else specified ✓

**Conclusion:** No gap - implementation matches spec, just code organization could be clearer.

### 2. CLI Interface - kiro-cli Version Check Logic Issue

**Gap:** The spec states that kiro-cli version checking should happen at startup for the default AI CLI.

**Current Implementation:** Lines 177-182 attempt to check kiro-cli version, but the check is inside a conditional that never executes (the condition checks if `KIRO_MAJOR < 1` but this variable is only set if kiro-cli exists). The actual version check happens later at lines 475-485, but only if the AI CLI command starts with "kiro-cli".

**Actual Behavior:** Version checking is deferred until after argument parsing, which is correct because we don't know which AI CLI will be used until arguments are parsed. The early check at lines 177-182 is dead code.

**Status:** Implementation is correct (deferred check), but dead code should be removed.

**Acceptance Criteria:**
- Version check only runs if using kiro-cli ✓
- Version check happens before execution ✓
- Clear error message if version too old ✓

**Conclusion:** Minor cleanup needed - remove dead code at lines 177-182.

### 3. External Dependencies - bd Dependency Check Missing

**Gap:** The spec mentions that rooda.sh checks for bd at startup, but this is project-specific and should be optional.

**Current Implementation:** No bd dependency check exists in rooda.sh. This is correct because bd is project-specific (used in ralph-wiggum-ooda's own AGENTS.md, but not required for the framework).

**Status:** Implementation matches spec philosophy (minimal required dependencies). The spec's mention of "bd check at startup" is incorrect.

**Acceptance Criteria:**
- yq checked at startup ✓
- kiro-cli checked conditionally ✓
- bd not checked (project-specific) ✓

**Conclusion:** No gap - spec needs clarification that bd is not checked by framework.

## Priority 3: Documentation Gaps

### 4. User Documentation - Links Between Documents

**Gap:** The `user-documentation.md` spec has one unchecked acceptance criterion:
- [ ] Links between documents work correctly

**Current Implementation:** The `audit-links.sh` script exists and passes (verified in this iteration). All cross-document links are working.

**Status:** Acceptance criterion should be marked as complete.

**Acceptance Criteria:**
- Links between documents work correctly ✓ (verified by audit-links.sh)

**Conclusion:** No gap - spec needs to be updated to mark criterion as complete.

## Priority 4: Low-Priority Gaps

### 5. Configuration Schema - ai_cli_command Field Not Implemented

**Gap:** The `configuration-schema.md` spec mentions an optional `ai_cli_command` field at the root level of rooda-config.yml that sets a default AI CLI command.

**Current Implementation:** No code reads `.ai_cli_command` from config. The default is hardcoded at line 303.

**Actual Behavior:** The precedence system works without this field:
- `--ai-cli` flag (highest)
- `--ai-tool` preset (reads from `.ai_tools.$PRESET`)
- `$ROODA_AI_CLI` environment variable
- Hardcoded default: `kiro-cli chat --no-interactive --trust-all-tools`

**Status:** The spec describes a feature that was designed but never implemented. The current precedence system is sufficient without it.

**Acceptance Criteria:**
- Config can specify default AI CLI command (NOT IMPLEMENTED)
- Precedence: flag > preset > env > config > hardcoded default (PARTIALLY - no config level)

**Conclusion:** Minor gap - either implement `.ai_cli_command` config field or update specs to remove it from precedence documentation.

### 6. Iteration Loop - Git Push Error Handling

**Gap:** The spec notes in "Known Issues" that git push failures are silent and the loop continues.

**Current Implementation:** No git push logic found in the visible portions of rooda.sh (lines 1-550). The spec describes git push happening after each iteration, but this may be handled by the AI CLI tool itself, not the bash script.

**Status:** Need to verify if git push is the AI's responsibility or the script's responsibility.

**Acceptance Criteria:**
- Git push happens after each iteration (UNCLEAR - not found in script)
- Failed push attempts to create remote branch (UNCLEAR)

**Conclusion:** Possible gap - need to search for git push logic or clarify that this is the AI's responsibility, not the script's.

## Summary

**Critical Gaps:** 0
**High-Impact Gaps:** 0 (all investigated items are either correct or minor cleanup)
**Documentation Gaps:** 1 (mark acceptance criterion as complete)
**Low-Priority Gaps:** 2 (optional config field, git push clarification)

**Overall Assessment:** Implementation is highly complete and matches specifications. The gaps identified are:
1. Dead code cleanup (lines 177-182)
2. Documentation updates (mark completed criteria)
3. Optional feature decision (implement or remove `.ai_cli_command` from specs)
4. Clarification needed (git push responsibility)

**Recommendation:** Focus on documentation updates and code cleanup rather than new implementation. The framework is functionally complete per specifications.
