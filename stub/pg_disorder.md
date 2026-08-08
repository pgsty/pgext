## Usage

Sources:

- [pg_disorder 0.1.0 README](https://api.pgxn.org/src/pg_disorder/pg_disorder-0.1.0/README.md)
- [pg_disorder 0.1.0 metadata](https://api.pgxn.org/src/pg_disorder/pg_disorder-0.1.0/META.json)
- [pg_disorder 0.1.0 Makefile](https://api.pgxn.org/src/pg_disorder/pg_disorder-0.1.0/Makefile)

`pg_disorder` is a test-only PostgreSQL loadable module that deliberately changes the output order of eligible `SELECT` queries. It helps find applications and tests that accidentally depend on unspecified row order. It is a headless module: there is no control file, SQL install script, or `CREATE EXTENSION pg_disorder` step.

### Enable It for a Test Database

Load the module at session start so its planner hook is available:

```sql
ALTER DATABASE regression_db
  SET session_preload_libraries = 'pg_disorder';

ALTER DATABASE regression_db
  SET pg_disorder.mode = 'reverse';
```

Reconnect after changing `session_preload_libraries`. Do not add this module to a production-wide `shared_preload_libraries` setting.

### Modes

```sql
SET pg_disorder.mode = 'off';
SET pg_disorder.mode = 'reverse';
SET pg_disorder.mode = 'shuffle';
SET pg_disorder.seed = 42;
SET pg_disorder.force_serial = on;
```

- `off` leaves plans unchanged.
- `reverse` deterministically reverses eligible output.
- `shuffle` produces a deterministic permutation for a fixed session seed, submitted query text, and plan. With the default seed of zero, each session first chooses and logs a random seed.
- `force_serial` suppresses parallel plans to make disorder tests reproducible.

Always fix a failing query by adding a semantically correct `ORDER BY`; do not encode the accidental order observed under `off`.

### Eligibility and Caveats

The hook targets top-level `SELECT` statements without `ORDER BY`. It deliberately skips query shapes where reordering is unsafe or changes SQL semantics, including aggregates, grouping, `DISTINCT`, set operations, window functions, recursive queries, row locks, and queries without a `FROM` relation.

- `pg_disorder` is fault-injection tooling, not a production query feature.
- Passing a disorder run does not prove every unordered query is safe; excluded query shapes and planner paths are not rewritten.
- The package installs a server module only. Verify enablement with the GUCs or module load state, not `pg_extension`.
