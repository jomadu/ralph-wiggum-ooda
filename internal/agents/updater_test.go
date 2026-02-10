package agents

import (
	"strings"
	"testing"
)

func TestUpdateAgentsMD_CommandDrift(t *testing.T) {
	agentsMD := &AgentsMD{
		RawContent: "# Agent Instructions\n\n## Build/Test/Lint Commands\n\n**Test:** `npm test`\n",
	}

	drifts := []DriftDetection{
		{
			Field:      "TestCommand",
			Expected:   "npm test",
			Actual:     "command not found",
			FixApplied: "go test ./...",
			Rationale:  "npm not installed, project uses Go",
		},
	}

	updated := UpdateAgentsMD(agentsMD, drifts)

	if !strings.Contains(updated, "go test ./...") {
		t.Error("Expected updated command 'go test ./...' not found")
	}

	if !strings.Contains(updated, "# Changed from npm test") {
		t.Error("Expected rationale comment not found")
	}

	// Old command should only appear in rationale comment, not as the active command
	if strings.Contains(updated, "**Test:** `npm test`") {
		t.Error("Old command should not be the active command")
	}
}

func TestUpdateAgentsMD_PathDrift(t *testing.T) {
	agentsMD := &AgentsMD{
		RawContent: "# Agent Instructions\n\n## Specification Definition\n\n**Location:** `specs/*.md`\n",
	}

	drifts := []DriftDetection{
		{
			Field:      "SpecPaths",
			Expected:   "specs/*.md",
			Actual:     "no files matched",
			FixApplied: "documentation/*.md",
			Rationale:  "specs moved to documentation/",
		},
	}

	updated := UpdateAgentsMD(agentsMD, drifts)

	if !strings.Contains(updated, "documentation/*.md") {
		t.Error("Expected updated path 'documentation/*.md' not found")
	}

	if !strings.Contains(updated, "# Changed from specs/*.md") {
		t.Error("Expected rationale comment not found")
	}
}

func TestUpdateAgentsMD_MultipleDrifts(t *testing.T) {
	agentsMD := &AgentsMD{
		RawContent: "# Agent Instructions\n\n## Build/Test/Lint Commands\n\n**Test:** `npm test`\n**Build:** `npm run build`\n",
	}

	drifts := []DriftDetection{
		{
			Field:      "TestCommand",
			Expected:   "npm test",
			FixApplied: "go test ./...",
			Rationale:  "npm not installed",
		},
		{
			Field:      "BuildCommand",
			Expected:   "npm run build",
			FixApplied: "go build",
			Rationale:  "npm not installed",
		},
	}

	updated := UpdateAgentsMD(agentsMD, drifts)

	if !strings.Contains(updated, "go test ./...") {
		t.Error("Expected updated test command not found")
	}

	if !strings.Contains(updated, "go build") {
		t.Error("Expected updated build command not found")
	}

	if strings.Count(updated, "# Changed from") != 2 {
		t.Errorf("Expected 2 rationale comments, found %d", strings.Count(updated, "# Changed from"))
	}
}

func TestUpdateAgentsMD_NoDrifts(t *testing.T) {
	original := "# Agent Instructions\n\n## Build/Test/Lint Commands\n\n**Test:** `go test ./...`\n"

	agentsMD := &AgentsMD{
		RawContent: original,
	}

	drifts := []DriftDetection{}

	updated := UpdateAgentsMD(agentsMD, drifts)

	if updated != original {
		t.Error("Expected content to remain unchanged when no drifts")
	}
}

func TestUpdateAgentsMD_PreservesOtherContent(t *testing.T) {
	agentsMD := &AgentsMD{
		RawContent: "# Agent Instructions\n\n## Build/Test/Lint Commands\n\n**Test:** `npm test`\n\n## Other Section\n\nThis content should be preserved.\n",
	}

	drifts := []DriftDetection{
		{
			Field:      "TestCommand",
			Expected:   "npm test",
			FixApplied: "go test ./...",
			Rationale:  "npm not installed",
		},
	}

	updated := UpdateAgentsMD(agentsMD, drifts)

	if !strings.Contains(updated, "## Other Section") {
		t.Error("Expected other sections to be preserved")
	}

	if !strings.Contains(updated, "This content should be preserved.") {
		t.Error("Expected other content to be preserved")
	}
}
