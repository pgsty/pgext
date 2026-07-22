## Usage

Sources:

- [Official README at the catalog source revision](https://github.com/pgspider/griddb_fdw/blob/a05fbf78a57d1941538aa11a0038037b08ff0e79/README.md)
- [Extension control file at the catalog source revision](https://github.com/pgspider/griddb_fdw/blob/a05fbf78a57d1941538aa11a0038037b08ff0e79/griddb_fdw.control)
- [Version 1.2 extension SQL](https://github.com/pgspider/griddb_fdw/blob/a05fbf78a57d1941538aa11a0038037b08ff0e79/griddb_fdw--1.2.sql)
- [GridDB C client repository](https://github.com/griddb/c_client)

griddb_fdw exposes GridDB containers as PostgreSQL foreign tables. It supports reads and writes plus selected expression, aggregate, ordering, and limit pushdown, but depends on the GridDB C client and does not provide PostgreSQL-local storage semantics.

### Core Workflow

Install the matching GridDB client library before building the extension, then create a server, user mapping, and foreign table. The row-key column must be identified explicitly when the remote container uses it.

```sql
CREATE EXTENSION griddb_fdw;

CREATE SERVER griddb_svr
FOREIGN DATA WRAPPER griddb_fdw
OPTIONS (
    host '239.0.0.1',
    port '31999',
    clustername 'myCluster',
    database 'public'
);

CREATE USER MAPPING FOR CURRENT_USER
SERVER griddb_svr
OPTIONS (username 'admin', password 'admin');

CREATE FOREIGN TABLE sensor_readings (
    ts timestamp OPTIONS (rowkey 'true'),
    sensor_id text,
    value double precision
)
SERVER griddb_svr
OPTIONS (table_name 'sensor_readings');

SELECT sensor_id, avg(value)
FROM sensor_readings
GROUP BY sensor_id;
```

For fixed-list notification mode, supply the member list instead of multicast host and port settings. Server-level controls include connection retention, updatability, and batch size; foreign-table controls include remote table name, updatability, cost estimates, and batch size. Column mappings support remote names and the row-key flag.

### Pushdown and Utilities

The cataloged 1.2 SQL installs connection inspection and cleanup functions, a version function, and local signatures for GridDB-specific functions used during deparsing. Those local stubs intentionally raise an error if executed in PostgreSQL; queries must be eligible for remote pushdown before relying on them.

Partial execution can be enabled for supported reads, and the README lists the functions and aggregates eligible for pushdown. Test the generated remote query with the deployed GridDB and extension versions rather than assuming every PostgreSQL expression is shippable.

### Limitations

The pinned upstream revision supports select, insert, update, and delete, and documents import of foreign schemas. It explicitly does not implement the foreign-data-wrapper truncate API. Bulk insert depends on the local PostgreSQL major version, and generated columns are evaluated by the wrapper because GridDB does not provide them.

Connection cache functions affect sessions that own those cached connections. Network credentials live in user mappings, so protect catalog visibility and use least-privileged GridDB users. No preload requirement is declared, but the server must be able to load the extension library and GridDB client library.
