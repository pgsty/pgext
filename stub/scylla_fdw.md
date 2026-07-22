## Usage

Sources:

- [Upstream README at the reviewed commit](https://github.com/GeoffMontee/scylla_fdw/blob/c4a977345d0367232ec9a5356e26097682a6cf2d/README.md)
- [Version 1.0 SQL objects](https://github.com/GeoffMontee/scylla_fdw/blob/c4a977345d0367232ec9a5356e26097682a6cf2d/scylla_fdw--1.0.sql)
- [Extension control file](https://github.com/GeoffMontee/scylla_fdw/blob/c4a977345d0367232ec9a5356e26097682a6cf2d/scylla_fdw.control)
- [FDW implementation](https://github.com/GeoffMontee/scylla_fdw/blob/c4a977345d0367232ec9a5356e26097682a6cf2d/scylla_fdw.c)
- [Connection implementation](https://github.com/GeoffMontee/scylla_fdw/blob/c4a977345d0367232ec9a5356e26097682a6cf2d/scylla_connection.cpp)
- [ScyllaDB driver dependency](https://github.com/scylladb/cpp-rs-driver)

`scylla_fdw` maps ScyllaDB tables into PostgreSQL through the foreign-data-wrapper API. The reviewed code implements scans, inserts, updates, deletes, and `IMPORT FOREIGN SCHEMA`, with limited predicate pushdown. It requires PostgreSQL 14 or later, a C++17 toolchain, and the separately built ScyllaDB C/C++ driver.

### Core Workflow

Create a server and a role-scoped user mapping, then define the remote table shape explicitly:

```sql
CREATE EXTENSION scylla_fdw;

CREATE SERVER scylla_server
  FOREIGN DATA WRAPPER scylla_fdw
  OPTIONS (host '127.0.0.1', port '9042', consistency 'local_quorum');

CREATE USER MAPPING FOR CURRENT_USER
  SERVER scylla_server
  OPTIONS (username 'app_reader', password 'replace-me');

CREATE FOREIGN TABLE users (
  user_id uuid,
  username text,
  created_at timestamptz
)
SERVER scylla_server
OPTIONS (keyspace 'app', table 'users', primary_key 'user_id');

EXPLAIN (VERBOSE)
SELECT * FROM users
WHERE user_id = '00000000-0000-0000-0000-000000000001';
```

`keyspace` and `table` identify the remote relation. Set `primary_key` when updates or deletes must build a remote key predicate. Confirm the actual ScyllaDB partition and clustering-key requirements rather than assuming that this option makes arbitrary predicates efficient.

### Object and Option Index

- `scylla_fdw_handler()` and `scylla_fdw_validator(text[], oid)` back the `scylla_fdw` foreign-data wrapper.
- Server options include `host`, `port`, `consistency`, connection-pool sizing, and TLS-related values.
- User-mapping options are `username` and `password`.
- Foreign-table options include `keyspace`, `table`, and `primary_key`.
- `IMPORT FOREIGN SCHEMA` derives PostgreSQL foreign-table definitions from ScyllaDB metadata.

### Query and Transaction Boundaries

- The deparser only pushes a subset of simple conditions. Joins and aggregates are not pushed down; expressions involving `OR`, `NOT`, or `NULL` can remain local. Use `EXPLAIN (VERBOSE)` to verify each important query.
- Collection-type handling is limited. Test every mapped type and its null behavior against representative data.
- Remote writes are executed through the driver and are not enlisted in PostgreSQL's transaction or rollback machinery. A later local error can therefore leave ScyllaDB effects committed.
- The source accepts a `request_timeout` option, but the reviewed connection implementation does not apply it to the driver. Do not rely on this setting as an enforced query deadline.
- The README describes `ssl_cert`, `ssl_key`, and `ssl_ca` as file paths, while the reviewed connection code passes their option strings directly to the driver's certificate/key setters rather than reading files. Validate TLS with the exact build and option representation before use.
- The reviewed `IMPORT FOREIGN SCHEMA` path establishes a separate default connection and does not carry the configured TLS, timeout, or consistency values into that metadata request. Test it separately, especially against TLS-only clusters.

### Operational Caveats

- Credentials live in PostgreSQL user-mapping catalogs. Limit catalog visibility and grant server usage only to a dedicated role.
- The repository's only release tag is `0.1.0-alpha`, while the extension control file reports version `1.0`. Treat both the API and on-disk metadata as alpha-quality until tested.
- Driver calls run inside a PostgreSQL backend. Exercise DNS failures, authentication errors, timeouts, remote-node loss, cancellation, and backend restart behavior before any non-test deployment.
