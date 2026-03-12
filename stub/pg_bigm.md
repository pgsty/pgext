

## Usage

> [pg_bigm Documentation](https://pgbigm.osdn.jp/pg_bigm_en-1-2.html) | [GitHub: pgbigm/pg_bigm](https://github.com/pgbigm/pg_bigm)

The `pg_bigm` module provides full text search capability in PostgreSQL. This module allows a user to create **2-gram** (bigram) index for faster full text search.

pg_bigm is released under the PostgreSQL License, a liberal Open Source license, similar to the BSD or MIT licenses.

## Features

- **Bigram indexing**: Creates 2-gram (bigram) GIN indexes for text columns
- **Faster LIKE searches**: Accelerates `LIKE` queries including prefix, suffix, and substring searches
- **All language support**: Works with any language including CJK (Chinese, Japanese, Korean) without additional configuration
- **Simple API**: Provides similarity search functions and operators

## Functions and Operators

### Functions

| Function | Return Type | Description |
|----------|-------------|-------------|
| `likequery(text)` | `text` | Generates a search query for full text search from a keyword |
| `show_bigm(text)` | `text[]` | Shows all 2-grams in the given string |
| `pg_gin_pending_stats(regclass)` | `record` | Returns the number of pages and tuples in the pending list of a GIN index |

### Operators

| Operator | Description |
|----------|-------------|
| `text =% text` | Returns true if the similarity between the left and right operands is greater than or equal to `pg_bigm.similarity_limit` |

## GUC Parameters

| Parameter | Type | Default | Description |
|-----------|------|---------|-------------|
| `pg_bigm.last_update` | `text` | - | Shows the last update date of the module (read-only) |
| `pg_bigm.enable_recheck` | `bool` | `on` | Controls whether recheck is performed during index scan |
| `pg_bigm.gin_key_limit` | `int` | `0` | Limits the maximum number of bigrams used for full text search. 0 means no limit |
| `pg_bigm.similarity_limit` | `real` | `0.3` | Sets the minimum similarity threshold for the `=%` operator |

## Examples

### Basic Full Text Search

```sql
-- Create extension
CREATE EXTENSION pg_bigm;

-- Create a table with text data
CREATE TABLE documents (id serial PRIMARY KEY, content text);
INSERT INTO documents (content) VALUES
  ('PostgreSQL is a powerful database'),
  ('Full text search with bigram indexing'),
  ('Japanese text: 日本語テキスト検索');

-- Create a bigram index
CREATE INDEX docs_bigm_idx ON documents USING gin (content gin_bigm_ops);

-- Search using LIKE
SELECT * FROM documents WHERE content LIKE '%search%';

-- Search using likequery function
SELECT * FROM documents WHERE content LIKE likequery('database');
```

### Similarity Search

```sql
-- Show bigrams for a string
SELECT show_bigm('PostgreSQL');

-- Similarity search
SET pg_bigm.similarity_limit = 0.2;
SELECT * FROM documents WHERE content =% 'database search';
```

### Comparison with pg_trgm

pg_bigm has the following advantages over the built-in `pg_trgm`:

| Feature | pg_bigm | pg_trgm |
|---------|---------|---------|
| N-gram type | 2-gram (bigram) | 3-gram (trigram) |
| Minimum search string | 1 character | 3 characters |
| Non-alphabetic languages | Full support | Limited support |
| LIKE search types | Prefix, suffix, and substring | Prefix, suffix, and substring |

For detailed documentation including advanced usage and performance tuning, see the [official pg_bigm documentation](https://pgbigm.osdn.jp/pg_bigm_en-1-2.html).
