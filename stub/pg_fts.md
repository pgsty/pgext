## Usage

Sources:

- [README.md](https://github.com/gburd/pg_fts/blob/d5c645f615b7a4f4e1a9374c30200a33858be401/README.md)
- [pg_fts.control](https://github.com/gburd/pg_fts/blob/d5c645f615b7a4f4e1a9374c30200a33858be401/pg_fts.control)
- [CHANGELOG.md](https://github.com/gburd/pg_fts/blob/d5c645f615b7a4f4e1a9374c30200a33858be401/CHANGELOG.md)
- [pg_fts--1.8.6--1.9.0.sql](https://github.com/gburd/pg_fts/blob/d5c645f615b7a4f4e1a9374c30200a33858be401/pg_fts--1.8.6--1.9.0.sql)

`pg_fts` 1.9.0 provides full-text search with BM25/BM25F ranking and the dedicated fts inverted index. PostgreSQL 17 and 18 are supported; PostgreSQL 19/master CI is best-effort. The control is trusted and relocatable, and core use needs no shared preload.

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

### Upgrade to 1.9.0

After installing matching files, run `ALTER EXTENSION pg_fts UPDATE TO '1.9.0'`. This release fixes incorrect ranking after deletes and missed top-k results at posting-block boundaries. There is no on-disk format change and no REINDEX requirement. `pg_fts.doclen_cache_mb` defaults to 64 (0 disables the cache); `pg_fts.dense_score_min_df` defaults to 32768 (0 disables that scoring path). Account for per-backend cache memory.
