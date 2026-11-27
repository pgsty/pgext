## Usage

Sources:

- [1.3.0 release notes](https://github.com/pgsty/vchord_bm25/blob/1.3.0/docs/release-notes/1.3.0.md)
- [1.3.0 control](https://github.com/pgsty/vchord_bm25/blob/1.3.0/vchord_bm25.control)

- [1.2.0 README](https://github.com/pgsty/vchord_bm25/blob/1.2.0/README.md)
- [1.2.0 release](https://github.com/pgsty/vchord_bm25/blob/1.2.0/docs/release-notes/1.2.0.md)
- [1.2.0 control](https://github.com/pgsty/vchord_bm25/blob/1.2.0/vchord_bm25.control)
- [Installation](https://vchord.pgsty.com/docs/bm25/getting-started/installation/)
- [Native tsvector quick start](https://vchord.pgsty.com/docs/bm25/getting-started/quick-start/)
- [1.0.0 migration and permissions](https://vchord.pgsty.com/docs/bm25/reference/migration/)
- [SQL reference](https://vchord.pgsty.com/docs/bm25/reference/sql/)
- [GUCs and index options](https://vchord.pgsty.com/docs/bm25/reference/gucs/)
- [Index maintenance](https://vchord.pgsty.com/docs/bm25/usage/indexing/)
- [1.0.0 control file](https://github.com/pgsty/vchord_bm25/blob/1.0.0/vchord_bm25.control)

`vchord_bm25` 1.3.0 adds BM25 relevance ranking over native PostgreSQL `tsvector` values on PostgreSQL 15–19. Its `bm25` access method and `bm25_ops` operator class use format-4 indexes. Standard text search needs no tokenizer extension; `pg_tokenizer` is optional for custom analyzers and multilingual models. Installation requires superuser privileges and shared preloading.

### Create and Query an Index

Add the library to the existing preload list and restart PostgreSQL:

```ini
shared_preload_libraries = 'vchord_bm25'
```

```sql
CREATE EXTENSION vchord_bm25;
CREATE TABLE documents (
    id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    passage text NOT NULL,
    terms tsvector GENERATED ALWAYS AS (to_tsvector('english', passage)) STORED
);
INSERT INTO documents (passage) VALUES
    ('PostgreSQL provides full text search.'),
    ('BM25 ranks PostgreSQL documents by relevance.'),
    ('A guide to database indexing.');
CREATE INDEX documents_bm25 ON documents USING bm25 (terms bm25_ops)
WITH (options = $$
k1 = 1.2
b = 0.75
$$, limit = 100);

SELECT id, passage,
       terms <&> to_bm25query(to_tsvector('english', 'PostgreSQL'), 'documents_bm25') AS rank
FROM documents
WHERE terms @@ plainto_tsquery('english', 'PostgreSQL')
ORDER BY rank LIMIT 10;
```

`to_bm25query(tsvector, regclass)` takes the query vector first and the target index second. `<&>` returns a negative BM25 score as `double precision`: ascending order puts greater relevance first. Negate the result only for display if a positive score is preferred. Ordering alone also includes zero-score nonmatches and NULLs; the example uses an explicit match predicate.

Keep positions in stored vectors: `strip(tsvector)` and `array_to_tsvector` cannot provide the term frequencies required by the index. Document weights are ignored, and native position limits can reduce frequency detail for long documents. Use the same analyzer for documents and queries. With custom tokenization, `analyze_to_tsvector` stores string lexemes and `tokenize_to_tsvector` stores model IDs as lexemes; both are `STABLE`, so maintain an ordinary column through the application or a trigger.

### Search Controls and Maintenance

| Setting or function | Meaning in 1.0.0 |
| --- | --- |
| `bm25.enable_scan` | Allow BM25 index plans; default true. |
| `bm25.limit`, index `limit` | Initial scan batch size, default 100; zero also selects 100. Scans continue for SQL limits, offsets, and filtering. |
| `bm25.prefilter`, index `prefilter` | Apply supported predicates before ranking candidates; default false. |
| `k1`, `b` | TOML scoring parameters; defaults 1.2 and 0.75, ranges 1.2–2.0 and 0–1. |
| `maintenance_min_rows`, `maintenance_scale_factor` | Automatic merge thresholds; defaults 1000 and 0.1. |
| `bm25_maintain(regclass)` | Request a merge without the automatic thresholds; returns whether a segment was published. |

Explicit GUC values, including false or the nominal default, override index options. Server, database, role, and startup settings also take precedence. Index defaults apply only when the GUC source is the built-in default.

```sql
SET "bm25.limit" = 200;
SET bm25.prefilter = true;
VACUUM documents;
SELECT bm25_maintain('documents_bm25');
```

`VACUUM` marks dead records and can merge segments; `ANALYZE` alone does not merge them. Manual maintenance requires table ownership or, on PostgreSQL 17+, `MAINTAIN` privilege, and does not replace dead-tuple detection. Deleted records stop appearing in results immediately but remain in corpus statistics until compaction. Long snapshots can delay page reuse and further merges. Standby readers encountering recycled pages can receive SQLSTATE `40001` and need to retry.

### Permissions and Migration

The extension is relocatable; applications need access to the selected installation schema and indexed data. Scoring requires `SELECT` on the indexed table, or column-level access for an ordinary non-partial single-column index. Expression and partial indexes require table-level access. With row-level security, scalar scoring and native scans reject queries by default because scores use whole-index corpus statistics; 1.3.0 adds the explicit opt-in below.

The index supports one indexed column and does not support unlogged tables. The access-method name is database-wide, so another extension defining `bm25` cannot coexist in that database.

Upgrades from 0.3.x are deliberately blocked. Keep the old database and library available, export source text, and create a fresh 1.0.0 installation with its own library directory. Recompute native vectors, rebuild indexes, and adapt query argument order, types, settings, and triggers before switching the application. Another database or port sharing the old library directory is insufficient isolation. Neither a binary index copy nor `REINDEX` on an old `bm25vector` column migrates the API.

### Version 1.2.0

Version 1.2.0 supports PostgreSQL 15–19. Install matching files, restart PostgreSQL and run `ALTER EXTENSION vchord_bm25 UPDATE;`. Upgrades from 1.0.0 and 1.1.0 retain their SQL APIs and format-4 indexes without routine REINDEX. PostgreSQL 14 requires a server upgrade first; the legacy 0.3.x migration boundary above still applies.

This release fixes error cleanup during prefiltering, planner costs and parallel-build snapshot protection. It also rejects invalid or partitioned indexes in `bm25_maintain`. These fixes do not replace normal vacuuming, access checks or application-level retry handling.

### Version 1.3.0: Shared Corpus Statistics

An index owner can opt into whole-index statistics under row security with the `shared_corpus_stats` index option, which defaults off. This exposes shared corpus information through scores; PostgreSQL still applies row policies before returning rows. Scoring checks corpus privileges on scalar, index and view paths. Do not enable this option when those statistics must remain isolated.

Install matching library and SQL files, restart and run `ALTER EXTENSION vchord_bm25 UPDATE;`. Versions 1.0.0–1.2.0 retain their format-4 indexes. Build sort memory is shared across participating workers, and ranked pagination/statistics and prefiltering are improved. Before reverting the library, reset `shared_corpus_stats` on every index where it was explicitly set, including explicit off values. SQL upgrades are forward-only.
