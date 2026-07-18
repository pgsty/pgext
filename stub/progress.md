## Usage

Sources:

- [Project README](https://github.com/wulczer/pg-progress/blob/5d8343e933ff18396abc33e2ecabf4bb2c11a1d1/README.rst)
- [Extension control file](https://github.com/wulczer/pg-progress/blob/5d8343e933ff18396abc33e2ecabf4bb2c11a1d1/progress.control)
- [Version 0.0.1 SQL API](https://github.com/wulczer/pg-progress/blob/5d8343e933ff18396abc33e2ecabf4bb2c11a1d1/sql/progress--0.0.1.sql)
- [Progress implementation](https://github.com/wulczer/pg-progress/blob/5d8343e933ff18396abc33e2ecabf4bb2c11a1d1/src/progress.c)

`progress` 0.0.1 is a 2013 prototype that estimates the fraction of a running query completed from executor instrumentation. It exposes `pg_progress_update(integer)`, `pg_progress()`, and `pg_progress_dot()` for a separate monitoring session.

### Prototype monitoring flow

The server must be built from the upstream project's modified PostgreSQL branch and preload the library:

```conf
shared_preload_libraries = 'progress'
```

After a restart and extension creation, a monitoring session signals the target backend PID and reads the shared snapshot:

```sql
CREATE EXTENSION progress;

SELECT pg_progress_update(12345);
SELECT pg_progress();
SELECT pg_progress_dot();
```

`pg_progress_update` asks the target backend to calculate a snapshot, `pg_progress` returns a double-precision estimate, and `pg_progress_dot` returns the instrumented executor tree in Graphviz DOT form. The repository's `show-progress.py` script coordinates the query and monitoring connections.

### Prototype-only boundary

This extension cannot be built against stock PostgreSQL: it depends on private executor, instrumentation, and process-signal hooks from a modified 2013 server tree. It has no current compatibility path and is cataloged as abandoned. Do not add it to a supported production PostgreSQL installation.

The estimate is heuristic and planner-cardinality errors affect it. A single cluster-wide shared snapshot is overwritten by monitors, so concurrent observations are not isolated by backend or database. The SQL functions are not declared superuser-only, while signaling another backend and collecting its executor shape can add work or reveal query structure. If preserving this prototype for research, isolate it on disposable data and restrict function execution explicitly.
