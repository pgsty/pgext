## Usage

Sources:

- [`documentdb_distributed` v0.117-0 control file](https://github.com/documentdb/documentdb/blob/v0.117-0/internal/pg_documentdb_distributed/documentdb_distributed.control)
- [DocumentDB package scope](https://github.com/documentdb/documentdb/blob/v0.117-0/packaging/README.md)
- [Official preload helper](https://github.com/documentdb/documentdb/blob/v0.117-0/scripts/preload_libraries.sh)
- [DocumentDB v0.117-0 changelog](https://github.com/documentdb/documentdb/blob/v0.117-0/CHANGELOG.md)

`documentdb_distributed` is DocumentDB's internal multi-node execution layer. It integrates the public `documentdb` API with Citus; it is not a standalone document API and does not add a separate client workflow.

### Prerequisites and Install

All nodes need matching builds of Citus, `documentdb_core`, `documentdb`, and `documentdb_distributed`. The official helper places the libraries in preload order; restart every node after changing it:

```conf
shared_preload_libraries = 'pg_cron, citus, pg_documentdb_core, pg_documentdb, pg_documentdb_distributed, pg_documentdb_extended_rum'
```

After the Citus topology and the base DocumentDB stack are configured, install the distributed component as a superuser:

```sql
CREATE EXTENSION documentdb CASCADE;
CREATE EXTENSION documentdb_extended_rum;
CREATE EXTENSION documentdb_distributed;

SELECT extname, extversion
FROM pg_extension
WHERE extname IN ('citus', 'documentdb_core', 'documentdb', 'documentdb_distributed');
```

Use the normal DocumentDB gateway or `documentdb_api` functions after installation. Collection placement, shard topology, worker availability, and metadata consistency must be managed as part of the cluster deployment.

### Important Boundaries

- The control file requires `citus`, `documentdb_core`, and `documentdb` and marks the extension superuser-only and non-relocatable.
- `documentdb_distributed` supplies distributed planner and execution support to existing DocumentDB commands; it is not an alternative to the public API extension.
- Release versions should remain synchronized across coordinators and workers before extension upgrades are attempted.
- Backup, restore, failover, and rolling-upgrade procedures must include both Citus metadata and DocumentDB data.

This component lives under the upstream internal directory. Confirm that the selected distribution actually ships its control file, SQL and library before enabling it; a base DocumentDB package alone is not evidence of distributed-component availability. Keep Citus and all DocumentDB components on compatible builds.
