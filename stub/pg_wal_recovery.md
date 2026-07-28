## Usage

Sources:

- [Official upstream README](https://github.com/misachi/pg_wal_recovery/blob/61ac3633d26bf414d9b145b153bc96d6b0e75a77/README.md)
- [Official extension control file (pg_wal_recovery.control)](https://github.com/misachi/pg_wal_recovery/blob/61ac3633d26bf414d9b145b153bc96d6b0e75a77/pg_wal_recovery.control)
- [Official extension SQL (pg_wal_recovery--1.0.sql)](https://github.com/misachi/pg_wal_recovery/blob/61ac3633d26bf414d9b145b153bc96d6b0e75a77/pg_wal_recovery--1.0.sql)

`pg_wal_recovery` — pg_wal_recovery is an educational PostgreSQL extension for database recovery, focusing on restoring databases from Write-Ahead Logs (WAL) and supporting point-in-time recovery. Check my post on it here. Use it when administering or automating the database behavior described above. Upstream describes it as a work in progress.

### Core Workflow

```sql
CREATE EXTENSION pg_wal_recovery;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `wal_list_records` is an extension function.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- The control file marks the extension as relocatable.
- Upstream describes the project as a work in progress.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
