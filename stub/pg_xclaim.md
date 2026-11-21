## Usage

Sources:

- [rc5 release](https://github.com/Fenoman/pg_xclaim/releases/tag/v1.0.0-rc5)

- [Official source](https://github.com/Fenoman/pg_xclaim/blob/v1.0.0-rc5/README.md)
- [Official source](https://github.com/Fenoman/pg_xclaim/blob/v1.0.0-rc5/pg_xclaim.control)
- [Official source](https://github.com/Fenoman/pg_xclaim/blob/v1.0.0-rc5/CHANGELOG.md)
- [Official source](https://github.com/Fenoman/pg_xclaim/blob/v1.0.0-rc5/sql/pg_xclaim--1.0.0-rc1--1.0.0-rc2.sql)
- [Official source](https://github.com/Fenoman/pg_xclaim/blob/v1.0.0-rc5/sql/pg_xclaim--1.0.0-rc4--1.0.0-rc5.sql)

`pg_xclaim` 1.0.0-rc5 is an experimental shared-memory claim mechanism for PostgreSQL 16–18. It supplies nonblocking, transaction-lifetime mutual exclusion in its own namespace. It is not a durable queue, and its keys do not coordinate with PostgreSQL advisory locks.

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

`xclaim.stats` and `xclaim.debug_snapshot` are monitoring functions with role restrictions; detailed debug functions require elevated privileges. The stale-claim injector is removed from the production library and schema in rc5. `xclaim.session_reset` is exceptional recovery, not a routine pool-reset command. Bound shared capacity and per-backend claims using the documented configuration limits; disabling the extension can remove mutual exclusion.

### rc5 Upgrade

rc5 pre-stages local claim entries before taking partition locks, fixes allocation-failure cleanup, and removes the production stale-entry injector. Install matching files, restart the server and run `ALTER EXTENSION pg_xclaim UPDATE TO '1.0.0-rc5'` in each database. The supplied upgrade chain covers rc1 through rc5. Exact SQL version strings retain the release-candidate suffix. This is experimental software; test cleanup, cancellation and capacity behavior before adoption.
