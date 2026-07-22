## Usage

Sources:

- [Official README](https://github.com/jezdez/pg_linegazer/blob/37456fde2b5c0100a4f25af56814f4fe92b71065/README.md)
- [Extension SQL](https://github.com/jezdez/pg_linegazer/blob/37456fde2b5c0100a4f25af56814f4fe92b71065/pg_linegazer--1.0.sql)
- [PL/pgSQL plugin implementation](https://github.com/jezdez/pg_linegazer/blob/37456fde2b5c0100a4f25af56814f4fe92b71065/src/pg_linegazer.c)
- [Preload configuration](https://github.com/jezdez/pg_linegazer/blob/37456fde2b5c0100a4f25af56814f4fe92b71065/conf.add)

`pg_linegazer` is an alpha-quality line-coverage collector for PL/pgSQL. It records statement hits without rewriting the function and includes nested calls, but its statistics exist only in the current backend process and are neither shared nor persistent.

### Setup and Workflow

Add the library to `shared_preload_libraries` and restart PostgreSQL before creating and using the extension:

```conf
shared_preload_libraries = 'pg_linegazer'
```

```sql
CREATE EXTENSION pg_linegazer;

SELECT test_func();
SELECT * FROM linegazer_simple_report('test_func'::regprocedure);
SELECT linegazer_simple_coverage('test_func'::regprocedure);

SELECT linegazer_clear();
```

Run the instrumented PL/pgSQL function and request its report in the same database session. With a connection pool, the report may be empty or incomplete when execution and reporting occur in different backends.

### Functions

- `linegazer_simple_report(regprocedure)` returns source line number, hit count, and source text.
- `linegazer_simple_coverage(regprocedure)` returns the fraction of executable lines whose hit count is greater than zero.
- `linegazer_clear()` clears all coverage data in the current backend.

Hit counts saturate at 9,999. The extension measures line execution only; branch coverage, shared storage, persistent storage, and LCOV output are roadmap items rather than implemented features.

### Operational Boundaries

Preloading installs a process-global `PLpgSQL_plugin` callback in every backend. The implementation replaces the single PL/pgSQL rendezvous pointer rather than composing multiple plugins, so it can conflict with other PL/pgSQL instrumentation depending on load order. Benchmark the overhead and avoid combining it with another plugin without source-level compatibility review.

Upstream labels the project alpha and last updated the implementation in 2016. Although the README claims PostgreSQL 9.5+, the code depends on internal PL/pgSQL structures and APIs that change across major versions. Build and test against the exact PostgreSQL major, use it only in development or controlled test environments, and do not treat its report as durable release evidence unless the relevant workload and report ran in the same backend.
