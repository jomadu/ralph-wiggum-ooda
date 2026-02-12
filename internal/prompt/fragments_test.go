package prompt

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestFragmentWordCount(t *testing.T) {
	const maxWords = 100
	
	phases := []string{"observe", "orient", "decide", "act"}
	var violations []string
	totalWords := 0
	totalFragments := 0
	
	for _, phase := range phases {
		phaseDir := filepath.Join("fragments", phase)
		entries, err := os.ReadDir(phaseDir)
		if err != nil {
			t.Fatalf("Failed to read %s directory: %v", phase, err)
		}
		
		for _, entry := range entries {
			if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".md") {
				continue
			}
			
			path := filepath.Join(phaseDir, entry.Name())
			content, err := os.ReadFile(path)
			if err != nil {
				t.Errorf("Failed to read %s: %v", path, err)
				continue
			}
			
			wordCount := countWords(string(content))
			totalWords += wordCount
			totalFragments++
			
			if wordCount > maxWords {
				firstLine := strings.TrimSpace(strings.Split(string(content), "\n")[0])
				violations = append(violations, filepath.Join(phase, entry.Name())+" ("+firstLine+"): "+intToString(wordCount)+" words")
			}
		}
	}
	
	if len(violations) > 0 {
		t.Errorf("Found %d fragments exceeding %d words:\n%s", len(violations), maxWords, strings.Join(violations, "\n"))
	}
	
	avgWords := 0
	if totalFragments > 0 {
		avgWords = totalWords / totalFragments
	}
	
	t.Logf("Fragment statistics:")
	t.Logf("  Total fragments: %d", totalFragments)
	t.Logf("  Total words: %d", totalWords)
	t.Logf("  Average words per fragment: %d", avgWords)
	t.Logf("  Violations (>%d words): %d", maxWords, len(violations))
}

func intToString(n int) string {
	if n == 0 {
		return "0"
	}
	
	var result []byte
	negative := n < 0
	if negative {
		n = -n
	}
	
	for n > 0 {
		result = append([]byte{byte('0' + n%10)}, result...)
		n /= 10
	}
	
	if negative {
		result = append([]byte{'-'}, result...)
	}
	
	return string(result)
}

func TestStudyAgentsMdPreservesTopics(t *testing.T) {
	content, err := os.ReadFile("fragments/observe/study_agents_md.md")
	if err != nil {
		t.Fatalf("failed to read study_agents_md.md: %v", err)
	}

	text := string(content)
	requiredTopics := []string{
		"Work Tracking",
		"Quick Reference",
		"Task Input",
		"Planning System",
		"Build",
		"Test",
		"Lint",
		"Specification",
		"Implementation",
		"Audit",
		"Quality Criteria",
		"Operational Learnings",
	}

	for _, topic := range requiredTopics {
		if !strings.Contains(text, topic) {
			t.Errorf("fragment missing required topic: %s", topic)
		}
	}
}

func countWords(s string) int {
	return len(strings.Fields(s))
}
