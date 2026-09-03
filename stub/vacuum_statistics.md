## Usage

Sources:

- [Official documentation](https://github.com/Alena0704/vacuum_statistics/blob/c40b73e3bcce0343b9e4c86eab353be26d6e006c/README.md)
- [Extension control file](https://github.com/Alena0704/vacuum_statistics/blob/c40b73e3bcce0343b9e4c86eab353be26d6e006c/vacuum_statistics.control)
- [Official repository](https://github.com/Alena0704/vacuum_statistics)

`vacuum_statistics` Shared-memory table, database, and index statistics for VACUUM activity.

### Enablement

Merge `vacuum_statistics` into the existing preload list, restart PostgreSQL, and then create `vacuum_statistics` in each database that needs its SQL objects:

```ini
shared_preload_libraries = 'vacuum_statistics'
```

```sql
CREATE EXTENSION vacuum_statistics;
```

### Core Workflow

The following example is taken from the reviewed upstream documentation. Adapt object names, paths, credentials, and workload values before use.

```sql
SELECT relname, tuples_deleted, pages_scanned, pages_removed, total_time
FROM pg_stats_vacuum_tables
ORDER BY total_time DESC
LIMIT 10;
```

### Main Objects

The reviewed install surface includes these prominent objects:

| Object | Kind | Operational role |
|---|---|---|
| `extvac_reset_entry` | FUNCTION | Callable function from the reviewed install surface. |
| `pg_stats_get_vacuum_database` | FUNCTION | Callable function from the reviewed install surface. |
| `pg_stats_get_vacuum_indexes` | FUNCTION | Callable function from the reviewed install surface. |
| `pg_stats_get_vacuum_tables` | FUNCTION | Callable function from the reviewed install surface. |
| `pg_stats_vacuum_database` | VIEW | Inspection or query view created by the extension. |
| `pg_stats_vacuum_indexes` | VIEW | Inspection or query view created by the extension. |
| `pg_stats_vacuum_tables` | VIEW | Inspection or query view created by the extension. |
| `extvac_reset_db_entry` | FUNCTION | Callable function from the reviewed install surface. |

### Operations and Boundaries

- Preloading `vacuum_statistics` changes cluster startup state; stage configuration and restart changes separately from `CREATE EXTENSION`.
