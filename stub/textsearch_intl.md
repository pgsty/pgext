## Usage

Sources:

- [README.md](https://github.com/benjaminsanborn/textsearch_intl/blob/6efb0a28a5f6643b46ecd2e9b8de75c30fb82680/README.md)
- [textsearch_intl--1.0.sql](https://github.com/benjaminsanborn/textsearch_intl/blob/6efb0a28a5f6643b46ecd2e9b8de75c30fb82680/textsearch_intl--1.0.sql)
- [textsearch_intl.control](https://github.com/benjaminsanborn/textsearch_intl/blob/6efb0a28a5f6643b46ecd2e9b8de75c30fb82680/textsearch_intl.control)
- [META.json](https://github.com/benjaminsanborn/textsearch_intl/blob/6efb0a28a5f6643b46ecd2e9b8de75c30fb82680/META.json)
- [LICENSE](https://github.com/benjaminsanborn/textsearch_intl/blob/6efb0a28a5f6643b46ecd2e9b8de75c30fb82680/LICENSE)
- [Rename commit](https://github.com/benjaminsanborn/textsearch_intl/commit/b023b27aa16b82fc3c1b80ef0928e9f185e7f43e)
- [PGXN textsearch_intl 1.0.1](https://pgxn.org/dist/textsearch_intl/1.0.1/)

`textsearch_intl` combines ICU word segmentation, normalization and n-gram matching for multilingual text search. It requires PostgreSQL 14 or later, a UTF8 database, ICU and `unaccent`; upstream CI covers PostgreSQL 14–18.

### Name and Version

The project renamed `icu_parser` to `textsearch_intl` on 2026-10-09 and restarted versioning. The renamed PGXN distribution 1.0.1 was published on 2026-10-22. Its SQL/control version remains `1.0`; the distribution number is not a SQL upgrade target. There is no direct SQL upgrade script from the old name: inspect dependent objects and plan their recreation before migrating.

### Core Workflow

```sql
CREATE EXTENSION textsearch_intl CASCADE;
SELECT alias, token FROM ts_debug('intl_simple', '東京駅前ラーメン店 ร้านกาแฟ Café');
SELECT to_tsvector('intl_search', 'Café') @@ phraseto_tsquery('intl_search', 'cafe');
SELECT to_tsvector('intl_search_ngram', '행복한편의점')
       @@ phraseto_tsquery('intl_search_ngram', '편의점');
SELECT intl_icu_version();
```

### Search Configurations

- `intl_simple` uses the `intl` parser and preserves accents and character widths.
- `intl_search` combines `intl_normalize`, `unaccent` and `intl_join` for normalized whole-word search. Word joining is off by default.
- `intl_search_ngram` uses `intl_runs` and `intl_ngram` to match inside Korean/CJK tokens and unspaced Thai, Lao, Khmer and Burmese runs. N-grams can overmatch and increase index size.

### Operational Boundaries

Install with a superuser; the extension is relocatable, is not marked trusted and needs no preload. Prefer `phraseto_tsquery()` when word order matters. `intl_icu_version()` reports the loaded ICU version; rebuild affected indexes after an ICU major-version change. ICU segmentation is not Japanese morphological analysis, and partial CJK prefix queries remain unreliable.
