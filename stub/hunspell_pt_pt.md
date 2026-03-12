

## Usage

> [hunspell_pt_pt: Portuguese Hunspell dictionary for PostgreSQL](https://github.com/postgrespro/hunspell_dicts)

Portuguese Hunspell dictionary and text search configuration for PostgreSQL full-text search.

```sql
CREATE EXTENSION hunspell_pt_pt;

SELECT ts_lexize('portuguese_hunspell', 'histórias');

SELECT to_tsvector('portuguese_hunspell', 'histórias');
```

This extension provides the `portuguese_hunspell` dictionary and text search configuration.
