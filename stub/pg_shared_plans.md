## Usage

Sources:

- [Official README](https://github.com/rjuju/pg_shared_plans/blob/master/README.md)
- [Extension SQL and privilege definitions](https://github.com/rjuju/pg_shared_plans/blob/master/pg_shared_plans--0.0.1.sql)
- [Preload and configuration implementation](https://github.com/rjuju/pg_shared_plans/blob/master/pg_shared_plans.c)

`pg_shared_plans` 0.0.1 is an explicit proof of concept for transparently caching execution plans in shared memory. It can reduce repeated planning work for eligible statements, but upstream says it is not production-ready and cannot reliably invalidate plans after DDL on a standby.

### Preload and Setup

The module supports PostgreSQL 12 and later, requires `pg_stat_statements`, and must be preloaded before the extension objects are created:

```conf
shared_preload_libraries = 'pg_stat_statements,pg_shared_plans'
```

```sql
CREATE EXTENSION pg_shared_plans CASCADE;

SELECT * FROM pg_shared_plans_info;
SELECT * FROM pg_shared_plans;
```

A postmaster restart is required after changing the preload list. PostgreSQL 14 and later also need query identifier computation enabled; the module calls the server hook that enables it when the core setting permits.

### Controls and Views

Important controls include `pg_shared_plans.enabled`, `pg_shared_plans.max`, `pg_shared_plans.min_plan_time`, `pg_shared_plans.threshold`, `pg_shared_plans.rdepend_max`, and `pg_shared_plans.read_only`. The current C default for `pg_shared_plans.max` is 1000 even though the README still says 200.

Use `pg_shared_plans` for summary data, `pg_shared_plans_relations` for relation OIDs, `pg_shared_plans_explain` for plan text, and `pg_shared_plans_all` for both. `pg_shared_plans_reset(userid, dbid, queryid)` removes matching cached entries and is revoked from PUBLIC by the install script.

### Security and Operational Notes

Cached query and plan text can expose schema names, relation details, and predicates. The low-level `pg_shared_plans(boolean, boolean, oid, oid)` function keeps default PUBLIC execution even though its curated views have narrower grants; revoke it and grant only the required monitoring views to trusted roles. Shared-memory contents disappear at restart. Treat DDL invalidation, RLS identity handling, prepared-statement behavior, eviction, memory sizing, and failover as experimental surfaces, and never enable this POC on a standby or production cluster without source-level validation.
