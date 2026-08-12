## Usage

Sources:

- [Biscuit 3.0.0 on PGXN](https://pgxn.org/dist/biscuit/3.0.0/)
- [Biscuit 3.0.0 release](https://github.com/CrystallineCore/Biscuit/releases/tag/v3.0.0)
- [Biscuit 3.0.0 README](https://github.com/CrystallineCore/Biscuit/blob/v3.0.0/README.md)
- [Biscuit 3.0.0 changelog](https://github.com/CrystallineCore/Biscuit/blob/v3.0.0/CHANGELOG.md)
- [Biscuit 3.0.0 metadata](https://api.pgxn.org/dist/biscuit/3.0.0/META.json)
- [Biscuit 3.0.0 control file](https://github.com/CrystallineCore/Biscuit/blob/v3.0.0/biscuit.control)
- [Biscuit 3.0.0 Makefile](https://github.com/CrystallineCore/Biscuit/blob/v3.0.0/Makefile)
- [Biscuit 3.0.0 installation SQL](https://github.com/CrystallineCore/Biscuit/blob/v3.0.0/sql/biscuit.sql)
- [Biscuit 2.5.0 to 3.0.0 upgrade SQL](https://github.com/CrystallineCore/Biscuit/blob/v3.0.0/sql/biscuit--2.5.0--3.0.0.sql)

`biscuit` 3.0.0 is a PostgreSQL 16+ positional-bitmap index access method for exact `LIKE` and `ILIKE` filtering. It is strongest for anchored patterns, `_` wildcards, length predicates, and multi-column conjunctions. Version 3.0.0 stores index state in WAL-logged relation pages, so crash recovery, point-in-time recovery, physical replication, and hot-standby reads use PostgreSQL's ordinary recovery path. It does not require `shared_preload_libraries` or a restart.

The project remains under active development and recommends representative staging tests. Its per-connection memory, write amplification, and cache-reload behavior make it best suited to read-mostly analytical workloads rather than continuously updated OLTP tables or very large connection pools.

### Build and Query an Index

Load the data first, then create the index. The default `biscuit_ops` supports both case-sensitive and case-insensitive predicates. Use `biscuit_like_ops` or `biscuit_ilike_ops` when only one mode is required, avoiding the unused structure set.

```sql
CREATE EXTENSION biscuit;

CREATE INDEX message_body_biscuit_idx
ON message USING biscuit (body biscuit_like_ops);

ANALYZE message;

EXPLAIN (ANALYZE, BUFFERS)
SELECT id, body
FROM message
WHERE body LIKE 'timeout%';
```

Expression and multi-column indexes are supported. The query must use expressions and operators compatible with the chosen operator class. Check representative plans after loading statistics, especially for unanchored patterns.

### Operator Classes and Query Boundaries

- `biscuit_ops` is the default text operator class and indexes `LIKE`, `NOT LIKE`, `ILIKE`, and `NOT ILIKE`.
- `biscuit_like_ops` indexes only `LIKE` and `NOT LIKE`.
- `biscuit_ilike_ops` indexes only `ILIKE` and `NOT ILIKE`.

Biscuit returns exact matches without a heap recheck, but it is a filtering index: it does not provide ordered, backward, index-only, or unique scans, cannot back `CLUSTER`, and does not support regular expressions, similarity search, fuzzy search, or locale-aware collation. A B-tree with `text_pattern_ops` is usually a better fit for selective prefix lookups, while `pg_trgm` is designed for unanchored substring, regular-expression, and similarity searches.

### Diagnostics and Configuration

Important inspection objects include `biscuit_indexes`, `biscuit_status`, `biscuit_index_stats(oid)`, `biscuit_index_memory_size()`, `biscuit_pending_list_stats(oid)`, and `biscuit_pending_list_usage`. The memory function reports the current backend's session-local copy. `total_pending_bytes` is refreshed during `VACUUM`, so pending-list figures can lag live writes by up to one vacuum cycle.

- `biscuit.delta_compaction_slots` defaults to 20000 and controls how many pending rows are tolerated before compaction. It is a privileged setting because raising it can increase reload work for other sessions.
- `biscuit.diag_scan_trace` defaults to off and emits verbose per-scan candidate accounting. Enable it only for a focused reproducer.

Every backend lazily loads its own copy of an index and keeps it for the connection lifetime. A committed write invalidates other cached copies; their next access reloads the index rather than refreshing it incrementally. Size pools for this memory behavior and avoid interleaving frequent writes with latency-sensitive reads.

Live-index `INSERT` and `UPDATE` generate substantial WAL; monitor `pg_wal`, replication lag, and replication-slot retention, and consider a bounded `max_slot_wal_keep_size`. Bulk loading before index creation is substantially cheaper. `VACUUM` drains pending work but does not shrink the index; use `REINDEX` to reclaim index space.

### Upgrade to 3.0.0

Version 3.0.0 is an incompatible on-disk format change. Updating the extension catalog does not convert existing index pages: every Biscuit index created under 2.x must be rebuilt. Plan enough maintenance time and WAL capacity for the rebuild.

```sql
ALTER EXTENSION biscuit UPDATE TO '3.0.0';

SELECT schema_name, index_name
FROM biscuit_indexes;

REINDEX INDEX CONCURRENTLY public.message_body_biscuit_idx;
```

The unpatched upstream 3.0.0 archive ships and installs only the `2.5.0--3.0.0` step, while earlier stable packages exposed catalog versions `2.4.0` or `2.4.1`. Pigsty's 3.0.0 RPM and DEB packages restore that missing catalog path before applying the upstream step. For another source build or package, inspect `pg_extension_update_paths('biscuit')` before `ALTER EXTENSION`; regardless of the available SQL path, the mandatory `REINDEX` or `REINDEX CONCURRENTLY` remains a separate manual operation.
