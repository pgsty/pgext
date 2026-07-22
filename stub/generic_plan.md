## Usage

Sources:

- [Official README](https://github.com/cybertec-postgresql/generic-plan/blob/f998a7eefc298a1e77f98e46464e20d96a2c1c1a/README.md)
- [Extension SQL implementation](https://github.com/cybertec-postgresql/generic-plan/blob/f998a7eefc298a1e77f98e46464e20d96a2c1c1a/generic_plan--1.0.sql)
- [PostgreSQL type-conversion rules](https://www.postgresql.org/docs/current/typeconv.html)

`generic_plan` 1.0 produces `EXPLAIN` output for a parameterized SQL statement without requiring parameter values. It is useful when investigating statements copied from `pg_stat_statements` or logs, but it returns only a generic estimated plan, never execution measurements.

### Core Workflow

```sql
CREATE EXTENSION generic_plan;

SELECT plan
FROM generic_plan(
  'SELECT * FROM orders WHERE customer_id = $1',
  verbose  => true,
  costs    => true,
  settings => true,
  format   => 'TEXT'
);
```

`generic_plan(query, verbose, costs, settings, format)` returns one text row per output line. Supported formats are `TEXT`, `XML`, `JSON`, and `YAML`. The function prepares the statement with `unknown` parameter types, sets `plan_cache_mode=force_generic_plan`, executes `EXPLAIN` without `ANALYZE`, and deallocates its temporary statement.

### Interpretation and Limitations

Because all parameter markers start as `unknown`, PostgreSQL may resolve a different overload or operator from the real client statement, or fail on ambiguity. The generated plan ignores actual parameter selectivity, runtime, row counts, buffers, and I/O. Treat it as a diagnostic lead and reproduce the real typed statement when possible.

The implementation counts parameter-marker occurrences rather than distinct parameter numbers. Repeated markers, gaps such as using `$2` without `$1`, dollar-quoted strings, and `$n` text inside SQL comments can produce an incorrect parameter list or an error. It recognizes ordinary single-quoted strings only, rejects embedded semicolons, and uses the session prepared-statement name `_stmt_`, which can conflict with an existing statement of that name. The function is `SECURITY INVOKER`; callers still need privileges needed to parse and plan the supplied SQL.
