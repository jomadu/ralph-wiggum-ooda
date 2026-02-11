package main

import (
	"bytes"
	"os"
	"os/exec"
	"strings"
	"testing"
)

// TestCLI runs the rooda binary and checks output/exit codes
func runRooda(t *testing.T, args ...string) (stdout, stderr string, exitCode int) {
	t.Helper()
	
	// Always rebuild the binary to ensure latest code
	binPath := "../../bin/rooda"
	buildCmd := exec.Command("go", "build", "-o", binPath, ".")
	if err := buildCmd.Run(); err != nil {
		t.Fatalf("Failed to build rooda binary: %v", err)
	}
	
	cmd := exec.Command(binPath, args...)
	var outBuf, errBuf bytes.Buffer
	cmd.Stdout = &outBuf
	cmd.Stderr = &errBuf
	
	err := cmd.Run()
	exitCode = 0
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			exitCode = exitErr.ExitCode()
		} else {
			t.Fatalf("Failed to run rooda: %v", err)
		}
	}
	
	return outBuf.String(), errBuf.String(), exitCode
}

// Test --version flag
func TestVersionFlag(t *testing.T) {
	stdout, _, exitCode := runRooda(t, "--version")
	
	if exitCode != ExitSuccess {
		t.Errorf("Expected exit code %d, got %d", ExitSuccess, exitCode)
	}
	
	if !strings.Contains(stdout, "rooda") {
		t.Errorf("Expected version output to contain 'rooda', got: %s", stdout)
	}
}

// Test --help flag (global help)
func TestGlobalHelp(t *testing.T) {
	stdout, _, exitCode := runRooda(t, "--help")
	
	if exitCode != ExitSuccess {
		t.Errorf("Expected exit code %d, got %d", ExitSuccess, exitCode)
	}
	
	requiredStrings := []string{
		"USAGE:",
		"rooda",
		"--help",
		"--version",
		"--list-procedures",
	}
	
	for _, required := range requiredStrings {
		if !strings.Contains(stdout, required) {
			t.Errorf("Expected help output to contain '%s', got: %s", required, stdout)
		}
	}
}

// Test --list-procedures flag
func TestListProcedures(t *testing.T) {
	stdout, _, exitCode := runRooda(t, "--list-procedures")
	
	if exitCode != ExitSuccess {
		t.Errorf("Expected exit code %d, got %d", ExitSuccess, exitCode)
	}
	
	// Check for some known procedures
	requiredProcedures := []string{
		"build",
		"agents-sync",
		"audit-spec",
		"audit-impl",
	}
	
	for _, proc := range requiredProcedures {
		if !strings.Contains(stdout, proc) {
			t.Errorf("Expected procedure list to contain '%s', got: %s", proc, stdout)
		}
	}
}

// Test no procedure specified error
func TestNoProcedureError(t *testing.T) {
	_, stderr, exitCode := runRooda(t)
	
	if exitCode != ExitUserError {
		t.Errorf("Expected exit code %d, got %d", ExitUserError, exitCode)
	}
	
	if !strings.Contains(stderr, "No procedure specified") {
		t.Errorf("Expected error message about no procedure, got: %s", stderr)
	}
	
	if !strings.Contains(stderr, "rooda --help") {
		t.Errorf("Expected error to suggest --help, got: %s", stderr)
	}
}

// Test unknown procedure error
func TestUnknownProcedureError(t *testing.T) {
	_, stderr, exitCode := runRooda(t, "nonexistent-procedure", "--ai-cmd", "echo", "--dry-run")
	
	if exitCode != ExitUserError {
		t.Errorf("Expected exit code %d, got %d", ExitUserError, exitCode)
	}
	
	if !strings.Contains(stderr, "Unknown procedure") {
		t.Errorf("Expected error about unknown procedure, got: %s", stderr)
	}
	
	if !strings.Contains(stderr, "nonexistent-procedure") {
		t.Errorf("Expected error to mention the procedure name, got: %s", stderr)
	}
	
	if !strings.Contains(stderr, "--list-procedures") {
		t.Errorf("Expected error to suggest --list-procedures, got: %s", stderr)
	}
}

// Test mutually exclusive flags: --verbose and --quiet
func TestMutuallyExclusiveVerboseQuiet(t *testing.T) {
	_, stderr, exitCode := runRooda(t, "build", "--verbose", "--quiet")
	
	if exitCode != ExitUserError {
		t.Errorf("Expected exit code %d, got %d", ExitUserError, exitCode)
	}
	
	if !strings.Contains(stderr, "mutually exclusive") {
		t.Errorf("Expected error about mutually exclusive flags, got: %s", stderr)
	}
	
	if !strings.Contains(stderr, "verbose") || !strings.Contains(stderr, "quiet") {
		t.Errorf("Expected error to mention both flags, got: %s", stderr)
	}
}

// Test mutually exclusive flags: --max-iterations and --unlimited
func TestMutuallyExclusiveMaxIterUnlimited(t *testing.T) {
	_, stderr, exitCode := runRooda(t, "build", "--max-iterations", "5", "--unlimited")
	
	if exitCode != ExitUserError {
		t.Errorf("Expected exit code %d, got %d", ExitUserError, exitCode)
	}
	
	if !strings.Contains(stderr, "mutually exclusive") {
		t.Errorf("Expected error about mutually exclusive flags, got: %s", stderr)
	}
	
	if !strings.Contains(stderr, "max-iterations") || !strings.Contains(stderr, "unlimited") {
		t.Errorf("Expected error to mention both flags, got: %s", stderr)
	}
}

