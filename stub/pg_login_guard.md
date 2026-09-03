## Usage

Sources:

- [Official documentation](https://github.com/mbpcore/pg_login_guard/blob/6f69b9eff49b3d03bc365314884b3fd4e182dd19/README.md)
- [Extension control file](https://github.com/mbpcore/pg_login_guard/blob/6f69b9eff49b3d03bc365314884b3fd4e182dd19/pg_login_guard.control)
- [Official repository](https://github.com/mbpcore/pg_login_guard)

`pg_login_guard` Fail2ban-style role lockout after repeated authentication failures.

### Enablement

Merge `pg_login_guard` into the existing preload list, restart PostgreSQL, and then create `pg_login_guard` in each database that needs its SQL objects:

```ini
shared_preload_libraries = 'pg_login_guard'
```

```sql
CREATE EXTENSION pg_login_guard;
```

### Core Workflow

The following example is taken from the reviewed upstream documentation. Adapt object names, paths, credentials, and workload values before use.

```sql
-- See everyone currently being tracked (has a recent failure, or is locked):
SELECT * FROM pg_login_guard_status();

--  role_name | failed_attempts | window_started_at      | locked_until
-- -----------+------------------+-------------------------+-------------------------
--  alice     |                3 | 2026-08-15 10:00:01+00  |
--  bob       |                5 | 2026-08-15 10:01:40+00  | 2026-08-15 10:16:40+00

-- Manually unlock a role before its lockout expires (superuser only):
SELECT pg_login_guard_unlock('bob');

-- Forget all tracking history for a role (superuser only):
SELECT pg_login_guard_reset('bob');
```

### Main Objects

The reviewed install surface includes these prominent objects:

| Object | Kind | Operational role |
|---|---|---|
| `pg_login_guard_status` | FUNCTION | Callable function from the reviewed install surface. |
| `pg_login_guard_reset` | FUNCTION | Callable function from the reviewed install surface. |
| `pg_login_guard_unlock` | FUNCTION | Callable function from the reviewed install surface. |

### Operations and Boundaries

- The verified PostgreSQL-major evidence covers 16, 17, 18; do not infer unlisted majors.
- Preloading `pg_login_guard` changes cluster startup state; stage configuration and restart changes separately from `CREATE EXTENSION`.
