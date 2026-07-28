## Usage

Sources:

- [Official upstream README](https://api.pgxn.org/src/dbpatch/dbpatch-1.3.0/README.md)
- [Official extension SQL (dbpatch--unpackaged--1.0.0.sql)](https://api.pgxn.org/src/dbpatch/dbpatch-1.3.0/sql/dbpatch--unpackaged--1.0.0.sql)

`dbpatch` — postgresql-dbpatch ==================. Use it when administering or automating the database behavior described above. Its extension dependencies must be installed and validated first.

### Core Workflow

The reviewed distribution uses a legacy SQL or non-control installation layout, so it does not establish a modern standalone `CREATE EXTENSION` and upgrade workflow. Follow the pinned upstream installation mechanism and verify the installed objects in an isolated database.

### Important Objects

- `IF` is a table installed or managed by the extension.

### Requirements and Caveats

- The catalog records version `1.3.0`.
- Install the confirmed extension dependencies first: `plpgsql`.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
