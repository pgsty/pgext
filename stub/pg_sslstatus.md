## Usage

Sources:

- [Official documentation](https://github.com/mhagander/pg_sslstatus/blob/906882efacf0c40907910dc09274a3447c6a58b8/README.md)
- [Extension control file](https://github.com/mhagander/pg_sslstatus/blob/906882efacf0c40907910dc09274a3447c6a58b8/pg_sslstatus.control)
- [Official repository](https://github.com/mhagander/pg_sslstatus)

`pg_sslstatus` Historical SSL connection-status views for PostgreSQL releases before 9.5.

### Enablement

Merge `pg_sslstatus` into the existing preload list, restart PostgreSQL, and then create `pg_sslstatus` in each database that needs its SQL objects:

```ini
shared_preload_libraries = 'pg_sslstatus'
```

```sql
CREATE EXTENSION pg_sslstatus;
```

### Core Workflow

The following example is taken from the reviewed upstream documentation. Adapt object names, paths, credentials, and workload values before use.

```sql
SELECT * FROM pg_sslstatus;
```

### Main Objects

The reviewed install surface includes these prominent objects:

| Object | Kind | Operational role |
|---|---|---|
| `pg_get_sslstatus` | FUNCTION | Callable function from the reviewed install surface. |

### Operations and Boundaries

- Preloading `pg_sslstatus` changes cluster startup state; stage configuration and restart changes separately from `CREATE EXTENSION`.
- Catalog lifecycle is deprecated; test upgrades, dump/restore, and server compatibility before production use.
