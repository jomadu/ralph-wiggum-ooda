# AI CLI Integration

## Job to be Done
Execute OODA loop prompts through a configurable AI CLI tool that can read files, modify code, run commands, and interact with the repository autonomously.

## Activities
1. Resolve AI CLI command from configuration (flag > config > default)
2. Pipe assembled OODA prompt to AI CLI via stdin
3. Pass flags to enable autonomous operation (no interactive prompts)
4. Trust all tool invocations without permission prompts
5. Allow AI to read/write files, execute commands, and commit changes
6. Capture AI CLI exit status for error handling

## Configuration

### ai_cli_command Field

The AI CLI command can be configured at the root level of `rooda-config.yml`:

```yaml
ai_cli_command: "kiro-cli chat --no-interactive --trust-all-tools"

procedures:
  bootstrap:
    # ... procedure config
```

**Field properties:**
- **Type:** String
- **Location:** Root level of rooda-config.yml (not per-procedure)
- **Purpose:** Specify which AI CLI tool to use for all procedures
- **Default:** `kiro-cli chat --no-interactive --trust-all-tools`
- **Validation:** Must be valid shell command

### --ai-cli Flag

Override the AI CLI command for a single execution:

```bash
./rooda.sh build --ai-cli "claude-cli --autonomous"
```

**Precedence rules:**
1. `--ai-cli` flag (highest priority)
2. `ai_cli_command` in rooda-config.yml
3. Default: `kiro-cli chat --no-interactive --trust-all-tools`

### Backward Compatibility

Existing installations continue to work without changes. The default AI CLI command remains `kiro-cli chat --no-interactive --trust-all-tools`, ensuring backward compatibility for users who don't specify configuration.

## Acceptance Criteria
- [x] Prompt piped to AI CLI via stdin
- [x] AI CLI command configurable via rooda-config.yml
- [x] AI CLI command overridable via --ai-cli flag
- [x] Precedence: flag > config > default
- [x] Default remains kiro-cli for backward compatibility
- [x] --no-interactive flag (or equivalent) disables interactive prompts
- [x] --trust-all-tools flag (or equivalent) bypasses permission prompts
- [x] AI can read files from repository
- [x] AI can write/modify files in repository
- [x] AI can execute bash commands
- [x] AI can commit changes to git
- [x] Script continues to next iteration regardless of AI CLI exit status

## Data Structures

### AI CLI Command Resolution
```bash
# Resolved from: --ai-cli flag > ai_cli_command config > default
AI_CLI_COMMAND="${AI_CLI_FLAG:-${AI_CLI_CONFIG:-kiro-cli chat --no-interactive --trust-all-tools}}"
```

### AI CLI Invocation
```bash
create_prompt | $AI_CLI_COMMAND
```

**Components:**
- `create_prompt` - Function that assembles OODA prompt from four phase files
- `$AI_CLI_COMMAND` - Resolved AI CLI command (configurable)
- Default: `kiro-cli chat --no-interactive --trust-all-tools`

**Common AI CLI tools:**
- `kiro-cli chat --no-interactive --trust-all-tools` (default)
- `claude-cli --autonomous --trust-tools`
- `aider --yes --auto-commits`
- Custom wrapper scripts

### Prompt Format
```markdown
# OODA Loop Iteration

## OBSERVE
[Content from observe phase file]

## ORIENT
[Content from orient phase file]

## DECIDE
[Content from decide phase file]

## ACT
[Content from act phase file]
```

## Algorithm

1. Resolve AI CLI command from configuration
   - Check for --ai-cli flag (highest priority)
   - Check for ai_cli_command in rooda-config.yml
   - Fall back to default: `kiro-cli chat --no-interactive --trust-all-tools`
2. Assemble OODA prompt using `create_prompt` function
3. Pipe prompt to AI CLI via stdin
4. AI CLI reads prompt and executes OODA phases
5. AI reads files, analyzes situation, makes decisions
6. AI executes actions (modify files, run commands, commit changes)
7. AI CLI exits (status ignored by script)
8. Script continues to git push and next iteration

**Pseudocode:**
```bash
# Resolve AI CLI command
if [ -n "$AI_CLI_FLAG" ]; then
    AI_CLI_COMMAND="$AI_CLI_FLAG"
elif [ -n "$AI_CLI_CONFIG" ]; then
    AI_CLI_COMMAND="$AI_CLI_CONFIG"
else
    AI_CLI_COMMAND="kiro-cli chat --no-interactive --trust-all-tools"
fi

create_prompt() {
    cat <<EOF
# OODA Loop Iteration

## OBSERVE
$(cat "$OBSERVE")

## ORIENT
$(cat "$ORIENT")

## DECIDE
$(cat "$DECIDE")

## ACT
$(cat "$ACT")
EOF
}

# Execute AI CLI
create_prompt | $AI_CLI_COMMAND
# Exit status not checked - script continues regardless
```

## Edge Cases

| Condition | Expected Behavior |
|-----------|-------------------|
| AI CLI not installed | Command fails, script exits with error |
| AI CLI exits with error | Script continues to git push (no error handling) |
| AI refuses to execute action | Iteration completes, next iteration may retry |
| AI modifies unexpected files | Changes committed and pushed (no validation) |
| AI executes dangerous command | Command runs (sandboxing required for safety) |
| Prompt exceeds token limit | AI CLI may truncate or fail (no size validation) |
| Network failure during AI call | AI CLI fails, script continues (no retry logic) |
| Invalid AI CLI command in config | Command fails at runtime, script exits |
| --ai-cli flag with invalid command | Command fails at runtime, script exits |
| AI CLI doesn't support stdin | Script fails, no fallback mechanism |

