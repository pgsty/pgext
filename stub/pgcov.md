## Usage

Sources:

- [Upstream README](https://github.com/johto/pgcov/blob/5f536efe07c91540bd149be152b65e80b75f3cf0/README.md)
- [Version 1.0 SQL API](https://github.com/johto/pgcov/blob/5f536efe07c91540bd149be152b65e80b75f3cf0/pgcov--1.0.sql)
- [Coverage implementation](https://github.com/johto/pgcov/blob/5f536efe07c91540bd149be152b65e80b75f3cf0/pgcov.c)

`pgcov` version `1.0` is an experimental PL/pgSQL statement-coverage collector. It uses preload hooks and shared memory to send execution reports from test backends to one listening session.

### Core Workflow

```conf
shared_preload_libraries = 'pgcov'
```

Restart PostgreSQL, install the extension, then dedicate one session to the blocking listener while tests run in other sessions.

```sql
CREATE EXTENSION pgcov;
SELECT pgcov_listen();
```

Cancel the listener after the test workload. In that same backend, inspect collected calls and per-line counts:

```sql
SELECT * FROM pgcov_called_functions();
SELECT * FROM pgcov_fn_line_coverage('public.my_function(integer)');
SELECT pgcov_fn_line_coverage_src('public.my_function(integer)');
```

Only one active listener is allowed. Collected state belongs to its backend; ending that session loses it. The `coverage` column returned by `pgcov_called_functions` is left NULL by the reviewed implementation, so derive coverage from the line rows instead of relying on that field.

### Operational Caveats

Upstream calls the project a work in progress and the catalog marks it abandoned. It depends on PostgreSQL internal PL/pgSQL execution hooks and old shared-memory APIs, so modern-version compatibility cannot be assumed. Use it only in an isolated test cluster: preload and instrumentation affect the whole server, and crashes or cancellation can interrupt collection. Verify function signatures, nested calls, exceptions, parallel tests, and listener cleanup before trusting a report.
