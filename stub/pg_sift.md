## Usage

Sources:

- [Official source](https://github.com/brauliobo/pg_sift/blob/232ee801ac3cee2af973003df11afb71f6138530/README.md)
- [Official source](https://github.com/brauliobo/pg_sift/blob/232ee801ac3cee2af973003df11afb71f6138530/docs/UPGRADING.md)
- [Official source](https://github.com/brauliobo/pg_sift/blob/232ee801ac3cee2af973003df11afb71f6138530/pg_sift.control)
- [Official source](https://github.com/brauliobo/pg_sift/blob/232ee801ac3cee2af973003df11afb71f6138530/sql/pg_sift--3.3.sql)
- [Usage and security](https://github.com/brauliobo/pg_sift/blob/232ee801ac3cee2af973003df11afb71f6138530/docs/USAGE.md)
- [Supported versions](https://github.com/brauliobo/pg_sift/blob/232ee801ac3cee2af973003df11afb71f6138530/docs/SUPPORT.md)

`pg_sift` release 3.3.0 installs SQL extension version 3.3. It combines BM25 text search with compressed column indexes and a columnar table access method. PostgreSQL 17–19 are supported upstream; this does not imply a packaged build for every major. Initial installation requires a superuser, but no preload or server restart.

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

The `sift` index handles a text column or expression, with boolean, phrase, prefix, wildcard, regex and fuzzy queries. Use `sift.score(ctid)` in the scanning query's output or ordering; it also supports comparisons with a constant score threshold in its filter. `sift.tokenize`, `sift.highlight` and `sift.facets` provide analysis, highlighting and facets. `sift.sections` and `sift.field_query` support the documented field-weight workflow; they do not create a multicolumn BM25 index.

### Column Analytics

```sql
CREATE TABLE orders (id bigint PRIMARY KEY, category text, amount numeric(10,2));
CREATE INDEX orders_col ON orders USING sift_col (category, amount);
SELECT category, count(*), sum(amount) FROM orders GROUP BY category;
CREATE TABLE events (id bigint, kind text, amount numeric(10,2)) USING sift_table;
```

`sift_col` supplements an ordinary table; `sift_table` stores the table itself in column format. Eligible aggregates and filters use `Custom Scan (SiftAgg)`; other cases retain ordinary plans. Inspect the actual plan. `sift.enable_agg = off` disables aggregate pushdown for comparison. Version 3.3 expands top-N ordering, point lookups and joins to small tables. New column indexes enable pre-aggregates by default. The optional Vulkan path is experimental and disabled by default, including in upstream packages.

### Security and Maintenance

Roles under row-level security do not get ranked scans by default. A superuser may opt in with `sift.rls_scores = on`; visible rows still obey the policy, but scores depend on statistics for the whole index, including hidden rows. Review that information boundary before enabling it. Security-barrier and other unsupported query shapes retain their documented limitations.

The 3.3 install and update grant schema usage on `sift` to `PUBLIC`; functions still enforce their own table/index permissions. `sift.verify`, `sift.segments`, `sift.merge` and related maintenance calls require table ownership or MAINTAIN privilege. Changing analyzers or scoring options requires rebuilding affected search indexes.

`sift.col_auto_merge` and `sift.table_auto_defrag` now default to on. Background maintenance can vacuum tables and rewrite column indexes, producing WAL. The column merge budget is controlled by `sift.col_auto_merge_wal_pct`, while `sift.col_maintenance_status` and `sift.col_index_status` expose maintenance state. Set the two automatic-maintenance switches to off to restore the older behavior.

`write_buffer_limit` replaces the older `mutable_limit` name, retained as an alias. Index truncation replay can cancel standby queries. Publishing a columnar table through logical replication requires the upstream preload configuration; basic text search and column analytics do not require preloading.

### Upgrade to 3.3

Use the upstream upgrade preflight, install matching files, update each database, and restart processes that loaded the old library so that all readers use the same release:

```sql
ALTER EXTENSION pg_sift UPDATE TO '3.3';
SELECT * FROM sift.upgrade_commands();
```

Rebuild every pre-3.3 text index except those using `analyzer = keyword`: the treatment of words longer than 255 bytes changed, so unreconstructed indexes can miss matching rows. The update prints commands; run concurrent index rebuilds individually outside a transaction. A suggested column-index rebuild may improve performance, while a suggested columnar-table rewrite can take an exclusive lock. Review each command before running it.

The default `sift.col_format_minor = 15` is unreadable by 3.2 libraries. Upgrade all readers, including standbys, before new-format writes. If compatibility with 3.2 is required, follow the versioned downgrade guide and set the format limit server-wide before upgrading; a session-only limit does not reach background workers. Opt-in point TID directories require format 16. There is no automatic downgrade of data files.

The 3.3 SQL migration fixes the earlier DROP event-trigger privilege defect; the prior manual workaround is no longer a separate installation step. It also handles a fresh backend's first query with bound parameters, extending the plain-query fix in 3.2.3.
