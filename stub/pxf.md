## Usage

Sources:

- [Official project page](https://github.com/greenplum-db/gpdb-archive)

`pxf` — Historical Greenplum extension for accessing unmanaged data through PXF. Use it when PostgreSQL must access the corresponding external data source through a foreign-data interface. Upstream describes it as a work in progress.

### Core Workflow

This component has no confirmed standalone `CREATE EXTENSION` workflow in the reviewed source. Build, load, or enable it only through the exact upstream mechanism, then verify the resulting server behavior in an isolated database.

### Requirements and Caveats

- The catalog records version `1.0`.
- Upstream describes the project as a work in progress.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
