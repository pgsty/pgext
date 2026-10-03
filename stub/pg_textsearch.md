## Usage

Sources:

- [Version 1.5.1 README](https://github.com/timescale/pg_textsearch/blob/v1.5.1/README.md)
- [Control file](https://github.com/timescale/pg_textsearch/blob/v1.5.1/pg_textsearch.control)
- [Versioned SQL](https://github.com/timescale/pg_textsearch/blob/v1.5.1/sql/pg_textsearch--1.5.1.sql)
- [Version 1.5.1 release](https://github.com/timescale/pg_textsearch/releases/tag/v1.5.1)

`pg_textsearch` provides BM25-ranked full-text search with the `bm25` access method and `<@>` operator. Version 1.5.1 supports PostgreSQL 17 and 18 and requires preloading and restart. Preserve other entries when updating the preload list.

### Build and Query

```conf
shared_preload_libraries = 'pg_textsearch'
```

```sql
CREATE EXTENSION pg_textsearch;
CREATE TABLE documents (id bigserial PRIMARY KEY, content text);
INSERT INTO documents(content) VALUES
    ('PostgreSQL is a database system'),
    ('BM25 ranks full text search results');
CREATE INDEX docs_idx ON documents USING bm25(content)
WITH (text_config = 'english');

SELECT id, content <@> 'database system' AS score
FROM documents
ORDER BY content <@> 'database system'
LIMIT 5;
```

Scores are negative BM25 values, so lower scores rank first. Use `ORDER BY` with `LIMIT` for top-k execution. Specify the index explicitly for standalone scoring, partial indexes, or PL/pgSQL:

```sql
SELECT id FROM documents
ORDER BY content <@> to_bm25query('database system', 'docs_idx')
LIMIT 5;
```

### Index and Query Options

`text_config` is required and names a PostgreSQL text search configuration. `k1` defaults to 1.2 and `b` to 0.75. The `bm25query` type and `to_bm25query(text, text)` carry explicit query/index context. Standalone scoring requires SELECT permission on the table or indexed columns.

The extension supports `text[]`, `varchar[]` and `bpchar[]`, immutable text expression indexes, partial indexes and partitioned tables. Repeat an indexed expression in the query. Partial indexes need the matching predicate and an explicit index name. Version 1.5.0 improves filtered top-k execution, but restrictive post-filters can still return fewer rows than requested.

For Chinese tokenization, configure a parser such as `zhparser` and use that text search configuration. This is an optional workflow dependency. For very large texts without whitespace word boundaries, upstream recommends application-controlled chunks in a text array.

### Maintenance and Configuration

Starting with 1.3.0, the durable memtable is stored in index pages and uses standard PostgreSQL WAL replay. Queries can still use a shared-memory read cache, controlled by the cache and memory-limit settings below. Automatic compaction runs during spills, so write-heavy workloads can observe synchronous compaction latency.

```sql
SELECT bm25_spill_index('docs_idx');
SELECT bm25_force_merge('docs_idx');
```

Use force-merge after bulk loading rather than during steady write traffic. VACUUM also spills pending memtable pages. Relevant settings are:

| Setting | Default | Purpose |
| --- | --- | --- |
| `pg_textsearch.default_limit` | 1000 | Scoring bound without a query limit |
| `pg_textsearch.compress_segments` | on | Posting-block compression |
| `pg_textsearch.segments_per_level` | 8 | Compaction threshold |
| `pg_textsearch.bulk_load_threshold` | 100000 | Terms per transaction before spilling |
| `pg_textsearch.memtable_pages_threshold` | 64 | Chain-page count before spilling |
| `pg_textsearch.memtable_cache_enabled` | on | Shared-memory read cache |
| `pg_textsearch.memory_limit` | 2GB | Approximate cache admission budget; 0 removes the limit |

Parallel builds require at least 64 MB of `maintenance_work_mem` and available parallel maintenance workers. For upgrades, install the matching binary, restart PostgreSQL and run the extension update according to the release instructions.

```sql
ALTER EXTENSION pg_textsearch UPDATE;
```

### Boundaries

The `bm25` access-method name conflicts with `pg_search` and `vchord_bm25`; do not install conflicting providers in the same database. Phrase matching uses conservative heap rechecks where needed. Partition scores use partition-local statistics and may not be comparable across partitions. Very long tokens inherit PostgreSQL text-search limits. Fixed LWLock tranche IDs can also conflict with another extension and produce misleading wait-event names.

### Version 1.5.0 Queries and Maintenance

Boolean filtering now accepts PostgreSQL `tsquery`, including AND/OR/NOT, pure-negative, prefix, weight and phrase conditions. Phrase/weight conditions may require heap rechecks. PostgreSQL 19 support is beta. Optional managed background compaction requires `pg_durable` 0.2.8+ in the same configured database, preloaded and configured according to upstream; inline/manual modes do not require it. `pg_textsearch.allow_rls` defaults on: BM25 statistics include all indexed rows, including rows hidden by RLS. Turning it off prevents new/rebuilt BM25 indexes on protected tables and related RLS enablement; it does not disable existing indexes.

### Upgrade to 1.5.1

Version 1.5.1 fixes concurrent spill/force-merge truncation corruption, expression-index and array/domain handling, and maintenance in databases without the extension. Install the matching library, restart PostgreSQL, then run the extension update in each database. The 1.5.0-to-1.5.1 migration checks that the library was preloaded; it adds no SQL object or explicit index-format migration.

```sql
ALTER EXTENSION pg_textsearch UPDATE TO '1.5.1';
```

The cache budget is approximate and concurrent work can exceed it. Hot standbys serving BM25 queries require `hot_standby_feedback = on` to delay physical page reuse while snapshots are active. Mutating maintenance helpers require index ownership and cannot run during recovery; published physical merge replacements are not undone by transaction rollback.
