

## Usage

> [dict_xsyn: extended synonym dictionary for text search](https://www.postgresql.org/docs/current/dict-xsyn.html)

Provides an extended synonym dictionary template for text search, replacing words with groups of synonyms.

```sql
CREATE EXTENSION dict_xsyn;
```

### Configuration Parameters

| Parameter | Description | Default |
|---|---|---|
| `matchorig` | Accept original word | `true` |
| `matchsynonyms` | Accept synonyms as input | `false` |
| `keeporig` | Include original in output | `true` |
| `keepsynonyms` | Include synonyms in output | `true` |
| `rules` | Base name of synonym file in `$SHAREDIR/tsearch_data/` (`.rules` extension) | -- |

### Rules File Format

```
word syn1 syn2 syn3
```

Lines starting with `#` are comments.

### Examples

```sql
-- Configure the dictionary
ALTER TEXT SEARCH DICTIONARY xsyn (RULES='my_rules', KEEPORIG=true);

-- Test the dictionary
SELECT ts_lexize('xsyn', 'word');
-- {word,syn1,syn2,syn3}

-- Match synonyms as input too
ALTER TEXT SEARCH DICTIONARY xsyn (RULES='my_rules', MATCHSYNONYMS=true);
SELECT ts_lexize('xsyn', 'syn1');
-- {syn1,syn2,syn3}

-- Use in a text search configuration
ALTER TEXT SEARCH CONFIGURATION english
  ALTER MAPPING FOR word, asciiword WITH xsyn, english_stem;
```
