package config

import (
	"strings"
	"testing"
)

func TestResolveAICommand_CLIFlagDirectCommand(t *testing.T) {
	config := Config{
		Loop: LoopConfig{AICmdAlias: "kiro-cli"},
		AICmdAliases: map[string]string{
			"kiro-cli": "kiro-cli chat --no-interactive --trust-all-tools",
		},
	}
	flags := CLIFlags{AICmd: "custom-tool --flag"}
	
	cmd, err := ResolveAICommand(config, "build", flags)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	if cmd.Command != "custom-tool --flag" {
		t.Errorf("expected command 'custom-tool --flag', got: %s", cmd.Command)
	}
	if cmd.Source != "--ai-cmd flag" {
		t.Errorf("expected source '--ai-cmd flag', got: %s", cmd.Source)
	}
}

func TestResolveAICommand_CLIFlagAlias(t *testing.T) {
	config := Config{
		AICmdAliases: map[string]string{
			"kiro-cli": "kiro-cli chat --no-interactive --trust-all-tools",
			"claude":   "claude -p --dangerously-skip-permissions",
		},
	}
	flags := CLIFlags{AICmdAlias: "claude"}
	
	cmd, err := ResolveAICommand(config, "build", flags)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	if cmd.Command != "claude -p --dangerously-skip-permissions" {
		t.Errorf("expected claude command, got: %s", cmd.Command)
	}
	if cmd.Source != "--ai-cmd-alias flag=claude" {
		t.Errorf("expected source '--ai-cmd-alias flag=claude', got: %s", cmd.Source)
	}
}

func TestResolveAICommand_ProcedureDirectCommand(t *testing.T) {
	config := Config{
		Procedures: map[string]Procedure{
			"build": {AICmd: "proc-tool --flag"},
		},
		Loop: LoopConfig{AICmdAlias: "kiro-cli"},
		AICmdAliases: map[string]string{
			"kiro-cli": "kiro-cli chat --no-interactive --trust-all-tools",
		},
	}
	
	cmd, err := ResolveAICommand(config, "build", CLIFlags{})
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	if cmd.Command != "proc-tool --flag" {
		t.Errorf("expected command 'proc-tool --flag', got: %s", cmd.Command)
	}
	if cmd.Source != "procedure.build.ai_cmd" {
		t.Errorf("expected source 'procedure.build.ai_cmd', got: %s", cmd.Source)
	}
}

func TestResolveAICommand_ProcedureAlias(t *testing.T) {
	config := Config{
		Procedures: map[string]Procedure{
			"build": {AICmdAlias: "claude"},
		},
		AICmdAliases: map[string]string{
			"claude": "claude -p --dangerously-skip-permissions",
		},
	}
	
	cmd, err := ResolveAICommand(config, "build", CLIFlags{})
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	if cmd.Command != "claude -p --dangerously-skip-permissions" {
		t.Errorf("expected claude command, got: %s", cmd.Command)
	}
	if cmd.Source != "procedure.build.ai_cmd_alias=claude" {
		t.Errorf("expected source 'procedure.build.ai_cmd_alias=claude', got: %s", cmd.Source)
	}
}

func TestResolveAICommand_LoopDirectCommand(t *testing.T) {
	config := Config{
		Loop: LoopConfig{AICmd: "loop-tool --flag"},
	}
	
	cmd, err := ResolveAICommand(config, "build", CLIFlags{})
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	if cmd.Command != "loop-tool --flag" {
		t.Errorf("expected command 'loop-tool --flag', got: %s", cmd.Command)
	}
	if cmd.Source != "loop.ai_cmd" {
		t.Errorf("expected source 'loop.ai_cmd', got: %s", cmd.Source)
	}
}

func TestResolveAICommand_LoopAlias(t *testing.T) {
	config := Config{
		Loop: LoopConfig{AICmdAlias: "kiro-cli"},
		AICmdAliases: map[string]string{
			"kiro-cli": "kiro-cli chat --no-interactive --trust-all-tools",
		},
	}
	
	cmd, err := ResolveAICommand(config, "build", CLIFlags{})
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	if cmd.Command != "kiro-cli chat --no-interactive --trust-all-tools" {
		t.Errorf("expected kiro-cli command, got: %s", cmd.Command)
	}
	if cmd.Source != "loop.ai_cmd_alias=kiro-cli" {
		t.Errorf("expected source 'loop.ai_cmd_alias=kiro-cli', got: %s", cmd.Source)
	}
}

func TestResolveAICommand_NoCommandConfigured(t *testing.T) {
	config := Config{
		AICmdAliases: map[string]string{
			"kiro-cli": "kiro-cli chat --no-interactive --trust-all-tools",
		},
	}
	
	_, err := ResolveAICommand(config, "build", CLIFlags{})
	if err == nil {
		t.Fatal("expected error when no AI command configured")
	}
	
	errMsg := err.Error()
	if !strings.Contains(errMsg, "no AI command configured") {
		t.Errorf("expected error to mention 'no AI command configured', got: %s", errMsg)
	}
	if !strings.Contains(errMsg, "--ai-cmd") {
		t.Errorf("expected error to mention '--ai-cmd', got: %s", errMsg)
	}
	if !strings.Contains(errMsg, "kiro-cli") {
		t.Errorf("expected error to list available aliases, got: %s", errMsg)
	}
}

func TestResolveAICommand_UnknownAlias(t *testing.T) {
	config := Config{
		AICmdAliases: map[string]string{
			"kiro-cli": "kiro-cli chat --no-interactive --trust-all-tools",
		},
	}
	flags := CLIFlags{AICmdAlias: "nonexistent"}
	
	_, err := ResolveAICommand(config, "build", flags)
	if err == nil {
		t.Fatal("expected error for unknown alias")
	}
	
	errMsg := err.Error()
	if !strings.Contains(errMsg, "unknown AI command alias") {
		t.Errorf("expected error to mention 'unknown AI command alias', got: %s", errMsg)
	}
	if !strings.Contains(errMsg, "nonexistent") {
		t.Errorf("expected error to mention alias name, got: %s", errMsg)
	}
}

func TestResolveAICommand_DirectCommandOverridesAlias(t *testing.T) {
	config := Config{
		Procedures: map[string]Procedure{
			"build": {
				AICmd:      "direct-cmd",
				AICmdAlias: "kiro-cli",
			},
		},
		AICmdAliases: map[string]string{
			"kiro-cli": "kiro-cli chat --no-interactive --trust-all-tools",
		},
	}
	
	cmd, err := ResolveAICommand(config, "build", CLIFlags{})
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	if cmd.Command != "direct-cmd" {
		t.Errorf("expected direct command to win, got: %s", cmd.Command)
	}
}

func TestResolveAICommand_AllBuiltinAliases(t *testing.T) {
	aliases := []string{"kiro-cli", "claude", "copilot", "cursor-agent"}
	
	for _, alias := range aliases {
		t.Run(alias, func(t *testing.T) {
			config := Config{
				Loop: LoopConfig{AICmdAlias: alias},
				AICmdAliases: map[string]string{
					"kiro-cli":     "kiro-cli chat --no-interactive --trust-all-tools",
					"claude":       "claude -p --dangerously-skip-permissions",
					"copilot":      "copilot --yolo",
					"cursor-agent": "cursor-wrapper.sh",
				},
			}
			
			cmd, err := ResolveAICommand(config, "build", CLIFlags{})
			if err != nil {
				t.Fatalf("expected no error for alias %s, got: %v", alias, err)
			}
			if cmd.Command == "" {
				t.Errorf("expected non-empty command for alias %s", alias)
			}
		})
	}
}
