## Usage

Sources:

- [Official v2.3.0 README](https://github.com/darkhanakh/pg-kazsearch/blob/v2.3.0/README.md)
- [v2.3.0 release](https://github.com/darkhanakh/pg-kazsearch/releases/tag/v2.3.0)
- [PostgreSQL extension control file](https://github.com/darkhanakh/pg-kazsearch/blob/v2.3.0/pg_ext/pg_kazsearch.control)
- [v2.2.0 to v2.3.0 upgrade SQL](https://github.com/darkhanakh/pg-kazsearch/blob/v2.3.0/pg_ext/sql/pg_kazsearch--2.2.0--2.3.0.sql)

`pg_kazsearch` provides Kazakh full-text stemming for PostgreSQL 16 through 18. It installs a ready-to-use `kazakh_cfg` configuration and `pg_kazsearch_dict` dictionary. Cyrillic and supported modern Latin-script Kazakh converge to canonical Cyrillic stems so documents and queries can match across scripts.

### Core Workflow

```sql
CREATE EXTENSION pg_kazsearch;

SELECT ts_lexize('pg_kazsearch_dict', 'алмаларымыздағы');
-- {алма}

SELECT to_tsvector('kazakh_cfg', 'мектептеріміздегі оқушылардың');
-- 'мектеп':1 'оқушы':2
```

Add a weighted stored vector and GIN index:

```sql
ALTER TABLE articles ADD COLUMN fts tsvector
GENERATED ALWAYS AS (
    setweight(to_tsvector('kazakh_cfg', title), 'A') ||
    setweight(to_tsvector('kazakh_cfg', body), 'B')
) STORED;

CREATE INDEX articles_fts_idx ON articles USING GIN (fts);

SELECT title
FROM articles
WHERE fts @@ websearch_to_tsquery('kazakh_cfg', 'президенттің жарлығы')
ORDER BY ts_rank_cd(
    fts,
    websearch_to_tsquery('kazakh_cfg', 'президенттің жарлығы')
) DESC;
```

### Dictionary Tuning

Penalty weights can be changed at runtime:

```sql
ALTER TEXT SEARCH DICTIONARY pg_kazsearch_dict
    (w_deriv = 3.5, w_short_char = 100.0);
```

The default `script_mode = auto` detects supported modern Kazakh Latin orthography and returns Cyrillic stems. Disable Latin handling when strict Cyrillic-only behavior is required:

```sql
ALTER TEXT SEARCH DICTIONARY pg_kazsearch_dict
    (script_mode = cyrillic_only);
```

### Upgrade and Search Caveats

- Stemmer upgrades change index terms. After upgrading to `2.3.0`, recompute stored `tsvector` columns or repopulate trigger-maintained vectors, then `VACUUM (ANALYZE)` the table.

```sql
ALTER EXTENSION pg_kazsearch UPDATE;
UPDATE articles SET title = title;
VACUUM (ANALYZE) articles;
```

- Long-lived sessions opened before an upgrade should reconnect so they load the new dictionary.
- Latin support targets the modern orthography. Mixed-script input, legacy apostrophe/acute/digraph spellings, and low-confidence ASCII tokens may remain unchanged.
- `websearch_to_tsquery` uses strict AND semantics for ordinary terms. Applications that need broader recall should deliberately implement and measure a fallback query rather than silently changing all searches to OR.
