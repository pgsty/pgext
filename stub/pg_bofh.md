## Usage

Sources:

- [Official README](https://github.com/rdunklau/pg_bofh/blob/master/README.md)
- [Extension control file](https://github.com/rdunklau/pg_bofh/blob/master/pg_bofh.control)
- [Planner-hook implementation](https://github.com/rdunklau/pg_bofh/blob/master/pg_bofh.c)

`pg_bofh` is a headless planner-hook demonstration that attempts to reject statements lacking a useful predicate. It installs no SQL objects and has no versioned extension script. Load the shared library only in a disposable test server; it is not a production query-safety or authorization system.

### Core Workflow

Build and install the library, add it to the server configuration, and restart PostgreSQL:

```conf
shared_preload_libraries = 'pg_bofh'
```

After restart, exercise representative read, write, maintenance, and application statements before enabling it anywhere else. There is no SQL function, GUC bypass, role allowlist, database allowlist, or status view.

### Actual Enforcement Boundary

The README summarizes the rule as forbidding queries without a WHERE clause, but the reviewed implementation is narrower and less reliable: its tree walker succeeds only after finding an operator expression in planner qualifications. Predicates such as `WHERE TRUE`, a bare Boolean column, or `IS NULL` can still be rejected, while the presence of an unrelated operator expression does not establish that a statement is safe.

The hook runs the standard planner first and errors afterward. It can affect planned statements beyond table scans, including ordinary SELECT statements and INSERT forms, while utility commands follow a different path. The old hook signature and direct planner-tree assumptions are PostgreSQL-version-specific. Because there is no built-in escape mechanism, an unexpected rejection can block normal operations; keep an out-of-band method to remove the library and restart the server.
