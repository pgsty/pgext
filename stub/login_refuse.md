## Usage

Sources:

- [Pinned upstream README](https://github.com/liang7878/login_refuse/blob/f2af59137c10e786b3abf9646fa11fcd8bf3f8d8/README.md)
- [Version 0.1 installation SQL](https://github.com/liang7878/login_refuse/blob/f2af59137c10e786b3abf9646fa11fcd8bf3f8d8/login_refuse--0.1.sql)
- [Pinned authentication-hook implementation](https://github.com/liang7878/login_refuse/blob/f2af59137c10e786b3abf9646fa11fcd8bf3f8d8/login_refuse.c)
- [Pinned extension control file](https://github.com/liang7878/login_refuse/blob/f2af59137c10e786b3abf9646fa11fcd8bf3f8d8/login_refuse.control)

login_refuse version 0.1 installs a client-authentication hook. For password, MD5, and SCRAM authentication it records failed attempts by username and refuses later connections for a configured interval after the threshold is reached. It can also mark a role as expired after a Unix timestamp.

### Configuration

The hook must be preloaded. Both numeric settings default to zero, so set explicit positive values before restarting PostgreSQL:

```conf
shared_preload_libraries = 'login_refuse'
login_refuse.minutes = 10
login_refuse.threshold = 5
```

```sql
CREATE EXTENSION login_refuse;

SELECT login_refuse_set_expire_time(
    'temporary_user',
    extract(epoch FROM timestamptz '2030-01-01 00:00:00+00')::bigint
);

SELECT login_refuse_reset_expire_time('temporary_user');
```

The two SQL functions check for superuser at runtime. The timestamp function makes the account fail this hook after the specified instant until reset; it is not a temporary lock-until timestamp.

### Authentication safety

Failure and expiry state are stored in two plain files directly under PGDATA. Every authenticating backend reads, rewrites, and appends these files without inter-process locking or durable transactional semantics. Concurrent attempts can lose or corrupt state. Usernames and lines are copied through small fixed C buffers, and several file and memory error paths are unsafe.

The project has only a one-line README, no test matrix, and no updates since 2018. Authentication hooks and internal enums are version-sensitive. A bug can lock out administrators or crash a connection process. Do not use login_refuse as the primary brute-force control; prefer a maintained proxy, identity provider, firewall, or PostgreSQL-supported authentication policy. If evaluating it, retain an independent local recovery path, restrict filesystem permissions, test concurrent failures and restart behavior, and monitor the state files.
