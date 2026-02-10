package config

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// TestBuiltInDefaults verifies built-in defaults are loaded correctly
func TestBuiltInDefaults(t *testing.T) {
	config := builtInDefaults()

	if config.Loop.IterationMode != DefaultIterationMode {
		t.Errorf("expected iteration_mode %s, got %s", DefaultIterationMode, config.Loop.IterationMode)
	}
	if *config.Loop.DefaultMaxIterations != DefaultMaxIterations {
		t.Errorf("expected default_max_iterations %d, got %d", DefaultMaxIterations, *config.Loop.DefaultMaxIterations)
	}
	if config.Loop.MaxOutputBuffer != DefaultMaxOutputBuffer {
		t.Errorf("expected max_output_buffer %d, got %d", DefaultMaxOutputBuffer, config.Loop.MaxOutputBuffer)
	}
	if config.Loop.FailureThreshold != DefaultFailureThreshold {
		t.Errorf("expected failure_threshold %d, got %d", DefaultFailureThreshold, config.Loop.FailureThreshold)
	}
	if config.Loop.LogLevel != DefaultLogLevel {
		t.Errorf("expected log_level %s, got %s", DefaultLogLevel, config.Loop.LogLevel)
	}
	if config.Loop.LogTimestampFormat != DefaultTimestampFormat {
		t.Errorf("expected log_timestamp_format %s, got %s", DefaultTimestampFormat, config.Loop.LogTimestampFormat)
	}
	if config.Loop.ShowAIOutput != DefaultShowAIOutput {
		t.Errorf("expected show_ai_output %v, got %v", DefaultShowAIOutput, config.Loop.ShowAIOutput)
	}

	// Verify built-in aliases exist
	aliases := builtInAliases()
	expectedAliases := []string{"kiro-cli", "claude", "copilot", "cursor-agent"}
	for _, name := range expectedAliases {
		if _, exists := aliases[name]; !exists {
			t.Errorf("expected built-in alias %s to exist", name)
		}
	}
}

// TestLoadConfigNoFiles verifies zero-config startup with built-in defaults
func TestLoadConfigNoFiles(t *testing.T) {
	// Create temp directory with no config files
	tmpDir := t.TempDir()
	origDir, _ := os.Getwd()
	defer os.Chdir(origDir)
	os.Chdir(tmpDir)

	config, err := LoadConfig(CLIFlags{})
	if err != nil {
		t.Fatalf("LoadConfig failed: %v", err)
	}

	// Should have built-in defaults
	if config.Loop.IterationMode != DefaultIterationMode {
		t.Errorf("expected default iteration_mode, got %s", config.Loop.IterationMode)
	}
	if *config.Loop.DefaultMaxIterations != DefaultMaxIterations {
		t.Errorf("expected default max_iterations %d, got %d", DefaultMaxIterations, *config.Loop.DefaultMaxIterations)
	}
}

// TestLoadConfigWorkspaceFile verifies workspace config loading
func TestLoadConfigWorkspaceFile(t *testing.T) {
	tmpDir := t.TempDir()
	origDir, _ := os.Getwd()
	defer os.Chdir(origDir)
	os.Chdir(tmpDir)

	// Create workspace config
	configYAML := `loop:
  default_max_iterations: 10
  failure_threshold: 5
`
	os.WriteFile("rooda-config.yml", []byte(configYAML), 0644)

	config, err := LoadConfig(CLIFlags{})
	if err != nil {
		t.Fatalf("LoadConfig failed: %v", err)
	}

	if *config.Loop.DefaultMaxIterations != 10 {
		t.Errorf("expected max_iterations 10, got %d", *config.Loop.DefaultMaxIterations)
	}
	if config.Loop.FailureThreshold != 5 {
		t.Errorf("expected failure_threshold 5, got %d", config.Loop.FailureThreshold)
	}
}

// TestLoadConfigCustomPath verifies --config flag
func TestLoadConfigCustomPath(t *testing.T) {
	tmpDir := t.TempDir()
	origDir, _ := os.Getwd()
	defer os.Chdir(origDir)
	os.Chdir(tmpDir)

	// Create custom config file
	customPath := filepath.Join(tmpDir, "custom-config.yml")
	configYAML := `loop:
  default_max_iterations: 15
`
	os.WriteFile(customPath, []byte(configYAML), 0644)

	config, err := LoadConfig(CLIFlags{ConfigPath: customPath})
	if err != nil {
		t.Fatalf("LoadConfig failed: %v", err)
	}

	if *config.Loop.DefaultMaxIterations != 15 {
		t.Errorf("expected max_iterations 15, got %d", *config.Loop.DefaultMaxIterations)
	}
}

