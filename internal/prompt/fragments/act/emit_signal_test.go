package act

import (
	"os"
	"strings"
	"testing"
)

func TestEmitSignalFragmentContent(t *testing.T) {
	content, err := os.ReadFile("emit_signal.md")
	if err != nil {
		t.Fatalf("Failed to read emit_signal.md: %v", err)
	}

	text := string(content)

	// Essential content checks
	essentialElements := []struct {
		name    string
		content string
	}{
		{"SUCCESS signal format", "<promise>SUCCESS</promise>"},
		{"FAILURE signal format", "<promise>FAILURE</promise>"},
		{"SUCCESS signal type", "SUCCESS"},
		{"FAILURE signal type", "FAILURE"},
		{"Continue behavior", "continue"},
		{"Signal format instruction", "signal"},
	}

	for _, elem := range essentialElements {
		if !strings.Contains(text, elem.content) {
			t.Errorf("Missing essential element '%s': expected to find '%s'", elem.name, elem.content)
		}
	}
}

func TestEmitSignalFragmentWordCount(t *testing.T) {
	content, err := os.ReadFile("emit_signal.md")
	if err != nil {
		t.Fatalf("Failed to read emit_signal.md: %v", err)
	}

	words := strings.Fields(string(content))
	wordCount := len(words)

	if wordCount >= 100 {
		t.Errorf("Fragment exceeds 100 words: got %d words, want <100", wordCount)
	}
}
