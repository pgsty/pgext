


## Usage

> [plpgsql_check: extended check for plpgsql functions](https://github.com/okbob/plpgsql_check)

`plpgsql_check` is a linter and checker for PL/pgSQL functions that detects errors at development time rather than runtime.

```sql
CREATE EXTENSION plpgsql_check;
```

### Check a Function

```sql
SELECT * FROM plpgsql_check_function('my_function()');
SELECT * FROM plpgsql_check_function('my_function(int, text)');
SELECT * FROM plpgsql_check_function('my_function()', fatal_errors := false);
```

### Output Formats

```sql
SELECT * FROM plpgsql_check_function('fx()', format := 'text');
SELECT * FROM plpgsql_check_function('fx()', format := 'json');
SELECT * FROM plpgsql_check_function('fx()', format := 'xml');
```

### Check Trigger Functions

```sql
-- Trigger functions need the associated table
SELECT * FROM plpgsql_check_function('my_trigger_func()', 'my_table');

-- With transition tables
SELECT * FROM plpgsql_check_function(
  'my_trigger_func()', 'my_table',
  newtable := 'newtab', oldtable := 'oldtab'
);
```

### Warning Categories

```sql
SELECT * FROM plpgsql_check_function('fx()',
  extra_warnings := true,         -- dead code, unused parameters
  performance_warnings := true,   -- index and casting issues
  security_warnings := true,      -- SQL injection checks
  compatibility_warnings := true  -- obsolete patterns
);
```

### Batch Check All Functions

```sql
SELECT p.oid, p.proname, plpgsql_check_function(p.oid)
FROM pg_catalog.pg_namespace n
JOIN pg_catalog.pg_proc p ON pronamespace = n.oid
JOIN pg_catalog.pg_language l ON p.prolang = l.oid
WHERE l.lanname = 'plpgsql' AND p.prorettype <> 2279;
```

### Passive Mode (Check on Execution)

```sql
LOAD 'plpgsql_check';
SET plpgsql_check.mode = 'every_start';  -- check before each execution
```

Or in `postgresql.conf`:

```
shared_preload_libraries = 'plpgsql,plpgsql_check'
plpgsql_check.mode = 'every_start'
```

### Profiler

```sql
-- Enable profiling
SELECT plpgsql_check_profiler(true);

-- Execute functions to collect data
SELECT my_function();

-- View per-line execution times
SELECT lineno, avg_time, source
FROM plpgsql_profiler_function_tb('my_function()');

-- Per-statement profile
SELECT stmtid, parent_stmtid, lineno, exec_stmts, stmtname
FROM plpgsql_profiler_function_statements_tb('my_function()');

-- All function statistics
SELECT * FROM plpgsql_profiler_functions_all();

-- Reset profiling data
SELECT plpgsql_profiler_reset_all();
```

### Dependency Tracking

```sql
SELECT * FROM plpgsql_show_dependency_tb('my_function(int)');
```

### Coverage Metrics

```sql
SELECT * FROM plpgsql_coverage_statements('my_function()');
SELECT * FROM plpgsql_coverage_branches('my_function()');
```

### Pragma Directives

Embed checking options in function comments:

```sql
CREATE OR REPLACE FUNCTION fx(anyelement) RETURNS text AS $$
BEGIN
  /* @plpgsql_check_options: anyelementtype = text */
  RETURN $1;
END;
$$ LANGUAGE plpgsql;
```