// TestLoadConfigEnvVars verifies environment variable overrides
func TestLoadConfigEnvVars(t *testing.T) {
	tmpDir := t.TempDir()
	origDir, _ := os.Getwd()
	defer os.Chdir(origDir)
	os.Chdir(tmpDir)

	// Set environment variables
	os.Setenv("ROODA_LOOP_DEFAULT_MAX_ITERATIONS", "20")
	os.Setenv("ROODA_LOOP_ITERATION_MODE", "unlimited")
	os.Setenv("ROODA_LOOP_AI_CMD", "test-ai-cmd")
	defer func() {
		os.Unsetenv("ROODA_LOOP_DEFAULT_MAX_ITERATIONS")
		os.Unsetenv("ROODA_LOOP_ITERATION_MODE")
		os.Unsetenv("ROODA_LOOP_AI_CMD")
	}()

	config, err := LoadConfig(CLIFlags{})
	if err != nil {
		t.Fatalf("LoadConfig failed: %v", err)
	}

	if *config.Loop.DefaultMaxIterations != 20 {
		t.Errorf("expected max_iterations 20 from env, got %d", *config.Loop.DefaultMaxIterations)
	}
	if config.Loop.IterationMode != "unlimited" {
		t.Errorf("expected iteration_mode unlimited from env, got %s", config.Loop.IterationMode)
	}
	if config.Loop.AICmd != "test-ai-cmd" {
		t.Errorf("expected ai_cmd from env, got %s", config.Loop.AICmd)
	}
}

// TestLoadConfigCLIFlags verifies CLI flag overrides (highest precedence)
func TestLoadConfigCLIFlags(t *testing.T) {
	tmpDir := t.TempDir()
	origDir, _ := os.Getwd()
	defer os.Chdir(origDir)
	os.Chdir(tmpDir)

	// Create workspace config
	configYAML := `loop:
  default_max_iterations: 10
  ai_cmd: "workspace-cmd"
`
	os.WriteFile("rooda-config.yml", []byte(configYAML), 0644)

	// Set env var
	os.Setenv("ROODA_LOOP_DEFAULT_MAX_ITERATIONS", "20")
	defer os.Unsetenv("ROODA_LOOP_DEFAULT_MAX_ITERATIONS")

	// CLI flags should override everything
	maxIter := 30
	config, err := LoadConfig(CLIFlags{
		MaxIterations: &maxIter,
		AICmd:         "cli-cmd",
		AICmdAlias:    "kiro-cli",
	})
	if err != nil {
		t.Fatalf("LoadConfig failed: %v", err)
	}

	if *config.Loop.DefaultMaxIterations != 30 {
		t.Errorf("expected max_iterations 30 from CLI, got %d", *config.Loop.DefaultMaxIterations)
	}
	if config.Loop.AICmd != "cli-cmd" {
		t.Errorf("expected ai_cmd from CLI, got %s", config.Loop.AICmd)
	}
	if config.Loop.AICmdAlias != "kiro-cli" {
		t.Errorf("expected ai_cmd_alias from CLI, got %s", config.Loop.AICmdAlias)
	}
}

// TestPrecedenceOrder verifies full precedence chain
func TestPrecedenceOrder(t *testing.T) {
	tmpDir := t.TempDir()
	origDir, _ := os.Getwd()
	defer os.Chdir(origDir)
	os.Chdir(tmpDir)

	// Workspace config
	configYAML := `loop:
  default_max_iterations: 10
  failure_threshold: 5
  ai_cmd: "workspace-cmd"
`
	os.WriteFile("rooda-config.yml", []byte(configYAML), 0644)

	// Env var overrides workspace for max_iterations
	os.Setenv("ROODA_LOOP_DEFAULT_MAX_ITERATIONS", "20")
	defer os.Unsetenv("ROODA_LOOP_DEFAULT_MAX_ITERATIONS")

	// CLI flag overrides env var
	maxIter := 30
	config, err := LoadConfig(CLIFlags{MaxIterations: &maxIter})
	if err != nil {
		t.Fatalf("LoadConfig failed: %v", err)
	}

	// CLI flag wins
	if *config.Loop.DefaultMaxIterations != 30 {
		t.Errorf("expected CLI flag to win, got %d", *config.Loop.DefaultMaxIterations)
	}
	// Workspace config wins (no override)
	if config.Loop.FailureThreshold != 5 {
		t.Errorf("expected workspace config to win, got %d", config.Loop.FailureThreshold)
	}
	if config.Loop.AICmd != "workspace-cmd" {
		t.Errorf("expected workspace ai_cmd, got %s", config.Loop.AICmd)
	}
}

