## Usage

Sources:

- [Official documentation](https://github.com/kmtr/pg_chaos_shuffle/blob/bdfe9a2268e79e59e54fb7409c2a054c7b77eefb/README.md)
- [Extension control file](https://github.com/kmtr/pg_chaos_shuffle/blob/bdfe9a2268e79e59e54fb7409c2a054c7b77eefb/pg_chaos_shuffle.control)
- [Build manifest](https://github.com/kmtr/pg_chaos_shuffle/blob/bdfe9a2268e79e59e54fb7409c2a054c7b77eefb/Cargo.toml)

`pg_chaos_shuffle` Chaos-testing hook that randomizes SELECT results lacking ORDER BY.

### Enablement

Merge `pg_chaos_shuffle` into the existing preload list, restart PostgreSQL, and then create `pg_chaos_shuffle` in each database that needs its SQL objects:

```ini
shared_preload_libraries = 'pg_chaos_shuffle'
```

```sql
CREATE EXTENSION pg_chaos_shuffle;
```

### Core Workflow

The following example is taken from the reviewed upstream documentation. Adapt object names, paths, credentials, and workload values before use.

```sql
LOAD 'pg_chaos_shuffle';

CREATE TEMP TABLE users AS
SELECT g AS id FROM generate_series(1, 100) AS g;

SELECT * FROM users LIMIT 5;
SELECT * FROM users ORDER BY id LIMIT 5;
```

### Main Objects

The official sources define the extension surface dynamically or through provider tooling; inspect the installed version before granting access.

### Operations and Boundaries

- The verified PostgreSQL-major evidence covers 13, 14, 15, 16, 17, 18; do not infer unlisted majors.
- Preloading `pg_chaos_shuffle` changes cluster startup state; stage configuration and restart changes separately from `CREATE EXTENSION`.
