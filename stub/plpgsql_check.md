## Usage

Sources: [official README](https://github.com/okbob/plpgsql_check), [v2.9.1 release notes](https://github.com/okbob/plpgsql_check/releases/tag/v2.9.1), [local package metadata](../db/extension.csv).

`plpgsql_check` is a PL/pgSQL checker, linter, profiler, tracer, and coverage tool. It detects many errors during development instead of waiting for runtime failures.

```sql
CREATE EXTENSION plpgsql_check;
```

### Check A Function

```sql
SELECT * FROM plpgsql_check_function('my_function()');
SELECT * FROM plpgsql_check_function('my_function(int, text)');
SELECT * FROM plpgsql_check_function('my_function()', fatal_errors := false);
```

The table-returning variant is useful for structured reports:

```sql
SELECT *
FROM plpgsql_check_function_tb('my_function()');
```

### Output Formats

```sql
SELECT * FROM plpgsql_check_function('fx()', format := 'text');
SELECT * FROM plpgsql_check_function('fx()', format := 'json');
SELECT * FROM plpgsql_check_function('fx()', format := 'xml');
```

### Check Trigger Functions

```sql
SELECT * FROM plpgsql_check_function('my_trigger_func()', 'my_table');

SELECT * FROM plpgsql_check_function(
  'my_trigger_func()',
  'my_table',
  newtable := 'newtab',
  oldtable := 'oldtab'
);
```

### Warning Categories

```sql
SELECT * FROM plpgsql_check_function(
  'fx()',
  extra_warnings := true,
  performance_warnings := true,
  security_warnings := true,
  compatibility_warnings := true
);
```

- `extra_warnings` covers issues such as missing returns, dead code, and unused arguments.
- `performance_warnings` covers performance-related checks.
- `security_warnings` includes checks such as SQL injection risk.
- `compatibility_warnings` reports obsolete or version-sensitive PL/pgSQL patterns.

### Batch Check All Functions

```sql
SELECT p.oid, p.proname, plpgsql_check_function(p.oid)
FROM pg_catalog.pg_namespace n
JOIN pg_catalog.pg_proc p ON p.pronamespace = n.oid
JOIN pg_catalog.pg_language l ON p.prolang = l.oid
WHERE l.lanname = 'plpgsql'
  AND p.prorettype <> 'pg_catalog.trigger'::pg_catalog.regtype;
```

### Passive Mode

Passive mode checks functions when they start. It is intended for development or preproduction because it adds overhead.

```sql
LOAD 'plpgsql_check';
SET plpgsql_check.mode = 'every_start';
```

Common settings:

```text
plpgsql_check.mode = [ disabled | by_function | fresh_start | every_start ]
plpgsql_check.fatal_errors = [ yes | no ]
plpgsql_check.show_nonperformance_warnings = false
plpgsql_check.show_performance_warnings = false
```

### Profiler

```sql
SELECT plpgsql_check_profiler(true);

SELECT my_function();

SELECT lineno, avg_time, source
FROM plpgsql_profiler_function_tb('my_function()');

SELECT stmtid, parent_stmtid, lineno, exec_stmts, stmtname
FROM plpgsql_profiler_function_statements_tb('my_function()');

SELECT * FROM plpgsql_profiler_functions_all();
SELECT plpgsql_profiler_reset_all();
```

For shared profiler statistics, preload `plpgsql` before `plpgsql_check`:

```text
shared_preload_libraries = 'plpgsql,plpgsql_check'
```

Without shared preload, profiler data is limited to the active session.

### Tracer

Tracing emits notices for function and statement entry/exit and can expose variable values. It is disabled by default and should be enabled only with superuser-controlled settings.

```sql
SET plpgsql_check.enable_tracer = on;
SELECT plpgsql_check_tracer(true);
SET plpgsql_check.tracer_verbosity = terse;
```

### Dependency Tracking

```sql
SELECT *
FROM plpgsql_show_dependency_tb('my_function(int)');
```

### Coverage Metrics

```sql
SELECT * FROM plpgsql_coverage_statements('my_function()');
SELECT * FROM plpgsql_coverage_branches('my_function()');
```

### Pragma Directives

Use pragma calls inside functions to tell `plpgsql_check` about dynamic SQL, temporary tables, inferred record types, or local check settings:

```sql
CREATE OR REPLACE FUNCTION fx(anyelement) RETURNS text AS $$
BEGIN
  PERFORM plpgsql_check_pragma('type: r (id int, processed bool)');
  RETURN $1::text;
END;
$$ LANGUAGE plpgsql;
```

### Caveats

- Pigsty packages `plpgsql_check` 2.9.1 for PostgreSQL 14-18 as RPMs; DEB packages come from PGDG.
- The extension requires `plpgsql`. Preloading is optional for active checking, but needed for shared profiler storage and reliable early tracer/profiler initialization.
- v2.9.1 fixes a possible crash when a traced inline block fails; no new user-facing SQL API was documented for this patch release.
- The tracer can expose function arguments or variable values. Use it carefully around security-definer functions or sensitive data.
