## Usage

Sources:

- [Official source](https://github.com/Fenoman/pg_xclaim/blob/49e5bc3982de87bf94b3108fa7ce29ae813a07b2/README.md)
- [Official source](https://github.com/Fenoman/pg_xclaim/blob/49e5bc3982de87bf94b3108fa7ce29ae813a07b2/pg_xclaim.control)
- [Official source](https://github.com/Fenoman/pg_xclaim/blob/49e5bc3982de87bf94b3108fa7ce29ae813a07b2/CHANGELOG.md)
- [Official source](https://github.com/Fenoman/pg_xclaim/blob/49e5bc3982de87bf94b3108fa7ce29ae813a07b2/sql/pg_xclaim--1.0.0-rc1--1.0.0-rc2.sql)
- [Official source](https://github.com/Fenoman/pg_xclaim/blob/49e5bc3982de87bf94b3108fa7ce29ae813a07b2/sql/pg_xclaim--1.0.0-rc2--1.0.0-rc3.sql)

`pg_xclaim` 1.0.0-rc3 is an experimental shared-memory claim mechanism for PostgreSQL 16–18. It supplies nonblocking, transaction-lifetime mutual exclusion in its own namespace. It is not a durable queue, and its keys do not coordinate with PostgreSQL advisory locks.

### Core Workflow

Add the library to the existing preload list, preserve other entries, and restart PostgreSQL before creating the extension as a superuser. Keep it preloaded on replicas: calls fail during recovery, but work after promotion.

```ini
shared_preload_libraries = 'pg_xclaim'
```

```sql
CREATE EXTENSION pg_xclaim;
BEGIN;
SELECT xclaim.try(1, 100);
SELECT xclaim.count();
COMMIT;
```

A true `xclaim.try` result means the current top-level transaction holds the claim. Repeated acquisition by the same owner succeeds; top-level commit or abort releases it. Claims survive rollback to a savepoint, unlike advisory transaction locks. Migrate all participating call sites together to avoid mixing disjoint lock namespaces.

### Objects and Limits

`xclaim.try` accepts either one bigint key or two int4 keys in separate namespaces. `xclaim.try_many` returns a boolean array in input order; partial acquisition is retained until transaction end. Roll back the entire transaction if all-or-nothing behavior is required. NULL inputs raise errors.

`xclaim.stats` and `xclaim.debug_snapshot` are monitoring functions with role restrictions; detailed debug functions and stale-claim injection are superuser tools, and injection is for tests only. `xclaim.session_reset` is exceptional recovery, not a routine pool-reset command. Bound shared capacity and per-backend claims using the documented configuration limits; disabling the extension can remove mutual exclusion.

### rc3 Upgrade

rc2 repairs backend-exit cleanup and allocation-failure handling; rc3 adds Tantor SE 18 build compatibility without new SQL objects. Install matching files, restart the server and run `ALTER EXTENSION pg_xclaim UPDATE TO '1.0.0-rc3'` in each database. This remains a release candidate; test failure and capacity behavior before adopting it.
