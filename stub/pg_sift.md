## Usage

Sources:

- [Current version source](https://github.com/brauliobo/pg_sift/blob/v5.0.1/README.md)
- [Current version source](https://github.com/brauliobo/pg_sift/blob/v5.0.1/pg_sift.control)
- [Current version source](https://github.com/brauliobo/pg_sift/blob/v5.0.1/CHANGELOG.md)
- [Current version source](https://github.com/brauliobo/pg_sift/blob/v5.0.1/docs/USAGE.md)
- [Current version source](https://github.com/brauliobo/pg_sift/blob/v5.0.1/docs/UPGRADING.md)
- [Current version source](https://github.com/brauliobo/pg_sift/blob/v5.0.1/docs/SUPPORT.md)
- [Current version source](https://github.com/brauliobo/pg_sift/blob/v5.0.1/sql/pg_sift--5.0--5.0.1.sql)

`pg_sift` 5.0.1 combines BM25 text search, compressed column indexes and a columnar table access method. Upstream supports PostgreSQL 17–19. Initial installation requires a superuser. Ordinary search and column analytics do not require preload; publishing columnar tables through logical replication has the separate preload requirements in the upstream guide.

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

Run the versioned preflight and install matching libraries on every reader, including standbys. Follow the required restart/reconnect procedure before updating each database:

```sql
ALTER EXTENSION pg_sift UPDATE TO '5.0.1';
SELECT * FROM sift.upgrade_commands();
SELECT * FROM sift.col_repair('public.orders_col', true);
```

The 5.0-to-5.0.1 SQL script is empty: it aligns the reported extension version but changes no objects. This patch fixes concurrent-build locking, memory retention, cancellation and crash paths without an on-disk format change or an additional REINDEX requirement.

Crossing to 5.0 still requires the full upgrade guide: move readers to at least 4.3.6 first and upgrade standbys before a primary writes the new format. A 4.x reader cannot interpret 5.0 index data. `sift_table` logical change records use WAL resource-manager ID 152; 5.0 also reads legacy ID 128, while 4.x cannot replay 152. Rollback needs the documented clean-shutdown and format-rebuild procedure. A format mismatch is not proof of corruption.

Older repair obligations remain: pre-3.3 text indexes other than `analyzer = keyword` need rebuilding for long-word correctness; affected 4.2.1–4.3.5 column segments may need `sift.col_repair` or REINDEX after liveness-page damage. Start with the dry run shown above. Concurrent rebuild commands run individually outside a transaction. Remove obsolete `sift.enable_col` and `sift.table_auto_defrag_dead_pct`; `sift.eval_methods` retains only `penalty_maintenance`. Review the current guide for option renames and reader format limits instead of downgrading libraries in place.
