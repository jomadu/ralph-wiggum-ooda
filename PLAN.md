# Draft Plan: Spec to Implementation Gap Analysis

## Priority 1: Implement --ai-cli Flag Support

**Gap:** `cli-interface.md` and `ai-cli-integration.md` specify --ai-cli flag to override AI CLI command, but implementation has no --ai-cli argument parsing or AI_CLI_COMMAND variable.

**Current State:**
- rooda.sh hardcodes `kiro-cli chat --no-interactive --trust-all-tools` at line 443
- No --ai-cli flag in argument parser (lines 218-289)
- No AI_CLI_COMMAND variable initialization or resolution

**Required Changes:**
- Add AI_CLI_COMMAND variable initialization
- Add --ai-cli flag parsing in argument parser
- Query ai_cli_command from config if procedure specified
- Implement precedence: --ai-cli flag > config ai_cli_command > default
- Replace hardcoded kiro-cli invocation with $AI_CLI_COMMAND variable

**Acceptance Criteria:**
- `./rooda.sh build --ai-cli "custom-cli"` uses custom-cli instead of kiro-cli
- Config ai_cli_command field overrides default when no flag specified
- Default remains kiro-cli for backward compatibility
- Precedence order enforced correctly

**Dependencies:** None

---

## Priority 2: Implement ai_cli_command Config Field Support

**Gap:** `configuration-schema.md` and `ai-cli-integration.md` specify ai_cli_command as root-level config field, but implementation doesn't query or use this field.

**Current State:**
- rooda-config.yml has no ai_cli_command field
- rooda.sh doesn't query .ai_cli_command from config
- No validation of ai_cli_command field type

**Required Changes:**
- Add ai_cli_command field to rooda-config.yml (optional, defaults to kiro-cli)
- Query .ai_cli_command from config when procedure specified
- Validate field is string type if present
- Use config value when --ai-cli flag not specified

**Acceptance Criteria:**
- Config with ai_cli_command: "claude-cli" uses claude-cli for all procedures
- Config without ai_cli_command defaults to kiro-cli
- Invalid ai_cli_command type produces clear error
- validate_config function checks ai_cli_command structure

**Dependencies:** Priority 1 (--ai-cli flag support)

---

## Priority 3: Remove Hardcoded Dependency Checks

**Gap:** `external-dependencies.md` specifies only yq is required, kiro-cli and bd are configurable/optional, but implementation hardcodes checks for all three at startup.

**Current State:**
- Lines 72-86: Hardcoded kiro-cli dependency check with exit on failure
- Lines 88-97: Hardcoded bd dependency check with exit on failure
- Lines 99-113: Version validation for kiro-cli and bd

**Required Changes:**
- Remove kiro-cli dependency check (configurable via ai_cli_command)
- Remove bd dependency check (project-specific work tracking)
- Keep yq dependency check (required for config parsing)
- Document that AI CLI and work tracking tools are project-specific

**Acceptance Criteria:**
- Script runs without kiro-cli installed if using different AI CLI
- Script runs without bd installed if using different work tracking
- yq check remains and exits with clear error if missing
- AGENTS.md updated to clarify dependency philosophy

**Dependencies:** Priority 2 (ai_cli_command config support)

---

## Priority 4: Add Short Flag Support

**Gap:** `cli-interface.md` examples show short flags (-o, -r, -d, -a, -m, -c, -h) but implementation only supports long flags (--observe, --orient, etc).

**Current State:**
- Argument parser has cases for long flags only
- No short flag alternatives defined
- Help text doesn't document short flags

**Required Changes:**
- Add short flag cases to argument parser: -o, -r, -d, -a, -m, -c, -h
- Update show_help to document short flags
- Test all short flags work identically to long flags

**Acceptance Criteria:**
- `-o file.md` works identically to `--observe file.md`
- `-m 5` works identically to `--max-iterations 5`
- `-h` works identically to `--help`
- Help text shows both short and long forms

**Dependencies:** None

---

## Priority 5: Implement Substep Numbering in act_build.md

**Gap:** `component-authoring.md` documents substep numbering (A3, A3.5, A3.6, A4) as valid pattern, but act_build.md may not follow this structure.

**Current State:** Need to verify act_build.md structure

