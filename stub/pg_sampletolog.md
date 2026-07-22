## Usage

Sources:

- [Official README](https://github.com/anayrat/pg_sampletolog/blob/478aa18638c987efba6403e619f2ff41ad0e1d48/README.md)
- [C implementation](https://github.com/anayrat/pg_sampletolog/blob/478aa18638c987efba6403e619f2ff41ad0e1d48/pg_sampletolog.c)
- [Regression scenarios](https://github.com/anayrat/pg_sampletolog/blob/478aa18638c987efba6403e619f2ff41ad0e1d48/sql/pg_sampletolog.sql)
- [Extension control file](https://github.com/anayrat/pg_sampletolog/blob/478aa18638c987efba6403e619f2ff41ad0e1d48/pg_sampletolog.control)

`pg_sampletolog` 2.0.0 is a headless PostgreSQL 9.4–11 module that samples statements or whole transactions into the server log through executor and utility hooks. PostgreSQL 12 added core statement-sampling settings, so this module is mainly relevant to legacy servers or historical replay experiments. It has no `CREATE EXTENSION` SQL surface.

### Load and Configure

Load it for one session or at every connection start:

```conf
session_preload_libraries = 'pg_sampletolog'

pg_sampletolog.statement_sample_rate = 0.10
pg_sampletolog.transaction_sample_rate = 0
pg_sampletolog.statement_sample_limit = '100ms'
pg_sampletolog.log_statement = 'ddl'
pg_sampletolog.log_before_execution = off
pg_sampletolog.log_level = 'log'
```

For a diagnostic session, use:

```sql
LOAD 'pg_sampletolog';
SET pg_sampletolog.statement_sample_rate = 0.10;
SET pg_sampletolog.statement_sample_limit = '100ms';
```

Both sample rates default to zero, so the module is disabled until configured. Statement sampling makes an independent choice per statement. Transaction sampling makes one choice at transaction start and, when selected, logs all statements in that transaction. A positive `statement_sample_limit` forces slow statements to be logged, but only when at least one sample rate is positive.

### Important Settings

- `pg_sampletolog.statement_sample_rate`: fraction of individual statements to log.
- `pg_sampletolog.transaction_sample_rate`: fraction of transactions whose complete statement set is logged.
- `pg_sampletolog.statement_sample_limit`: always log executions longer than the duration threshold within enabled sampling.
- `pg_sampletolog.log_statement`: `none`, `ddl`, or `mod`, paralleling the historical core `log_statement` categories.
- `pg_sampletolog.log_before_execution`: emit before execution for replay tooling, rather than after with duration.
- `pg_sampletolog.disable_log_duration`: omit duration, primarily for tests.
- `pg_sampletolog.log_level`: select the PostgreSQL message severity used for sampled records.

If `pg_stat_statements` is loaded, the module can add a query identifier to log output. It logs only the outer function call for a SQL function containing nested queries. PREPARE is unsupported, and a multi-statement query string can be repeated in misleading per-statement log records.

### Logging and Compatibility Boundaries

Sampled SQL can contain passwords, tokens, personal data, and literal values. Apply the same file permissions, transport encryption, retention, redaction, and access controls as full statement logging. Sampling is probabilistic, not an audit trail; `ddl` or `mod` modes still require validation if completeness matters. Logging before execution records statements that may later fail or roll back, while logging after execution misses a statement if the backend crashes first.

The hooks are tied to PostgreSQL 9.4–11 internals. Do not infer support for the catalog's modern-version matrix. On PostgreSQL 12+, prefer core `log_statement_sample_rate`, `log_transaction_sample_rate`, `log_min_duration_sample`, and related current settings. Before legacy use, test hook coexistence, standby behavior, prepared and multi-statement traffic, nested queries, error paths, log volume, latency, random-sampling distribution, rotation, and replay compatibility.
