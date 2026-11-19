## Usage

Sources:

- [README.md](https://github.com/erinshek/pg-kaasearch/blob/v1.0.0/README.md)
- [pg_kaasearch.control](https://github.com/erinshek/pg-kaasearch/blob/v1.0.0/pg_kaasearch.control)
- [pg_kaasearch--1.0.sql](https://github.com/erinshek/pg-kaasearch/blob/v1.0.0/pg_kaasearch--1.0.sql)
- [LICENSE](https://github.com/erinshek/pg-kaasearch/blob/v1.0.0/LICENSE)

`pg_kaasearch` provides Karakalpak full-text search, including stemming, stop words and normalization of Latin, Cyrillic and apostrophe spellings. Distribution 1.0.0 installs SQL version 1.0. PostgreSQL 15–18 and a UTF8 database are required.

### Enable and Search

Installation requires a privileged extension installer; the control file does not mark the extension trusted. It is relocatable and needs no preload.

```sql
CREATE EXTENSION pg_kaasearch;
SELECT to_tsvector('kaasearch', 'Мектепке барған балалар китап оқыйды.')
    @@ to_tsquery('kaasearch', 'mektep & kitap');
SELECT kaa_stem('mektebimizde');
```

Use the same text-search configuration in an expression index and its queries:

```sql
CREATE TABLE articles (id bigint PRIMARY KEY, title text, body text);
CREATE INDEX articles_fts ON articles USING gin (to_tsvector('kaasearch', body));
SELECT title FROM articles
WHERE to_tsvector('kaasearch', body)
   @@ websearch_to_tsquery('kaasearch', 'mektep kitap');
```

### Objects and Dictionaries

`kaasearch` is the text-search configuration; `kaasearch_parser`, `kaasearch_template` and `kaasearch_dict` provide its parser and dictionary. `kaa_stem` returns a stem, `kaa_normalize` normalizes spelling, and `kaa_analyze` returns the normalized word, stem, segments, classes and source. Standard `ts_rank` and `ts_headline` work with the configuration.

Stop words come from the installed `kaasearch.stop` file. A custom lexicon can be configured with `ALTER TEXT SEARCH DICTIONARY kaasearch_dict (Lexicon = name)`; the corresponding server-side dictionary file must exist. Rebuild affected text-search indexes after changing tokenization or dictionary behavior. Schema-qualify the configuration when installing outside the search path.
