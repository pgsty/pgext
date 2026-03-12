

## Usage

> [hunspell_ru_ru: Russian Hunspell dictionary for PostgreSQL](https://github.com/postgrespro/hunspell_dicts)

Russian Hunspell dictionary and text search configuration for PostgreSQL full-text search.

```sql
CREATE EXTENSION hunspell_ru_ru;

SELECT ts_lexize('russian_hunspell', 'рассказы');

SELECT to_tsvector('russian_hunspell', 'рассказы');
```

This extension provides the `russian_hunspell` dictionary and text search configuration.
