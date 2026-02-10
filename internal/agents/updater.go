package agents

import (
	"fmt"
	"regexp"
)

// UpdateAgentsMD updates AGENTS.md content in-place with drift fixes
// Returns updated content with inline rationale comments
func UpdateAgentsMD(agentsMD *AgentsMD, drifts []DriftDetection) string {
	if len(drifts) == 0 {
		return agentsMD.RawContent
	}

	content := agentsMD.RawContent

	for _, drift := range drifts {
		content = applyDriftFix(content, drift)
	}

	return content
}

// applyDriftFix applies a single drift fix to the content
func applyDriftFix(content string, drift DriftDetection) string {
	switch drift.Field {
	case "TestCommand":
		return updateCommand(content, "Test", drift)
	case "BuildCommand":
		return updateCommand(content, "Build", drift)
	case "LintCommands":
		return updateCommand(content, "Lint", drift)
	case "SpecPaths":
		return updatePath(content, "Location", drift)
	case "ImplPaths":
		return updatePath(content, "Location", drift)
	case "WorkTracking.QueryCommand":
		return updateCommand(content, "Query ready work", drift)
	default:
		return content
	}
}

// updateCommand updates a command line with rationale
func updateCommand(content, label string, drift DriftDetection) string {
	// Pattern: **Label:** `command`
	pattern := fmt.Sprintf(`\*\*%s:\*\*\s+` + "`" + `([^` + "`" + `]+)` + "`", regexp.QuoteMeta(label))
	re := regexp.MustCompile(pattern)

	if drift.FixApplied == "" {
		// No fix available - just add comment
		return re.ReplaceAllString(content, fmt.Sprintf("**%s:** `$1`  # %s", label, drift.Rationale))
	}

	// Replace with fix and add rationale
	replacement := fmt.Sprintf("**%s:** `%s`  # Changed from %s - %s", 
		label, drift.FixApplied, drift.Expected, drift.Rationale)
	
	return re.ReplaceAllString(content, replacement)
}

// updatePath updates a path pattern with rationale
func updatePath(content, label string, drift DriftDetection) string {
	// Pattern: **Label:** `path`
	pattern := fmt.Sprintf(`\*\*%s:\*\*\s+` + "`" + `([^` + "`" + `]+)` + "`", regexp.QuoteMeta(label))
	re := regexp.MustCompile(pattern)

	if drift.FixApplied == "" {
		// No fix available - just add comment
		return re.ReplaceAllString(content, fmt.Sprintf("**%s:** `$1`  # %s", label, drift.Rationale))
	}

	// Replace with fix and add rationale
	replacement := fmt.Sprintf("**%s:** `%s`  # Changed from %s - %s", 
		label, drift.FixApplied, drift.Expected, drift.Rationale)
	
	return re.ReplaceAllString(content, replacement)
}
