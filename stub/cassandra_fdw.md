## Usage

Sources:

- [Upstream reference manual](https://github.com/pgsql-io/cassandra_fdw/blob/2ee3d7950954f99464c6f3249d622136e558cfc8/doc.pdf)
- [Extension control file](https://github.com/pgsql-io/cassandra_fdw/blob/2ee3d7950954f99464c6f3249d622136e558cfc8/cassandra_fdw.control)
- [Official upstream repository](https://github.com/pgsql-io/cassandra_fdw)

`cassandra_fdw` exposes Cassandra 3+ data as PostgreSQL foreign tables through the Cassandra C/C++ driver. The authoritative server options are `host` (required, optionally comma-separated), `port` (default 9042), and `protocol` (documented default 4).

```sql
CREATE EXTENSION cassandra_fdw;
CREATE SERVER cassandra_server
  FOREIGN DATA WRAPPER cassandra_fdw
  OPTIONS (host '10.0.0.10,10.0.0.11', port '9042', protocol '4');

CREATE USER MAPPING FOR app_user SERVER cassandra_server;
```

Define foreign tables with exact Cassandra keyspace, table, column, and type mappings from the installed version. Pushdown, writes, authentication, TLS, consistency level, paging, and timeout support are not described by the reviewed 2016 reference, so verify each capability from the exact SQL and C driver build rather than assuming normal PostgreSQL semantics.

Remote reads do not participate in PostgreSQL MVCC snapshots or atomic cross-system transactions. Plan for Cassandra consistency, schema changes, partial failures, retries, duplicate effects, driver ABI, node discovery, and network latency. Store credentials in restricted user mappings when the build supports them, restrict server usage, and benchmark predicate pushdown with `EXPLAIN` before production use.
