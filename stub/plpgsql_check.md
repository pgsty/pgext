


## Usage

Sources:

- [PGXN plpgsql_check 2.9.2](https://pgxn.org/dist/plpgsql_check/2.9.2/)
- [plpgsql_check README](https://github.com/okbob/plpgsql_check)
- [plpgsql_check control file](https://pgxn.org/dist/plpgsql_check/2.9.2/)

`plpgsql_check` is a PL/pgSQL checker, linter, profiler, tracer, and coverage tool. It analyzes PL/pgSQL function bodies with PostgreSQL's own parser and executor infrastructure, so many problems that would otherwise appear only at runtime can be found during development or CI.

The PGXN distribution version is 2.9.2, while the extension control file still declares SQL `default_version = '2.9'`. PostgreSQL 14-18 are documented as supported in the upstream README.

```sql
CREATE EXTENSION IF NOT EXISTS plpgsql_check;
```

### Check a Function

```sql
SELECT *
FROM plpgsql_check_function('public.refresh_totals()');

SELECT *
FROM plpgsql_check_function('public.refresh_totals(int, text)', fatal_errors := false);
```

The table-returning variant is easier to filter, store, or use in CI output:

```sql
SELECT functionid, lineno, statement, sqlstate, message, level
FROM plpgsql_check_function_tb('public.refresh_totals()');
```

Output formats include text, JSON, and XML:

```sql
SELECT * FROM plpgsql_check_function('fx()', format := 'text');
SELECT * FROM plpgsql_check_function('fx()', format := 'json');
SELECT * FROM plpgsql_check_function('fx()', format := 'xml');
```

### Trigger Functions

Trigger functions need the relation they operate on:

```sql
SELECT *
FROM plpgsql_check_function('public.audit_trigger()', 'public.accounts');

SELECT *
FROM plpgsql_check_function(
  'public.audit_trigger()',
  'public.accounts',
  newtable := 'new_rows',
  oldtable := 'old_rows'
);
```

### Warning Levels

```sql
SELECT *
FROM plpgsql_check_function(
  'fx()',
  extra_warnings         := true,
  performance_warnings   := true,
  security_warnings      := true,
  compatibility_warnings := true
);
```

- `extra_warnings` covers missing returns, dead code, shadowed variables, and unused arguments.
- `performance_warnings` covers hidden casts, type modifiers, and patterns that can block index usage.
- `security_warnings` includes dynamic SQL and SQL-injection risk checks.
- `compatibility_warnings` reports obsolete or version-sensitive PL/pgSQL patterns.

### Batch Checks

```sql
SELECT n.nspname, p.proname, c.*
FROM pg_catalog.pg_namespace n
JOIN pg_catalog.pg_proc p ON p.pronamespace = n.oid
JOIN pg_catalog.pg_language l ON l.oid = p.prolang
CROSS JOIN LATERAL plpgsql_check_function_tb(p.oid) AS c
WHERE l.lanname = 'plpgsql'
  AND n.nspname NOT IN ('pg_catalog', 'information_schema');
```

Use this pattern in migration pipelines to catch changed dependencies, dropped columns, unsafe casts, and PL/pgSQL mistakes before release.

### Passive Checking

Passive mode checks functions when they start. It is useful in development and preproduction, but it adds overhead.

```sql
LOAD 'plpgsql_check';
SET plpgsql_check.mode = 'fresh_start';
```

Common settings:

```text
plpgsql_check.mode = disabled | by_function | fresh_start | every_start
plpgsql_check.fatal_errors = yes | no
plpgsql_check.show_nonperformance_warnings = false
plpgsql_check.show_performance_warnings = false
```

### Profiler

```sql
SELECT plpgsql_check_profiler(true);

SELECT public.refresh_totals();

SELECT lineno, exec_stmts, total_time, avg_time, source
FROM plpgsql_profiler_function_tb('public.refresh_totals()');

SELECT stmtid, parent_stmtid, lineno, exec_stmts, stmtname
FROM plpgsql_profiler_function_statements_tb('public.refresh_totals()');

SELECT * FROM plpgsql_profiler_functions_all();
SELECT plpgsql_profiler_reset_all();
```

For shared profiler statistics and reliable early initialization, preload `plpgsql` before `plpgsql_check`:

```conf
shared_preload_libraries = 'plpgsql,plpgsql_check'
```

Without shared preload, profiler data is limited to the active session.

### Tracer and Coverage

Tracing emits notices for function and statement entry/exit and can expose variable values. It is disabled by default and must be enabled by a superuser-controlled setting.

```sql
SET plpgsql_check.enable_tracer = on;
SELECT plpgsql_check_tracer(true, 'terse');

SELECT * FROM plpgsql_coverage_statements('public.refresh_totals()');
SELECT * FROM plpgsql_coverage_branches('public.refresh_totals()');
```

### Pragmas

Use pragma calls inside functions to describe dynamic SQL, temporary tables, inferred record types, or local check settings:

```sql
CREATE OR REPLACE FUNCTION fx(anyelement) RETURNS text AS $$
DECLARE
  r record;
BEGIN
  PERFORM plpgsql_check_pragma('type: r (id int, processed bool)');
  RETURN $1::text;
END;
$$ LANGUAGE plpgsql;
```

### Caveats

- `plpgsql_check` requires `plpgsql`.
- Preloading is optional for active checks, but required for shared profiler storage and robust tracer/profiler initialization.
- Tracer output can include function arguments and local variable values; do not enable it broadly on sensitive production workloads.
- The checker cannot perfectly understand every dynamic SQL string. Use pragmas to document expected dynamic objects and reduce false positives.
