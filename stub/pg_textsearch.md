## Usage

Sources:

- [1.6.0 README](https://github.com/timescale/pg_textsearch/blob/v1.6.0/README.md)
- [1.6.0 control](https://github.com/timescale/pg_textsearch/blob/v1.6.0/pg_textsearch.control)
- [1.6.0 SQL](https://github.com/timescale/pg_textsearch/blob/v1.6.0/sql/pg_textsearch--1.6.0.sql)
- [1.6.0 upgrade guide](https://github.com/timescale/pg_textsearch/blob/v1.6.0/UPGRADING.md)
- [1.6.0 release](https://github.com/timescale/pg_textsearch/releases/tag/v1.6.0)

`pg_textsearch` 1.6.0 provides BM25-ranked text search on PostgreSQL 17–19 through the `bm25` access method and `<@>` operator.

### Enable and Query

An administrator must add the library to the existing preload list and restart PostgreSQL. Installation requires elevated privileges.

```ini
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

Scores are negative BM25 values: lower values rank first. Use an explicit index for partial indexes, ambiguous expressions, prepared or dynamic query text, and PL/pgSQL. Standalone scoring needs SELECT permission on the indexed table or columns.

```sql
SELECT id FROM documents
ORDER BY content <@> to_bm25query('database system', 'docs_idx')
LIMIT 5;
```

### Filtering and Index Options

`text_config` is required; `k1` defaults to 1.2 and `b` to 0.75. Immutable text expressions, text-like arrays, partial indexes and partitioned tables are supported. Repeat indexed expressions in queries and supply a matching predicate for partial indexes. A parser such as `zhparser` is an optional dependency for Chinese text.

```sql
SET default_text_search_config = 'english';
SELECT id FROM documents
WHERE content @@ to_tsquery('english', 'database & !mysql')
ORDER BY content <@> 'database'
LIMIT 5;
```

Boolean queries support AND/OR/NOT, phrases, prefixes and weights. Phrase and weight checks can require table-row rechecks. The text-search configuration must match the index. If it changes while using a cached Boolean plan, deallocate and prepare the statement again.

In 1.6.0, eligible constant Boolean predicates run inside ranked scans when every match is guaranteed to contain a ranking term. Other forms, including generic plans with unknown terms, retain PostgreSQL filtering and sorting. Post-filtered ranked scans expand their scoring batch until enough rows qualify or matches run out. Inspect EXPLAIN; a small LIMIT does not guarantee little work.

### Maintenance and Boundaries

The durable memtable lives in WAL-logged index pages. Shared-memory caching is optional; `pg_textsearch.memory_limit` defaults to 2GB and is an approximate admission budget, not a hard bound. `pg_textsearch.default_limit` defaults to 1000 and sets the initial batch when no SQL limit is available. Spills and serial VACUUM can perform synchronous compaction.

```sql
SELECT bm25_spill_index('docs_idx');
SELECT bm25_force_merge('docs_idx');
```

Mutating maintenance functions require index ownership and do not operate on partitioned parent indexes or during recovery. Published physical merges are not undone by ROLLBACK. Optional managed background compaction has additional upstream setup requirements; inline/manual modes do not require it. Querying hot standbys need `hot_standby_feedback = on` to delay physical page reuse for active snapshots.

The `bm25` access-method name conflicts with other providers such as `pg_search` and `vchord_bm25`. RLS hides rows but BM25 corpus statistics still reflect all indexed rows. Superusers can set `pg_textsearch.allow_rls` to off to reject new or rebuilt combinations with RLS; this does not disable existing indexes.

### Upgrade to 1.6.0

Install matching libraries and SQL files, restart PostgreSQL, then update every database:

```sql
ALTER EXTENSION pg_textsearch UPDATE TO '1.6.0';
```

Version 1.6.0 keeps the 1.5.1 index format and adds partition-aware scoring functions. Earlier parallel builds involving HOT chains could omit rows, and interrupted VACUUM could leave incorrect counts or bitmaps. Rebuild affected indexes; upgrading binaries and SQL cannot repair their contents.

```sql
REINDEX INDEX docs_idx;
```

Standard REINDEX blocks writes. Check all intervening upgrade notes for older indexes: pre-1.3 unspilled data needs handling under the old binary, pre-1.5 indexes can omit empty or stopword-only documents from negative Boolean queries, and version-9 headers cannot be read by pre-1.5 servers. Upgrade replicas before new binaries modify index headers; a binary downgrade does not reverse that change.
