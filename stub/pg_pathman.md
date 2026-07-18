## Usage

Sources:

- [Project README and compatibility matrix](https://github.com/postgrespro/pg_pathman/blob/09c0afe25154a034f5ae9cf5b15f977ce8a7a1e9/README.md)
- [Extension control file](https://github.com/postgrespro/pg_pathman/blob/09c0afe25154a034f5ae9cf5b15f977ce8a7a1e9/pg_pathman.control)
- [Hook implementation](https://github.com/postgrespro/pg_pathman/blob/09c0afe25154a034f5ae9cf5b15f977ce8a7a1e9/src/hooks.c)
- [Partition functions](https://github.com/postgrespro/pg_pathman/blob/09c0afe25154a034f5ae9cf5b15f977ce8a7a1e9/src/pl_funcs.c)

`pg_pathman` 1.5 provides hash and range partition management, row routing, runtime partition selection, and data-migration workers. Upstream states that the project is no longer under development and recommends PostgreSQL native partitioning. Its documented ceiling is PostgreSQL 15, with patched core builds required for several versions; use native partitioning for every new design.

### Load an existing deployment

The library must be preloaded and PostgreSQL restarted:

```conf
shared_preload_libraries = 'pg_pathman'
```

Install it in a dedicated schema where untrusted roles cannot create objects:

```sql
CREATE SCHEMA pathman;
REVOKE CREATE ON SCHEMA pathman FROM PUBLIC;
GRANT USAGE ON SCHEMA pathman TO PUBLIC;
CREATE EXTENSION pg_pathman WITH SCHEMA pathman;

SELECT pathman.pathman_version();
```

The clean schema matters because calls that do not exactly match a function signature may otherwise be exposed to malicious overloads. Review the extension's schema privileges and callers' `search_path` settings.

### Migration and retirement risks

Functions such as `create_hash_partitions()` and `create_range_partitions()` change table structure and, with `partition_data` enabled, copy rows while holding a lock until commit. `partition_table_concurrently()` moves batches in short transactions but still changes live data and requires operational monitoring. Rehearse constraints, indexes, triggers, sequences, foreign keys, replication, rollback, and application behavior on a restored copy before any conversion.

The module uses planner, executor, utility, and shared-memory hooks. Upstream specifically warns that `ProcessUtility_hook` ordering may interfere with extensions such as `pg_stat_statements`; preload order is not a substitute for compatibility testing. Do not load this build on an unlisted PostgreSQL major. For an existing deployment, plan a measured migration to declarative partitioning before the operating system or PostgreSQL upgrade removes the compatible build environment.
