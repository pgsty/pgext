## Usage

Sources: [README v1.1.0](https://github.com/timescale/pg_textsearch/blob/v1.1.0/README.md), [v1.1.0 release notes](https://github.com/timescale/pg_textsearch/releases/tag/v1.1.0)

`pg_textsearch` provides BM25-ranked full-text search for PostgreSQL with a `bm25` access method and the `<@>` scoring operator. Upstream marks `v1.1.0` as production ready.

`v1.1.0` supports PostgreSQL 17 and 18. Prebuilt release assets are published for both PostgreSQL versions on Linux and macOS. The extension must be loaded through `shared_preload_libraries` before `CREATE EXTENSION`.

### Enable the Extension

```ini
shared_preload_libraries = 'pg_textsearch'  # add to any existing list
```

```sql
CREATE EXTENSION pg_textsearch;
```

Install the new binary and restart PostgreSQL before running an extension upgrade:

```sql
ALTER EXTENSION pg_textsearch UPDATE;
```

Upstream says upgrading from 1.0.0 to 1.1.0 does not require `REINDEX`.

### Build and Query BM25 Indexes

```sql
CREATE TABLE documents (id bigserial PRIMARY KEY, content text);

CREATE INDEX docs_idx
ON documents USING bm25(content)
WITH (text_config = 'english');

SELECT *
FROM documents
ORDER BY content <@> 'database system'
LIMIT 5;
```

`<@>` returns the negative BM25 score because PostgreSQL operator index scans are ascending; lower values are better matches. Use `ORDER BY ... LIMIT` for fast top-k searches.

For an explicit index reference, use `to_bm25query()`:

```sql
SELECT *
FROM documents
ORDER BY content <@> to_bm25query('database system', 'docs_idx')
LIMIT 5;
```

The main documented SQL surface is:

- `text <@> 'query'` to score text with planner-detected index context.
- `text <@> bm25query` to score with an explicit `bm25query`.
- `to_bm25query(text)` for ORDER BY use with planner-selected index context.
- `to_bm25query(text, text)` for query text plus index name.
- `bm25query = bm25query` for equality checks.

### Index Options and Data Shapes

```sql
CREATE INDEX ON documents USING bm25(content)
WITH (text_config = 'english', k1 = 1.5, b = 0.8);
```

Index options are `text_config` (required), `k1` (default `1.2`), and `b` (default `0.75`). Text search configurations such as `english`, `simple`, `french`, and `german` use PostgreSQL text search configuration names.

`v1.1.0` adds native array input support for `text[]`, `varchar[]`, and `bpchar[]` columns; array elements are concatenated before tokenization.

```sql
CREATE TABLE posts (id serial PRIMARY KEY, tags text[]);
CREATE INDEX posts_tags_idx ON posts USING bm25(tags)
WITH (text_config = 'english');

SELECT *
FROM posts
ORDER BY tags <@> 'database'
LIMIT 10;
```

Expression indexes support immutable text expressions, including JSONB extraction, text transformations, and multi-column concatenation:

```sql
CREATE INDEX events_msg_idx ON events USING bm25 ((data->>'message'))
WITH (text_config = 'english');

SELECT *
FROM events
ORDER BY (data->>'message') <@> to_bm25query('network error', 'events_msg_idx')
LIMIT 10;
```

Partial indexes scope search to a subset of rows. Query them with an explicit index name:

```sql
CREATE INDEX docs_en_idx ON docs USING bm25(content)
WITH (text_config = 'english')
WHERE lang = 'en';

SELECT *
FROM docs
WHERE lang = 'en'
ORDER BY content <@> to_bm25query('databases', 'docs_en_idx')
LIMIT 10;
```

### Operations and GUCs

```sql
SELECT bm25_force_merge('docs_idx');
SELECT * FROM bm25_memory_usage();
```

`bm25_force_merge(index_name)` consolidates all segments into one and is best used after bulk loads, not during steady write traffic. `bm25_memory_usage()` reports shared memory usage for memtables.

Documented `pg_textsearch` GUCs in v1.1.0 include:

- `pg_textsearch.default_limit`
- `pg_textsearch.compress_segments`
- `pg_textsearch.segments_per_level`
- `pg_textsearch.memory_limit`
- `pg_textsearch.bulk_load_threshold`
- `pg_textsearch.memtable_spill_threshold` (deprecated; use `memory_limit` for new deployments)

`pg_textsearch.memory_limit` defaults to `2GB` and caps dynamic shared memory used by memtables. The release notes also call out improved concurrent insert throughput, faster VACUUM via segment alive bitsets, subtransaction cleanup, and parallel build race fixes.

### Caveats

- `pg_textsearch` requires `shared_preload_libraries = 'pg_textsearch'` and a PostgreSQL restart before `CREATE EXTENSION`.
- Inside PL/pgSQL and stored procedures, the implicit `text <@> 'query'` form does not use planner hooks; upstream says to use `to_bm25query()` with an explicit index name there.
- Phrase queries are not native because the index stores term frequencies, not term positions; use BM25 ranking plus a post-filter for phrase-like matching.
- Partial indexes require `to_bm25query()` with the index name because the implicit query form skips them.
- BM25 indexes on partitioned tables use partition-local statistics, so cross-partition scores may not be directly comparable.
- Words longer than PostgreSQL's `tsvector` word length limit are ignored during tokenization.
