
## Usage

Sources: [README](https://github.com/timescale/pg_textsearch/blob/main/README.md), [release notes](https://github.com/timescale/pg_textsearch/releases), [Timescale changelog](https://github.com/timescale/docs/blob/latest/about/changelog.md)

`pg_textsearch` provides BM25-ranked full-text search for PostgreSQL with a `bm25` access method and the `<@>` scoring operator. Upstream marks `v1.0.0` as production ready.

### Enable the extension

```ini
shared_preload_libraries = 'pg_textsearch'
```

```sql
CREATE EXTENSION pg_textsearch;
```

### Build a BM25 index and query it

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

The README notes that `<@>` returns the negative BM25 score, so lower values are better matches.

### Explicit queries and index options

```sql
SELECT *
FROM documents
ORDER BY content <@> to_bm25query('database system', 'docs_idx')
LIMIT 5;

CREATE INDEX ON documents USING bm25(content)
WITH (text_config = 'english', k1 = 1.5, b = 0.8);
```

The README also documents expression indexes, partial indexes, and multilingual partial indexes.

### Useful functions and GUCs

```sql
SELECT bm25_force_merge('docs_idx');
```

Documented GUCs in current upstream docs include:

- `pg_textsearch.default_limit`
- `pg_textsearch.segments_per_level`
- `pg_textsearch.memory_limit`
- `pg_textsearch.log_scores`

### Caveats

- `pg_textsearch` currently supports PostgreSQL 17 and 18 upstream.
- Inside PL/pgSQL and stored procedures, the implicit `text <@> 'query'` form does not use planner hooks; upstream says to use `to_bm25query()` with an explicit index name there.
- `v1.0.0` adds the production-ready status, `bm25_force_merge()`, and newer GUCs documented in the official changelog and README.
