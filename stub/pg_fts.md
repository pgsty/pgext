## Usage

Sources:

- [Official v0.2.0 README](https://codeberg.org/gregburd/pg_fts/src/tag/v0.2.0/README.md)
- [v0.2.0 changelog](https://codeberg.org/gregburd/pg_fts/src/tag/v0.2.0/CHANGELOG.md)
- [v0.2.0 SQL API](https://codeberg.org/gregburd/pg_fts/src/tag/v0.2.0/pg_fts--0.2.0.sql)
- [v0.2.0 control file](https://codeberg.org/gregburd/pg_fts/src/tag/v0.2.0/pg_fts.control)

`pg_fts` provides BM25/BM25F full-text ranking through dedicated `ftsdoc` and `ftsquery` types and an `fts` inverted-index access method. It supports boolean, phrase, NEAR, prefix, fuzzy, and regular-expression terms while keeping corpus statistics in the index for relevance scoring. Version `0.2.0` requires PostgreSQL 17 or newer.

### Create and Query an Index

```sql
CREATE EXTENSION pg_fts;

CREATE TABLE docs (
    id bigint PRIMARY KEY,
    body text NOT NULL
);

CREATE INDEX docs_fts
ON docs USING fts (to_ftsdoc('english', body));
```

Use the same text-search configuration for documents and ordinary query terms:

```sql
WITH q AS (
    SELECT to_ftsquery('english', 'postgres & "query planner" & index*') AS query
)
SELECT d.id,
       fts_snippet(d.body, q.query) AS excerpt
FROM docs AS d
CROSS JOIN q
WHERE to_ftsdoc('english', d.body) @@@ q.query
ORDER BY to_ftsdoc('english', d.body) <=> q.query
LIMIT 10;
```

`@@@` matches, while ascending `<=>` distance orders rows by descending relevance and can drive an index ordering scan for top-k queries.

### Query Language and API Index

- `to_ftsdoc([regconfig,] text)` and `to_ftsquery([regconfig,] text)`: analyze documents and parse queries.
- `quick brown`, `quick & brown`, `quick | brown`, and `!slow`: implicit/explicit AND, OR, and NOT.
- `"quick brown"`, `NEAR(...)`, `term*`, `term~2`, and `/regular-expression/`: phrase, proximity, prefix, fuzzy, and regex terms.
- `fts_bm25`, `fts_bm25_opts`, and `fts_bm25f`: explicit BM25 scoring variants and multi-field scoring.
- `fts_index_stats(index)` and `fts_index_df(index, query)`: index-maintained document count, average length, vocabulary size, and term frequencies.
- `fts_highlight` and `fts_snippet`: present matching text.
- `fts_search(index, query, k)` and `fts_count(index, query)`: index-native top-k and MVCC-aware count operations.
- `tsquery_to_ftsquery(tsquery)`: migration helper; it does not make `pg_fts` a transparent replacement for `tsvector`/GIN.

### Maintenance and Version Caveats

```sql
SELECT fts_merge('docs_fts');
SELECT fts_vacuum('docs_fts');
```

- Inserts enter an immediately matchable pending list, but ranked `<=>` and `fts_search` results cover merged segments. Run `fts_merge()` when newly inserted documents must participate in ranking immediately.
- `fts_vacuum()` compacts segments and truncates reclaimable index pages; ordinary `VACUUM` also participates in pending-list and tombstone maintenance.
- Version `0.2.0` renamed the access method from `bm25` to `fts`. Indexes created by `0.1.0` with `USING bm25` must be recreated.
- If the library reports an on-disk format mismatch, follow its `REINDEX` hint rather than attempting to read the index with a different format version.
- The access method is non-covering and does not provide parallel scans in this release. Provision the extension and index separately on logical-replication subscribers; indexes themselves are not logically replicated.

