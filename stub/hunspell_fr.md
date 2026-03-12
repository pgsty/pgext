

## Usage

> [hunspell_fr: French Hunspell dictionary for PostgreSQL](https://github.com/postgrespro/hunspell_dicts)

French Hunspell dictionary and text search configuration for PostgreSQL full-text search.

```sql
CREATE EXTENSION hunspell_fr;

SELECT ts_lexize('french_hunspell', 'histoires');

SELECT to_tsvector('french_hunspell', 'histoires');
```

This extension provides the `french_hunspell` dictionary and text search configuration.
