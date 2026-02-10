package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/jomadu/rooda/internal/config"
)

var (
	Version   = "dev"
	CommitSHA = "unknown"
	BuildDate = "unknown"
)

const (
	ExitSuccess        = 0
	ExitUserError      = 1
	ExitConfigError    = 2
	ExitExecutionError = 3
)

func main() {
	flags := parseFlags()

	if flags.ShowVersion {
		fmt.Printf("rooda %s\n", Version)
		fmt.Printf("Commit: %s\n", CommitSHA)
		fmt.Printf("Built: %s\n", BuildDate)
		os.Exit(ExitSuccess)
	}

	if flags.ShowHelp {
		if flags.ProcedureName == "" {
			printGlobalHelp()
		} else {
			printProcedureHelp(flags.ProcedureName)
		}
		os.Exit(ExitSuccess)
	}

	if flags.ListProcedures {
		listProcedures(flags)
		os.Exit(ExitSuccess)
	}

	// Validate mutually exclusive flags before checking procedure name
	if flags.Verbose && flags.Quiet {
		fmt.Fprintln(os.Stderr, "Error: --verbose and --quiet are mutually exclusive.")
		os.Exit(ExitUserError)
	}

	if flags.MaxIterations != nil && flags.Unlimited {
		fmt.Fprintln(os.Stderr, "Error: --max-iterations and --unlimited are mutually exclusive.")
		os.Exit(ExitUserError)
	}

	// Validate max iterations value
	if flags.MaxIterations != nil && *flags.MaxIterations < 1 {
		fmt.Fprintln(os.Stderr, "Error: --max-iterations must be >= 1.")
		os.Exit(ExitUserError)
	}

	if flags.ProcedureName == "" {
		fmt.Fprintln(os.Stderr, "Error: No procedure specified. Run 'rooda --help' for usage.")
		os.Exit(ExitUserError)
	}

	fmt.Fprintln(os.Stderr, "rooda: OODA loop implementation not yet available")
	os.Exit(ExitExecutionError)
}

func parseFlags() config.CLIFlags {
	var flags config.CLIFlags
	var maxIterLong, maxIterShort int

	// Define flags
	fs := flag.NewFlagSet("rooda", flag.ContinueOnError)
	fs.SetOutput(os.Stderr)

	// Info flags
	fs.BoolVar(&flags.ShowVersion, "version", false, "Print version information")
	fs.BoolVar(&flags.ShowHelp, "help", false, "Show help")
	fs.BoolVar(&flags.ListProcedures, "list-procedures", false, "List all available procedures")

	// Loop control flags
	fs.IntVar(&maxIterLong, "max-iterations", -1, "Maximum iterations (>= 1)")
	fs.IntVar(&maxIterShort, "n", -1, "Maximum iterations (short form)")
	fs.BoolVar(&flags.Unlimited, "unlimited", false, "Run unlimited iterations")
	fs.BoolVar(&flags.Unlimited, "u", false, "Run unlimited iterations (short form)")
	fs.BoolVar(&flags.DryRun, "dry-run", false, "Validate without executing")
	fs.BoolVar(&flags.DryRun, "d", false, "Validate without executing (short form)")

	// AI command flags
	fs.StringVar(&flags.AICmd, "ai-cmd", "", "Override AI command")
	fs.StringVar(&flags.AICmdAlias, "ai-cmd-alias", "", "Override AI command using alias")

	// Output control flags
	fs.BoolVar(&flags.Verbose, "verbose", false, "Enable verbose output")
	fs.BoolVar(&flags.Verbose, "v", false, "Enable verbose output (short form)")
	fs.BoolVar(&flags.Quiet, "quiet", false, "Suppress non-error output")
	fs.BoolVar(&flags.Quiet, "q", false, "Suppress non-error output (short form)")
	fs.StringVar(&flags.LogLevel, "log-level", "", "Set log level (debug, info, warn, error)")

	// Configuration flags
	fs.StringVar(&flags.ConfigPath, "config", "", "Alternate workspace config file path")

	// Custom flag parsing for repeatable flags
	fs.Func("context", "Context file path or inline text (repeatable)", func(s string) error {
		flags.Contexts = append(flags.Contexts, s)
		return nil
	})
	fs.Func("c", "Context file path or inline text (short form, repeatable)", func(s string) error {
		flags.Contexts = append(flags.Contexts, s)
		return nil
	})

	fs.Func("observe", "Observe phase fragment (repeatable)", func(s string) error {
		flags.ObserveFragments = append(flags.ObserveFragments, s)
		return nil
	})
	fs.Func("orient", "Orient phase fragment (repeatable)", func(s string) error {
		flags.OrientFragments = append(flags.OrientFragments, s)
		return nil
	})
	fs.Func("decide", "Decide phase fragment (repeatable)", func(s string) error {
		flags.DecideFragments = append(flags.DecideFragments, s)
		return nil
	})
	fs.Func("act", "Act phase fragment (repeatable)", func(s string) error {
		flags.ActFragments = append(flags.ActFragments, s)
		return nil
	})

	// Parse flags
	if err := fs.Parse(os.Args[1:]); err != nil {
		if err == flag.ErrHelp {
			flags.ShowHelp = true
			return flags
		}
		os.Exit(ExitUserError)
	}

	// Handle max-iterations from both long and short forms
	if maxIterLong >= 0 {
		flags.MaxIterations = &maxIterLong
	} else if maxIterShort >= 0 {
		flags.MaxIterations = &maxIterShort
	}

	// Get procedure name from remaining args
	if fs.NArg() > 0 {
		flags.ProcedureName = fs.Arg(0)
	}

	return flags
}

