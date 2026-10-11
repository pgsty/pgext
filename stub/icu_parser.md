## Usage

Sources:

- [README.md](https://github.com/benjaminsanborn/textsearch_intl/blob/d24405defbbfa249af420436b6b2442d8342246c/README.md)
- [icu_parser--1.1.sql](https://github.com/benjaminsanborn/textsearch_intl/blob/d24405defbbfa249af420436b6b2442d8342246c/icu_parser--1.1.sql)
- [LICENSE](https://github.com/benjaminsanborn/textsearch_intl/blob/d24405defbbfa249af420436b6b2442d8342246c/LICENSE)
- [icu_parser.control](https://github.com/benjaminsanborn/textsearch_intl/blob/d24405defbbfa249af420436b6b2442d8342246c/icu_parser.control)

`icu_parser` splits unspaced languages such as Chinese, Japanese and Thai using ICU word boundaries. The release distribution is 1.1.0; the SQL/control version is 1.1. It requires PostgreSQL 14 or later, a UTF8 database, ICU and `unaccent`.

### Core Workflow

```sql
CREATE EXTENSION icu_parser CASCADE;
SELECT alias, token FROM ts_debug('icu_simple', '東京駅前ラーメン店 ร้านกาแฟ Café');
SELECT to_tsvector('icu_search', 'Café') @@ phraseto_tsquery('icu_search', 'cafe');
SELECT icu_parser_icu_version();
```

### Operational Boundaries

The `icu` parser exposes word, number, kana, ideograph and blank tokens. `icu_simple` lowercases while preserving accents and widths. `icu_search` chains `icu_normalize`, `unaccent` and `simple`: normalization applies NFKC and apostrophe removal while preserving relevant Thai/Lao vowels. Prefer `phraseto_tsquery()` when token order matters. `icu_parser_icu_version()` reports the runtime ICU version. Rebuild affected indexes after changing ICU major versions, because segmentation can change. Dictionary segmentation is not Japanese morphological analysis or arbitrary substring search; partial CJK prefixes can fail to match. Native installation is not marked trusted and no preload is required.
