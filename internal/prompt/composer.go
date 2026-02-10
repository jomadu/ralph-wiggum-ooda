package prompt

import (
	"fmt"
	"strings"

	"github.com/jomadu/rooda/internal/config"
)

// AssemblePrompt assembles a complete prompt from a procedure definition.
// It concatenates fragments from each OODA phase with section markers and
// optionally injects user context at the top.
func AssemblePrompt(procedure config.Procedure, userContext string, configDir string) (string, error) {
	var prompt strings.Builder

	// Inject user context first if provided
	if userContext != "" {
		prompt.WriteString("# CONTEXT\n")
		prompt.WriteString(userContext)
		prompt.WriteString("\n\n")
	}

	// Process each OODA phase in order
	phases := []struct {
		name      string
		fragments []config.FragmentAction
	}{
		{"OBSERVE", procedure.Observe},
		{"ORIENT", procedure.Orient},
		{"DECIDE", procedure.Decide},
		{"ACT", procedure.Act},
	}

	for _, phase := range phases {
		phaseContent, err := ComposePhasePrompt(phase.fragments, configDir)
		if err != nil {
			return "", fmt.Errorf("failed to compose %s phase: %v", phase.name, err)
		}

		// Add section marker and content if phase has content
		trimmed := strings.TrimSpace(phaseContent)
		if trimmed != "" {
			prompt.WriteString("# ")
			prompt.WriteString(phase.name)
			prompt.WriteString("\n")
			prompt.WriteString(trimmed)
			prompt.WriteString("\n\n")
		}
	}

	return prompt.String(), nil
}

// ComposePhasePrompt composes a single phase prompt from an array of fragment actions.
// It loads fragments, processes templates if parameters are provided, and concatenates
// with double newlines.
func ComposePhasePrompt(fragments []config.FragmentAction, configDir string) (string, error) {
	if len(fragments) == 0 {
		return "", nil
	}

	var parts []string

	for _, fragment := range fragments {
		var content string

		// Determine content source (inline content or file path)
		if fragment.Content != "" {
			content = fragment.Content
		} else if fragment.Path != "" {
			var err error
			content, err = LoadFragment(fragment.Path, configDir)
			if err != nil {
				return "", err
			}
		} else {
			return "", fmt.Errorf("fragment must specify either content or path")
		}

		// Process template if parameters provided
		if len(fragment.Parameters) > 0 {
			var err error
			content, err = ProcessTemplate(content, fragment.Parameters)
			if err != nil {
				return "", err
			}
		}

		// Append to parts
		parts = append(parts, strings.TrimSpace(content))
	}

	// Concatenate with double newlines
	return strings.Join(parts, "\n\n"), nil
}
