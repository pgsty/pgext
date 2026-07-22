## Usage

Sources:

- [Official README](https://github.com/fairingrey/pg_simple_parser/blob/cdeaa0f1ed3b89d1961312c01e529f1c7c068084/README.md)
- [Extension control file](https://github.com/fairingrey/pg_simple_parser/blob/cdeaa0f1ed3b89d1961312c01e529f1c7c068084/simple_parser.control)
- [Version 1.0 extension SQL](https://github.com/fairingrey/pg_simple_parser/blob/cdeaa0f1ed3b89d1961312c01e529f1c7c068084/simple_parser--1.0.sql)
- [C tokenizer implementation](https://github.com/fairingrey/pg_simple_parser/blob/cdeaa0f1ed3b89d1961312c01e529f1c7c068084/simple_parser.c)

`simple_parser` 1.0 installs the full-text-search parser `simpleparser`. It treats every run of bytes between ASCII space characters as one `word`, preserving punctuation such as underscores and hyphens inside the token.

### Create a Text Search Configuration

The extension installs a parser, not a ready-to-use text search configuration. Create one and map its `word` token type to a dictionary:

```sql
CREATE EXTENSION simple_parser;

CREATE TEXT SEARCH CONFIGURATION simple_space (
  PARSER = simpleparser
);

ALTER TEXT SEARCH CONFIGURATION simple_space
ADD MAPPING FOR word WITH simple;

SELECT to_tsvector('simple_space', 'alpha_beta gamma-delta');
```

The parser also emits a `blank` token type so PostgreSQL's standard headline function can retain spacing, but it normally needs no dictionary mapping.

### Installed Objects

- `simpleparser` is the public text search parser.
- `simpleprs_start`, `simpleprs_getlexeme`, `simpleprs_end`, and `simpleprs_lextype` are internal C parser callbacks; call PostgreSQL text-search functions rather than invoking them directly.
- The parser reuses PostgreSQL token IDs 3 (`word`) and 12 (`blank`) and delegates headline generation to `pg_catalog.prsd_headline`.

### Tokenization Caveats

Only byte value ASCII space (`0x20`) is a separator. Tabs, newlines, non-breaking spaces, and other Unicode whitespace remain inside a `word`, as do punctuation and control bytes. The parser does not perform linguistic segmentation, Unicode normalization, case folding, stemming, or stop-word removal; those behaviors come only from the mapped dictionaries.

Inspect actual tokens before indexing application data:

```sql
SELECT t.alias, p.token
FROM ts_parse('simpleparser', E'alpha_beta\tgamma delta') AS p
JOIN ts_token_type('simpleparser') AS t USING (tokid);
```

Changing a configuration or dictionary after building a `tsvector` index requires regenerating stored vectors and rebuilding the affected index. Use this parser only when literal ASCII-space token boundaries match the application's search semantics.
