## Usage

Sources:

- [Pinned upstream README](https://github.com/pgq/londiste-sql/blob/55085fc4e8e6cb00e2b6255427fd86c50cd6fef3/README.rst)
- [Pinned extension control file](https://github.com/pgq/londiste-sql/blob/55085fc4e8e6cb00e2b6255427fd86c50cd6fef3/londiste.control)
- [Pinned schema installation script](https://github.com/pgq/londiste-sql/blob/55085fc4e8e6cb00e2b6255427fd86c50cd6fef3/londiste--3.8.sql)
- [Pinned function sources](https://github.com/pgq/londiste-sql/tree/55085fc4e8e6cb00e2b6255427fd86c50cd6fef3/functions)
- [Pinned privilege grants](https://github.com/pgq/londiste-sql/blob/55085fc4e8e6cb00e2b6255427fd86c50cd6fef3/structure/grants.sql)

londiste version 3.8 supplies the PostgreSQL-side support schema for Londiste logical replication. It is not a standalone replication daemon. Its tables and functions track registered tables and sequences, merge state, pending foreign keys, partition maintenance, replication triggers, and replicated DDL execution. The control file requires pgq_node and requires a superuser to install the extension.

### Install and inspect

Install the PGQ components from matching packages before Londiste, then let the Londiste administration process register queues and tables:

```sql
CREATE EXTENSION pgq;
CREATE EXTENSION pgq_node;
CREATE EXTENSION londiste;

SELECT extname, extversion
FROM pg_extension
WHERE extname IN ('pgq', 'pgq_node', 'londiste')
ORDER BY extname;

SELECT queue_name, table_name, local, merge_state
FROM londiste.table_info
ORDER BY queue_name, table_name;
```

The SQL repository contains server-side support code only. Queue topology, provider/subscriber setup, daemon configuration, copy, catch-up, failover, and maintenance must follow the documentation for the matching Londiste/SkyTools client package.

### Destructive DDL and privilege boundaries

The API deliberately creates and removes replication triggers, drops and restores foreign keys, changes session_replication_role, drops obsolete partition tables, and records or emits DDL scripts. These operations can suppress normal triggers and constraints or destroy tables. Use a dedicated database owner and daemon role, keep the replication topology under configuration control, and test repair procedures on a copy.

The reviewed set_session_replication_role function is SECURITY DEFINER and has no protected search_path. PostgreSQL functions are executable by PUBLIC by default, while the supplied grants script grants public schema usage and read access to several state tables without revoking function execution. Revoke execution on every londiste function from PUBLIC after installation, then grant only the exact administration entry points required by the trusted daemon role. Review all SECURITY DEFINER functions and schema ownership before allowing untrusted users into the database.

Londiste state must remain consistent with PGQ events and the external daemon. Do not edit table_info, seq_info, pending_fkeys, or applied_execute manually. Take coordinated backups, monitor queue lag and failed events, and rehearse upgrades on every node. The extension is fixed to the pg_catalog installation schema, is not relocatable, and its SQL must match the installed pgq_node version; mixing independently sourced versions can leave unusable functions or replication metadata.
