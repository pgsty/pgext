## Usage

Sources:

- [Upstream README](https://github.com/harukat/pg_rlimit/blob/51aec672a347f92fe3e08ed2be75c8ee6b3427e2/README.md)
- [Extension SQL](https://github.com/harukat/pg_rlimit/blob/51aec672a347f92fe3e08ed2be75c8ee6b3427e2/pg_rlimit--1.0.sql)
- [Resource-limit implementation](https://github.com/harukat/pg_rlimit/blob/51aec672a347f92fe3e08ed2be75c8ee6b3427e2/pg_rlimit.c)

`pg_rlimit` version `1.0` exposes the POSIX address-space soft limit for each PostgreSQL backend. It can cap a session's virtual address space, but it is not a per-query memory quota.

### Core Workflow

Preload the library to enable the startup hook and `pg_rlimit.v` GUC, then restart PostgreSQL.

```conf
shared_preload_libraries = 'pg_rlimit'
pg_rlimit.v = '1GB'
```

```sql
CREATE EXTENSION pg_rlimit;
SELECT pg_getrlimit('v');
SELECT pg_setrlimit('v', 1073741824);
SET pg_rlimit.v = '768MB';
```

The only supported resource selector is `v`, corresponding to `RLIMIT_AS`. `pg_getrlimit` returns the current soft limit in bytes, with `-1` representing infinity. `pg_setrlimit` sets the current backend's soft limit and returns the resulting value. A backend cannot raise its soft limit above the operating-system hard limit.

### Operational Caveats

The GUC is measured in kilobytes and accepts PostgreSQL memory units. The client-authentication hook applies its configured value when a session starts; later `SET` changes the current backend.

Address space includes mapped regions and PostgreSQL shared memory, not just private allocations or `work_mem`. A low value can produce out-of-memory failures in unrelated operations. Test authentication, connection pooling, parallel workers, large mappings, and error recovery on the target OS. The SQL functions are not declared superuser-only upstream, so manage EXECUTE privileges if ordinary sessions must not alter their limits.
