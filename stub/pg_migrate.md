## Usage

Sources:

- [Official upstream README](https://api.pgxn.org/src/pg_migrate/pg_migrate-0.1.1/README.md)

`pg_migrate` — pg_migrate is a PostgreSQL extension and CLI which lets you make schema changes to tables and indexes. Unlike ALTER TABLE it works online, without holding a long lived exclusive lock on the processed tables during the migration. It builds a copy of the target table and swaps them. Use it when moving, transforming, or integrating the corresponding data from PostgreSQL. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

The reviewed distribution uses a legacy SQL or non-control installation layout, so it does not establish a modern standalone `CREATE EXTENSION` and upgrade workflow. Follow the pinned upstream installation mechanism and verify the installed objects in an isolated database.

### Requirements and Caveats

- The catalog records version `0.1.1`.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
