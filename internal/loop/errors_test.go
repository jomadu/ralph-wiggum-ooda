package loop

import "testing"

func TestDetectIterationFailure_ExitZero_NoSignal(t *testing.T) {
	result := IterationResult{
		ExitCode: 0,
		Output:   "some output without signals",
	}

	outcome := DetectIterationFailure(result)

	if outcome != OutcomeSuccess {
		t.Errorf("expected OutcomeSuccess, got %v", outcome)
	}
}

func TestDetectIterationFailure_ExitZero_SuccessSignal(t *testing.T) {
	result := IterationResult{
		ExitCode: 0,
		Output:   "some output\n<promise>SUCCESS</promise>\nmore output",
	}

	outcome := DetectIterationFailure(result)

	if outcome != OutcomeJobDone {
		t.Errorf("expected OutcomeJobDone, got %v", outcome)
	}
}

func TestDetectIterationFailure_ExitZero_FailureSignal(t *testing.T) {
	result := IterationResult{
		ExitCode: 0,
		Output:   "some output\n<promise>FAILURE</promise>\nmore output",
	}

	outcome := DetectIterationFailure(result)

	if outcome != OutcomeFailure {
		t.Errorf("expected OutcomeFailure, got %v", outcome)
	}
}

func TestDetectIterationFailure_ExitZero_BothSignals(t *testing.T) {
	result := IterationResult{
		ExitCode: 0,
		Output:   "<promise>SUCCESS</promise>\n<promise>FAILURE</promise>",
	}

	outcome := DetectIterationFailure(result)

	if outcome != OutcomeFailure {
		t.Errorf("expected OutcomeFailure (FAILURE wins), got %v", outcome)
	}
}

func TestDetectIterationFailure_NonZeroExit_NoSignal(t *testing.T) {
	result := IterationResult{
		ExitCode: 1,
		Output:   "error output without signals",
	}

	outcome := DetectIterationFailure(result)

	if outcome != OutcomeFailure {
		t.Errorf("expected OutcomeFailure, got %v", outcome)
	}
}

func TestDetectIterationFailure_NonZeroExit_SuccessSignal(t *testing.T) {
	result := IterationResult{
		ExitCode: 1,
		Output:   "error but\n<promise>SUCCESS</promise>\njob done",
	}

	outcome := DetectIterationFailure(result)

	if outcome != OutcomeJobDone {
		t.Errorf("expected OutcomeJobDone (signal wins), got %v", outcome)
	}
}

func TestDetectIterationFailure_NonZeroExit_FailureSignal(t *testing.T) {
	result := IterationResult{
		ExitCode: 1,
		Output:   "error and\n<promise>FAILURE</promise>\nblocked",
	}

	outcome := DetectIterationFailure(result)

	if outcome != OutcomeFailure {
		t.Errorf("expected OutcomeFailure, got %v", outcome)
	}
}

func TestDetectIterationFailure_NonZeroExit_BothSignals(t *testing.T) {
	result := IterationResult{
		ExitCode: 1,
		Output:   "<promise>SUCCESS</promise>\n<promise>FAILURE</promise>",
	}

	outcome := DetectIterationFailure(result)

	if outcome != OutcomeFailure {
		t.Errorf("expected OutcomeFailure (FAILURE wins), got %v", outcome)
	}
}
