## Usage

Sources:

- [Official README at the catalog source revision](https://github.com/luxms/greenplum-fdw/blob/fc8f03883391b72940ff8ceb9dc0aa56644a3c67/README.md)
- [Extension control file at the catalog source revision](https://github.com/luxms/greenplum-fdw/blob/fc8f03883391b72940ff8ceb9dc0aa56644a3c67/greenplum_fdw.control)
- [Version 1.0 extension SQL](https://github.com/luxms/greenplum-fdw/blob/fc8f03883391b72940ff8ceb9dc0aa56644a3c67/greenplum_fdw--1.0.sql)
- [PostgreSQL postgres_fdw documentation](https://www.postgresql.org/docs/current/postgres-fdw.html)

greenplum_fdw is a fork of PostgreSQL's foreign-data wrapper for querying a remote Greenplum Database over libpq. The cataloged source changes remote transactions to serializable isolation because older Greenplum releases reject the repeatable-read transaction started by the corresponding upstream wrapper.

### Core Workflow

Create the wrapper, describe the remote Greenplum endpoint, map a local role to remote credentials, and define or import foreign tables.

```sql
CREATE EXTENSION greenplum_fdw;

CREATE SERVER gp_sales
FOREIGN DATA WRAPPER greenplum_fdw
OPTIONS (host 'greenplum.example.net', port '5432', dbname 'warehouse');

CREATE USER MAPPING FOR app_user
SERVER gp_sales
OPTIONS (user 'reporter', password 'secret');

CREATE FOREIGN TABLE gp_orders (
    order_id bigint,
    ordered_at timestamp,
    total numeric
)
SERVER gp_sales
OPTIONS (schema_name 'public', table_name 'orders');

SELECT order_id, total
FROM gp_orders
WHERE ordered_at >= current_date - interval '7 days';
```

### Important Options

Connection keywords accepted by libpq belong on the foreign server; credentials belong on the user mapping. Table and column mappings can override remote names. The wrapper source also recognizes remote-estimate, startup-cost, tuple-cost, fetch-size, updatability, and shippable-extension settings inherited from its postgres_fdw code base.

### Compatibility and Caveats

Use a greenplum_fdw source branch built for the exact local PostgreSQL major version. The upstream README describes separate PostgreSQL-major branches and an older Greenplum compatibility problem; do not assume the cataloged 1.0 source follows the feature set or version range of current postgres_fdw.

The wrapper's defining behavior is the remote transaction isolation change. Verify it against the target Greenplum version and workload, especially for writes and multi-statement transactions. Remote permissions, network reachability, TLS parameters, and credential handling remain ordinary libpq and foreign-server responsibilities. No preload or server restart is declared by the extension.
