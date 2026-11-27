## Usage

Sources:

- [Release 3.3.0](https://github.com/matroidbe/pg_extensions-releases/releases/tag/pg_aurum-v3.3.0)
- [Official control, SQL and manifest archive](https://github.com/matroidbe/pg_extensions-releases/releases/download/pg_aurum-v3.3.0/pg_aurum-3.3.0-pg18-linux-amd64.tar.gz)
- [SHA256SUMS](https://github.com/matroidbe/pg_extensions-releases/releases/download/pg_aurum-v3.3.0/SHA256SUMS)
- [MIT license](https://github.com/matroidbe/pg_extensions-releases/releases/download/pg_aurum-v3.3.0/LICENSE-pg_aurum)

`pg_aurum` 3.3.0 provides the WAL worker and in-database compiler for Aurum incremental pipelines. The reviewed release publishes PostgreSQL 18 binaries for Linux amd64 and macOS arm64, with control, installation SQL, upgrade SQL and an MIT license. Implementation source is not published in this release repository.

### Enable and Inspect

Use a verified upstream archive for the exact server platform. The archive manifest requires preload, logical WAL and a restart. Preserve existing preload entries and set the worker database explicitly:

```ini
shared_preload_libraries = 'pg_aurum'
wal_level = logical
aurum.database = 'postgres'
```

Create the extension as a superuser in that database. It is nonrelocatable and creates `aurum` and `_aurum` schemas, an FDW/server, event triggers and pipeline catalogs.

```sql
CREATE EXTENSION pg_aurum;
SELECT aurum_worker_status();
SELECT * FROM aurum_wal_status();
SELECT * FROM aurum_checkpoint_status();
```

### Pipeline Lifecycle

For a separately reviewed Aurum pipeline definition, bind its name and source text as SQL parameters:

```sql
SELECT * FROM aurum.install($1::text, $2::text);
SELECT * FROM aurum.start($1::text);
```

`aurum.install` compiles and installs a definition; `aurum.install_compiled` accepts already compiled SQL. Foreign-table declarations using server `aurum` can instead be staged and inspected through `aurum.deploy_preview`, then deployed with `aurum.deploy`. The default existing-object policy is `error`.

`aurum.stop`, `aurum.restart`, `aurum.recompile` and `aurum.uninstall` change pipeline state. Their SQL return columns include dropped publications, truncation and initial loads: review the requested operation before executing it against retained data. `aurum_process_pending`, `aurum_stats` and `aurum_force_checkpoint` expose worker progress and checkpoint controls.

### Search and Boundaries

For configured search targets, `aurum.search` returns keys, scores and highlights; `aurum.search_rows` also returns documents. `aurum.search_chunks` retrieves passages for RAG, while `aurum.search_facets` and `aurum.search_suggest` expose facets and suggestions. These wrappers accept a target relation and query; their defaults and return columns are defined in the installation SQL.

The SQL wrappers document caller-based row rechecks. The installation SQL does not establish a restrictive PUBLIC privilege policy for every native function. Restrict pipeline compilation, deployment and worker administration to trusted roles and audit grants before exposing search to applications. Logical replication slots can retain WAL; monitor slot lag and worker health. These binary and SQL checks establish identity and interface, not production qualification of the unpublished implementation.
