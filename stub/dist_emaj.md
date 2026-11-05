## Usage

Sources:

- [Upstream source 1](https://api.pgxn.org/src/dist_emaj/dist_emaj-0.9.0/README.md)
- [Upstream source 2](https://api.pgxn.org/src/dist_emaj/dist_emaj-0.9.0/dist_emaj.control)
- [Upstream source 3](https://api.pgxn.org/src/dist_emaj/dist_emaj-0.9.0/sql/dist_emaj--0.9.0.sql)
- [Upstream source 4](https://api.pgxn.org/src/dist_emaj/dist_emaj-0.9.0/client/distEmaj.pl)

`dist_emaj` 0.9.0 coordinates E-Maj operations across databases so marked groups can be rolled back consistently.

### Core Workflow

Install the coordinator with its `dblink` dependency. Register connection definitions with `dist_emaj.create_database`, define a logical cluster with `dist_emaj.create_cluster`, then assign E-Maj groups using `dist_emaj.assign_group`. Run `dist_emaj.verify_cluster` before coordinated operations. The supplied `distEmaj.pl` and `distEmajRollback.pl` clients drive distributed marks and rollbacks.

```sql
CREATE EXTENSION dist_emaj CASCADE;
SELECT extversion FROM pg_extension WHERE extname = 'dist_emaj';
```

### Operational Boundaries

Every managed database needs compatible E-Maj and authorized remote connections. Distributed rollback uses two-phase commit: configure prepared-transaction capacity and retain a recovery procedure for interrupted operations. Treat connection secrets as privileged data and follow the versioned SQL function signatures. This extension does not create distributed transaction semantics for unrelated application writes.