// TestMergeAICmdAliases verifies AI command aliases merge correctly
func TestMergeAICmdAliases(t *testing.T) {
	tmpDir := t.TempDir()
	origDir, _ := os.Getwd()
	defer os.Chdir(origDir)
	os.Chdir(tmpDir)

	// Workspace config adds custom alias
	configYAML := `ai_cmd_aliases:
  custom: "custom-ai-cmd"
  kiro-cli: "overridden-kiro-cmd"
`
	os.WriteFile("rooda-config.yml", []byte(configYAML), 0644)

	config, err := LoadConfig(CLIFlags{})
	if err != nil {
		t.Fatalf("LoadConfig failed: %v", err)
	}

	// Custom alias added
	if config.AICmdAliases["custom"] != "custom-ai-cmd" {
		t.Errorf("expected custom alias, got %s", config.AICmdAliases["custom"])
	}
	// Built-in alias overridden
	if config.AICmdAliases["kiro-cli"] != "overridden-kiro-cmd" {
		t.Errorf("expected overridden kiro-cli, got %s", config.AICmdAliases["kiro-cli"])
	}
	// Other built-in aliases still exist
	if _, exists := config.AICmdAliases["claude"]; !exists {
		t.Error("expected claude alias to still exist")
	}
}

// TestMergeProcedures verifies procedure merging
func TestMergeProcedures(t *testing.T) {
	tmpDir := t.TempDir()
	origDir, _ := os.Getwd()
	defer os.Chdir(origDir)
	os.Chdir(tmpDir)

	// Workspace config adds custom procedure
	configYAML := `procedures:
  custom-proc:
    display: "Custom Procedure"
    observe:
      - path: "prompts/observe.md"
    orient:
      - path: "prompts/orient.md"
    decide:
      - path: "prompts/decide.md"
    act:
      - path: "prompts/act.md"
    default_max_iterations: 3
`
	os.WriteFile("rooda-config.yml", []byte(configYAML), 0644)

	config, err := LoadConfig(CLIFlags{})
	if err != nil {
		t.Fatalf("LoadConfig failed: %v", err)
	}

	// Custom procedure added
	proc, exists := config.Procedures["custom-proc"]
	if !exists {
		t.Fatal("expected custom-proc to exist")
	}
	if proc.Display != "Custom Procedure" {
		t.Errorf("expected display name, got %s", proc.Display)
	}
	if *proc.DefaultMaxIterations != 3 {
		t.Errorf("expected max_iterations 3, got %d", *proc.DefaultMaxIterations)
	}
}

// TestResolveGlobalConfigDir verifies global config directory resolution
func TestResolveGlobalConfigDir(t *testing.T) {
	// Test ROODA_CONFIG_HOME
	os.Setenv("ROODA_CONFIG_HOME", "/custom/rooda")
	defer os.Unsetenv("ROODA_CONFIG_HOME")
	dir := resolveGlobalConfigDir()
	if dir != "/custom/rooda" {
		t.Errorf("expected ROODA_CONFIG_HOME to win, got %s", dir)
	}

	// Test XDG_CONFIG_HOME
	os.Unsetenv("ROODA_CONFIG_HOME")
	os.Setenv("XDG_CONFIG_HOME", "/custom/xdg")
	defer os.Unsetenv("XDG_CONFIG_HOME")
	dir = resolveGlobalConfigDir()
	if dir != "/custom/xdg/rooda" {
		t.Errorf("expected XDG_CONFIG_HOME/rooda, got %s", dir)
	}
}

// TestInvalidYAML verifies error handling for invalid YAML
func TestInvalidYAML(t *testing.T) {
	tmpDir := t.TempDir()
	origDir, _ := os.Getwd()
	defer os.Chdir(origDir)
	os.Chdir(tmpDir)

	// Create invalid YAML
	invalidYAML := `loop:
  default_max_iterations: [invalid
`
	os.WriteFile("rooda-config.yml", []byte(invalidYAML), 0644)

	_, err := LoadConfig(CLIFlags{})
	if err == nil {
		t.Fatal("expected error for invalid YAML")
	}
}

