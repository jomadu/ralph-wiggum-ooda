package prompt

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestLoadFragment_Builtin_Success(t *testing.T) {
	// Use a known embedded fragment
	content, err := LoadFragment("builtin:fragments/observe/study_agents_md.md", "")
	if err != nil {
		t.Fatalf("expected success, got error: %v", err)
	}
	if content == "" {
		t.Fatal("expected non-empty content")
	}
	// Verify it looks like a markdown fragment
	if !strings.Contains(content, "#") {
		t.Error("expected markdown content with headers")
	}
}

func TestLoadFragment_Builtin_NotFound(t *testing.T) {
	_, err := LoadFragment("builtin:fragments/nonexistent.md", "")
	if err == nil {
		t.Fatal("expected error for missing embedded fragment")
	}
	if !strings.Contains(err.Error(), "embedded fragment not found") {
		t.Errorf("expected 'embedded fragment not found' error, got: %v", err)
	}
}

func TestLoadFragment_Filesystem_Success(t *testing.T) {
	// Create temp directory with test fragment
	tmpDir := t.TempDir()
	fragmentPath := filepath.Join(tmpDir, "test.md")
	testContent := "# Test Fragment\n\nThis is a test."
	if err := os.WriteFile(fragmentPath, []byte(testContent), 0644); err != nil {
		t.Fatalf("failed to create test file: %v", err)
	}

	// Load fragment using relative path
	content, err := LoadFragment("test.md", tmpDir)
	if err != nil {
		t.Fatalf("expected success, got error: %v", err)
	}
	if content != testContent {
		t.Errorf("expected %q, got %q", testContent, content)
	}
}

func TestLoadFragment_Filesystem_NotFound(t *testing.T) {
	tmpDir := t.TempDir()
	_, err := LoadFragment("nonexistent.md", tmpDir)
	if err == nil {
		t.Fatal("expected error for missing filesystem fragment")
	}
	if !strings.Contains(err.Error(), "fragment file not found") {
		t.Errorf("expected 'fragment file not found' error, got: %v", err)
	}
	// Verify error includes resolved path
	if !strings.Contains(err.Error(), "resolved to") {
		t.Error("expected error to include resolved path")
	}
}

func TestLoadFragment_Filesystem_RelativePath(t *testing.T) {
	// Create nested directory structure
	tmpDir := t.TempDir()
	subDir := filepath.Join(tmpDir, "fragments", "observe")
	if err := os.MkdirAll(subDir, 0755); err != nil {
		t.Fatalf("failed to create subdirectory: %v", err)
	}
	
	fragmentPath := filepath.Join(subDir, "custom.md")
	testContent := "# Custom Fragment"
	if err := os.WriteFile(fragmentPath, []byte(testContent), 0644); err != nil {
		t.Fatalf("failed to create test file: %v", err)
	}

	// Load using relative path from tmpDir
	content, err := LoadFragment("fragments/observe/custom.md", tmpDir)
	if err != nil {
		t.Fatalf("expected success, got error: %v", err)
	}
	if content != testContent {
		t.Errorf("expected %q, got %q", testContent, content)
	}
}

func TestLoadFragment_EmptyPath(t *testing.T) {
	_, err := LoadFragment("", "")
	if err == nil {
		t.Fatal("expected error for empty path")
	}
}
