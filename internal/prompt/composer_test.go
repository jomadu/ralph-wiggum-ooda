package prompt

import (
	"strings"
	"testing"

	"github.com/jomadu/rooda/internal/config"
)

func TestComposePhasePrompt_SingleFragment(t *testing.T) {
	fragments := []config.FragmentAction{
		{Path: "builtin:fragments/observe/read_agents_md.md"},
	}

	result, err := ComposePhasePrompt(fragments, "")
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	if !strings.Contains(result, "# Read AGENTS.md") {
		t.Errorf("expected fragment content, got: %s", result)
	}
}

func TestComposePhasePrompt_MultipleFragments(t *testing.T) {
	fragments := []config.FragmentAction{
		{Path: "builtin:fragments/observe/read_agents_md.md"},
		{Path: "builtin:fragments/observe/read_specs.md"},
	}

	result, err := ComposePhasePrompt(fragments, "")
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	// Should contain content from both fragments separated by double newlines
	if !strings.Contains(result, "# Read AGENTS.md") {
		t.Errorf("expected first fragment content")
	}
	if !strings.Contains(result, "# Read Specifications") {
		t.Errorf("expected second fragment content")
	}
	if !strings.Contains(result, "\n\n") {
		t.Errorf("expected double newlines between fragments")
	}
}

func TestComposePhasePrompt_InlineContent(t *testing.T) {
	fragments := []config.FragmentAction{
		{Content: "This is inline content."},
	}

	result, err := ComposePhasePrompt(fragments, "")
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	if result != "This is inline content." {
		t.Errorf("expected inline content, got: %s", result)
	}
}

func TestComposePhasePrompt_WithTemplate(t *testing.T) {
	fragments := []config.FragmentAction{
		{
			Content:    "Hello {{.name}}",
			Parameters: map[string]interface{}{"name": "World"},
		},
	}

	result, err := ComposePhasePrompt(fragments, "")
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	if result != "Hello World" {
		t.Errorf("expected 'Hello World', got: %s", result)
	}
}

func TestComposePhasePrompt_EmptyFragments(t *testing.T) {
	fragments := []config.FragmentAction{}

	result, err := ComposePhasePrompt(fragments, "")
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	if result != "" {
		t.Errorf("expected empty string, got: %s", result)
	}
}

func TestAssemblePrompt_AllPhases(t *testing.T) {
	procedure := config.Procedure{
		Observe: []config.FragmentAction{
			{Content: "Observe content"},
		},
		Orient: []config.FragmentAction{
			{Content: "Orient content"},
		},
		Decide: []config.FragmentAction{
			{Content: "Decide content"},
		},
		Act: []config.FragmentAction{
			{Content: "Act content"},
		},
	}

	result, err := AssemblePrompt(procedure, "", "")
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	// Check section markers
	if !strings.Contains(result, "# OBSERVE") {
		t.Errorf("expected OBSERVE section marker")
	}
	if !strings.Contains(result, "# ORIENT") {
		t.Errorf("expected ORIENT section marker")
	}
	if !strings.Contains(result, "# DECIDE") {
		t.Errorf("expected DECIDE section marker")
	}
	if !strings.Contains(result, "# ACT") {
		t.Errorf("expected ACT section marker")
	}

	// Check content
	if !strings.Contains(result, "Observe content") {
		t.Errorf("expected observe content")
	}
	if !strings.Contains(result, "Orient content") {
		t.Errorf("expected orient content")
	}
	if !strings.Contains(result, "Decide content") {
		t.Errorf("expected decide content")
	}
	if !strings.Contains(result, "Act content") {
		t.Errorf("expected act content")
	}
}

func TestAssemblePrompt_WithUserContext(t *testing.T) {
	procedure := config.Procedure{
		Observe: []config.FragmentAction{
			{Content: "Observe content"},
		},
		Orient: []config.FragmentAction{
			{Content: "Orient content"},
		},
		Decide: []config.FragmentAction{
			{Content: "Decide content"},
		},
		Act: []config.FragmentAction{
			{Content: "Act content"},
		},
	}

	userContext := "Focus on authentication module"
	result, err := AssemblePrompt(procedure, userContext, "")
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	// Check context appears first with marker
	if !strings.HasPrefix(result, "# CONTEXT\n") {
		t.Errorf("expected CONTEXT section marker at start")
	}
	if !strings.Contains(result, "Focus on authentication module") {
		t.Errorf("expected user context content")
	}

	// Context should appear before OBSERVE
	contextIdx := strings.Index(result, "# CONTEXT")
	observeIdx := strings.Index(result, "# OBSERVE")
	if contextIdx == -1 || observeIdx == -1 || contextIdx >= observeIdx {
		t.Errorf("expected CONTEXT before OBSERVE")
	}
}

func TestAssemblePrompt_EmptyPhases(t *testing.T) {
	procedure := config.Procedure{
		Observe: []config.FragmentAction{},
		Orient:  []config.FragmentAction{},
		Decide:  []config.FragmentAction{},
		Act:     []config.FragmentAction{},
	}

	result, err := AssemblePrompt(procedure, "", "")
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	// Empty phases should not produce section markers
	if strings.Contains(result, "# OBSERVE") {
		t.Errorf("expected no OBSERVE marker for empty phase")
	}
	if result != "" {
		t.Errorf("expected empty result, got: %s", result)
	}
}
