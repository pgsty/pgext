## Usage

Sources:

- [Official documentation](https://github.com/maheshk32/pg_table_guard/blob/65a85d8f451986ee662bb379ee631d159c1d51fa/README.md)
- [Extension control file](https://github.com/maheshk32/pg_table_guard/blob/65a85d8f451986ee662bb379ee631d159c1d51fa/table_guard.control)
- [Official repository](https://github.com/maheshk32/pg_table_guard)

`table_guard` Configurable protection against DROP, TRUNCATE, and unqualified whole-table DELETE.

### Enablement

Merge `table_guard` into the existing preload list, restart PostgreSQL, and then create `table_guard` in each database that needs its SQL objects:

```ini
shared_preload_libraries = 'table_guard'
```

```sql
CREATE EXTENSION table_guard;
```

### Core Workflow

The following example is taken from the reviewed upstream documentation. Adapt object names, paths, credentials, and workload values before use.

```sql
SELECT * FROM table_guard_settings;
```

### Main Objects

The reviewed install surface includes these prominent objects:

| Object | Kind | Operational role |
|---|---|---|
| `table_guard_help` | FUNCTION | Callable function from the reviewed install surface. |
| `table_guard_settings` | VIEW | Inspection or query view created by the extension. |
| `table_guard_version` | FUNCTION | Callable function from the reviewed install surface. |

### Operations and Boundaries

- The verified PostgreSQL-major evidence covers 14, 15, 16, 17, 18; do not infer unlisted majors.
- Preloading `table_guard` changes cluster startup state; stage configuration and restart changes separately from `CREATE EXTENSION`.
