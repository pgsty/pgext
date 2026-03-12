

## Usage

> [pg_pinyin: Pinyin romanization and search helpers for PostgreSQL](https://github.com/aiyou178/pg_pinyin)

Convert Chinese characters to Pinyin romanization for search and indexing. Works well with `pg_trgm` for fuzzy Pinyin search or `pg_search` for word-based search.

```sql
CREATE EXTENSION pg_pinyin;
```

### Functions

| Function | Description |
|----------|-------------|
| `pinyin_char_romanize(text)` | Character-level Pinyin romanization |
| `pinyin_char_romanize(text, suffix)` | With custom dictionary suffix |
| `pinyin_word_romanize(text)` | Word-level Pinyin romanization |
| `pinyin_word_romanize(text, suffix)` | With custom dictionary suffix |

### Generated Column + Trigram Search

```sql
CREATE EXTENSION IF NOT EXISTS pg_pinyin;
CREATE EXTENSION IF NOT EXISTS pg_trgm;

CREATE TABLE voice (
  id bigserial PRIMARY KEY,
  description text NOT NULL,
  pinyin text GENERATED ALWAYS AS (public.pinyin_char_romanize(description)) STORED
);

CREATE INDEX voice_pinyin_trgm_idx ON voice USING gin (pinyin gin_trgm_ops);

INSERT INTO voice (description) VALUES ('郑爽ABC');
SELECT id, description, pinyin FROM voice;
```

### Custom Dictionary

Provide custom dictionary tables in schema `pinyin` with a suffix:

```sql
CREATE TABLE IF NOT EXISTS pinyin.pinyin_mapping_suffix1 (
  character text PRIMARY KEY,
  pinyin text NOT NULL
);

CREATE TABLE IF NOT EXISTS pinyin.pinyin_words_suffix1 (
  word text PRIMARY KEY,
  pinyin text NOT NULL
);

INSERT INTO pinyin.pinyin_mapping_suffix1 (character, pinyin)
VALUES ('郑', '|zhengx|')
ON CONFLICT (character) DO UPDATE SET pinyin = EXCLUDED.pinyin;

-- Use custom dictionary
SELECT public.pinyin_char_romanize('郑爽ABC', '_suffix1');
```
