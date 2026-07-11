/*
Copyright 2018-2025 Ruohang Feng <rh@vonng.com>
*/
package cli

import (
	"context"
	"errors"
	"fmt"
	"strings"
)

// RunFailure records one failed item in a multi-item operation.
type RunFailure struct {
	Item string
	Err  error
}

// RunReport summarizes a multi-item operation such as fetch, scan, or parse.
// Skipped items count as successful outcomes because their cached data remains
// valid and no work was required.
type RunReport struct {
	Operation string
	Total     int
	Succeeded int
	Skipped   int
	Failures  []RunFailure
}

// AddFailure appends a named failure to the report.
func (r *RunReport) AddFailure(item string, err error) {
	if err == nil {
		return
	}
	r.Failures = append(r.Failures, RunFailure{Item: item, Err: err})
}

// Failed returns the number of failed items.
func (r RunReport) Failed() int {
	return len(r.Failures)
}

// Err enforces strict-by-default semantics. Best-effort mode accepts a partial
// result only when at least one item succeeded or was safely skipped; a run in
// which every item failed is always an error.
func (r RunReport) Err(bestEffort bool) error {
	if len(r.Failures) == 0 {
		return nil
	}
	for _, failure := range r.Failures {
		if errors.Is(failure.Err, context.Canceled) || errors.Is(failure.Err, context.DeadlineExceeded) {
			item := strings.TrimSpace(failure.Item)
			if item == "" {
				return failure.Err
			}
			return fmt.Errorf("%s: %w", item, failure.Err)
		}
	}
	if bestEffort && r.Succeeded+r.Skipped > 0 {
		return nil
	}

	op := strings.TrimSpace(r.Operation)
	if op == "" {
		op = "operation"
	}

	errs := make([]error, 0, len(r.Failures))
	for _, failure := range r.Failures {
		item := strings.TrimSpace(failure.Item)
		if item == "" {
			errs = append(errs, failure.Err)
		} else {
			errs = append(errs, fmt.Errorf("%s: %w", item, failure.Err))
		}
	}

	return fmt.Errorf("%s failed for %d of %d items: %w",
		op, len(r.Failures), r.Total, errors.Join(errs...))
}
