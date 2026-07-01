


## Usage

Sources: [README](https://github.com/darkhanakh/pg-kazsearch/blob/v2.2.0/README.md), [v2.2.0 release](https://github.com/darkhanakh/pg-kazsearch/releases/tag/v2.2.0), [v2.1.0 release](https://github.com/darkhanakh/pg-kazsearch/releases/tag/v2.1.0)

`pg_kazsearch` is a PostgreSQL full-text search extension for the Kazakh language. The README says it supports PostgreSQL 16-18 and creates a ready-to-use text search configuration `kazakh_cfg` and dictionary `pg_kazsearch_dict`. Version 2.2.0 adds Latin-script Kazakh support in the core stemmer; successful Latin and Cyrillic inputs converge to canonical Cyrillic stems.

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

Control Latin handling with `script_mode`; the default `auto` mode detects supported modern Kazakh Latin orthography and normalizes to Cyrillic output:

```sql
ALTER TEXT SEARCH DICTIONARY pg_kazsearch_dict
  (script_mode = cyrillic_only);
```

### Release and packaging notes

- This project's CSV tracks package/source version `2.2.0`, extension control version `0.1.0`, `pgrx` `0.18.1`, and PostgreSQL versions 16-18.
- Upstream release `v2.2.0` adds Latin-script Kazakh support to the core stemmer.
- Upstream release `v2.0.0` introduced the current Rust / `pgrx` PostgreSQL extension packaging.
- Upstream release `v2.1.0` adds an Elasticsearch plugin alongside the PostgreSQL extension; the PostgreSQL SQL usage in the README is unchanged.

### Caveat

The PostgreSQL-facing docs are concise and focused on stemming plus FTS usage. For this stub, avoid inferring extra SQL objects beyond `kazakh_cfg`, `pg_kazsearch_dict`, and the documented examples above.
