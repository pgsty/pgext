## Usage

Sources:

- [Official extension SQL](https://github.com/snaga/monetdb_fdw/blob/fbff386be673bb6f1d43cddbceb0081a20e9b459/monetdb_fdw--0.0.sql)
- [Official implementation](https://github.com/snaga/monetdb_fdw/blob/fbff386be673bb6f1d43cddbceb0081a20e9b459/monetdb_fdw.c)
- [Official upstream README](https://github.com/snaga/monetdb_fdw/blob/fbff386be673bb6f1d43cddbceb0081a20e9b459/README.md)

`monetdb_fdw` is a version 0.0 prototype for reading MonetDB data through the MonetDB MAPI client library. It implements foreign scans only: there is no DML, join/aggregate pushdown, remote costing, or production-grade option and credential model.

### Core Workflow

Connection credentials and the remote relation or query are configured directly on each foreign table:

```sql
CREATE EXTENSION monetdb_fdw;
CREATE SERVER monet_remote FOREIGN DATA WRAPPER monetdb_fdw;

CREATE FOREIGN TABLE public.remote_lineitem (
  orderkey bigint,
  quantity numeric
)
SERVER monet_remote
OPTIONS (
  host 'monetdb.example.com',
  port '50000',
  user 'monet_user',
  passwd 'secret',
  dbname 'demo',
  table 'sys.lineitem'
);

SELECT * FROM public.remote_lineitem;
```

Use either `table` or `query`. With `table`, the wrapper constructs a simple `SELECT * FROM ...`; with `query`, it executes the supplied text verbatim.

### Options and Build Boundary

Foreign-table options are `host`, `port`, `user`, `passwd`, `dbname`, `table`, and `query`. The project must be compiled against MonetDB's MAPI headers and `libmapi`. There is no user-mapping credential layer, and result values arrive as strings which PostgreSQL converts to the locally declared column types.

### Safety and Limitations

The validator restricts changes to the foreign-table options to a superuser, but credentials remain in foreign-table catalog metadata visible to sufficiently privileged roles. Restrict catalog and DDL access and do not place production secrets in this prototype. The table-name path uses minimal string construction and the query path is arbitrary remote SQL; never pass untrusted identifiers or text. Test the exact PostgreSQL ABI and MonetDB MAPI version, type/NULL conversion, long identifiers and queries, encoding, cancellation, partial reads, and connection loss in an isolated environment. Use a mature alternative for production workloads.
