## Usage

Source: [README](https://github.com/valehdba/pgclone/blob/main/README.md), [Usage guide](https://github.com/valehdba/pgclone/blob/main/docs/USAGE.md), [Async guide](https://github.com/valehdba/pgclone/blob/main/docs/ASYNC.md), [Release v4.0.0](https://github.com/valehdba/pgclone/releases/tag/v4.0.0), [SQL install script](https://github.com/valehdba/pgclone/blob/main/sql/pgclone--4.0.0.sql)

`pgclone` clones tables, schemas, functions, roles, and whole databases directly from SQL. In v4.0.0 the public API is namespaced under the `pgclone` schema.

### Core clone functions

```sql
CREATE EXTENSION pgclone;

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
- Upstream documents masking, auto-discovery of sensitive columns, static masking, dynamic masking, clone verification, and GDPR/compliance reporting in the usage guide.

```sql
SELECT pgclone.table(
  'host=source-server dbname=mydb user=postgres',
  'public', 'users', true, 'users_lite',
  '{"columns":["id","name","email"],"where":"status = ''active''"}'
);
```

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

### Caveats

- Upstream requires PostgreSQL 14+.
- The usage guide states the extension requires superuser privileges to install and use.
- Async features need `shared_preload_libraries = 'pgclone'`; worker-pool parallelism also depends on `max_worker_processes`.
