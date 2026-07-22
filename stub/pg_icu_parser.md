## Usage

Sources:

- [Official PGXN documentation](https://pgxn.org/dist/pg_icu_parser/1.0.0/)

`pg_icu_parser` supplies the `icu_parser` full-text-search parser, which uses ICU Unicode word-boundary analysis. It is useful for languages such as Chinese, Japanese, Korean, and Hebrew where PostgreSQL's whitespace- and pattern-oriented default parser may not produce suitable word boundaries.

### Core Workflow

The server must be PostgreSQL 10 or later and must have been built with ICU support. Create the extension, then create a text-search configuration that selects its parser and maps the token types you need:

```sql
CREATE EXTENSION pg_icu_parser;

CREATE TEXT SEARCH CONFIGURATION sample (PARSER = icu_parser);
ALTER TEXT SEARCH CONFIGURATION sample
  ADD MAPPING FOR word WITH english_stem;
ALTER TEXT SEARCH CONFIGURATION sample
  ADD MAPPING FOR number, ideo, kana WITH simple;

SELECT to_tsvector('sample', '東京でPostgreSQLを使う');
```

Choose dictionaries for the actual language; `english_stem` above only demonstrates the mapping syntax.

### Parser Surface

- `word`, `number`, `kana`, `ideo`, and `blank` are the token types emitted by `icu_parser`.
- `pg_icu_parser.locale` selects the ICU locale and defaults to `en`. Upstream notes that word-boundary rules are usually not locale-sensitive, so change it only after testing a concrete need.
- Any PostgreSQL text-search dictionary can be mapped to the emitted token types.

### Trade-offs

- The parser requires ICU in the PostgreSQL build; creating the extension cannot add ICU to a server built without it.
- It emits fewer and coarser token types than the default parser and does not recognize special patterns such as URLs or email addresses as distinct categories.
- The reviewed release does not support `ts_headline`.
- Upstream reports that it is slower than the default parser, with a lower conversion overhead when the database encoding is UTF-8. Benchmark representative documents and query shapes on the target system.
- When parser mappings or locale behavior change, rebuild stored `tsvector` values and dependent indexes so old and new rows do not use different tokenization.
