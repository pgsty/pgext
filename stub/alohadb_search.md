## Usage

Sources:

- [README.md](https://github.com/Vern-AllWorks-LLC/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/README.md)
- [contrib/alohadb_search/alohadb_search--1.0.sql](https://github.com/Vern-AllWorks-LLC/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/contrib/alohadb_search/alohadb_search--1.0.sql)
- [contrib/alohadb_search/search_bm25.c](https://github.com/Vern-AllWorks-LLC/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/contrib/alohadb_search/search_bm25.c)
- [contrib/alohadb_search/search_suggest.c](https://github.com/Vern-AllWorks-LLC/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/contrib/alohadb_search/search_suggest.c)
- [contrib/alohadb_search/meson.build](https://github.com/Vern-AllWorks-LLC/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/contrib/alohadb_search/meson.build)
- [contrib/alohadb_search/alohadb_search.control](https://github.com/Vern-AllWorks-LLC/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/contrib/alohadb_search/alohadb_search.control)

`alohadb_search` 1.0 is an AlohaDB module for text relevance, fuzzy and phonetic matching, suggestions and geographic distance. The reviewed source belongs to AlohaDB, based on PostgreSQL 18; portability to a stock PostgreSQL installation is not established.

### Core Workflow

```sql
CREATE EXTENSION alohadb_search;
SELECT search_bm25(to_tsvector('english', 'database search'),
                   to_tsquery('english', 'database'));
SELECT search_fuzzy_match('alohadb', 'alohdb', 2);
SELECT search_haversine(21.3, -157.8, 21.31, -157.81);
```

### Operational Boundaries

The module installs in `public`, is not relocatable and requires privileged installation. Its source defines no preload hook. `search_bm25()` scores a `tsvector` against a `tsquery` with optional `k1` and `b`, using fixed corpus-size and average-length estimates in this revision; it is not a new index access method or a corpus-statistics engine. `search_edit_distance()` and `search_fuzzy_match()` compare strings; `search_phonetic()` computes a phonetic code. `search_autocomplete()` reads a supplied table/column; `search_expand_synonyms()` uses `alohadb_search_synonyms`, and `search_analyze()` exposes text analysis. `search_haversine()` returns meters, while `search_nearby()` queries a table within a radius. Only pass trusted table expressions to the dynamic-query helpers: the source interpolates the table argument into SQL. Review privileges and the implementation before production use.
