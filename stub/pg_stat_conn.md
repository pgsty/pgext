## Usage

Sources:

- [Official documentation](https://github.com/guedes/pg_stat_conn/blob/8e5576e55e43238e94f42712ea3c0b138b64632b/README.md)
- [Extension control file](https://github.com/guedes/pg_stat_conn/blob/8e5576e55e43238e94f42712ea3c0b138b64632b/pg_stat_conn.control)
- [Official repository](https://github.com/guedes/pg_stat_conn)

`pg_stat_conn` Cumulative connection, disconnection, and authentication-failure counts by database and role.

### Enablement

Merge `pg_stat_conn` into the existing preload list, restart PostgreSQL, and then create `pg_stat_conn` in each database that needs its SQL objects:

```ini
shared_preload_libraries = 'pg_stat_conn'
```

```sql
CREATE EXTENSION pg_stat_conn;
```

### Core Workflow

The following example is taken from the reviewed upstream documentation. Adapt object names, paths, credentials, and workload values before use.

```sql
psql -c "SELECT pg_stat_conn_reset('postgres', 'postgres');"
psql -c "SELECT pg_stat_conn_reset();"
```

### Main Objects

The reviewed install surface includes these prominent objects:

| Object | Kind | Operational role |
|---|---|---|
| `pg_stat_conn_reset` | FUNCTION | Callable function from the reviewed install surface. |

### Operations and Boundaries

- The verified PostgreSQL-major evidence covers 15, 16, 17, 18; do not infer unlisted majors.
- Preloading `pg_stat_conn` changes cluster startup state; stage configuration and restart changes separately from `CREATE EXTENSION`.
