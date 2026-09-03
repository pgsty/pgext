## Usage

Sources:

- [Official documentation](https://github.com/MasaoFujii/pg_exclusive_backup/blob/d6c38fc0e00a7fdbb9f5a9604c66d10a857c3d19/README.md)
- [Extension control file](https://github.com/MasaoFujii/pg_exclusive_backup/blob/d6c38fc0e00a7fdbb9f5a9604c66d10a857c3d19/pg_exclusive_backup.control)
- [Official repository](https://github.com/MasaoFujii/pg_exclusive_backup)

`pg_exclusive_backup` Compatibility functions for exclusive-style physical backup on PostgreSQL 15 or later.

### Enablement

Install the files for the intended server, then create `pg_exclusive_backup` in the target database:

```sql
CREATE EXTENSION pg_exclusive_backup;
```

### Core Workflow

The following example is taken from the reviewed upstream documentation. Adapt object names, paths, credentials, and workload values before use.

```sql
SELECT pg_start_backup('nightly', true);
SELECT pg_is_in_backup();

-- Copy the physical backup while backup mode is active.
SELECT pg_stop_backup();
```

### Main Objects

The reviewed install surface includes these prominent objects:

| Object | Kind | Operational role |
|---|---|---|
| `pg_catalog.pg_start_backup` | FUNCTION | Callable function from the reviewed install surface. |
| `pg_catalog.pg_stop_backup` | FUNCTION | Callable function from the reviewed install surface. |
| `pg_catalog.pg_backup_start_time` | FUNCTION | Callable function from the reviewed install surface. |
| `pg_catalog.pg_is_in_backup` | FUNCTION | Callable function from the reviewed install surface. |

### Operations and Boundaries

- The verified PostgreSQL-major evidence covers 15, 16, 17, 18; do not infer unlisted majors.
- The extension fixes or creates schema objects under `pg_catalog`; include them in privilege and backup review.
- Catalog lifecycle is deprecated; test upgrades, dump/restore, and server compatibility before production use.
