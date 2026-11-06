## Usage

Sources:

- [README](https://github.com/cgkanchi/pg_lion/blob/v0.3.0/README.md)
- [Reference](https://github.com/cgkanchi/pg_lion/blob/v0.3.0/docs/REFERENCE.md)
- [Control file](https://github.com/cgkanchi/pg_lion/blob/v0.3.0/pg_lion.control)

`pg_lion` 0.3.0 adds roaring bitmap inverted indexes for filtering, counts and grouped aggregates on PostgreSQL 16 and later. It works on heap tables and uses ordinary PostgreSQL operators; it does not require preloading for its default generic WAL mode.

### Core Workflow

Create `lion` indexes on the columns used by filters and groups. Vacuumed, all-visible pages let counts avoid heap visits; recently changed pages may need visibility checks in the heap.

```sql
CREATE EXTENSION pg_lion;
CREATE TABLE lion_events (country text, amount integer);
INSERT INTO lion_events VALUES ('JP', 10), ('JP', 20), ('NZ', 30);
CREATE INDEX lion_country ON lion_events USING lion (country);
CREATE INDEX lion_amount ON lion_events USING lion (amount) WITH (slices = on);
VACUUM ANALYZE lion_events;
SELECT count(*), sum(amount) FROM lion_events WHERE country = 'JP';
SELECT country, count(*) FROM lion_events GROUP BY country;
```

Scalar equality, ranges, `IN`, null tests and combined predicates are supported. `lion` also indexes arrays, `jsonb` and `tsvector`; `lion_bm25_score` ranks full-text matches. An expression index on a JSON field can support grouping that a default `jsonb` index cannot.

### Options and Diagnostics

| Object | Purpose |
| --- | --- |
| `summaries = auto` | Pre-merge bitmaps to accelerate ranges with many distinct values. |
| `slices = on` | Store bit slices for filtered aggregates on integer, date and time values; adds write and storage cost. |
| `lion_index_count(index, key)` | Count a key through a named index. |
| `lion_index_stats(index)` | Inspect index statistics; initially granted to `pg_stat_scan_tables`. |
| `lion_index_verify(index, true)` | Check structure and heap coverage; restricted diagnostic function. |
| `lion_index_wal_mode(index)` | Report `generic` or `rmgr`; restricted diagnostic function. |

`pg_lion_citext` adds the `citext` operator class. `pg_lion_vector` adds the `lionvec` access method and requires pgvector 0.7 or later. Both companions ship with the parent distribution and are enabled separately.

### WAL and Operational Boundaries

Creating the extensions requires a superuser. Optional `shared_preload_libraries = 'pg_lion'` plus a server restart enables the custom WAL resource manager (default ID 170). Indexes built in that mode require the preload for writes and `VACUUM` until rebuilt; check `lion_index_wal_mode` before changing the setting.

Keep the library preloaded wherever old custom WAL may still be replayed, including standbys and archive recovery. `REINDEX` does not rewrite historical WAL. Before removing the preload, rebuild every `rmgr` index with `wal_mode = generic` while the library is still preloaded, and account for WAL replay consumers.

There are no unique indexes or `INCLUDE` columns, and only heap tables are supported. Keep B-tree indexes for uniqueness and ordinary ordered scans. Loading the library installs planner hooks and can affect estimates for other queries. Measure write-heavy workloads before adding indexes or slices.

The on-disk format may change before 1.0; follow release-specific `REINDEX` requirements. Builds predating 0.3.0 have no extension update path: upstream requires recreating the extensions and their indexes. Version 0.3.0 supplies installation scripts only.
