## Usage

Sources:

- [Official README](https://github.com/postgrespro/pg_tsparser/blob/11ed232e83dfe7e900de084f09219041f880d1a2/README.md)
- [Official extension SQL](https://github.com/postgrespro/pg_tsparser/blob/11ed232e83dfe7e900de084f09219041f880d1a2/pg_tsparser--1.0.sql)
- [Official parser implementation](https://github.com/postgrespro/pg_tsparser/blob/11ed232e83dfe7e900de084f09219041f880d1a2/tsparser.c)

`pg_tsparser` provides the `tsparser` full-text parser, a modified PostgreSQL 9.6 default parser that additionally emits underscore-linked words and mixed letter/number hyphenated words as whole tokens. Use it to make identifiers such as `pg_trgm` or `123-abc` searchable without losing their components.

### Create a Configuration

Installing the extension creates the parser but not a complete text search configuration. Map its token types to dictionaries explicitly:

```sql
CREATE EXTENSION pg_tsparser;

CREATE TEXT SEARCH CONFIGURATION english_ts (
  PARSER = tsparser
);

ALTER TEXT SEARCH CONFIGURATION english_ts
  ADD MAPPING FOR email, file, float, host, hword_numpart, int,
  numhword, numword, sfloat, uint, url, url_path, version
  WITH simple;

ALTER TEXT SEARCH CONFIGURATION english_ts
  ADD MAPPING FOR asciiword, asciihword, hword_asciipart,
  word, hword, hword_part
  WITH english_stem;
```

Compare tokenization before adopting the configuration:

```sql
SELECT to_tsvector('english', 'pg_trgm 123-abc') AS default_tokens,
       to_tsvector('english_ts', 'pg_trgm 123-abc') AS linked_tokens;
```

### Object Index

- `tsparser` is the text search parser.
- `tsparser_start()`, `tsparser_nexttoken()`, `tsparser_end()`, `tsparser_lextype()`, and `tsparser_headline()` implement the parser callbacks.

### Caveats

The implementation is forked from PostgreSQL 9.6 parser internals, so binary/source compatibility must be verified for every target major version. Extra whole-word tokens change positions, ranking, phrase behavior, index size, and query results; rebuild dependent `tsvector` columns and indexes after configuration changes. Test punctuation, version strings, paths, URLs, email, non-ASCII text, stemming, dictionaries, headlines, and phrase queries against a representative corpus before replacing an existing configuration.
