## Usage

Sources:

- [README.md](https://github.com/ColtWindy/pg-kiwi/blob/24fba80bb601dc177cf6b851693258e7e2924d73/README.md)
- [pg_kiwi.control](https://github.com/ColtWindy/pg-kiwi/blob/24fba80bb601dc177cf6b851693258e7e2924d73/pg_kiwi.control)
- [pg_kiwi--1.0.sql](https://github.com/ColtWindy/pg-kiwi/blob/24fba80bb601dc177cf6b851693258e7e2924d73/pg_kiwi--1.0.sql)
- [pg_kiwi.c](https://github.com/ColtWindy/pg-kiwi/blob/24fba80bb601dc177cf6b851693258e7e2924d73/pg_kiwi.c)
- [Changelog](https://github.com/ColtWindy/pg-kiwi/blob/24fba80bb601dc177cf6b851693258e7e2924d73/CHANGELOG.md)
- [Kiwi 0.22.2 license](https://raw.githubusercontent.com/bab2min/Kiwi/v0.22.2/LICENSE)

`pg_kiwi` 1.0 integrates the Kiwi Korean morphological analyzer with PostgreSQL text search. Upstream source instructions target PostgreSQL 18 and Kiwi 0.22.2. The server also needs readable Kiwi model files.

### Core workflow

```sql
CREATE EXTENSION pg_kiwi WITH SCHEMA public;
CREATE TABLE korean_docs (
 id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
 body text,
 terms tsvector GENERATED ALWAYS AS
   (to_tsvector('public.korean', body)) STORED
);
CREATE INDEX ON korean_docs USING gin (terms);
SELECT id FROM korean_docs
WHERE terms @@ to_tsquery('public.korean', '불국사');
```

### Configuration and maintenance

The extension installs the `kiwi_parser` text-search parser and `korean` configuration, mapping its token classes to the simple dictionary. Store a generated `tsvector` and index it with GIN to avoid re-analyzing every document for each query.

`pg_kiwi.model_path` selects model files and is reloadable; `pg_kiwi.pos_filter` controls part-of-speech filtering per session. Changing tokenization or models requires reconsidering stored vectors and rebuilding affected search data. Creation requires a superuser; no shared preload is declared for this extension. The README also shows optional `pg_textsearch` BM25 integration: that extension has its own compatibility and preload requirements. The core GIN workflow does not depend on it.

The control version is still 1.0; the 1.1.0 changelog entry changes the bundled BM25 image rather than the Kiwi parser. The extension source is MIT-licensed; the dynamically linked Kiwi 0.22.2 library is LGPL-2.1-or-later.
