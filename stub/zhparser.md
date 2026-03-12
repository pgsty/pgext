

## Usage

> [GitHub: amutu/zhparser](https://github.com/amutu/zhparser)

`zhparser` is a PostgreSQL extension for full-text search of Chinese, based on the Simple Chinese Word Segmentation (SCWS) library.

## Features

- Chinese text segmentation for PostgreSQL full-text search
- Built on the SCWS (Simple Chinese Word Segmentation) library
- Supports custom dictionaries (TXT and XDB formats)
- Database-level custom word tables (since v2.1)
- Multiple tunable parameters for segmentation behavior

## Quick Start

```sql
-- Create the extension
CREATE EXTENSION zhparser;

-- Create a text search configuration using zhparser
CREATE TEXT SEARCH CONFIGURATION chinese (PARSER = zhparser);

-- Add token type mappings
ALTER TEXT SEARCH CONFIGURATION chinese ADD MAPPING FOR n,v,a,i,e,l WITH simple;

-- Test Chinese text segmentation
SELECT to_tsvector('chinese', '小明硕士毕业于中国科学院计算所，后在日本京都大学深造');

-- Create a table and index for Chinese full text search
CREATE TABLE articles (id serial PRIMARY KEY, title text, body text);

CREATE INDEX articles_body_idx ON articles
  USING gin (to_tsvector('chinese', body));

-- Query with Chinese full text search
SELECT * FROM articles
  WHERE to_tsvector('chinese', body) @@ to_tsquery('chinese', '中国');
```

## Configuration Parameters

zhparser provides several GUC parameters to control segmentation behavior:

| Parameter | Default | Description |
|-----------|---------|-------------|
| `zhparser.punctuation_ignore` | `off` | Ignore all punctuation |
| `zhparser.seg_with_duality` | `off` | Perform duality segmentation on long words |
| `zhparser.dict_in_memory` | `off` | Load the whole dictionary into memory |
| `zhparser.multi_short` | `off` | Short word compound segmentation |
| `zhparser.multi_duality` | `off` | Duality compound segmentation |
| `zhparser.multi_zmain` | `off` | Key word in first compound segmentation |
| `zhparser.multi_zall` | `off` | Use all compound segmentation |

## Token Types

zhparser supports the following token types from SCWS:

| Code | Description |
|------|-------------|
| `a` | Adjective |
| `b` | Differentiation (区别词) |
| `c` | Conjunction |
| `d` | Adverb |
| `e` | Exclamation |
| `f` | Position word (方位词) |
| `g` | Root word (词根) |
| `h` | Prefix |
| `i` | Idiom |
| `j` | Abbreviation |
| `k` | Suffix |
| `l` | Temporary idiom |
| `m` | Numeral |
| `n` | Noun |
| `o` | Onomatopoeia |
| `p` | Preposition |
| `q` | Classifier |
| `r` | Pronoun |
| `s` | Space word (处所词) |
| `t` | Time word |
| `u` | Auxiliary |
| `v` | Verb |
| `w` | Punctuation |
| `x` | Unknown |
| `y` | Modal particle |
| `z` | Status word (状态词) |

## Custom Dictionaries

### File-based Dictionaries

Place custom dictionary files in the share directory (typically `$SHAREDIR/tsearch_data/`):

- TXT format: one word per line
- XDB format: compiled SCWS dictionary format

Custom dictionaries take precedence over built-in dictionaries.

### Database-level Custom Words (v2.1+)

```sql
-- Add custom words via zhparser's built-in table
INSERT INTO zhparser.zhprs_custom_word VALUES ('中国科学院计算所');

-- Reload custom dictionary (reconnect after sync to take effect)
SELECT sync_zhprs_custom_word();

-- Verify segmentation with custom word
SELECT to_tsvector('chinese', '小明硕士毕业于中国科学院计算所');
```

## Docker Quick Start

```bash
docker run --name pgzhparser -d \
  -e POSTGRES_PASSWORD=somepassword \
  zhparser/zhparser:bookworm-16
```
