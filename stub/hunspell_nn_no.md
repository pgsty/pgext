

## Usage

> [hunspell_nn_no: Norwegian Hunspell dictionary for PostgreSQL](https://github.com/postgrespro/hunspell_dicts)

Norwegian (Nynorsk) Hunspell dictionary and text search configuration for PostgreSQL full-text search.

```sql
CREATE EXTENSION hunspell_nn_no;

SELECT ts_lexize('norwegian_hunspell', 'historier');

SELECT to_tsvector('norwegian_hunspell', 'historier');
```

This extension provides the `norwegian_hunspell` dictionary and text search configuration.
