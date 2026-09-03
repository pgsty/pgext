## Usage

Sources:

- [Official documentation](https://github.com/dalibo/pg_restore_points/blob/d1c8b71c10cd223b601d231bf2b726605cce5371/README.md)
- [Extension control file](https://github.com/dalibo/pg_restore_points/blob/d1c8b71c10cd223b601d231bf2b726605cce5371/pg_restore_points.control)
- [Official repository](https://github.com/dalibo/pg_restore_points)

`pg_restore_points` Create, catalog, list, and remove named PostgreSQL restore points.

### Enablement

Install the files for the intended server, then create `pg_restore_points` in the target database:

```sql
CREATE EXTENSION pg_restore_points;
```

### Core Workflow

The following example is taken from the reviewed upstream documentation. Adapt object names, paths, credentials, and workload values before use.

```sql
SELECT pg_purge_restore_points('interval_value');
```

### Main Objects

The reviewed install surface includes these prominent objects:

| Object | Kind | Operational role |
|---|---|---|
| `restore_points` | TABLE | Extension-owned table; account for its data in backup and upgrades. |
| `pg_extend_create_restore_point` | FUNCTION | Callable function from the reviewed install surface. |
| `pg_purge_restore_points` | FUNCTION | Callable function from the reviewed install surface. |
| `restore_point_mode` | TYPE | User-facing data type created by the extension. |

### Operations and Boundaries

- Treat extension upgrades as database changes: review the upstream upgrade path, privileges, locks, and backup/restore behavior first.
