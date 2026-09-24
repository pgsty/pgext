## Usage

Sources:

- [Standalone FDW guide](https://github.com/mkleczek/pgwrh/blob/v1.0.0-alpha1/pgwrh_fdw/README.md)
- [Control file](https://github.com/mkleczek/pgwrh/blob/v1.0.0-alpha1/pgwrh_fdw/pgwrh_fdw.control)
- [Extension SQL](https://github.com/mkleczek/pgwrh/blob/v1.0.0-alpha1/pgwrh_fdw/pgwrh_fdw--1.0.0-alpha1.sql)
- [Licensing](https://github.com/mkleczek/pgwrh/blob/v1.0.0-alpha1/pgwrh_fdw/LICENSING.md)

`pgwrh_fdw` 1.0.0-alpha1 is a PostgreSQL 18 foreign data wrapper derived from `postgres_fdw`. It adds selected transaction-setting propagation and routing among replica servers. It works independently of the pgwrh core and is licensed AGPL-3.0-only, with inherited PostgreSQL notices preserved. This alpha supports fresh installation only.

### Access a Remote Table

Enable the extension as an administrator; no shared preload is required. Create a server and arrange a user mapping whose remote role can read the target table. This example assumes the mapping is already configured, the remote database contains `public.items`, and the local schema has no conflicting table.

```sql
CREATE EXTENSION pgwrh_fdw;
CREATE SERVER replica_a FOREIGN DATA WRAPPER pgwrh_fdw
    OPTIONS (host 'replica-a', dbname 'replica');
IMPORT FOREIGN SCHEMA public LIMIT TO (items)
    FROM SERVER replica_a INTO public;
SELECT * FROM items;
```

### Propagate Transaction Settings

```sql
ALTER SERVER replica_a OPTIONS (ADD transaction_parameters 'app.request_id');
BEGIN;
SET LOCAL app.request_id = 'request-42';
SELECT * FROM items;
COMMIT;
```

`transaction_parameters` is a foreign-server option containing a nonempty comma-separated list of dotted custom parameter names. Core settings, duplicates, and the `postgres_fdw.*` and `pgwrh_fdw.*` namespaces are rejected. Names are normalized to lowercase and limited to 63 bytes.

Set context before planning or accessing foreign tables. The first participating remote transaction captures the visible custom settings; every participating server uses that capture until the local top-level transaction ends. Later local changes are not propagated. Remote estimation can trigger capture during planning. Remote roles and parameter privileges still apply, and setting values may appear in logs.

### Virtual Servers and Helpers

The `members` server option lists ordinary servers behind a virtual server. A virtual server needs an empty user mapping; credentials come from the selected member. Selection favors reusable connections, then uses `load_balance_weight`, and remains fixed for the transaction. Only initial connection failure can try another member; established transactions and query errors never fail over.

`pgwrh_fdw_set_members` changes membership while waiting for users of the previous routing configuration. `pgwrh_fdw_get_connections`, `pgwrh_fdw_disconnect`, and `pgwrh_fdw_disconnect_all` manage cached connections. `pgwrh_fdw_scram_verifier` creates a newly salted SCRAM verifier; protect its inputs and result as authentication material.

### Consistency Boundaries

Propagation itself does not wait for replication, interpret LSNs, provide a global snapshot, or coordinate distributed commits. For replication barriers, configure `pgwrh_wait` on every actual receiving server and configure propagation explicitly. Successful setting assignment alone does not prove the receiver implements the requested behavior. Existing foreign servers are not converted automatically.
