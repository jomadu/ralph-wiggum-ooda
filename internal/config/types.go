package config

// LogLevel defines the verbosity of loop logging.
type LogLevel string

const (
	LogLevelDebug LogLevel = "debug"
	LogLevelInfo  LogLevel = "info"
	LogLevelWarn  LogLevel = "warn"
	LogLevelError LogLevel = "error"
)

// TimestampFormat defines the format for log timestamps.
type TimestampFormat string

const (
	TimestampTime     TimestampFormat = "time"     // [HH:MM:SS.mmm]
	TimestampTimeMs   TimestampFormat = "time-ms"  // [HH:MM:SS.mmm] (alias for time)
	TimestampRelative TimestampFormat = "relative" // [+0.123s] relative to loop start
	TimestampISO      TimestampFormat = "iso"      // 2026-02-08T20:59:35.877-08:00
	TimestampNone     TimestampFormat = "none"     // No timestamp
)

// IterationMode controls whether iterations are limited or unlimited.
type IterationMode string

const (
	ModeMaxIterations IterationMode = "max-iterations" // Run up to DefaultMaxIterations
	ModeUnlimited     IterationMode = "unlimited"      // Run until SUCCESS signal, failure threshold, or Ctrl+C
)

// ConfigTier identifies which configuration source provided a value.
type ConfigTier string

const (
	TierBuiltIn   ConfigTier = "built-in"
	TierGlobal    ConfigTier = "global"    // <config_dir>/rooda-config.yml
	TierWorkspace ConfigTier = "workspace" // ./rooda-config.yml
	TierEnvVar    ConfigTier = "env"       // ROODA_* environment variables
	TierCLIFlag   ConfigTier = "cli"       // --flag values
)

// Built-in defaults
const (
	DefaultMaxIterations  = 5
	DefaultMaxOutputBuffer = 10485760 // 10MB
	DefaultFailureThreshold = 3
)

var (
	DefaultLogLevel          = LogLevelInfo
	DefaultTimestampFormat   = TimestampTime
	DefaultIterationMode     = ModeMaxIterations
	DefaultShowAIOutput      = false
)

// FragmentAction specifies a prompt fragment with optional inline content or file path.
type FragmentAction struct {
	Content    string                 // Inline prompt content (optional)
	Path       string                 // Path to fragment file (optional)
	Parameters map[string]interface{} // Template parameters (optional)
}

// Procedure defines an OODA loop procedure with fragments for each phase.
type Procedure struct {
	Display              string           // Human-readable name (optional)
	Summary              string           // One-line description (optional)
	Description          string           // Detailed description (optional)
	Observe              []FragmentAction // Array of observe phase fragments
	Orient               []FragmentAction // Array of orient phase fragments
	Decide               []FragmentAction // Array of decide phase fragments
	Act                  []FragmentAction // Array of act phase fragments
	IterationMode        IterationMode    // Override loop iteration mode ("" = inherit from loop)
	DefaultMaxIterations *int             // Override loop.default_max_iterations (nil = inherit from loop). Must be >= 1 when set.
	IterationTimeout     *int             // Override loop.iteration_timeout (nil = inherit from loop). Must be >= 1 when set. Seconds.
	MaxOutputBuffer      *int             // Override loop.max_output_buffer (nil = inherit from loop). Must be >= 1024 when set. Bytes.
	AICmd                string           // Override AI command for this procedure (optional)
	AICmdAlias           string           // Override AI command alias for this procedure (optional)
}

// LoopConfig defines global loop settings.
type LoopConfig struct {
	IterationMode        IterationMode   // Iteration mode (built-in default: ModeMaxIterations)
	DefaultMaxIterations *int            // Global default (built-in default: 5). Must be >= 1 when set. nil = not set (inherit).
	IterationTimeout     *int            // Per-iteration timeout in seconds (built-in default: nil). nil = no timeout.
	MaxOutputBuffer      int             // Max AI CLI output buffer in bytes (built-in default: 10485760 = 10MB). Must be >= 1024.
	FailureThreshold     int             // Consecutive failures before abort (built-in default: 3)
	LogLevel             LogLevel        // Loop log level (built-in default: LogLevelInfo)
	LogTimestampFormat   TimestampFormat // Log timestamp format (built-in default: TimestampTime)
	ShowAIOutput         bool            // Stream AI CLI output to terminal (built-in default: false)
	AICmd                string          // Default AI command (direct command string, optional)
	AICmdAlias           string          // Default AI command alias name (resolved from AICmdAliases, optional)
}

// ConfigSource tracks which tier provided a configuration value.
type ConfigSource struct {
	Tier  ConfigTier // Which tier provided this value
	File  string     // File path (for workspace/global tiers) or "" for built-in/env/cli
	Value any        // The resolved value
}

// Config is the fully resolved configuration after merging all tiers.
type Config struct {
	Loop         LoopConfig              // Global loop settings
	Procedures   map[string]Procedure    // Named procedure definitions
	AICmdAliases map[string]string       // AI command alias name -> command string
	Provenance   map[string]ConfigSource // Setting path -> source that provided it
}

// AICommand represents a resolved AI command with provenance.
type AICommand struct {
	Command string // Full command string to execute
	Source  string // Provenance: where this command came from
}
