package prompt

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// TestFragmentsUseImperativeVoice validates that all fragments use imperative voice
// and follow the established pattern for executable procedures.
func TestFragmentsUseImperativeVoice(t *testing.T) {
	fragmentsDir := "fragments"
	
	// Strong imperative patterns that indicate direct commands
	strongImperativePatterns := []string{
		"you must",
		"your task is",
		"use the",
		"execute the",
		"load the",
		"modify the",
		"create the",
		"update the",
	}
	
	// Passive patterns that indicate documentation rather than instructions
	passivePatterns := []string{
		"this phase",
		"this section",
		"the agent should",
		"the system will",
		"is used to",
		"can be used",
		"allows you to",
	}
	
	phases := []string{"observe", "orient", "decide", "act"}
	
	for _, phase := range phases {
		phaseDir := filepath.Join(fragmentsDir, phase)
		entries, err := os.ReadDir(phaseDir)
		if err != nil {
			t.Fatalf("Failed to read %s directory: %v", phaseDir, err)
		}
		
		for _, entry := range entries {
			if !strings.HasSuffix(entry.Name(), ".md") {
				continue
			}
			
			fragmentPath := filepath.Join(phaseDir, entry.Name())
			content, err := os.ReadFile(fragmentPath)
			if err != nil {
				t.Errorf("Failed to read fragment %s: %v", fragmentPath, err)
				continue
			}
			
			contentLower := strings.ToLower(string(content))
			
			// Check for strong imperative patterns
			hasStrongImperative := false
			for _, pattern := range strongImperativePatterns {
				if strings.Contains(contentLower, pattern) {
					hasStrongImperative = true
					break
				}
			}
			
			if !hasStrongImperative {
				t.Errorf("Fragment %s does not use strong imperative voice (should start with 'You must...' or 'Your task is...')", fragmentPath)
			}
			
			// Check for passive patterns (should not be present)
			for _, pattern := range passivePatterns {
				if strings.Contains(contentLower, pattern) {
					t.Errorf("Fragment %s uses passive documentation language: contains '%s'", fragmentPath, pattern)
				}
			}
			
			// Check that content is substantial (not just a title)
			lines := strings.Split(string(content), "\n")
			nonEmptyLines := 0
			for _, line := range lines {
				if strings.TrimSpace(line) != "" && !strings.HasPrefix(strings.TrimSpace(line), "#") {
					nonEmptyLines++
				}
			}
			
			if nonEmptyLines < 2 {
				t.Errorf("Fragment %s appears to be incomplete (less than 2 non-empty, non-header lines)", fragmentPath)
			}
		}
	}
}
