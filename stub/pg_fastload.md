## Usage

Sources:

- [Official source](https://github.com/yuansong/pg_fastload/blob/6e50f18ce90554dddfeeb1a0c27e854158b3ea41/README.md)
- [Official source](https://github.com/yuansong/pg_fastload/blob/6e50f18ce90554dddfeeb1a0c27e854158b3ea41/docs/USER_MANUAL.md)
- [Official source](https://github.com/yuansong/pg_fastload/blob/6e50f18ce90554dddfeeb1a0c27e854158b3ea41/RELEASE_NOTES_v1.0.0.md)
- [Official source](https://github.com/yuansong/pg_fastload/blob/6e50f18ce90554dddfeeb1a0c27e854158b3ea41/include/common.h)
- [Official source](https://github.com/yuansong/pg_fastload/blob/6e50f18ce90554dddfeeb1a0c27e854158b3ea41/lib/Makefile)
- [Official source](https://github.com/yuansong/pg_fastload/blob/6e50f18ce90554dddfeeb1a0c27e854158b3ea41/lib/pg_fastload.control.in)
- [Official source](https://github.com/yuansong/pg_fastload/blob/6e50f18ce90554dddfeeb1a0c27e854158b3ea41/lib/pg_fastload.sql.in)
- [Official source](https://github.com/yuansong/pg_fastload/blob/6e50f18ce90554dddfeeb1a0c27e854158b3ea41/lib/pg_fastload.c)
- [Official source](https://github.com/yuansong/pg_fastload/blob/6e50f18ce90554dddfeeb1a0c27e854158b3ea41/lib/export.c)
- [Official source](https://github.com/yuansong/pg_fastload/blob/6e50f18ce90554dddfeeb1a0c27e854158b3ea41/bin/recovery.c)
- [Official source](https://github.com/yuansong/pg_fastload/blob/6e50f18ce90554dddfeeb1a0c27e854158b3ea41/LICENSE)

`pg_fastload` `1.0.0` is a distinct fork of pg_bulkload with parallel parsing, bulk loading and file export. The inspected release source supports PostgreSQL 15–19; upstream's complete binary verification used PostgreSQL 18.6 on x86_64 Linux. Its control and versioned installation SQL are generated from templates, and objects live in the fixed `pgfastload` schema.

### Load and export

Installation and both load/export functions require a database superuser. Server-side input files must be readable, and output directories writable, by the PostgreSQL OS user. Schema grants alone do not bypass the functions' superuser checks.

For a CSV file without a header, this example uses the normal buffer/WAL writer:

```sql
CREATE EXTENSION pg_fastload;
CREATE TABLE load_target (id integer, label text);
SELECT * FROM pgfastload.pg_fastload(ARRAY[
  'INPUT=/srv/import/data.csv',
  'OUTPUT=public.load_target',
  'TYPE=CSV',
  'WRITER=BUFFERED',
  'PARALLEL=ON',
  'PARALLEL_WORKERS=4'
]);
```

No `shared_preload_libraries` entry is needed. Parallel loading uses dynamic background workers and requires `max_worker_processes` to accommodate the parser workers plus one reader and any other cluster workers. Changing that server setting requires a restart; installing the extension alone does not.

The CLI accepts the same options through a control file: `pg_fastload control.ctl`. The SQL loader returns counts, rejected-row totals and timing. Export requires an absolute server-side output path that does not already exist:

```sql
SELECT pgfastload.pg_fastload_export(ARRAY[
  'TABLE=public.load_target',
  'OUTPUT=/srv/export/load_target.csv',
  'TYPE=CSV',
  'HEADER=ON',
  'PARALLEL=OFF'
]);
SELECT * FROM pgfastload.pg_fastload_progress;
```

`TABLE=` exports can run in parallel under a shared MVCC snapshot; output order is unspecified. `QUERY=` accepts one read-only query and always uses the serial path. A long export can delay vacuum cleanup. The progress view exposes database, relation and file names; limit monitoring grants. Its underlying `fl_progress_rows()` function has public execution revoked.

### Writer and recovery boundaries

The default `WRITER=DIRECT` writes data files outside the normal buffer path and keeps load-status files in `$PGDATA/pg_fastload/`. An interrupted direct load can leave the table unusable. Rehearse the documented offline recovery with `pg_fastload -r -D "$PGDATA"` after stopping the database server; the recovery program checks that server processes are absent. Do not delete its status files to suppress recovery. Choose `WRITER=BUFFERED` when the normal buffer/WAL path is required.

`PARALLEL=ON` parallelizes reading/parsing while insertion stays in the caller's transaction. The separate `WRITER=PARALLEL` mode instead inserts through another connection and transaction; it is not atomic with the caller and is never combined with the parallel reader.

### Constraints and indexes

This is a privileged bulk loader, not a general substitute for ordinary application inserts. The target is exclusively locked during a parallel load. Unsupported conditions—including filters, encoding overrides, constraint checks, tolerated parse errors and several input modes—fall back to serial execution with a diagnostic. `CHECK_CONSTRAINTS` defaults to `NO`; set and verify the desired validation contract explicitly. Parallel loading does not preserve row order and has no per-row error tolerance.

`SKIP_INDEX=ON` drops secondary indexes within the load transaction and prints reconstruction statements; it does not rebuild them automatically. Constraint-backed indexes are refused, and `DUP_ERRORS` cannot be combined with this option. A standalone unique index can be dropped, removing its duplicate protection. Preserve the emitted statements and verify rebuilt indexes before returning the table to service.
