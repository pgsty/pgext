## Usage

Sources:

- [README.md](https://github.com/pgexperts/pgx-bm25/blob/e0e1820586a53b1683c616f16ee0e64efcebdfff/README.md)
- [bm25_native--1.0.sql](https://github.com/pgexperts/pgx-bm25/blob/e0e1820586a53b1683c616f16ee0e64efcebdfff/bm25_native--1.0.sql)
- [Makefile](https://github.com/pgexperts/pgx-bm25/blob/e0e1820586a53b1683c616f16ee0e64efcebdfff/Makefile)
- [bm25_native.control](https://github.com/pgexperts/pgx-bm25/blob/e0e1820586a53b1683c616f16ee0e64efcebdfff/bm25_native.control)

`bm25_native` supplies a native full-text index with BM25 ranking, stored in PostgreSQL index pages and maintained through normal WAL and vacuum machinery. Release v1.0.0 installs SQL extension version 1.0; supported majors are PostgreSQL 17 and 18.

### Core Workflow

```sql
CREATE EXTENSION bm25_native;
CREATE TABLE bm25_demo (id integer PRIMARY KEY, body text);
INSERT INTO bm25_demo VALUES (1, 'quick fox'), (2, 'quiet database');
CREATE INDEX ON bm25_demo USING bm25_native (body);
SELECT id, bm25_score(ctid) AS score
FROM bm25_demo
WHERE body @@@ 'quick fox'
ORDER BY body &@@ 'quick fox'
LIMIT 10;
```

### Operational Boundaries

Pair `@@@` matching with `&@@` ordering for indexed top-k search. Ordering alone does not produce a useful ranked index scan. `bm25_score` reads the current ranked scan; use its query-qualified form, and table identity when required, if several ranked scans or partition children can be involved. The JSONB query builders include `bm25_term`, `bm25_boolean`, `bm25_phrase`, `bm25_wildcard` and `bm25_boost`. Multi-column index matching can differ from applying the operator as a scalar filter, including under row security; unsupported scoped or phrase filters raise an error. Installation is not trusted, relocation after installation is disabled, and PostgreSQL 19 CI is experimental. Follow the versioned format-compatibility and maintenance instructions before upgrading indexes.
