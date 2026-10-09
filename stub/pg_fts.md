## Usage

Sources:

- [README.md](https://github.com/gburd/pg_fts/blob/8bccf0ce75f04b3af559b34d635a5f71e0d2d04e/README.md)
- [pg_fts.control](https://github.com/gburd/pg_fts/blob/8bccf0ce75f04b3af559b34d635a5f71e0d2d04e/pg_fts.control)
- [CHANGELOG.md](https://github.com/gburd/pg_fts/blob/8bccf0ce75f04b3af559b34d635a5f71e0d2d04e/CHANGELOG.md)
- [pg_fts--1.9.1--1.10.0.sql](https://github.com/gburd/pg_fts/blob/8bccf0ce75f04b3af559b34d635a5f71e0d2d04e/pg_fts--1.9.1--1.10.0.sql)

`pg_fts` 1.10.0 provides full-text search with BM25/BM25F ranking and the dedicated fts inverted index. PostgreSQL 17 and 18 are supported; PostgreSQL 19/master CI is best-effort. The control is trusted and relocatable, and core use needs no shared preload.

### Search and Rank

```sql
CREATE EXTENSION pg_fts;
CREATE TABLE docs (id bigint, body text);
CREATE INDEX docs_fts ON docs USING fts (to_ftsdoc('english', body));
SELECT id FROM docs
 WHERE to_ftsdoc('english', body) @@@ to_ftsquery('english', 'quick fox')
 ORDER BY to_ftsdoc('english', body) <=> to_ftsquery('english', 'quick fox')
 LIMIT 10;
SELECT fts_merge('docs_fts');
SELECT fts_vacuum('docs_fts');
```

### Objects and Queries

`ftsdoc` and `ftsquery` represent analyzed documents and queries. `to_ftsdoc()` and `to_ftsquery()` construct them; `@@@` matches documents and `<=>` orders relevance distance. The match predicate is required for an indexed KNN ordering scan. Queries support boolean terms, phrases, prefix, fuzzy and regular-expression matching. Multi-column documents support BM25F field weighting. `fts_count()` and ordinary count queries can use exact index counting; `fts_search()` exposes direct ranked results. Regex/long-fuzzy acceleration is opt-in through `trigrams = on`.

### Maintenance and Privileges

Pending inserts are immediately searchable; merging is not a visibility prerequisite. `fts_merge()` compacts segments and `fts_vacuum()` reclaims physical space. Both write WAL, require index ownership and must run on a primary. Large pending documents can temporarily consume substantial space, so merge and reclaim periodically during bulk ingestion. `fts_search()` and `fts_anomalous_docs()` expose indexed content and are revoked from PUBLIC by default; widen access only deliberately. Regular table-query visibility and direct helper access are different permission surfaces.

### Upgrade to 1.10.0

After installing matching files, run `ALTER EXTENSION pg_fts UPDATE TO '1.10.0'`. Version 1.10.0 shares each segment's document-length array across backends by default and adds `fts_shared_doclen_stats()` to inspect the shared cache. A superuser can change `pg_fts.shared_doclen`; unavailable shared memory falls back to a private copy. There is no on-disk format change and no REINDEX requirement. Earlier delete-ranking and posting-block fixes remain included; distinguish shared-cache memory from fallback per-backend copies when planning capacity.
