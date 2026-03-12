

## Usage

> [unaccent: text search dictionary for accent removal](https://www.postgresql.org/docs/current/unaccent.html)

Removes accents (diacritic signs) from text. Can be used as a text search dictionary or as a standalone function.

```sql
CREATE EXTENSION unaccent;
```

### Functions

| Function | Description |
|---|---|
| `unaccent(text)` | Remove accents using the default dictionary |
| `unaccent(dictionary regdictionary, text)` | Remove accents using a specific dictionary |

### Text Search Usage

```sql
-- Test the dictionary
SELECT ts_lexize('unaccent', 'Hôtel');
-- {Hotel}

-- Create an accent-insensitive French text search config
CREATE TEXT SEARCH CONFIGURATION fr (COPY = french);
ALTER TEXT SEARCH CONFIGURATION fr
  ALTER MAPPING FOR hword, hword_part, word
  WITH unaccent, french_stem;

SELECT to_tsvector('fr', 'Hôtels de la Mer');
-- 'hotel':1 'mer':4

-- Accent-insensitive search
SELECT to_tsvector('fr', 'Hôtel de la Mer') @@ to_tsquery('fr', 'Hotels');
-- true
```

### Standalone Function Usage

```sql
SELECT unaccent('Hôtel');           -- Hotel
SELECT unaccent('café');            -- cafe
SELECT unaccent('résumé');          -- resume
SELECT unaccent('niño');            -- nino
```
