# Task: Make rooda.sh Agnostic of External Tools

## Story

As a user of ralph-wiggum-ooda, I want the framework to be agnostic of specific external tools so that I can use it with different AI CLI tools and work tracking systems without modifying the core script.

## Desired Changes

### 1. AI CLI Tool Agnosticism

**Current State:**
- rooda.sh has hardcoded dependency on `kiro-cli`
- Script checks for kiro-cli at startup and exits if not found
- Prompt is piped to `kiro-cli chat --no-interactive --trust-all-tools`

**Desired State:**
- rooda.sh should be agnostic of which AI CLI tool is used
- Users should be able to specify the AI CLI command via:
  - Configuration file (e.g., `ai_cli_command: "kiro-cli chat --no-interactive --trust-all-tools"`)
  - CLI flag (e.g., `--ai-cli "claude-cli --no-prompt"`)
- Script should not check for kiro-cli specifically at startup
- Default to kiro-cli for backward compatibility if not specified

**Benefits:**
- Users can use Claude CLI, OpenAI CLI, or any other AI tool
- Framework becomes more portable and flexible
- Removes hard dependency on AWS-specific tooling

**Examples:**
```bash
# Using Claude CLI
./rooda.sh build --ai-cli "claude-cli --no-prompt"

# Using custom AI tool via config
# rooda-config.yml:
ai_cli_command: "my-ai-tool --batch-mode"
```

### 2. Work Tracking System Agnosticism

**Current State:**
- rooda.sh has hardcoded dependency on `bd` (beads)
- Script checks for bd at startup and exits if not found
- Not every project uses beads for work tracking

**Desired State:**
- rooda.sh should not check for bd at startup
- Work tracking system is defined in AGENTS.md (already the case for usage)
- Dependency checking should be optional or removed entirely
- Let procedures fail gracefully if work tracking commands don't work

**Benefits:**
- Projects using GitHub Issues, Linear, Jira, or file-based tracking can use rooda
- Framework doesn't impose work tracking system choice
- Reduces installation friction

**Rationale:**
- AGENTS.md already defines the work tracking system per project
- Not all procedures require work tracking (bootstrap, some planning procedures)
- Dependency checking should happen when commands are actually used, not at startup

## Acceptance Criteria

- [ ] rooda.sh accepts `--ai-cli` flag to specify AI CLI command
- [ ] rooda-config.yml supports `ai_cli_command` field for default AI CLI
- [ ] rooda.sh defaults to kiro-cli if no AI CLI specified (backward compatibility)
- [ ] rooda.sh does not check for kiro-cli at startup
- [ ] rooda.sh does not check for bd at startup
- [ ] Documentation updated to explain AI CLI configuration
- [ ] Examples provided for common AI CLI tools (kiro-cli, claude-cli, etc.)
- [ ] AGENTS.md specification updated to clarify work tracking is project-specific

## Notes

**Minimal Dependency Philosophy:**
The only truly required dependency should be `yq` (for config parsing). Everything else should be configurable or optional based on what the project actually uses.

**Backward Compatibility:**
Existing installations should continue to work without changes. Default to kiro-cli if no AI CLI specified.

**Error Messages:**
When AI CLI or work tracking commands fail, error messages should be helpful and point to AGENTS.md or configuration options.
