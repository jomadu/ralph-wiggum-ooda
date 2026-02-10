package agents

import (
	"os"
	"path/filepath"
	"testing"
)

func TestVerifyAgentsMD_CommandSuccess(t *testing.T) {
	agentsMD := &AgentsMD{
		TestCommand: "echo test",
	}

	drifts := VerifyAgentsMD(agentsMD)

	if len(drifts) != 0 {
		t.Errorf("Expected no drifts for successful command, got %d", len(drifts))
	}
}

func TestVerifyAgentsMD_CommandFailure(t *testing.T) {
	agentsMD := &AgentsMD{
		TestCommand: "nonexistent-command-xyz",
	}

	drifts := VerifyAgentsMD(agentsMD)

	if len(drifts) == 0 {
		t.Fatal("Expected drift for failed command, got none")
	}

	drift := drifts[0]
	if drift.Field != "TestCommand" {
		t.Errorf("Expected Field=TestCommand, got %s", drift.Field)
	}
	if drift.Expected != "nonexistent-command-xyz" {
		t.Errorf("Expected Expected=nonexistent-command-xyz, got %s", drift.Expected)
	}
}

func TestVerifyAgentsMD_PathExists(t *testing.T) {
	tmpDir := t.TempDir()
	testFile := filepath.Join(tmpDir, "test.md")
	if err := os.WriteFile(testFile, []byte("test"), 0644); err != nil {
		t.Fatal(err)
	}

	agentsMD := &AgentsMD{
		SpecPaths: []string{testFile},
	}

	drifts := VerifyAgentsMD(agentsMD)

	if len(drifts) != 0 {
		t.Errorf("Expected no drifts for existing path, got %d", len(drifts))
	}
}

func TestVerifyAgentsMD_PathMissing(t *testing.T) {
	agentsMD := &AgentsMD{
		SpecPaths: []string{"/nonexistent/path/*.md"},
	}

	drifts := VerifyAgentsMD(agentsMD)

	if len(drifts) == 0 {
		t.Fatal("Expected drift for missing path, got none")
	}

	drift := drifts[0]
	if drift.Field != "SpecPaths" {
		t.Errorf("Expected Field=SpecPaths, got %s", drift.Field)
	}
}

func TestVerifyAgentsMD_MultipleDrifts(t *testing.T) {
	agentsMD := &AgentsMD{
		TestCommand: "nonexistent-command",
		SpecPaths:   []string{"/nonexistent/*.md"},
	}

	drifts := VerifyAgentsMD(agentsMD)

	if len(drifts) < 2 {
		t.Errorf("Expected at least 2 drifts, got %d", len(drifts))
	}
}

func TestVerifyAgentsMD_EmptyCommands(t *testing.T) {
	agentsMD := &AgentsMD{
		TestCommand:  "",
		BuildCommand: "",
	}

	drifts := VerifyAgentsMD(agentsMD)

	// Empty commands should not be verified (not an error)
	if len(drifts) != 0 {
		t.Errorf("Expected no drifts for empty commands, got %d", len(drifts))
	}
}
