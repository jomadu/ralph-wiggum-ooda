package agents

import (
	"os/exec"
	"path/filepath"
)

// VerifyAgentsMD verifies AGENTS.md claims against reality
// Returns list of detected drifts
func VerifyAgentsMD(agentsMD *AgentsMD) []DriftDetection {
	var drifts []DriftDetection

	// Verify build command
	if agentsMD.BuildCommand != "" {
		if drift := verifyCommand("BuildCommand", agentsMD.BuildCommand); drift != nil {
			drifts = append(drifts, *drift)
		}
	}

	// Verify test command
	if agentsMD.TestCommand != "" {
		if drift := verifyCommand("TestCommand", agentsMD.TestCommand); drift != nil {
			drifts = append(drifts, *drift)
		}
	}

	// Verify lint commands
	for _, cmd := range agentsMD.LintCommands {
		if cmd != "" {
			if drift := verifyCommand("LintCommands", cmd); drift != nil {
				drifts = append(drifts, *drift)
			}
		}
	}

	// Verify spec paths
	for _, pattern := range agentsMD.SpecPaths {
		if drift := verifyPath("SpecPaths", pattern); drift != nil {
			drifts = append(drifts, *drift)
		}
	}

	// Verify impl paths
	for _, pattern := range agentsMD.ImplPaths {
		if drift := verifyPath("ImplPaths", pattern); drift != nil {
			drifts = append(drifts, *drift)
		}
	}

	// Verify work tracking query command
	if agentsMD.WorkTracking.QueryCommand != "" {
		if drift := verifyCommand("WorkTracking.QueryCommand", agentsMD.WorkTracking.QueryCommand); drift != nil {
			drifts = append(drifts, *drift)
		}
	}

	return drifts
}

// verifyCommand checks if a command can be executed
func verifyCommand(field, command string) *DriftDetection {
	// Parse command to get binary name
	cmd := exec.Command("sh", "-c", "command -v "+command+" >/dev/null 2>&1")
	if err := cmd.Run(); err != nil {
		return &DriftDetection{
			Field:    field,
			Expected: command,
			Actual:   "command not found or failed",
			Rationale: "Command verification failed",
		}
	}
	return nil
}

// verifyPath checks if a path pattern matches any files
func verifyPath(field, pattern string) *DriftDetection {
	// Handle absolute paths
	if filepath.IsAbs(pattern) {
		matches, err := filepath.Glob(pattern)
		if err != nil || len(matches) == 0 {
			return &DriftDetection{
				Field:     field,
				Expected:  pattern,
				Actual:    "no files matched",
				Rationale: "Path pattern matched no files",
			}
		}
		return nil
	}

	// Handle relative paths - check if file exists
	matches, err := filepath.Glob(pattern)
	if err != nil || len(matches) == 0 {
		return &DriftDetection{
			Field:     field,
			Expected:  pattern,
			Actual:    "no files matched",
			Rationale: "Path pattern matched no files",
		}
	}

	return nil
}
