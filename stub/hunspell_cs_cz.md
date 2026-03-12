

## Usage

> [hunspell_cs_cz: Czech Hunspell dictionary for PostgreSQL](https://github.com/postgrespro/hunspell_dicts)

Czech Hunspell dictionary and text search configuration for PostgreSQL full-text search.

```sql
CREATE EXTENSION hunspell_cs_cz;

SELECT ts_lexize('czech_hunspell', 'příběhy');

SELECT to_tsvector('czech_hunspell', 'příběhy');
```

This extension provides the `czech_hunspell` dictionary and text search configuration.
