## Usage

Sources:

- [Official EDB SPL Check documentation](https://www.enterprisedb.com/docs/pg_extensions/spl_check/)
- [Official installation guide](https://www.enterprisedb.com/docs/pg_extensions/spl_check/installing/)
- [Official configuration guide](https://www.enterprisedb.com/docs/pg_extensions/spl_check/configuring/)

`spl_check` statically checks EPAS SPL and PL/pgSQL routines for type errors, invalid object references, dead code, missing returns, suspicious casts, and some SQL-injection patterns.

### Enablement

Install the matching EPAS package and create the extension:

```sql
CREATE EXTENSION spl_check;
```

No preload is required for active/manual checks. The provider-only extension supports EDB Postgres Advanced Server.

### Check a Routine

```sql
SELECT *
FROM spl_check_function_tb('public.f1()');

SELECT *
FROM spl_check_function(
  'public.f1()',
  fatal_errors := false
);
```

`spl_check_function` supports text, JSON, and XML output. Pass a complete signature or `regprocedure` when a routine name is overloaded.

### Passive Mode and Review

Passive mode checks routines before execution and can surface compatibility warnings, but static analysis cannot prove runtime correctness or safety of dynamic SQL. Treat findings as review input, not automatic rewrites. Restrict who can inspect routine definitions, because diagnostics can reveal SQL text and object names; test warning-level changes before enabling passive checks across application workloads.