func printGlobalHelp() {
	fmt.Println(`rooda - OODA Loop Framework

USAGE:
  rooda <procedure> [flags]
  rooda --help
  rooda --version
  rooda --list-procedures

LOOP CONTROL FLAGS:
  -n, --max-iterations <n>    Maximum iterations (>= 1)
  -u, --unlimited              Run unlimited iterations
  -d, --dry-run                Validate without executing

AI COMMAND FLAGS:
  --ai-cmd <command>           Override AI command
  --ai-cmd-alias <alias>       Override AI command using alias

PROMPT OVERRIDE FLAGS:
  --observe <value>            Observe phase fragment (repeatable)
  --orient <value>             Orient phase fragment (repeatable)
  --decide <value>             Decide phase fragment (repeatable)
  --act <value>                Act phase fragment (repeatable)

OUTPUT CONTROL FLAGS:
  -v, --verbose                Enable verbose output
  -q, --quiet                  Suppress non-error output
  --log-level <level>          Set log level (debug, info, warn, error)

CONFIGURATION FLAGS:
  --config <path>              Alternate workspace config file path
  -c, --context <value>        Context file path or inline text (repeatable)

INFO FLAGS:
  --help                       Show this help
  --version                    Print version information
  --list-procedures            List all available procedures

EXAMPLES:
  rooda build                           # Run build procedure
  rooda build --max-iterations 5        # Run with iteration limit
  rooda build --dry-run                 # Validate without executing
  rooda build --context task.md         # Add context from file
  rooda build --ai-cmd-alias claude     # Use Claude AI
  rooda --list-procedures               # List available procedures

For procedure-specific help:
  rooda <procedure> --help`)
}

func printProcedureHelp(procedureName string) {
	fmt.Printf("rooda %s - Procedure-specific help not yet implemented\n", procedureName)
	fmt.Println("\nRun 'rooda --help' for general usage.")
}

func listProcedures(flags config.CLIFlags) {
	cfg, err := config.LoadConfig(flags)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading configuration: %v\n", err)
		os.Exit(ExitConfigError)
	}

	if len(cfg.Procedures) == 0 {
		fmt.Println("No procedures defined.")
		return
	}

	fmt.Println("Available procedures:")
	for name, proc := range cfg.Procedures {
		desc := proc.Description
		if desc == "" {
			desc = "(no description)"
		}
		// Truncate long descriptions
		if len(desc) > 80 {
			desc = desc[:77] + "..."
		}
		fmt.Printf("  %-20s %s\n", name, desc)
	}
}
