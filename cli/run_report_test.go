package cli

import (
	"context"
	"errors"
	"strings"
	"testing"
)

func TestRunReportStrictRejectsPartialFailure(t *testing.T) {
	report := RunReport{Operation: "fetch", Total: 2, Succeeded: 1}
	report.AddFailure("repo-b", errors.New("unavailable"))

	err := report.Err(false)
	if err == nil {
		t.Fatal("strict run accepted a partial failure")
	}
	for _, want := range []string{"fetch", "1 of 2", "repo-b", "unavailable"} {
		if !strings.Contains(err.Error(), want) {
			t.Fatalf("error %q does not contain %q", err, want)
		}
	}
}

func TestRunReportBestEffortRequiresSomeSuccess(t *testing.T) {
	partial := RunReport{Operation: "parse", Total: 2, Succeeded: 1}
	partial.AddFailure("repo-b", errors.New("invalid metadata"))
	if err := partial.Err(true); err != nil {
		t.Fatalf("best-effort partial run returned error: %v", err)
	}

	failed := RunReport{Operation: "parse", Total: 1}
	failed.AddFailure("repo-a", errors.New("invalid metadata"))
	if err := failed.Err(true); err == nil {
		t.Fatal("best-effort run accepted an all-failed result")
	}
}

func TestRunReportTreatsSafeSkipAsSuccess(t *testing.T) {
	report := RunReport{Operation: "fetch", Total: 2, Skipped: 1}
	report.AddFailure("repo-b", errors.New("unavailable"))
	if err := report.Err(true); err != nil {
		t.Fatalf("best-effort run rejected a valid cached result: %v", err)
	}
}

func TestRunReportBestEffortStillPropagatesCancellation(t *testing.T) {
	for _, canceled := range []error{context.Canceled, context.DeadlineExceeded} {
		report := RunReport{Operation: "fetch", Total: 2, Succeeded: 1}
		report.AddFailure("repo-b", canceled)
		err := report.Err(true)
		if !errors.Is(err, canceled) {
			t.Fatalf("best-effort cancellation = %v, want %v", err, canceled)
		}
	}
}

func TestLegacyKeepOptionMapsToKeepTemp(t *testing.T) {
	parser := NewParser(ParseOptions{Keep: true})
	if !parser.opts.KeepTemp {
		t.Fatal("deprecated ParseOptions.Keep did not enable KeepTemp")
	}
}
