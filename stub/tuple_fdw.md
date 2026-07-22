## Usage

Sources:

- [Official README](https://github.com/zilder/tuple_fdw/blob/9c737a2043bfd78fe085f0966e3a58b62e38fd70/README.md)
- [FDW implementation](https://github.com/zilder/tuple_fdw/blob/9c737a2043bfd78fe085f0966e3a58b62e38fd70/tuple_fdw.c)
- [File storage implementation](https://github.com/zilder/tuple_fdw/blob/9c737a2043bfd78fe085f0966e3a58b62e38fd70/storage.c)
- [Extension SQL for version 0.1](https://github.com/zilder/tuple_fdw/blob/9c737a2043bfd78fe085f0966e3a58b62e38fd70/tuple_fdw--0.1.sql)

`tuple_fdw` stores PostgreSQL heap tuples in an append-only external binary file made of checksummed, LZ4-compressed blocks. It supports sequential reads and inserts while bypassing PostgreSQL's buffer cache and autovacuum. It is experimental cold storage, not a transactional table replacement.

### Core Workflow

Install liblz4 on the PostgreSQL server, create the extension and server, then define a foreign table with an absolute server-side file path. The validator creates a missing file using the PostgreSQL operating-system account.

```sql
CREATE EXTENSION tuple_fdw;
CREATE SERVER tuple_store FOREIGN DATA WRAPPER tuple_fdw;

CREATE FOREIGN TABLE archive_event (
  event_id bigint,
  occurred_at timestamptz,
  payload text
)
SERVER tuple_store
OPTIONS (
  filename '/var/lib/postgresql/tuple_fdw/archive_event.bin',
  use_mmap 'false',
  lz4_acceleration '1'
);
```

Bulk insert in the intended physical order, then query normally. The FDW implements INSERT but not UPDATE or DELETE.

```sql
INSERT INTO archive_event
SELECT id, occurred_at, payload
FROM live_event
ORDER BY occurred_at, id;

SELECT *
FROM archive_event
WHERE occurred_at >= timestamptz '2026-01-01 00:00:00+00';
```

### Foreign-Table Options

- `filename`: required file path; created if absent.
- `use_mmap`: use memory-mapped reads instead of `fread`; it can help read-heavy concurrency but does not make writes concurrent-safe.
- `sorted`: space-separated column names that tell the planner the file is ordered.
- `lz4_acceleration`: LZ4 speed/compression tradeoff, default `1`; higher values favor speed over compression ratio.

The `sorted` option is a promise, not an enforcement mechanism. If the physical file is not actually ordered by those columns, the planner can choose an invalid order-dependent plan and return wrong results.

### Durability, Concurrency, and Portability

The external file is outside PostgreSQL WAL, MVCC, transaction rollback, logical/physical replication, and ordinary base-backup management. Although writes are fsynced and the FDW takes a relation lock, an aborted transaction does not undo file changes, and another database, foreign table, process, or host can access the same path without that lock. Upstream explicitly disallows concurrent writes and mixed concurrent reads/writes.

Single-row inserts repeatedly decompress and recompress the last block, so use bulk loads. The file stores raw heap-tuple bytes and depends on the foreign-table row layout and PostgreSQL binary representation; do not attach a different schema, architecture, or PostgreSQL build to an existing file without a proven migration. Back up and restore the file with an external, quiesced procedure, and verify its checksum and row count before exposing it again. `mmap` readers also require the file not to be replaced or truncated underneath them.
