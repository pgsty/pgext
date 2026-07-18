## Usage

Sources:

- [Upstream README and option reference](https://github.com/slamdata/quasar-fdw/blob/23ece25661b2071644179e079134f4041e099026/README.md)
- [Extension control file](https://github.com/slamdata/quasar-fdw/blob/23ece25661b2071644179e079134f4041e099026/quasar_fdw.control)
- [Option validator](https://github.com/slamdata/quasar-fdw/blob/23ece25661b2071644179e079134f4041e099026/src/quasar_options.c)
- [Archived upstream repository](https://github.com/slamdata/quasar-fdw)

`quasar_fdw` is a read-only foreign data wrapper for the historical SlamData Quasar query engine. It forwards `SELECT` queries and can push supported `WHERE`, `ORDER BY`, and join clauses into Quasar.

### Define a legacy Quasar table

```sql
CREATE EXTENSION quasar_fdw;

CREATE SERVER quasar
  FOREIGN DATA WRAPPER quasar_fdw
  OPTIONS (
    server 'http://localhost:8080',
    path '/local/quasar',
    timeout_ms '1000',
    use_remote_estimate 'true'
  );

CREATE FOREIGN TABLE public.zips (
  city text,
  population integer OPTIONS (map 'pop'),
  state char(2)
)
SERVER quasar
OPTIONS (table 'zips');

SELECT city, population
FROM public.zips
WHERE population > 100000
ORDER BY population DESC;
```

Server options also include `fdw_startup_cost` and `fdw_tuple_cost`. Table options are `table` and `use_remote_estimate`; column options are `map`, `nopushdown`, and `join_rowcount_estimate`. Use `nopushdown` when remote strings are locally cast to types for which a pushed predicate would change semantics.

The upstream repository is archived and its README supports only PostgreSQL 9.4/9.5 with Quasar 9 through 13. The control/catalog version is `1.4.1`, while the README still calls `1.4.0` latest. No write operations or current compatibility path are documented. Use this extension only to preserve a tested legacy deployment; for new systems, choose a maintained data source and FDW.
