

## Usage

> [hunspell_nl_nl: Dutch Hunspell dictionary for PostgreSQL](https://github.com/postgrespro/hunspell_dicts)

Dutch Hunspell dictionary and text search configuration for PostgreSQL full-text search.

```sql
CREATE EXTENSION hunspell_nl_nl;

SELECT ts_lexize('dutch_hunspell', 'verhalen');

SELECT to_tsvector('dutch_hunspell', 'verhalen');
```

This extension provides the `dutch_hunspell` dictionary and text search configuration.
