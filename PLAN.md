# Draft Plan: Make rooda.sh Agnostic of External Tools

## Priority Tasks

1. **Update external-dependencies.md - Change dependency model**
   - Change kiro-cli from "required" to "configurable with default"
   - Change bd from "required" to "optional (project-specific)"
   - Document that only yq is truly required
   - Add section explaining dependency philosophy: minimal required, everything else configurable
   - Update acceptance criteria to reflect optional dependencies
   - Acceptance: Spec clearly states kiro-cli is default but configurable, bd is optional

2. **Update ai-cli-integration.md - Generic AI CLI integration**
   - Change "Job to be Done" from kiro-cli specific to generic AI CLI tool
   - Add "Configuration" section documenting `ai_cli_command` config field and `--ai-cli` flag
   - Update "Data Structures" section to show configurable AI CLI command instead of hardcoded kiro-cli
   - Add examples for multiple AI CLI tools (kiro-cli, claude-cli, custom tools)
   - Update "Algorithm" section to show configuration resolution (flag > config > default)
   - Add "Backward Compatibility" subsection explaining kiro-cli default
   - Update acceptance criteria to include configuration support
   - Acceptance: Spec documents generic AI CLI integration with configuration mechanism

3. **Update configuration-schema.md - Add ai_cli_command field**
   - Add `ai_cli_command` field to root-level configuration structure
   - Document field type (string), purpose (specify AI CLI command), and default (kiro-cli chat --no-interactive --trust-all-tools)
   - Add example showing custom AI CLI configuration
   - Update "Algorithm" section to show how ai_cli_command is queried
   - Add validation rules (optional field, string type, must be valid shell command)
   - Acceptance: Spec documents ai_cli_command field with examples and validation

4. **Update cli-interface.md - Add --ai-cli flag**
   - Add `--ai-cli <command>` flag to flags table
   - Document flag purpose: override AI CLI command for this execution
   - Document precedence: CLI flag overrides config file, config file overrides default
   - Add example usage showing custom AI CLI via flag
   - Update "Algorithm" section to show flag parsing and precedence resolution
   - Add to acceptance criteria: --ai-cli flag supported and takes precedence
   - Acceptance: Spec documents --ai-cli flag with precedence rules and examples

## Dependencies

- Task 1 should complete first (establishes dependency philosophy)
- Tasks 2-4 can be done in parallel (independent spec updates)
- All tasks are spec updates only (no implementation changes)
