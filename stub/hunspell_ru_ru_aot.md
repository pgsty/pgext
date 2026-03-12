

## Usage

> [hunspell_ru_ru_aot: Russian (AOT) Hunspell dictionary for PostgreSQL](https://github.com/postgrespro/hunspell_dicts)

Russian (AOT variant) Hunspell dictionary and text search configuration for PostgreSQL full-text search.

```sql
CREATE EXTENSION hunspell_ru_ru_aot;

SELECT ts_lexize('russian_aot_hunspell', 'рассказы');

SELECT to_tsvector('russian_aot_hunspell', 'рассказы');
```

This extension provides the `russian_aot_hunspell` dictionary and text search configuration.
