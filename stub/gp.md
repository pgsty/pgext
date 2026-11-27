## Usage

Sources:

- [Official source](https://github.com/RekGRpth/gp/blob/a52d096479a185f1ad362c1588e80efa466a5583/contrib/gp/README.md)
- [Official source](https://github.com/RekGRpth/gp/blob/a52d096479a185f1ad362c1588e80efa466a5583/contrib/gp/gp.control)
- [Official source](https://github.com/RekGRpth/gp/blob/a52d096479a185f1ad362c1588e80efa466a5583/contrib/gp/gp--1.0.sql)
- [Official source](https://github.com/RekGRpth/gp/blob/a52d096479a185f1ad362c1588e80efa466a5583/contrib/gp/gp.c)
- [Official source](https://github.com/RekGRpth/gp/blob/a52d096479a185f1ad362c1588e80efa466a5583/contrib/gp/LICENSE)

`gp` 1.0 is an experimental coordinator/segment extension built with the project's patched PostgreSQL 20 development tree. It requires the `new_oid_hook` patch; a released stock server is not a compatible replacement. The repository explicitly marks it as under development and unsuitable for production.

### Form a Test Cluster

Use the same patched server and extension binaries on every node. Preload the library on all nodes and restart them. A minimal coordinator configuration is:

```ini
shared_preload_libraries = 'gp'
gp.role = 'coordinator'
gp.database = 'postgres'
```

```sql
CREATE EXTENSION gp;
SELECT * FROM gp.cluster_segments();
```

Segments use `gp.role = 'segment'`, a positive unique `gp.segment_id`, and a reachable `gp.coordinator` connection string. Configure `max_prepared_transactions` above zero on segments. The registration worker creates its local SQL objects and registers in the coordinator's configured database. Installation needs a superuser; the fixed schema is `gp`, and its SQL uses `plpgsql`.

### Scope and Privileges

The coordinator stores segment metadata and keeps a shared-memory copy. DDL propagation preserves catalog OIDs across nodes, so all participating extension versions and roles must match. `gp.cluster_segments()` inspects registered nodes. Administrative `gp.exec_on_all()` can execute commands on segments; its SQL contract warns that coordinator rollback does not generally undo independently committed remote commands.

Distributed table placement, query dispatch, prepared transactions and recovery are experimental parts of this source tree. Its README's opening summary lags its detailed roadmap and SQL surface; qualify the exact revision and workload rather than assuming complete Greenplum compatibility. Mirrors and automatic redistribution are not established by the reviewed documentation.

### Connection and Recovery Boundaries

Node-to-node sessions need explicit authentication and access restrictions. Preserve the coordinator metadata, segment data, prepared-transaction state, interconnect secret and configuration together. Public schema access does not grant administrative function execution; the installation SQL revokes privileged control functions from PUBLIC. There is no versioned SQL upgrade chain in this revision. Changing the core patch or module requires a coordinated cluster migration and restart.
