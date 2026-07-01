


## Usage

Source: [README](https://github.com/valehdba/pgclone/blob/main/README.md), [Usage guide](https://github.com/valehdba/pgclone/blob/main/docs/USAGE.md), [Async guide](https://github.com/valehdba/pgclone/blob/main/docs/ASYNC.md), [Release v4.3.2](https://github.com/valehdba/pgclone/releases/tag/v4.3.2), [changelog](https://github.com/valehdba/pgclone/blob/main/CHANGELOG.md), [SQL install script](https://github.com/valehdba/pgclone/blob/main/sql/pgclone--4.3.2.sql)

`pgclone` clones tables, schemas, functions, roles, and whole databases directly from SQL. In v4.x the public API is namespaced under the `pgclone` schema; upstream and Pigsty currently track PostgreSQL 14-18.

### Core clone functions

```sql
CREATE EXTENSION pgclone;
SELECT pgclone.version();

SELECT pgclone.table(
  'host=source-server dbname=mydb user=postgres password=secret',
  'public',
  'customers',
  true
);

SELECT pgclone.schema(
  'host=source-server dbname=mydb user=postgres password=secret',
  'sales',
  true
);

SELECT pgclone.database(
  'host=source-server dbname=mydb user=postgres password=secret',
  true
);
```

- `pgclone.table(...)`, `pgclone.schema(...)`, `pgclone.functions(...)`, `pgclone.database(...)`
- `pgclone.database_create(...)` creates a local target database and clones into it.
- `_ex` variants expose explicit booleans for indexes, constraints, and triggers.

### Options and masking

- JSON options support `columns`, `where`, `conflict`, and object toggles such as `indexes`, `constraints`, and `triggers`.
- JSON options also include `consistent`; it defaults to cross-table consistent snapshots in v4.3.0+ and can be disabled per call with `{"consistent": false}`.
- Upstream documents masking, auto-discovery of sensitive columns, static masking, dynamic masking, clone verification, and GDPR/compliance reporting in the usage guide.

```sql
SELECT pgclone.table(
  'host=source-server dbname=mydb user=postgres',
  'public', 'users', true, 'users_lite',
  '{"columns":["id","name","email"],"where":"status = ''active''"}'
);
```

### Consistency, diff, and preflight

```sql
SELECT pgclone.diff(
  'host=source-server dbname=prod user=postgres',
  'app_schema'
)::jsonb;

SELECT pgclone.preflight(
  'host=source-server dbname=prod user=postgres',
  'app_schema'
)::jsonb;
```

- `pgclone.diff(conninfo, schema)` reports read-only DDL drift for tables, columns, indexes, constraints, triggers, views, and sequences.
- `pgclone.preflight(conninfo, schema)` checks source and target readiness before a clone, including connection, version, permission, capacity, naming-conflict, missing-role, missing-extension, and tablespace issues.
- v4.3.0+ clones read the source under `REPEATABLE READ READ ONLY` by default. Multi-connection schema, database, and parallel-pool clones share one exported snapshot, preserving parent/child consistency while a live source is taking writes.
- Long clones hold a source transaction open, which can delay vacuum cleanup and WAL recycling; use `{"consistent": false}` when that tradeoff matters more than cross-table consistency.

### Async and progress

```sql
-- postgresql.conf
shared_preload_libraries = 'pgclone'

SELECT pgclone.schema_async(
  'host=source-server dbname=mydb user=postgres',
  'sales', true, '{"parallel":4}'
);

SELECT * FROM pgclone.jobs_view;
SELECT pgclone.progress(1);
SELECT pgclone.cancel(1);
```

- `pgclone.table_async(...)` and `pgclone.schema_async(...)` run in background workers.
- `pgclone.jobs_view`, `pgclone.progress_detail()`, `pgclone.resume()`, and `pgclone.clear_jobs()` provide job tracking and recovery.
- v4.3.2 ports the snapshot-keeper resilience fixes to async/background-worker paths, including keepalive injection and timeout protection for networked source connections.

### Caveats

- Upstream requires PostgreSQL 14+.
- The usage guide states the extension requires superuser privileges to install and use.
- Async features need `shared_preload_libraries = 'pgclone'`; worker-pool parallelism also depends on `max_worker_processes`.
- Consistent async clones may still be opted out with `{"consistent": false}` if a source-side snapshot issue must be bypassed.
- Pigsty packages `4.3.2` for PostgreSQL 14-18. The June 2026 RPM rebuild used an `LLVM_BINPATH` build fix; reviewed upstream, no material stub delta beyond the package caveat and existing v4.3.2 async-snapshot note.