// TestMissingConfigFile verifies missing files are silently skipped
func TestMissingConfigFile(t *testing.T) {
	tmpDir := t.TempDir()
	origDir, _ := os.Getwd()
	defer os.Chdir(origDir)
	os.Chdir(tmpDir)

	// No config file exists
	config, err := LoadConfig(CLIFlags{})
	if err != nil {
		t.Fatalf("expected no error for missing config, got %v", err)
	}

	// Should have built-in defaults
	if config.Loop.IterationMode != DefaultIterationMode {
		t.Error("expected built-in defaults when config missing")
	}
}

// TestProvenance verifies provenance tracking
func TestProvenance(t *testing.T) {
	tmpDir := t.TempDir()
	origDir, _ := os.Getwd()
	defer os.Chdir(origDir)
	os.Chdir(tmpDir)

	// Workspace config
	configYAML := `loop:
  default_max_iterations: 10
`
	os.WriteFile("rooda-config.yml", []byte(configYAML), 0644)

	// Env var
	os.Setenv("ROODA_LOOP_FAILURE_THRESHOLD", "7")
	defer os.Unsetenv("ROODA_LOOP_FAILURE_THRESHOLD")

	// CLI flag
	maxIter := 30
	config, err := LoadConfig(CLIFlags{MaxIterations: &maxIter})
	if err != nil {
		t.Fatalf("LoadConfig failed: %v", err)
	}

	// Check provenance
	if src, exists := config.Provenance["loop.default_max_iterations"]; !exists {
		t.Error("expected provenance for max_iterations")
	} else if src.Tier != TierCLIFlag {
		t.Errorf("expected CLI flag tier, got %s", src.Tier)
	}

	if src, exists := config.Provenance["loop.failure_threshold"]; !exists {
		t.Error("expected provenance for failure_threshold")
	} else if src.Tier != TierEnvVar {
		t.Errorf("expected env var tier, got %s", src.Tier)
	}

	if src, exists := config.Provenance["loop.iteration_mode"]; !exists {
		t.Error("expected provenance for iteration_mode")
	} else if src.Tier != TierBuiltIn {
		t.Errorf("expected built-in tier, got %s", src.Tier)
	}
}

// TestBackwardCompatibility_V01StringFormat verifies v0.1.0 string format loads correctly
func TestBackwardCompatibility_V01StringFormat(t *testing.T) {
	tmpDir := t.TempDir()
	configPath := filepath.Join(tmpDir, "rooda-config.yml")

	// Write v0.1.0 format config (string paths instead of arrays)
	configContent := `
procedures:
  bootstrap:
    display: "Bootstrap Repository"
    observe: prompts/observe_bootstrap.md
    orient: prompts/orient_bootstrap.md
    decide: prompts/decide_bootstrap.md
    act: prompts/act_bootstrap.md
    default_max_iterations: 1
`
	if err := os.WriteFile(configPath, []byte(configContent), 0644); err != nil {
		t.Fatal(err)
	}

	origDir, _ := os.Getwd()
	defer os.Chdir(origDir)
	os.Chdir(tmpDir)

	config, err := LoadConfig(CLIFlags{})
	if err != nil {
		t.Fatalf("LoadConfig failed: %v", err)
	}

	proc, exists := config.Procedures["bootstrap"]
	if !exists {
		t.Fatal("expected bootstrap procedure to exist")
	}

	// Verify string paths were converted to FragmentAction arrays
	if len(proc.Observe) != 1 {
		t.Errorf("expected 1 observe fragment, got %d", len(proc.Observe))
	}
	if !strings.HasSuffix(proc.Observe[0].Path, "prompts/observe_bootstrap.md") {
		t.Errorf("expected observe path to end with prompts/observe_bootstrap.md, got %s", proc.Observe[0].Path)
	}

	if len(proc.Orient) != 1 {
		t.Errorf("expected 1 orient fragment, got %d", len(proc.Orient))
	}
	if !strings.HasSuffix(proc.Orient[0].Path, "prompts/orient_bootstrap.md") {
		t.Errorf("expected orient path to end with prompts/orient_bootstrap.md, got %s", proc.Orient[0].Path)
	}

	if len(proc.Decide) != 1 {
		t.Errorf("expected 1 decide fragment, got %d", len(proc.Decide))
	}
	if !strings.HasSuffix(proc.Decide[0].Path, "prompts/decide_bootstrap.md") {
		t.Errorf("expected decide path to end with prompts/decide_bootstrap.md, got %s", proc.Decide[0].Path)
	}

	if len(proc.Act) != 1 {
		t.Errorf("expected 1 act fragment, got %d", len(proc.Act))
	}
	if !strings.HasSuffix(proc.Act[0].Path, "prompts/act_bootstrap.md") {
		t.Errorf("expected act path to end with prompts/act_bootstrap.md, got %s", proc.Act[0].Path)
	}

	if *proc.DefaultMaxIterations != 1 {
		t.Errorf("expected default_max_iterations 1, got %d", *proc.DefaultMaxIterations)
	}
}

