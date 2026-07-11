package cmd

import "testing"

func TestPGXNUsesIndependentWorkerDefaults(t *testing.T) {
	parallel := pgxnCmd.Flags().Lookup("parallel")
	retryFlag := pgxnCmd.Flags().Lookup("retry")
	if parallel == nil || retryFlag == nil {
		t.Fatal("pgxn parallel/retry flags are not registered")
	}
	if parallel.DefValue != "16" {
		t.Fatalf("pgxn --parallel default = %q, want 16", parallel.DefValue)
	}
	if retryFlag.DefValue != "2" {
		t.Fatalf("pgxn --retry default = %q, want 2", retryFlag.DefValue)
	}

	originalPGXNWorkers, originalPGXNRetry := pgxnWorkers, pgxnRetry
	originalWorkers, originalRetry := workers, retry
	t.Cleanup(func() {
		pgxnWorkers, pgxnRetry = originalPGXNWorkers, originalPGXNRetry
		workers, retry = originalWorkers, originalRetry
	})

	pgxnWorkers, pgxnRetry = 16, 2
	workers, retry = 8, 1
	if pgxnWorkers != 16 || pgxnRetry != 2 {
		t.Fatalf("pgxn defaults were polluted: workers=%d retry=%d", pgxnWorkers, pgxnRetry)
	}
}

func TestDefaultGeneratorCommandsRejectUnexpectedArgs(t *testing.T) {
	for _, cmd := range []struct {
		name string
		args func([]string) error
	}{
		{name: "gen", args: func(args []string) error { return genCmd.Args(genCmd, args) }},
		{name: "cc", args: func(args []string) error { return ccCmd.Args(ccCmd, args) }},
		{name: "io", args: func(args []string) error { return ioCmd.Args(ioCmd, args) }},
		{name: "pgxn", args: func(args []string) error { return pgxnCmd.Args(pgxnCmd, args) }},
		{name: "status", args: func(args []string) error { return statusCmd.Args(statusCmd, args) }},
		{name: "repo", args: func(args []string) error { return repoCmd.Args(repoCmd, args) }},
	} {
		t.Run(cmd.name, func(t *testing.T) {
			if err := cmd.args([]string{"typo"}); err == nil {
				t.Fatalf("%s accepted an unexpected positional argument", cmd.name)
			}
		})
	}
}
