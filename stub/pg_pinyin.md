


## Usage

> Sources: [pg_pinyin upstream README](https://github.com/aiyou178/pg_pinyin), [Chinese README](https://github.com/aiyou178/pg_pinyin/blob/main/README.zh-CN.md).

`pg_pinyin` converts Chinese text to Pinyin, either character by character or by word. It is useful for generated search columns, trigram search, and `pg_search` BM25 queries that need Pinyin input.

```sql
CREATE EXTENSION pg_pinyin;
```

### Functions

| Function | Description |
|----------|-------------|
| `pinyin_char_romanize(text)` | Character-level Pinyin romanization |
| `pinyin_char_romanize(text, suffix text)` | Character-level romanization with a custom dictionary suffix |
| `pinyin_word_romanize(text)` | Word-level Pinyin romanization |
| `pinyin_word_romanize(text, suffix text)` | Word-level romanization with a custom dictionary suffix |
| `pinyin_word_romanize(tokenizer_input anyelement)` | Word-level romanization from a `pg_search` tokenizer input such as `name::pdb.icu::text[]` |
| `pinyin_word_romanize(tokenizer_input anyelement, suffix text)` | Tokenizer-input romanization with a custom dictionary suffix |
| `pinyin_regex_phrase(text, slope integer DEFAULT NULL, max_expansions integer DEFAULT NULL, generated_pinyin boolean DEFAULT false)` | `pg_search` query helper returning `pdb.query`, available when `pg_search` was enabled before `CREATE EXTENSION pg_pinyin` |
| `pinyin_regex_phrase_patterns(text, generated_pinyin boolean DEFAULT false)` | Internal helper returning regex phrase tokens as `text[]` |

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

### Word Tokenization + pg_search

For word-oriented search, use `pinyin_word_romanize`. When `pg_search` is available, it can consume tokenizer input such as `pdb.icu::text[]`.

```sql
CREATE EXTENSION IF NOT EXISTS pg_search;
CREATE EXTENSION IF NOT EXISTS pg_pinyin;

CREATE TABLE voice (
  id bigserial PRIMARY KEY,
  description text NOT NULL,
  pinyin text GENERATED ALWAYS AS (public.pinyin_word_romanize(description)) STORED
);

CREATE INDEX voice_pinyin_bm25_idx
ON voice
USING bm25 (id, pinyin)
WITH (key_field='id');

SELECT *
FROM voice
WHERE pinyin @@@ public.pinyin_regex_phrase('zhengshuang');

SELECT public.pinyin_word_romanize('郑爽ABC'::pdb.icu::text[]);
```

`pinyin_regex_phrase` has return type `pdb.query`, so `pg_search` must be enabled in the database before `pg_pinyin` is created. If `pg_pinyin` is created first, upstream documents that the romanization functions are installed, but `pinyin_regex_phrase` is installed as an error stub with a clear exception.

### Dictionary Tables

The extension seeds bundled dictionary tables under schema `pinyin` during `CREATE EXTENSION pg_pinyin`; no separate data-load step is needed for normal extension usage. The bundled data covers character mappings, word tokens, and word mappings.

Provide custom dictionary tables in schema `pinyin` with a suffix. Calls using that suffix merge the base dictionary with the suffix tables, and suffix entries take priority.

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

INSERT INTO pinyin.pinyin_words_suffix1 (word, pinyin)
VALUES ('郑爽', '|zhengx| |shuangx|')
ON CONFLICT (word) DO UPDATE SET pinyin = EXCLUDED.pinyin;

SELECT public.pinyin_char_romanize('郑爽ABC', '_suffix1');
SELECT public.pinyin_word_romanize('郑爽ABC'::pdb.icu::text[], '_suffix1');
```