// TestBackwardCompatibility_V2ArrayFormat verifies v2 array format loads correctly
func TestBackwardCompatibility_V2ArrayFormat(t *testing.T) {
	tmpDir := t.TempDir()
	configPath := filepath.Join(tmpDir, "rooda-config.yml")

	// Write v2 format config (array of FragmentAction)
	configContent := `
procedures:
  build:
    display: "Build from Plan"
    observe:
      - path: prompts/observe_plan.md
      - path: prompts/observe_specs.md
    orient:
      - path: prompts/orient_build.md
    decide:
      - path: prompts/decide_build.md
    act:
      - path: prompts/act_build.md
    default_max_iterations: 5
`
	if err := os.WriteFile(configPath, []byte(configContent), 0644); err != nil {
		t.Fatal(err)
	}

	origDir, _ := os.Getwd()
	defer os.Chdir(origDir)
	os.Chdir(tmpDir)

	config, err := LoadConfig(CLIFlags{})
	if err != nil {
		t.Fatalf("LoadConfig failed: %v", err)
	}

	proc, exists := config.Procedures["build"]
	if !exists {
		t.Fatal("expected build procedure to exist")
	}

	// Verify array format loaded correctly
	if len(proc.Observe) != 2 {
		t.Errorf("expected 2 observe fragments, got %d", len(proc.Observe))
	}
	if !strings.HasSuffix(proc.Observe[0].Path, "prompts/observe_plan.md") {
		t.Errorf("expected first observe path to end with prompts/observe_plan.md, got %s", proc.Observe[0].Path)
	}
	if !strings.HasSuffix(proc.Observe[1].Path, "prompts/observe_specs.md") {
		t.Errorf("expected second observe path to end with prompts/observe_specs.md, got %s", proc.Observe[1].Path)
	}

	if len(proc.Orient) != 1 {
		t.Errorf("expected 1 orient fragment, got %d", len(proc.Orient))
	}

	if *proc.DefaultMaxIterations != 5 {
		t.Errorf("expected default_max_iterations 5, got %d", *proc.DefaultMaxIterations)
	}
}

// TestBackwardCompatibility_MixedFormat verifies mixed v0.1.0 and v2 formats work
func TestBackwardCompatibility_MixedFormat(t *testing.T) {
	tmpDir := t.TempDir()
	configPath := filepath.Join(tmpDir, "rooda-config.yml")

	// Write config with both formats
	configContent := `
procedures:
  bootstrap:
    observe: prompts/observe_bootstrap.md
    orient: prompts/orient_bootstrap.md
    decide: prompts/decide_bootstrap.md
    act: prompts/act_bootstrap.md
  build:
    observe:
      - path: prompts/observe_plan.md
      - path: prompts/observe_specs.md
    orient:
      - path: prompts/orient_build.md
    decide:
      - path: prompts/decide_build.md
    act:
      - path: prompts/act_build.md
`
	if err := os.WriteFile(configPath, []byte(configContent), 0644); err != nil {
		t.Fatal(err)
	}

	origDir, _ := os.Getwd()
	defer os.Chdir(origDir)
	os.Chdir(tmpDir)

	config, err := LoadConfig(CLIFlags{})
	if err != nil {
		t.Fatalf("LoadConfig failed: %v", err)
	}

	// Verify v0.1.0 format procedure
	bootstrap, exists := config.Procedures["bootstrap"]
	if !exists {
		t.Fatal("expected bootstrap procedure to exist")
	}
	if len(bootstrap.Observe) != 1 {
		t.Errorf("expected 1 observe fragment for bootstrap, got %d", len(bootstrap.Observe))
	}

	// Verify v2 format procedure
	build, exists := config.Procedures["build"]
	if !exists {
		t.Fatal("expected build procedure to exist")
	}
	if len(build.Observe) != 2 {
		t.Errorf("expected 2 observe fragments for build, got %d", len(build.Observe))
	}
}