**Required Changes:**
- Review act_build.md for substep usage
- Ensure substeps follow documented pattern (A3, A3.5, A3.6, A4)
- Verify conditional steps clearly marked
- Confirm backpressure steps properly labeled

**Acceptance Criteria:**
- act_build.md uses substep numbering where appropriate
- Conditional steps have "If X Modified" markers
- Critical warnings use **bold** emphasis
- Backpressure concept in step names

**Dependencies:** None

---

## Priority 6: Add Cross-Document Link Validation

**Gap:** `user-documentation.md` acceptance criteria includes "All cross-document links work correctly" but no validation mechanism exists.

**Current State:**
- No automated link checking
- Links may break when files renamed
- Quality criteria doesn't include link validation

**Required Changes:**
- Add script to validate markdown links in docs/ and specs/
- Check internal links (relative paths) resolve correctly
- Check external links return 200 status
- Add to quality criteria: "All cross-document links work correctly (PASS/FAIL)"

**Acceptance Criteria:**
- Script validates all markdown links in repository
- Broken internal links reported with file and line number
- Broken external links reported (with timeout handling)
- Can be run as part of quality assessment procedures

**Dependencies:** None

---

## Priority 7: Document Verbose/Quiet Modes

**Gap:** Implementation has --verbose and --quiet flags (lines 283-289) but specs don't document these features.

**Current State:**
- rooda.sh implements VERBOSE variable (0=default, 1=verbose, -1=quiet)
- --verbose shows full prompt before execution
- --quiet suppresses progress output
- No documentation in cli-interface.md or user-documentation.md

**Required Changes:**
- Add --verbose and --quiet to cli-interface.md data structures
- Document behavior in cli-interface.md examples
- Add to README.md usage examples
- Update show_help to include these flags

**Acceptance Criteria:**
- cli-interface.md documents --verbose and --quiet flags
- Examples show verbose mode displaying full prompt
- Examples show quiet mode suppressing progress
- Help text includes these options

**Dependencies:** None

---

## Priority 8: Implement Help Text Generation from Config

**Gap:** `configuration-schema.md` specifies display, summary, description fields for procedures but notes "reserved for future help text generation" - not yet implemented.

**Current State:**
- Config has display/summary/description fields
- show_help function doesn't use these fields
- No procedure listing command

**Required Changes:**
- Add --list-procedures flag to show available procedures
- Query config for all procedure names
- Display procedure name, display field, and summary
- Consider --help <procedure> for detailed description

**Acceptance Criteria:**
- `./rooda.sh --list-procedures` shows all available procedures
- Output includes procedure name and summary from config
- `./rooda.sh --help bootstrap` shows detailed description
- Graceful handling if display/summary/description missing

**Dependencies:** None

---

## Priority 9: Add Validation for Prompt File Structure

**Gap:** `component-authoring.md` documents prompt file structure (phase header, step codes, prose) but no validation ensures files follow this structure.

**Current State:**
- No automated validation of prompt files
- Malformed prompts may cause confusing AI behavior
- No linting for step code consistency

**Required Changes:**
- Add script to validate prompt file structure
- Check for phase header: `# [Phase]: [Purpose]`
- Check for step headers: `## [Code]: [Name]`
- Validate step codes match phase (O1-O15, R1-R22, D1-D15, A1-A9)
- Add to quality criteria for implementation

**Acceptance Criteria:**
- Script validates all prompt files in src/prompts/
- Reports missing phase headers
- Reports invalid step codes
- Reports step codes that don't match phase
- Can be run as part of quality assessment

**Dependencies:** None

---

## Priority 10: Document Version Flag

**Gap:** Implementation has --version flag (lines 237-240, 251-254) but cli-interface.md lists this as "Areas for Improvement" suggesting it's not implemented.

**Current State:**
- rooda.sh implements --version flag
- VERSION variable defined (need to verify)
- cli-interface.md incorrectly states "No --version flag"

**Required Changes:**
- Update cli-interface.md to document --version flag
- Verify VERSION variable is defined in rooda.sh
- Add example to cli-interface.md showing version output
- Update "Areas for Improvement" to remove this item

**Acceptance Criteria:**
- cli-interface.md documents --version flag
- Example shows expected output format
- VERSION variable properly defined in script
- "Areas for Improvement" section updated

**Dependencies:** None
