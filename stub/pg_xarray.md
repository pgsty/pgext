## Usage

Sources:

- [CHANGELOG.md](https://github.com/matroidbe/pg_extensions-releases/blob/76ba568872dcf082e73690057e5404beed27b189/CHANGELOG.md)
- [LICENSE](https://github.com/matroidbe/pg_extensions-releases/blob/76ba568872dcf082e73690057e5404beed27b189/LICENSE)
- [extensions/pg_xarray/pg_xarray.control](https://github.com/matroidbe/pg_extensions-releases/blob/76ba568872dcf082e73690057e5404beed27b189/extensions/pg_xarray/pg_xarray.control)
- [extensions/pg_xarray/Cargo.toml](https://github.com/matroidbe/pg_extensions-releases/blob/76ba568872dcf082e73690057e5404beed27b189/extensions/pg_xarray/Cargo.toml)
- [extensions/pg_xarray/src/lib.rs](https://github.com/matroidbe/pg_extensions-releases/blob/76ba568872dcf082e73690057e5404beed27b189/extensions/pg_xarray/src/lib.rs)
- [extensions/pg_xarray/sql/pg_xarray--0.4.2--0.5.0.sql](https://github.com/matroidbe/pg_extensions-releases/blob/76ba568872dcf082e73690057e5404beed27b189/extensions/pg_xarray/sql/pg_xarray--0.4.2--0.5.0.sql)
- [crates/pg_bgworker/src/supervision.rs](https://github.com/matroidbe/pg_extensions-releases/blob/76ba568872dcf082e73690057e5404beed27b189/crates/pg_bgworker/src/supervision.rs)
- [extensions/pg_xarray/src/server/worker.rs](https://github.com/matroidbe/pg_extensions-releases/blob/76ba568872dcf082e73690057e5404beed27b189/extensions/pg_xarray/src/server/worker.rs)
- [extensions/pg_xarray/pgbrew.toml](https://github.com/matroidbe/pg_extensions-releases/blob/76ba568872dcf082e73690057e5404beed27b189/extensions/pg_xarray/pgbrew.toml)
- [extensions/pg_xarray/demo/01_register_local.sql](https://github.com/matroidbe/pg_extensions-releases/blob/76ba568872dcf082e73690057e5404beed27b189/extensions/pg_xarray/demo/01_register_local.sql)

`pg_xarray` 0.5.0 catalogs scientific datasets, variables, chunks and meshes in `pgx`, with scientific-array queries and optional WMS serving.

### Core Workflow

```sql
CREATE EXTENSION postgis;
CREATE EXTENSION pg_xarray;
SELECT pgx.register_dataset('weather', 'zarr');
SELECT * FROM pgx.list_datasets();
```

### Operational Boundaries

Requires `postgis` and superuser installation. Dataset/file registration points at server-readable local or object-storage resources; restrict those paths and credentials. The SQL surface includes `register_dataset`, `register_file`, `register_variable`, `register_chunk` and `list_datasets`. WMS requires preload/restart and explicit `pg_xarray.wms_enabled`; it defaults off. Enabling the SQL catalog does not start a public tile service. Preserve the external array data and catalog together when backing up. This is an unsupported proof of concept under Matroid Source Available License 1.0. APIs may change. Version 0.3.0 is the new upgrade baseline: earlier 0.2.0 installations require a rehearsed data migration/recreation, not ordinary ALTER EXTENSION UPDATE. Back up data and dependencies before following that destructive upstream path.

### Current Release and Upgrade

Extension version 0.5.0 uses `pg_xarray.database` for background workers. Create the extension in that database; a worker now waits instead of repeatedly exiting when it is absent. Changing restart-sensitive worker configuration requires a restart. The chunk catalog now includes `chunk_key` in deduplication; after upgrading, re-register affected Zarr variables to recover missing levels or tiles. From extension 0.3.0 onward, install matching files and use ALTER EXTENSION UPDATE. Earlier versions still require the migration described above.

### Worker Supervision

Version 0.4.2 adds `pgx.worker_status()` and the superuser-only `pgx.reset_workers()`. Status reports worker name, state, failures, restarts, PID, time of entry into the current state and last failure. `pg_xarray.max_worker_failures` defaults to 10 consecutive failures; 0 retries indefinitely. Restart delays back off from 5 to 60 seconds, and reaching the limit leaves a failed worker idle until reset.

Install matching libraries, restart PostgreSQL to initialize the new shared memory, and update the extension SQL in the configured worker database. Inspect the failure before resetting; a reset does not repair its cause.

```sql
ALTER EXTENSION pg_xarray UPDATE;
SELECT * FROM pgx.worker_status();
```

### Changes in This Release

Source feature declarations cover PostgreSQL 14–18. The repository release is 0.7.0, while each extension retains its own SQL version; this does not establish Pigsty package availability.

Version 0.5.0 applies the SELAFIN IPARAM origin. The new `register_file` argument `transform` accepts a 2D affine in `ST_Affine` order, mapping model-local coordinates to the chosen SRID. Re-registering moves the mesh and its extent, so check coordinate parameters first. WMS GetMap accepts `COLORSCALERANGE`, `BELOWMINCOLOR` and `ABOVEMAXCOLOR`; `xarray_to_png` adds `below_min` and `above_max`. Capabilities now reports geographic extents in degrees, and responses carry `Access-Control-Allow-Origin: *`; review the external access boundary when deploying.
