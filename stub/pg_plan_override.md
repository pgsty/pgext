## Usage

Sources:

- [Official source](https://github.com/aaaristo/pg_plan_override/blob/e0215040e42047f957d0e301ec40600c2b61ce61/README.md)
- [Official source](https://github.com/aaaristo/pg_plan_override/blob/e0215040e42047f957d0e301ec40600c2b61ce61/src/pg_plan_override.control)
- [Official source](https://github.com/aaaristo/pg_plan_override/blob/e0215040e42047f957d0e301ec40600c2b61ce61/src/pg_plan_override--1.0.sql)
- [Official source](https://github.com/aaaristo/pg_plan_override/blob/e0215040e42047f957d0e301ec40600c2b61ce61/src/pg_plan_override.c)

`pg_plan_override` 1.0 applies per-query planner settings matched by query ID or SQL text pattern. Upstream claims PostgreSQL 12 and later, tests PostgreSQL 12, and includes compatibility guards for the newer planner-hook signature; later majors have not been independently validated here.

### Enable and Add a Rule

An administrator configures preloading and restarts PostgreSQL:

```ini
shared_preload_libraries = 'pg_plan_override'
```

Install as a superuser and create a narrowly matched rule:

```sql
CREATE EXTENSION pg_plan_override;
SELECT plan_override.add_by_pattern(
    '%slow_report%',
    '{"enable_seqscan": "off"}'::jsonb,
    'Review report plan'
);
SELECT plan_override.refresh_cache();
SELECT * FROM plan_override.override_rules;
```

`plan_override.add_by_query_id` matches a numeric fingerprint instead. The highest-priority enabled rule wins. Query-ID availability depends on the PostgreSQL version and fingerprinting configuration; upstream notes the statement-statistics extension for older majors and native computation for newer majors.

### Scope and Cache

The hook temporarily changes settings while planning and restores them afterwards, including on errors. It does not retain those settings during execution. In particular, changing an execution-time resource setting through a rule does not give the executing query that resource budget.

`pg_plan_override.enabled` is the master switch; `pg_plan_override.debug` logs applied overrides. `pg_plan_override.cache_ttl` defaults to 60 seconds. Each backend has its own cache: refreshing one connection does not refresh others, and cached prepared plans need separate replanning when testing a new rule.

### Permissions and Maintenance

The installation grants `PUBLIC` read access to the rules table and its schema. Patterns and descriptions therefore must not contain secrets. Restrict rule modification to trusted administrators; invalid or overly broad rules can worsen plans, and pattern-check cost grows with the rule count. Validate with actual execution plans and disable or remove individual rules when they no longer help. The fixed schema is `plan_override`; this extension is not relocatable.
