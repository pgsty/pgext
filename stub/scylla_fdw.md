## Usage

Sources:

- [Upstream README at the reviewed commit](https://github.com/GeoffMontee/scylla_fdw/blob/c4a977345d0367232ec9a5356e26097682a6cf2d/README.md)
- [Version 1.0 SQL objects](https://github.com/GeoffMontee/scylla_fdw/blob/c4a977345d0367232ec9a5356e26097682a6cf2d/scylla_fdw--1.0.sql)
- [Extension control file](https://github.com/GeoffMontee/scylla_fdw/blob/c4a977345d0367232ec9a5356e26097682a6cf2d/scylla_fdw.control)
- [ScyllaDB driver dependency](https://github.com/scylladb/cpp-rs-driver)

`scylla_fdw` exposes ScyllaDB tables through PostgreSQL's foreign-data-wrapper API. It depends on the separately built ScyllaDB C/C++ driver and supports reads, writes, simple predicate pushdown, TLS, connection pooling, and `IMPORT FOREIGN SCHEMA`.

```sql
CREATE EXTENSION scylla_fdw;
CREATE SERVER scylla_server
  FOREIGN DATA WRAPPER scylla_fdw
  OPTIONS (host '127.0.0.1', port '9042', consistency 'local_quorum');
CREATE USER MAPPING FOR CURRENT_USER
  SERVER scylla_server
  OPTIONS (username 'cassandra', password 'replace-me');
```

Foreign tables require `keyspace` and `table` options. Add `primary_key` for `UPDATE` and `DELETE`; ScyllaDB usually also requires equality on the partition key for efficient access. Keep credentials out of broadly visible catalogs by restricting access and using a dedicated mapping role.

The project is an alpha: its only release is tagged 0.1.0-alpha even though the control version is 1.0. It does not push down joins or aggregates, only simple conditions are eligible for pushdown, and collection types remain limited. Validate query plans and failure behavior before any non-test use.
