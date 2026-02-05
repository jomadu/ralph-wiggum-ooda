# Task: Rename AI CLI Flags for Clarity

## Problem

The `--ai-cli` and `--ai-tool` flags are too similar in naming but serve different purposes:

- `--ai-cli <command>` - Direct command override (full command string with flags)
- `--ai-tool <preset>` - Named preset that resolves to a command

Users may confuse these flags because:
1. Both names suggest "specifying an AI CLI tool"
2. The distinction between "direct command" vs "preset name" isn't clear from the flag names
3. `--ai-cli` sounds like it should accept a tool name, not a full command
4. Presets aren't just about which tool - they can specify models, flags, and other configuration

## Proposed Solution

Rename both flags to use `ai-cmd` prefix for clarity:

**New naming:**
- `--ai-cmd <command>` - Direct command override (full command string)
- `--ai-cmd-preset <name>` - Named preset (resolves to command via config)

This makes the relationship clear: presets are named configurations for commands.

**Precedence remains the same:**
1. `--ai-cmd` flag (renamed from --ai-cli)
2. `--ai-cmd-preset` (renamed from --ai-tool)
3. `$ROODA_AI_CMD` environment variable (renamed from $ROODA_AI_CLI)
4. Default: `kiro-cli chat --no-interactive --trust-all-tools`

## Acceptance Criteria

- [ ] `--ai-cli` flag renamed to `--ai-cmd` in rooda.sh
- [ ] `--ai-tool` flag renamed to `--ai-cmd-preset` in rooda.sh
- [ ] `$ROODA_AI_CLI` environment variable renamed to `$ROODA_AI_CMD`
- [ ] Internal bash variables renamed (AI_CLI_COMMAND → AI_CMD_COMMAND, etc.)
- [ ] Short flags removed (avoid conflicts, long flags are clearer)
- [ ] Help text updated to reflect new flag names
- [ ] Error messages updated to use new flag names
- [ ] README.md updated with new flag names and environment variable
- [ ] specs/cli-interface.md updated
- [ ] specs/ai-cli-integration.md updated
- [ ] All examples in documentation updated
- [ ] No backward compatibility (clean break, beta software)

## Alternative Solutions

**Option 1: Keep --ai-cli, rename --ai-tool**
- `--ai-cli <command>` - Keep as-is
- `--ai-preset <name>` - Rename from --ai-tool

**Option 2: Use "command" terminology**
- `--ai-command <command>` - Direct command
- `--ai-command-preset <name>` - Named preset

**Option 3: Keep current names, improve documentation**
- Add clear examples showing the difference
- Update help text to emphasize "full command" vs "preset name"

## Recommendation

**Proposed solution** (`--ai-cmd` + `--ai-cmd-preset`) provides:
- Clear relationship: presets are named configurations for commands
- Shorter than "command" while still clear
- Consistent prefix makes the pairing obvious
- Accurately reflects that presets can specify tool + model + flags

## Impact

**Files to modify:**
- `src/rooda.sh` - Argument parsing, help text, error messages, variable names (AI_CLI_* → AI_CMD_*)
- `README.md` - All examples using --ai-cli, --ai-tool, and $ROODA_AI_CLI
- `specs/cli-interface.md` - Flag documentation
- `specs/ai-cli-integration.md` - Configuration precedence documentation
- `AGENTS.md` - If any examples use the flags or environment variable

**Breaking change:** Yes. Clean break, no backward compatibility (beta software).

**Renamed identifiers:**
- Flag: `--ai-cli` → `--ai-cmd`
- Flag: `--ai-tool` → `--ai-cmd-preset`
- Environment variable: `$ROODA_AI_CLI` → `$ROODA_AI_CMD`
- Internal variables: `AI_CLI_COMMAND` → `AI_CMD_COMMAND`, `AI_CLI_FLAG` → `AI_CMD_FLAG`, etc.

## Notes

**Scope:** This is purely a naming/UX change. No behavioral changes.

**Example usage after change:**
```bash
# Direct command override
./rooda.sh build --ai-cmd "claude-cli --model claude-3-opus --no-interactive"

# Using preset (could specify tool + model + flags)
./rooda.sh build --ai-cmd-preset fast

# Environment variable
export ROODA_AI_CMD="aider --yes"
./rooda.sh build

# Config defines preset
ai_tools:
  fast: "kiro-cli chat --no-interactive --trust-all-tools --model claude-3-5-haiku-20241022"
  thorough: "kiro-cli chat --no-interactive --trust-all-tools --model claude-3-7-sonnet-20250219"
```
