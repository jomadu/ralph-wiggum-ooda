package loop

import "strings"

// IterationOutcome represents the result of analyzing an iteration
type IterationOutcome string

const (
	OutcomeSuccess IterationOutcome = "success"  // Exit 0, no signal - reset failures
	OutcomeJobDone IterationOutcome = "job-done" // SUCCESS signal - terminate loop
	OutcomeFailure IterationOutcome = "failure"  // FAILURE signal or non-zero exit - increment failures
)

// IterationResult holds the output and exit code from an AI CLI execution
type IterationResult struct {
	ExitCode int
	Output   string
}

// DetectIterationFailure analyzes iteration result per the outcome matrix
// from iteration-loop.md. Promise signals override exit code.
// FAILURE takes precedence over SUCCESS when both present.
func DetectIterationFailure(result IterationResult) IterationOutcome {
	hasSuccess := strings.Contains(result.Output, "<promise>SUCCESS</promise>")
	hasFailure := strings.Contains(result.Output, "<promise>FAILURE</promise>")

	// FAILURE wins if both present
	if hasFailure {
		return OutcomeFailure
	}

	// SUCCESS signal terminates loop regardless of exit code
	if hasSuccess {
		return OutcomeJobDone
	}

	// No signals - exit code determines outcome
	if result.ExitCode == 0 {
		return OutcomeSuccess
	}

	return OutcomeFailure
}
