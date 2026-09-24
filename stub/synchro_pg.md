## Usage

Sources:

- [README](https://github.com/trainstar/synchro/blob/3de1e69cfeb7a093b43d7e986ead8cdae8eab4f2/README.md)
- [Control file](https://github.com/trainstar/synchro/blob/3de1e69cfeb7a093b43d7e986ead8cdae8eab4f2/extensions/synchro-pg/synchro_pg.control)
- [extensions/synchro-pg/sql/synchro_pg--0.3.0.sql](https://github.com/trainstar/synchro/blob/3de1e69cfeb7a093b43d7e986ead8cdae8eab4f2/extensions/synchro-pg/sql/synchro_pg--0.3.0.sql)
- [extensions/synchro-pg/src/registry.rs](https://github.com/trainstar/synchro/blob/3de1e69cfeb7a093b43d7e986ead8cdae8eab4f2/extensions/synchro-pg/src/registry.rs)
- [docs/src/content/docs/operations/configuration.mdx](https://github.com/trainstar/synchro/blob/3de1e69cfeb7a093b43d7e986ead8cdae8eab4f2/docs/src/content/docs/operations/configuration.mdx)

`synchro_pg` is the database component of Synchro’s offline-first synchronization stack. The pinned extension is 0.3.0: PostgreSQL captures and materializes scoped changes, an authenticated host adapter serves clients, and native clients maintain local SQLite state.

### Enablement and Readiness

The documented target is PostgreSQL 18. Install as a superuser, preload `synchro_pg`, enable logical WAL, and configure replication slots. Set `synchro.database` and `synchro.worker_login` to the intended database and dedicated worker login; postmaster settings require restart. Provision that login with the documented `synchro_worker` role and replication privileges.

```sql
CREATE EXTENSION synchro_pg;
SELECT synchro.synchro_readiness();
```

### Synchronization Workflow

Define application tables with stable primary keys and a deterministic membership function mapping rows to server-controlled scopes. Register each table using `synchro.synchro_register_table`, supplying the relation, membership function, composition, timestamp/deletion columns, synced columns, and affected scopes. Registration checks ownership and function dependencies; application migrations must preserve that contract.

Expose the extension through the Go adapter or another implementation of the same authenticated protocol. Clients connect, push local intent, pull scoped changes using opaque cursors, and rebuild when required. The extension owns change capture and CRUD semantics; the host adapter owns HTTP and authentication. Portable seed databases provide only server-declared portable data and do not grant authorization.

### Operations and Limits

`synchro.replication_slot` and `synchro.publication_name` select logical-capture objects. `synchro.max_worker_heartbeat_age_seconds`, `synchro.max_wal_lag_bytes`, and `synchro.max_wal_lag_seconds` set positive readiness limits. Monitor readiness, retained WAL, and source-transaction size. Decoding is bounded to 16 MiB and 10,000 payload records per source transaction; an oversized transaction blocks progress until handled.

The extension is non-relocatable. Installation creates dedicated roles and revokes public access to its schemas, tables, and functions. Grant the documented operator, adapter, worker, monitor, and seed roles according to function. Older repository v0.1.x releases do not identify this extension version.
