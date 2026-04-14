## Usage
> Sources: [README](https://github.com/darkhanakh/pg-kazsearch/blob/main/README.md) and [project repo](https://github.com/darkhanakh/pg-kazsearch).

`pg_kazsearch` is a PostgreSQL full-text search extension for the Kazakh language.
The upstream README describes it as a Rust extension built with `pgrx` that plugs into PostgreSQL's text search pipeline.

It creates a ready-to-use configuration named `kazakh_cfg` and the supporting dictionary `pg_kazsearch_dict`.

### Quick Start

```sql
CREATE EXTENSION pg_kazsearch;

SELECT to_tsvector('kazakh_cfg', 'президенттің жарлығы');
-- 'жарлық':2 'президент':1

SELECT ts_lexize('pg_kazsearch_dict', 'алмаларымыздағы');
-- {алма}
```

### Use Cases

The README shows the extension being used for:

- stemming individual Kazakh words
- building `tsvector` values with `to_tsvector('kazakh_cfg', ...)`
- adding generated `tsvector` columns to a table
- indexing those columns with GIN
- searching with `websearch_to_tsquery('kazakh_cfg', ...)`

Example table workflow:

```sql
ALTER TABLE articles ADD COLUMN fts tsvector
    GENERATED ALWAYS AS (
        setweight(to_tsvector('kazakh_cfg', title), 'A') ||
        setweight(to_tsvector('kazakh_cfg', body), 'B')
    ) STORED;

CREATE INDEX idx_fts ON articles USING GIN (fts);

SELECT title FROM articles
WHERE fts @@ websearch_to_tsquery('kazakh_cfg', 'президенттің жарлығы')
ORDER BY ts_rank_cd(fts, websearch_to_tsquery('kazakh_cfg', 'президенттің жарлығы')) DESC
LIMIT 10;
```

### Tuning

Penalty weights are adjustable at runtime:

```sql
ALTER TEXT SEARCH DICTIONARY pg_kazsearch_dict (w_deriv = 3.5, w_short_char = 100.0);
```

### Deployment

The README documents three supported paths:

- pre-built Debian/Ubuntu packages
- a Docker image based on `ghcr.io/darkhanakh/pg-kazsearch`
- source builds with `cargo pgrx install`

The repository metadata in this project matches PostgreSQL 16-18.
