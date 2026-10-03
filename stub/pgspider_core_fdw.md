## Usage

Sources:

- [Official README](https://github.com/pgspider/pgspider/blob/d2ad2403600b3ff35fd35acf32864fabf958cb9b/README.md)
- [Extension control file](https://github.com/pgspider/pgspider/blob/d2ad2403600b3ff35fd35acf32864fabf958cb9b/contrib/pgspider_core_fdw/pgspider_core_fdw.control)
- [Installation SQL](https://github.com/pgspider/pgspider/blob/d2ad2403600b3ff35fd35acf32864fabf958cb9b/contrib/pgspider_core_fdw/pgspider_core_fdw--1.0.sql)
- [Build configuration](https://github.com/pgspider/pgspider/blob/d2ad2403600b3ff35fd35acf32864fabf958cb9b/contrib/pgspider_core_fdw/Makefile)

`pgspider_core_fdw` coordinates queries over child foreign tables inside the PGSpider database fork. It is distinct from the portable `pgspider_ext` project and requires PGSpider kernel changes and its `pgspider_keepalive` support library.

### Core Workflow

On a matching PGSpider server, install the wrapper and a child FDW. This example assumes the remote database already contains a table named t1 with matching columns; replace the example connection values.

```sql
CREATE EXTENSION pgspider_core_fdw;
CREATE EXTENSION postgres_fdw;
CREATE SERVER parent FOREIGN DATA WRAPPER pgspider_core_fdw;
CREATE SERVER postgres_svr FOREIGN DATA WRAPPER postgres_fdw
  OPTIONS (host '127.0.0.1', port '5432', dbname 'postgres');
CREATE USER MAPPING FOR CURRENT_USER SERVER postgres_svr
  OPTIONS (user 'app_user', password 'replace-me');
CREATE FOREIGN TABLE t1(i integer, t text, __spd_url text) SERVER parent;
CREATE FOREIGN TABLE t1__postgres_svr__0(i integer, t text)
  SERVER postgres_svr OPTIONS (table_name 't1');
SELECT * FROM t1;
```

### Mapping and Objects

Parent tables include `__spd_url` to identify the source. Child tables follow the documented parent-table/server-name/suffix naming convention; a second child FDW can expose another source with the same logical shape. The module installs its handler, validator, foreign data wrapper, and `pgspider_core_fdw_version()`, which returns an integer version identifier.

### Operational Boundaries

Provision FDW/server privileges and user mappings for each source. INSERT routes to a suitable child; UPDATE/DELETE behavior depends on all participating FDWs. Upstream excludes `RETURNING`, `WITH CHECK OPTION`, `ON CONFLICT`, and modifications through foreign partitions. COPY is also limited. Plan remote failure and transaction handling; this is not a cross-system atomic-commit guarantee. No preload requirement is declared for this module. SQL/control version `1.0` identifies the module, not a stock PostgreSQL version.