## Dependencies

- AI CLI tool (configurable, defaults to kiro-cli)
- AI CLI must support:
  - Reading prompts from stdin
  - Non-interactive operation mode
  - Tool invocation without permission prompts
  - File read/write capabilities
  - Command execution capabilities

## Implementation Mapping

**Source files:**
- `src/rooda.sh` - Lines 143-159 implement `create_prompt` function
- `src/rooda.sh` - Line 169 implements AI CLI invocation

**Related specs:**
- `component-authoring.md` - Defines how OODA phases are assembled
- `iteration-loop.md` - Defines loop execution behavior
- `cli-interface.md` - Defines command-line argument parsing

## Examples

### Example 1: Default AI CLI (kiro-cli)

**Input:**
```bash
create_prompt | kiro-cli chat --no-interactive --trust-all-tools
```

**Expected Output:**
```
[AI reads files, analyzes, makes decisions, executes actions]
[AI commits changes]
[AI CLI exits with status 0]
```

**Verification:**
- Files modified by AI exist on disk
- Git commits created by AI
- Script continues to next iteration

### Example 2: Custom AI CLI via Config

**Config (rooda-config.yml):**
```yaml
ai_cli_command: "claude-cli --autonomous --trust-tools"

procedures:
  build:
    # ... procedure config
```

**Input:**
```bash
./rooda.sh build
```

**Expected Output:**
```
[Prompt piped to claude-cli]
[AI executes OODA loop]
```

**Verification:**
- claude-cli invoked instead of kiro-cli
- Iteration completes successfully

### Example 3: Override via --ai-cli Flag

**Input:**
```bash
./rooda.sh build --ai-cli "aider --yes --auto-commits"
```

**Expected Output:**
```
[Prompt piped to aider]
[AI executes OODA loop]
```

**Verification:**
- aider invoked (flag overrides config and default)
- Iteration completes successfully

### Example 4: AI CLI Not Installed

**Input:**
```bash
./rooda.sh build --ai-cli "nonexistent-cli"
```

**Expected Output:**
```
bash: nonexistent-cli: command not found
```

**Verification:**
- Script exits with error
- No iteration executed

### Example 5: AI Refuses Action

**Input:**
```bash
create_prompt | kiro-cli chat --no-interactive --trust-all-tools
```

**Expected Output:**
```
[AI analyzes situation]
[AI responds: "I cannot complete this action because..."]
[AI CLI exits]
```

**Verification:**
- No files modified
- No commits created
- Script continues to next iteration (may retry)

## Notes

**Design Rationale:**

The AI CLI integration is designed for autonomous operation with minimal human intervention. Configuration support enables users to choose their preferred AI CLI tool while maintaining backward compatibility with kiro-cli as the default.

**Configuration Flexibility:**

The three-tier precedence system (flag > config > default) provides flexibility:
1. **--ai-cli flag** - Quick experimentation or one-off overrides
2. **ai_cli_command config** - Project-specific AI CLI preference
3. **Default (kiro-cli)** - Backward compatibility for existing users

**Security Implications:**

The AI CLI must support autonomous operation (no interactive prompts, no permission prompts for tool invocations). This is inherently risky and requires sandboxed execution environments (Docker, Fly Sprites, E2B) to limit blast radius.

**Error Handling:**

The script does not check AI CLI exit status. This design choice allows the loop to continue even if the AI encounters errors or refuses actions. The assumption is that subsequent iterations can self-correct through empirical feedback.

**AI CLI Requirements:**

Any AI CLI tool can be used if it supports:
- Reading prompts from stdin
- Non-interactive operation mode
- Tool invocation without permission prompts
- File read/write capabilities
- Command execution capabilities

**Token Limits:**

The script does not validate prompt size before piping to AI CLI. Large OODA phase files or extensive file contents could exceed token limits. The AI CLI is responsible for handling this (truncation, error, or chunking).

## Known Issues

**No error handling:** Script continues to git push even if AI CLI fails. This could result in pushing incomplete or invalid changes.

**No retry logic:** If AI CLI fails due to transient issues (network, rate limits), the iteration is lost. No automatic retry mechanism exists.

**No validation:** Script does not validate that AI CLI command is valid or that the tool supports required capabilities before invocation. Incompatible tools will fail at runtime.

**No timeout:** If AI CLI hangs, the script waits indefinitely. No timeout mechanism exists.

## Areas for Improvement

**Dependency checking:** Add validation that configured AI CLI is installed and accessible before starting loop.

**Capability detection:** Detect if AI CLI supports required features (stdin, non-interactive mode, tool invocation) and provide clear error messages if not.

**Error handling:** Check AI CLI exit status and handle failures gracefully (retry, skip push, abort loop).

**Timeout mechanism:** Add timeout for AI CLI invocation to prevent indefinite hangs.

**Prompt size validation:** Check assembled prompt size before piping to AI CLI, warn if approaching token limits.

**AI CLI profiles:** Support multiple AI CLI configurations (e.g., "fast" vs "thorough" models) selectable per procedure or via flag.

**Version requirements:** Document minimum version requirements for supported AI CLI tools.
