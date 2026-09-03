## Usage

Sources:

- [Official documentation](https://github.com/vladich/auto_explain_z/blob/b281da41ca5b2fdb03adaaa82a700f5ac75822cc/README.md)
- [Extension control file](https://github.com/vladich/auto_explain_z/blob/b281da41ca5b2fdb03adaaa82a700f5ac75822cc/auto_explain_z.control)
- [Official repository](https://github.com/vladich/auto_explain_z)

`auto_explain_z` Compressed binary execution-plan logging with rotation and offline decoding.

### Enablement

Merge `auto_explain_z` into the existing preload list, restart PostgreSQL, and then create `auto_explain_z` in each database that needs its SQL objects:

```ini
shared_preload_libraries = 'auto_explain_z'
```

```sql
CREATE EXTENSION auto_explain_z;
```

### Core Workflow

The following example is taken from the reviewed upstream documentation. Adapt object names, paths, credentials, and workload values before use.

```sql
SELECT auto_explain_z_rotate_logfile();
```

### Main Objects

The reviewed install surface includes these prominent objects:

| Object | Kind | Operational role |
|---|---|---|
| `auto_explain_z_rotate_logfile` | FUNCTION | Callable function from the reviewed install surface. |

### Operations and Boundaries

- The verified PostgreSQL-major evidence covers 14, 15, 16, 17, 18; do not infer unlisted majors.
- Preloading `auto_explain_z` changes cluster startup state; stage configuration and restart changes separately from `CREATE EXTENSION`.
