## Usage

Sources:

- [README.md](https://github.com/cmsedore/pgtrino/blob/acf9f98087b14ed7c93a380b761190d2705554d1/README.md)
- [trino_fdw--1.0.sql](https://github.com/cmsedore/pgtrino/blob/acf9f98087b14ed7c93a380b761190d2705554d1/trino_fdw--1.0.sql)
- [Makefile](https://github.com/cmsedore/pgtrino/blob/acf9f98087b14ed7c93a380b761190d2705554d1/Makefile)
- [trino_fdw.control](https://github.com/cmsedore/pgtrino/blob/acf9f98087b14ed7c93a380b761190d2705554d1/trino_fdw.control)

`trino_fdw` exposes Trino tables as PostgreSQL foreign tables using libcurl. Upstream documents PostgreSQL 14–17. Install the extension as a superuser, then configure the server and caller mapping. It requires no shared preload.

### Core Workflow

```sql
CREATE EXTENSION trino_fdw;
CREATE SERVER trino FOREIGN DATA WRAPPER trino_fdw
  OPTIONS (host 'trino.example.com', port '8443', ssl 'true', catalog 'hive');
CREATE USER MAPPING FOR CURRENT_USER SERVER trino OPTIONS (user 'analytics');
CREATE SCHEMA sales;
IMPORT FOREIGN SCHEMA sales FROM SERVER trino INTO sales;
EXPLAIN SELECT * FROM sales.orders WHERE order_id = 42;
```

### Operational Boundaries

Server options include host, port, ssl, ssl_verify, catalog, schema, spooling and batch_size. Use HTTPS and certificate verification for remote credentials. Mapping options user and password select the Trino identity; PostgreSQL foreign-table privileges apply in addition to Trino authorization. Projection, eligible filters, joins, aggregates and ordering can be pushed down; EXPLAIN shows the remote SQL and ANALYZE collects local planner statistics. `trino_query` returns records and `trino_execute` returns an update count; both require USAGE on the foreign server. UPDATE/DELETE require a fully shippable statement and connector support; RETURNING, joined modifications and row triggers are unsupported. Each remote statement or INSERT batch autocommits: a PostgreSQL rollback cannot undo it. Spooling also requires the PostgreSQL server to reach Trino object storage; disable spooling if that path is unavailable.
