## Usage

Sources:

- [Pinned upstream README](https://github.com/d-e-n-t-y/pg_fdw_mv_rewrite/blob/0b5ed382c3c956309a7378f74c8619478bb3b3b0/README.md)
- [Version 0.6.2 installation SQL](https://github.com/d-e-n-t-y/pg_fdw_mv_rewrite/blob/0b5ed382c3c956309a7378f74c8619478bb3b3b0/mv_rewrite--0.6.2.sql)
- [Pinned planner-hook implementation](https://github.com/d-e-n-t-y/pg_fdw_mv_rewrite/blob/0b5ed382c3c956309a7378f74c8619478bb3b3b0/mv_rewrite.c)
- [Pinned extension entry point](https://github.com/d-e-n-t-y/pg_fdw_mv_rewrite/blob/0b5ed382c3c956309a7378f74c8619478bb3b3b0/extension.c)

mv_rewrite version 0.6.2 is an experimental planner extension that considers registered materialized views as alternative paths. When it finds a semantically usable view with a lower estimated cost, EXPLAIN shows a Custom Scan (MVRewriteScan) and the query reads the materialized view transparently.

### Register a materialized view

```sql
CREATE EXTENSION mv_rewrite;

CREATE TABLE rewrite_sales (region text, amount numeric);
INSERT INTO rewrite_sales VALUES
    ('east', 10),
    ('east', 20),
    ('west', 30);

CREATE MATERIALIZED VIEW rewrite_sales_mv AS
SELECT region, count(amount) AS n
FROM rewrite_sales
GROUP BY region;

SELECT mv_rewrite.enable_rewrite('rewrite_sales_mv');

EXPLAIN (VERBOSE)
SELECT region, count(amount)
FROM rewrite_sales
GROUP BY region;
```

The install script changes the current database's session_preload_libraries setting to a versioned module name and loads it in the installing session. New sessions inherit the database setting. This persistent database-level side effect should be reviewed before installation.

### Correctness and support boundary

Upstream explicitly says the code is not production ready. It considers selected GROUP BY aggregates, DISTINCT, ordering, simple SELECTs, joins, WHERE, and HAVING shapes; recursive queries and lateral joins are not properly supported. Planner matching is costly, so the extension provides a minimum-cost threshold, an enabled-table list, progress logging, and a global enable switch.

Registration does not refresh a materialized view. A stale view can therefore yield stale rewritten results even when matching is otherwise correct. Establish refresh and dependency procedures and compare plans and results with rewriting disabled. The source targets old PostgreSQL planner internals and has not changed since 2019. Test every query family, extension interaction, and PostgreSQL major version in isolation, and retain a reliable way to remove the database preload setting.
