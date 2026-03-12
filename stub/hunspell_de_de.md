

## Usage

> [hunspell_de_de: German Hunspell dictionary for PostgreSQL](https://github.com/postgrespro/hunspell_dicts)

German Hunspell dictionary and text search configuration for PostgreSQL full-text search.

```sql
CREATE EXTENSION hunspell_de_de;

SELECT ts_lexize('german_hunspell', 'Geschichten');

SELECT to_tsvector('german_hunspell', 'Geschichten');
```

This extension provides the `german_hunspell` dictionary and text search configuration.
