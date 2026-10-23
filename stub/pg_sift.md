## Usage

Sources:

- [Official source](https://github.com/brauliobo/pg_sift/blob/3e7fbf4aab7e4b1042ac83b106677dbc7377d501/README.md)
- [Official source](https://github.com/brauliobo/pg_sift/blob/3e7fbf4aab7e4b1042ac83b106677dbc7377d501/CHANGELOG.md)
- [Official source](https://github.com/brauliobo/pg_sift/blob/3e7fbf4aab7e4b1042ac83b106677dbc7377d501/pg_sift.control)
- [Official source](https://github.com/brauliobo/pg_sift/blob/3e7fbf4aab7e4b1042ac83b106677dbc7377d501/sql/pg_sift--3.2.sql)

`pg_sift` release 3.2.2 installs SQL extension version 3.2. It combines BM25 text search with compressed column indexes and a columnar table access method. PostgreSQL 17–18 are supported; PG19 remains experimental. Initial installation requires a superuser, but no preload or server restart.

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

The `sift` index handles a text column or expression, with boolean, phrase, prefix, wildcard, regex and fuzzy queries. Use `sift.score(ctid)` in the scanning query's output or ordering; 3.2 also supports comparisons with a constant score threshold in its filter. `sift.tokenize`, `sift.highlight` and `sift.facets` provide analysis, highlighting and facets. `sift.sections` and `sift.field_query` support the documented field-weight workflow; they do not create a multicolumn BM25 index.

### Column Analytics

```sql
CREATE TABLE orders (id bigint PRIMARY KEY, category text, amount numeric(10,2));
CREATE INDEX orders_col ON orders USING sift_col (category, amount);
SELECT category, count(*), sum(amount) FROM orders GROUP BY category;
CREATE TABLE events (id bigint, kind text, amount numeric(10,2)) USING sift_table;
```

`sift_col` supplements an ordinary table; `sift_table` stores the table itself in column format. Eligible aggregates and filters use `Custom Scan (SiftAgg)`; other cases retain ordinary plans. Inspect the actual plan. `sift.enable_agg = off` disables aggregate pushdown for comparison. Version 3.1 extends calendar grouping, supported JSON shapes and element rollups. The optional Vulkan path is experimental and disabled by default, including in upstream packages.

### Security and Maintenance

Roles constrained by row security or security-barrier views cannot use index-wide BM25 scores for those relations. The older facets API rejects RLS tables; the newer SQL-based facet API has its own rules. `sift.verify`, `sift.segments`, `sift.merge` and related maintenance calls require table ownership or MAINTAIN privilege. Changing analyzers or scoring options requires rebuilding affected search indexes.

`write_buffer_limit` replaces the older `mutable_limit` name, retained as an alias. Vacuum manages dead rows and segments; index truncation replay can cancel standby queries. The prototype test companion is not a production extension. Check encoding, block-size, supported column types and access-method restrictions in the versioned documentation.

### Upgrade to 3.2

Install matching files, then update every database:

```sql
ALTER EXTENSION pg_sift UPDATE TO '3.2';
```

Released 2.x/3.0 structures remain readable without a routine rebuild. New writes introduce newer BM25, column-index and table formats that older libraries refuse. Upgrade every server that opens the data, including standbys, before the first new-format write. Downgrading requires a deliberate rewrite/rebuild procedure; there is no reverse SQL update script. Reconnect sessions that loaded the older library. When crossing older 1.x versions, also check the removed phrase-index option and function/view dependency changes in the changelog.

### 3.2.2 Fix and Known Limitation

3.2.2 fixes leakage of hidden-row statistics through planner estimates on RLS-protected partition or inheritance children, affecting upstream 3.0.0–3.2.1. Replacing the library fixes this issue without rebuilding indexes. 3.2 expands aggregate pushdown and supported column types; new storage formats may be unreadable by older libraries, so check every node that reads the data before upgrading.

The event trigger in 3.2.0–3.2.2 also has a known privilege defect: roles other than superusers or the extension owner can be unable to execute DROP. Upstream provides this temporary correction, run by a superuser in each database; recheck after restoring a backup. Do not assume the later 3.3 correction is already installed.

```sql
ALTER FUNCTION sift.table_paths_drop() SECURITY DEFINER SET search_path = pg_catalog, pg_temp;
```
