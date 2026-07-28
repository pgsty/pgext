## Usage

Sources:

- [Official upstream README](https://api.pgxn.org/src/tblsize_nolock/tblsize_nolock-1.0.0/README)

`tblsize_nolock` — Functions to calculate relation size without grabbing AccessShareLock on a target relation. Use it when administering or automating the database behavior described above. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

This component has no confirmed standalone `CREATE EXTENSION` workflow in the reviewed source. Build, load, or enable it only through the exact upstream mechanism, then verify the resulting server behavior in an isolated database.

### Requirements and Caveats

- The catalog records version `1.0.0`.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