// Test invalid --max-iterations value (< 1)
func TestInvalidMaxIterations(t *testing.T) {
	_, stderr, exitCode := runRooda(t, "build", "--max-iterations", "0")
	
	if exitCode != ExitUserError {
		t.Errorf("Expected exit code %d, got %d", ExitUserError, exitCode)
	}
	
	if !strings.Contains(stderr, "max-iterations must be >= 1") {
		t.Errorf("Expected error about max-iterations constraint, got: %s", stderr)
	}
}

// Test procedure-specific help
func TestProcedureHelp(t *testing.T) {
	stdout, _, exitCode := runRooda(t, "build", "--help")
	
	if exitCode != ExitSuccess {
		t.Errorf("Expected exit code %d, got %d", ExitSuccess, exitCode)
	}
	
	if !strings.Contains(stdout, "build") {
		t.Errorf("Expected help to mention procedure name, got: %s", stdout)
	}
}

// Test dry-run mode with valid procedure
func TestDryRunMode(t *testing.T) {
	// Create a minimal config file for testing
	tmpDir := t.TempDir()
	configPath := tmpDir + "/rooda-config.yml"
	configContent := `
loop:
  ai_cmd: "echo"
procedures:
  test-proc:
    display: "Test Procedure"
    observe:
      - content: "Test observe"
    orient:
      - content: "Test orient"
    decide:
      - content: "Test decide"
    act:
      - content: "Test act"
`
	if err := os.WriteFile(configPath, []byte(configContent), 0644); err != nil {
		t.Fatalf("Failed to create test config: %v", err)
	}
	
	stdout, _, exitCode := runRooda(t, "test-proc", "--dry-run", "--config", configPath)
	
	if exitCode != ExitSuccess {
		t.Errorf("Expected exit code %d for valid dry-run, got %d", ExitSuccess, exitCode)
	}
	
	// Dry-run should show the assembled prompt
	if !strings.Contains(stdout, "OBSERVE") || !strings.Contains(stdout, "ORIENT") {
		t.Errorf("Expected dry-run to show OODA phases, got: %s", stdout)
	}
}

// Test short flags
func TestShortFlags(t *testing.T) {
	tests := []struct {
		name      string
		args      []string
		shouldErr bool
	}{
		{"verbose short", []string{"build", "-v", "--dry-run"}, false},
		{"quiet short", []string{"build", "-q", "--dry-run"}, false},
		{"help short", []string{"-h"}, false},
		{"dry-run short", []string{"build", "-d"}, false},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, _, exitCode := runRooda(t, tt.args...)
			
			if tt.shouldErr && exitCode == ExitSuccess {
				t.Errorf("Expected error for %v, but got success", tt.args)
			}
			if !tt.shouldErr && exitCode != ExitSuccess && exitCode != ExitConfigError {
				// ExitConfigError is acceptable for some tests (missing AI command)
				t.Errorf("Expected success for %v, but got exit code %d", tt.args, exitCode)
			}
		})
	}
}

// Test flag formats: --flag=value and --flag value
func TestFlagFormats(t *testing.T) {
	tests := []struct {
		name string
		args []string
	}{
		{"equals format", []string{"build", "--max-iterations=5", "--dry-run"}},
		{"space format", []string{"build", "--max-iterations", "5", "--dry-run"}},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, _, exitCode := runRooda(t, tt.args...)
			
			// Both formats should parse successfully (may fail later due to missing config)
			if exitCode == ExitUserError {
				t.Errorf("Flag parsing failed for %v", tt.args)
			}
		})
	}
}

// Test context flag accumulation
func TestContextAccumulation(t *testing.T) {
	// This test verifies that multiple --context flags are accepted
	// We can't easily verify they're all used without mocking, but we can verify parsing
	_, _, exitCode := runRooda(t, "build", "--context", "first", "--context", "second", "--dry-run")
	
	// Should not fail with user error (may fail with config error)
	if exitCode == ExitUserError {
		t.Errorf("Multiple --context flags should be accepted")
	}
}

// Test OODA phase override flags
func TestOODAPhaseOverrides(t *testing.T) {
	tests := []struct {
		name string
		flag string
	}{
		{"observe", "--observe"},
		{"orient", "--orient"},
		{"decide", "--decide"},
		{"act", "--act"},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, _, exitCode := runRooda(t, "build", tt.flag, "custom content", "--dry-run")
			
			// Should not fail with user error (may fail with config error)
			if exitCode == ExitUserError {
				t.Errorf("OODA phase override flag %s should be accepted", tt.flag)
			}
		})
	}
}

// Test empty inline content rejection
func TestEmptyInlineContent(t *testing.T) {
	_, stderr, exitCode := runRooda(t, "build", "--observe", "", "--dry-run", "--ai-cmd", "echo")
	
	if exitCode != ExitUserError {
		t.Errorf("Expected exit code %d for empty inline content, got %d", ExitUserError, exitCode)
	}
	
	if !strings.Contains(stderr, "Empty inline content") || !strings.Contains(stderr, "observe") {
		t.Errorf("Expected error about empty inline content, got: %s", stderr)
	}
}

// Test invalid log level
func TestInvalidLogLevel(t *testing.T) {
	_, stderr, exitCode := runRooda(t, "build", "--log-level", "invalid", "--ai-cmd", "echo")
	
	if exitCode != ExitUserError {
		t.Errorf("Expected exit code %d for invalid log level, got %d", ExitUserError, exitCode)
	}
	
	if !strings.Contains(stderr, "Invalid log level") {
		t.Errorf("Expected error about invalid log level, got: %s", stderr)
	}
}
