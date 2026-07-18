## Usage

Sources:

- [Archived upstream README at the reviewed commit](https://github.com/ildus/clickhouse_fdw/blob/008644003bf7f5dbde838123f717b395951a4b34/README.md)
- [Extension control file](https://github.com/ildus/clickhouse_fdw/blob/008644003bf7f5dbde838123f717b395951a4b34/src/clickhouse_fdw.control)
- [FDW implementation](https://github.com/ildus/clickhouse_fdw/blob/008644003bf7f5dbde838123f717b395951a4b34/src/clickhouse_fdw.c)

`clickhouse_fdw` is an archived foreign-data wrapper for ClickHouse. The reviewed release supports HTTP and binary drivers, remote-schema import, reads, inserts, and selected aggregate pushdown; upstream directs new users to its maintained successor.

```sql
CREATE EXTENSION clickhouse_fdw;
CREATE SERVER clickhouse_svr
  FOREIGN DATA WRAPPER clickhouse_fdw
  OPTIONS (dbname 'analytics', host 'clickhouse.internal', port '8123');
CREATE USER MAPPING FOR CURRENT_USER
  SERVER clickhouse_svr OPTIONS (user 'reporter', password 'secret');
IMPORT FOREIGN SCHEMA "analytics"
  FROM SERVER clickhouse_svr INTO clickhouse_ext;
```

Do not keep a literal password in deployment SQL; use restricted provisioning and catalog access. The HTTP driver sends queries to a remote service, so require TLS where supported, least-privilege ClickHouse credentials, network allowlists, timeouts, and resource limits. Keep ClickHouse in UTC as upstream requires and verify type, null, timestamp, and aggregation mappings.

Because this repository is archived, use its successor for new deployments. For legacy systems, pin dependencies such as libcurl and UUID libraries, inspect `EXPLAIN VERBOSE` for remote SQL, and test transaction semantics, partial writes, retries, cancellation, failover, predicate pushdown, and major-version upgrades before relying on it.
