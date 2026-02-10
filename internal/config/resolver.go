package config

import (
	"fmt"
	"sort"
	"strings"
)

// ResolveAICommand resolves the AI command from the precedence chain.
// Precedence: CLI flags > procedure config > loop config > error
func ResolveAICommand(config Config, procedureName string, cliFlags CLIFlags) (AICommand, error) {
	// 1. --ai-cmd flag (direct command, highest precedence)
	if cliFlags.AICmd != "" {
		return AICommand{
			Command: cliFlags.AICmd,
			Source:  "--ai-cmd flag",
		}, nil
	}

	// 2. --ai-cmd-alias flag (alias from merged config)
	if cliFlags.AICmdAlias != "" {
		return resolveAlias(config, cliFlags.AICmdAlias, "--ai-cmd-alias flag")
	}

	// 3. procedure.ai_cmd (direct command)
	if proc, exists := config.Procedures[procedureName]; exists {
		if proc.AICmd != "" {
			return AICommand{
				Command: proc.AICmd,
				Source:  fmt.Sprintf("procedure.%s.ai_cmd", procedureName),
			}, nil
		}

		// 4. procedure.ai_cmd_alias (alias from merged config)
		if proc.AICmdAlias != "" {
			return resolveAlias(config, proc.AICmdAlias, fmt.Sprintf("procedure.%s.ai_cmd_alias", procedureName))
		}
	}

	// 5. loop.ai_cmd (already merged from config tiers + env vars)
	if config.Loop.AICmd != "" {
		return AICommand{
			Command: config.Loop.AICmd,
			Source:  "loop.ai_cmd",
		}, nil
	}

	// 6. loop.ai_cmd_alias (already merged from config tiers + env vars)
	if config.Loop.AICmdAlias != "" {
		return resolveAlias(config, config.Loop.AICmdAlias, "loop.ai_cmd_alias")
	}

	// 7. No AI command configured — error with guidance
	aliases := make([]string, 0, len(config.AICmdAliases))
	for alias := range config.AICmdAliases {
		aliases = append(aliases, alias)
	}
	sort.Strings(aliases)

	return AICommand{}, fmt.Errorf(`no AI command configured

Set one via:
  --ai-cmd "your-command"           CLI flag (direct command)
  --ai-cmd-alias <name>             CLI flag (alias from config)
  ROODA_LOOP_AI_CMD=your-command    Environment variable
  ROODA_LOOP_AI_CMD_ALIAS=<name>    Environment variable
  loop.ai_cmd or loop.ai_cmd_alias  rooda-config.yml
  procedure.ai_cmd or ai_cmd_alias  rooda-config.yml

Available aliases: %s`, strings.Join(aliases, ", "))
}

// resolveAlias resolves an alias name to a command string.
func resolveAlias(config Config, aliasName string, source string) (AICommand, error) {
	command, exists := config.AICmdAliases[aliasName]
	if !exists {
		aliases := make([]string, 0, len(config.AICmdAliases))
		for alias := range config.AICmdAliases {
			aliases = append(aliases, alias)
		}
		sort.Strings(aliases)

		return AICommand{}, fmt.Errorf("unknown AI command alias: %s (from %s)\nAvailable: %s",
			aliasName, source, strings.Join(aliases, ", "))
	}

	return AICommand{
		Command: command,
		Source:  fmt.Sprintf("%s=%s", source, aliasName),
	}, nil
}
