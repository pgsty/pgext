## Usage

Sources: [README](https://github.com/darkhanakh/pg-kazsearch/blob/main/README.md), [releases](https://github.com/darkhanakh/pg-kazsearch/releases)

`pg_kazsearch` is a PostgreSQL full-text search extension for the Kazakh language. The README says it creates a ready-to-use text search configuration `kazakh_cfg` and dictionary `pg_kazsearch_dict`.

### Quick start

```sql
CREATE EXTENSION pg_kazsearch;

SELECT ts_lexize('pg_kazsearch_dict', 'алмаларымыздағы');
-- {алма}

SELECT to_tsvector('kazakh_cfg', 'президенттің жарлығы');
-- 'жарлық':2 'президент':1
```

### Add Kazakh FTS to a table

```sql
ALTER TABLE articles ADD COLUMN fts tsvector
    GENERATED ALWAYS AS (
        setweight(to_tsvector('kazakh_cfg', title), 'A') ||
        setweight(to_tsvector('kazakh_cfg', body), 'B')
    ) STORED;

CREATE INDEX idx_fts ON articles USING GIN (fts);

SELECT title
FROM articles
WHERE fts @@ websearch_to_tsquery('kazakh_cfg', 'президенттің жарлығы')
ORDER BY ts_rank_cd(fts, websearch_to_tsquery('kazakh_cfg', 'президенттің жарлығы')) DESC
LIMIT 10;
```

### Tuning

The README documents runtime dictionary tuning without restart:

```sql
ALTER TEXT SEARCH DICTIONARY pg_kazsearch_dict
  (w_deriv = 3.5, w_short_char = 100.0);
```

### Release and packaging notes

- Upstream release `v2.0.0` introduced the current Rust / `pgrx` architecture.
- Upstream release `v2.1.0` adds an Elasticsearch plugin alongside the PostgreSQL extension; the PostgreSQL SQL usage in the README is unchanged.
- The repository README publishes Debian packages as `2.x` releases, while this project's CSV note separately tracks the extension control version.

### Caveat

The PostgreSQL-facing docs are concise and focused on stemming plus FTS usage. For this stub, avoid inferring extra SQL objects beyond `kazakh_cfg`, `pg_kazsearch_dict`, and the documented examples above.
