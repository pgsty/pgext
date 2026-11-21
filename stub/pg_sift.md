## Usage

Sources:

- [Official source](https://github.com/brauliobo/pg_sift/releases/tag/v5.1.0)
- [Official source](https://github.com/brauliobo/pg_sift/blob/v5.0.2/sql/pg_sift--5.0.1--5.0.2.sql)

- [Current version source](https://github.com/brauliobo/pg_sift/blob/v5.1.0/README.md)
- [Current version source](https://github.com/brauliobo/pg_sift/blob/v5.1.0/pg_sift.control)
- [Current version source](https://github.com/brauliobo/pg_sift/blob/v5.1.0/CHANGELOG.md)
- [Current version source](https://github.com/brauliobo/pg_sift/blob/v5.1.0/docs/USAGE.md)
- [Current version source](https://github.com/brauliobo/pg_sift/blob/v5.1.0/docs/UPGRADING.md)
- [Current version source](https://github.com/brauliobo/pg_sift/blob/v5.1.0/docs/SUPPORT.md)
- [Current version source](https://github.com/brauliobo/pg_sift/blob/v5.0.2/sql/pg_sift--5.0--5.0.1.sql)

`pg_sift` 5.1.0 combines BM25 text search, compressed column indexes and a columnar table access method. Upstream supports PostgreSQL 17–19. Initial installation requires a superuser. Ordinary search and column analytics do not require preload; publishing columnar tables through logical replication has the separate preload requirements in the upstream guide.

### Text Search

```sql
CREATE EXTENSION pg_sift;
CREATE TABLE articles (id bigint PRIMARY KEY, body text);
INSERT INTO articles VALUES
  (1, 'PostgreSQL extensions add full text search'),
  (2, 'PostgreSQL indexing and query planning');
CREATE INDEX articles_search ON articles USING sift (body);
SELECT id, sift.score(ctid) AS score
FROM articles
WHERE body @@> 'postgresql AND search'
ORDER BY score DESC LIMIT 10;
```

The `sift` access method indexes one text column or expression. It supports boolean, phrase, prefix, wildcard, regex and fuzzy queries. Use `sift.score(ctid)` in the scanning query. `sift.tokenize`, `sift.highlight` and `sift.facets` provide analysis, highlighting and facets; `sift.sections` and `sift.field_query` implement the documented field-weight workflow.

### Column Analytics

```sql
CREATE TABLE orders (id bigint PRIMARY KEY, category text, amount numeric(10,2));
CREATE INDEX orders_col ON orders USING sift_col (category, amount);
SELECT category, count(*), sum(amount) FROM orders GROUP BY category;
CREATE TABLE events (id bigint, kind text, amount numeric(10,2)) USING sift_table;
```

`sift_col` supplements an ordinary table; `sift_table` stores the table itself in column format. Eligible plans use `Custom Scan (SiftAgg)`; other shapes keep ordinary PostgreSQL plans. Inspect the actual plan. `sift.enable_agg = off` disables aggregate pushdown, while `sift.enable_col_scan = off` disables column scans. The optional Vulkan path is experimental and disabled by default.

New or rebuilt eligible column indexes use `track_changes = 'auto'`, which installs statement triggers and adds write work. Incomplete tracking falls back to heap checks. Consult the upgrade guide before changing partition topology or restoring with triggers disabled.

### Security and Maintenance

Ranked scans are disabled by default for roles subject to row-level security. A superuser may enable `sift.rls_scores = on`; returned rows obey RLS, but their scores depend on whole-index statistics, including hidden rows. Review that disclosure boundary. Maintenance functions such as `sift.verify`, `sift.segments` and `sift.merge` require table ownership or MAINTAIN privilege. Analyzer or scoring-option changes can require a search-index rebuild.

Version 5.0.1 stops automatic full rewrites of `sift_col` indexes. `sift.col_maintenance_status` and `sift.col_maintenance_due` report when a rebuild is recommended; an operator chooses `REINDEX INDEX CONCURRENTLY` or schedules `sift.col_maintain(index)`. Incremental write-buffer promotion, small-segment merges, log folding, vacuum and value-store maintenance remain automatic; `sift.col_auto_merge` still controls merges. Remove retired `sift.col_auto_merge_max_active` and `sift.col_auto_merge_max_workers` settings. Some build/merge memory scopes remain unbounded in upstream qualification; do not assume a hard memory ceiling.

With the default `sift.on_corruption = quarantine`, the detecting statement fails and a damaged `sift` or `sift_col` index is taken out of service. Later queries use exact heap answers; an old plan can raise `40001` once. Inspect `sift.quarantine` and rebuild deliberately. This does not repair PostgreSQL system catalogs or cover `sift_table` corruption.

### Upgrade Boundaries

Run the versioned preflight and install matching libraries on every reader, including standbys. Upgrade standbys first, restart the servers with matching files, and reconnect pooled clients before updating each database:

```sql
ALTER EXTENSION pg_sift UPDATE TO '5.1';
SELECT * FROM sift.upgrade_commands();
SELECT * FROM sift.col_repair('public.orders_col', true);
```

The 5.0-to-5.0.1 SQL script is empty: it aligns the reported extension version but changes no objects. This patch fixes concurrent-build locking, memory retention, cancellation and crash paths without an on-disk format change or an additional REINDEX requirement.

Crossing to 5.0 still requires the full upgrade guide: move readers to at least 4.3.6 first and upgrade standbys before a primary writes the new format. A 4.x reader cannot interpret 5.0 index data. `sift_table` logical change records use WAL resource-manager ID 152; 5.0 also reads legacy ID 128, while 4.x cannot replay 152. Rollback needs the documented clean-shutdown and format-rebuild procedure. A format mismatch is not proof of corruption.

Older repair obligations remain: pre-3.3 text indexes other than `analyzer = keyword` need rebuilding for long-word correctness; affected 4.2.1–4.3.5 column segments may need `sift.col_repair` or REINDEX after liveness-page damage. Start with the dry run shown above. Concurrent rebuild commands run individually outside a transaction. Remove obsolete `sift.enable_col` and `sift.table_auto_defrag_dead_pct`; `sift.eval_methods` retains only `penalty_maintenance`. Review the current guide for option renames and reader format limits instead of downgrading libraries in place.

### Version 5.0.2

Version 5.0.2 terminates repeated column-merge passes that could grow an index until the disk filled, fixes planner and shutdown crashes, and avoids false duplicate reports while VACUUM and inserts run. Maintenance inspectors reject a partitioned parent index; inspect its storage-bearing children instead.

Remove `sift.col_build_key_ranges` from saved settings: the option no longer exists. A key-range build estimates temporary space and fails with SQLSTATE 54000 if it would exceed `temp_file_limit` or half the temporary tablespace's free space. Review `temp_tablespaces` and available capacity before retrying.

The library and empty 5.0.1-to-5.0.2 update script preserve the disk format; healthy 5.0 indexes need no routine REINDEX. An index inflated by a runaway merge retains that space until deliberately rebuilt. `REINDEX INDEX CONCURRENTLY` temporarily needs space for both copies. Install matching files, restart or recycle all old-library sessions, then update each SQL extension. Earlier 4.x migration and repair requirements still apply.

### Version 5.1.0

Release 5.1.0 uses SQL extension version 5.1. Existing 5.0 indexes remain readable without a mandatory REINDEX. New writes can introduce deletion records, key-range write buffers and chained directories that 5.0 refuses. Rebuilding is optional to adopt new layouts; it is not a compatibility prerequisite for upgrading.

`sift.col_build_source` defaults to auto and chooses among the heap, a suitable btree and the old index being replaced; heap forces the earlier strategy. Builds that cannot fit key-range temporary files can use a suitable btree or open units instead of always failing. `sift.col_wbuf_key_range` enables key ranges by default and is superuser-controlled. `sift.backend_cache_kb` now defaults to 128 MB per backend.

Before downgrading, inspect `sift.upgrade_check('5.0')`. Follow the guide's VACUUM, SQL-version rollback, shutdown and older-library installation sequence, then rebuild the listed indexes with the older release. REINDEX under 5.1 can write the incompatible formats again. The earlier repair obligations still apply to indexes carried forward from affected releases.
