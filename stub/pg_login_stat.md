## Usage

Sources:

- [Official documentation](https://github.com/asotolongo/pg_login_stat/blob/a0fbc06b6e3d7358561647a3893a702faa2a4774/README.md)
- [Extension control file](https://github.com/asotolongo/pg_login_stat/blob/a0fbc06b6e3d7358561647a3893a702faa2a4774/pg_login_stat.control)
- [Official repository](https://github.com/asotolongo/pg_login_stat)

`pg_login_stat` Per-user and per-database successful and failed authentication counters.

### Enablement

Merge `pg_login_stat` into the existing preload list, restart PostgreSQL, and then create `pg_login_stat` in each database that needs its SQL objects:

```ini
shared_preload_libraries = 'pg_login_stat'
```

```sql
CREATE EXTENSION pg_login_stat;
```

### Core Workflow

The following example is taken from the reviewed upstream documentation. Adapt object names, paths, credentials, and workload values before use.

```sql
SELECT * FROM pg_login_stats;
```

### Main Objects

The reviewed install surface includes these prominent objects:

| Object | Kind | Operational role |
|---|---|---|
| `pg_login_stats` | VIEW | Inspection or query view created by the extension. |
| `pg_login_stat_reset` | FUNCTION | Callable function from the reviewed install surface. |

### Operations and Boundaries

- The verified PostgreSQL-major evidence covers 15, 16, 17; do not infer unlisted majors.
- Preloading `pg_login_stat` changes cluster startup state; stage configuration and restart changes separately from `CREATE EXTENSION`.
