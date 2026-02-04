# Draft Plan: Spec to Implementation Gap Analysis

## Priority 1: Critical Missing Features

### 1. AI CLI Configuration Support (ai_cli_command)
**Gap:** Specs define `ai_cli_command` field in rooda-config.yml with three-tier precedence (--ai-cli flag > config > default), but implementation hardcodes kiro-cli

**Evidence:**
- `ai-cli-integration.md` and `configuration-schema.md` specify ai_cli_command field
- rooda-config.yml has no ai_cli_command field
- rooda.sh has no --ai-cli flag parsing
- rooda.sh line 436 hardcodes: `kiro-cli chat --no-interactive --trust-all-tools`

**Tasks:**
- Add ai_cli_command field to rooda-config.yml root level (optional, defaults to kiro-cli)
- Add --ai-cli flag to argument parser (lines 245-285)
- Add AI_CLI_COMMAND variable initialization
- Query config for ai_cli_command when procedure specified (lines 286-340)
- Implement three-tier precedence: flag > config > default
- Replace hardcoded kiro-cli invocation with $AI_CLI_COMMAND variable (line 436)
- Update show_help() to document --ai-cli flag

**Acceptance:**
- rooda-config.yml has ai_cli_command field at root level
- ./rooda.sh build --ai-cli "custom-cli" uses custom-cli instead of kiro-cli
- Config ai_cli_command overrides default but not --ai-cli flag
- Default remains kiro-cli for backward compatibility

---

### 2. Dependency Philosophy Alignment
**Gap:** Specs say only yq is required (kiro-cli configurable, bd project-specific), but implementation requires all three

**Evidence:**
- `external-dependencies.md` states "Minimal Required, Maximum Flexibility" philosophy
- rooda.sh lines 60-88 exit if kiro-cli or bd not installed
- Contradicts spec: "kiro-cli (default, configurable)" and "bd (project-specific, optional)"

**Tasks:**
- Make kiro-cli check conditional on ai_cli_command not being set to alternative
- Make bd check conditional or remove entirely (work tracking is project-specific)
- Add comments explaining dependency philosophy
- Update error messages to reflect optional nature

**Acceptance:**
- Script runs with only yq installed if ai_cli_command points to installed alternative
- Script provides helpful error only when trying to use missing tool
- Dependency checks align with "only yq required" philosophy

**Dependencies:** Implement after task #1 (needs ai_cli_command support)

---

## Priority 2: Documentation Sync (Specs Outdated)

### 3. Update cli-interface.md Spec
**Gap:** Spec lists help/version/short-flags/verbose as "Areas for Improvement" but they're implemented

**Evidence:**
- rooda.sh lines 233-242 handle --help, --version
- rooda.sh lines 255-277 support short flags (-c, -o, -r, -d, -a, -m, -h)
- rooda.sh lines 278-284 handle --verbose, --quiet
- rooda.sh lines 426-432 implement verbose mode

**Tasks:**
- Move implemented features from "Areas for Improvement" to acceptance criteria
- Mark acceptance criteria as [x] completed
- Add examples for --help, --version, --verbose, --quiet flags
- Update "Known Issues" section to remove implemented items

**Acceptance:**
- cli-interface.md accurately reflects current implementation
- All implemented features documented with examples
- No contradictions between spec and code

---

### 4. Update ai-cli-integration.md Examples
**Gap:** Spec shows examples with ai_cli_command configuration that doesn't exist yet

**Evidence:**
- Examples show ai_cli_command in config (not implemented)
- Examples show --ai-cli flag usage (not implemented)
- Examples show precedence resolution (not implemented)

**Tasks:**
- Add note to examples: "Requires task #1 implementation"
- Or update examples to show current hardcoded behavior
- Update after task #1 is complete

**Acceptance:**
- Examples match implementation reality
- No misleading examples showing unimplemented features

**Dependencies:** Update after task #1 (AI CLI configuration)

---

### 5. Update external-dependencies.md
**Gap:** Spec describes dependency philosophy that implementation contradicts

**Evidence:**
- Spec: "kiro-cli (default, configurable)" but implementation requires it
- Spec: "bd (project-specific, optional)" but implementation requires it
- Philosophy section contradicts actual behavior

**Tasks:**
- Update dependency tables to reflect current requirements
- Or note that philosophy will be implemented in task #2
- Ensure consistency between spec and implementation

**Acceptance:**
- Spec accurately describes which dependencies are required vs optional
- Philosophy section matches implementation behavior

**Dependencies:** Update after task #2 (dependency philosophy)

---

## Priority 3: Spec Clarifications (Already Implemented)

### 6. Document Implemented Features
**Gap:** Specs note several features as "missing" that are actually implemented

**Already Implemented:**
- Help flag support (--help, -h)
- Version flag support (--version)
- Short flag alternatives (-c, -o, -r, -d, -a, -m, -h)
- Verbose/quiet modes (--verbose, --quiet)
- Config validation function (validate_config)
- Git push error handling

**Tasks:**
- Update specs to reflect these implementations
- Remove from "Known Issues" or "Areas for Improvement"
- Add to acceptance criteria as [x] completed

**Acceptance:**
- Specs accurately document all implemented features
- No features listed as missing that actually exist

---

## Summary

**Implementation tasks (2):**
1. AI CLI configuration support (ai_cli_command + --ai-cli flag)
2. Dependency philosophy alignment (make kiro-cli/bd optional)

**Documentation tasks (3):**
3. Update cli-interface.md (implemented features)
4. Update ai-cli-integration.md (after task #1)
5. Update external-dependencies.md (after task #2)

**Key insight:** Implementation is ahead of specs in many areas (help, version, verbose, validation), but behind in AI CLI configuration. Priority is to implement missing features, then sync all documentation.
