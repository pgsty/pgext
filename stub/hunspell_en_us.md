

## Usage

> [hunspell_en_us: English (US) Hunspell dictionary for PostgreSQL](https://github.com/postgrespro/hunspell_dicts)

English (US) Hunspell dictionary and text search configuration for PostgreSQL full-text search.

```sql
CREATE EXTENSION hunspell_en_us;

SELECT ts_lexize('english_hunspell', 'stories');
-- {story}

SELECT to_tsvector('english_hunspell', 'stories');
-- 'story':1
```

This extension provides the `english_hunspell` dictionary and text search configuration.
