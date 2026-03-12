

## Usage

> [hunspell_ne_np: Nepali Hunspell dictionary for PostgreSQL](https://github.com/postgrespro/hunspell_dicts)

Nepali Hunspell dictionary and text search configuration for PostgreSQL full-text search.

```sql
CREATE EXTENSION hunspell_ne_np;

SELECT ts_lexize('nepali_hunspell', 'कथाहरू');

SELECT to_tsvector('nepali_hunspell', 'कथाहरू');
```

This extension provides the `nepali_hunspell` dictionary and text search configuration.
