## Usage

Sources:

- [README.md](https://github.com/adept/pg_relusage/blob/017ca90bc08aa131d074017511ea9ed6ab89d6d3/README.md)
- [Control file](https://github.com/adept/pg_relusage/blob/017ca90bc08aa131d074017511ea9ed6ab89d6d3/pg_relusage.control)
- [Implementation](https://github.com/adept/pg_relusage/blob/017ca90bc08aa131d074017511ea9ed6ab89d6d3/pg_relusage.c)

`pg_relusage` 1.0.2 logs relations used by statements after plan simplification. It supports PostgreSQL 10–18 and defines no SQL objects. Loading the library enables it; optional CREATE EXTENSION only records installation.

### Core Workflow

```sql
LOAD 'pg_relusage';
SET pg_relusage.nested_statements = 'fold';
SET pg_relusage.log_level = 'notice';
CREATE TABLE relusage_demo (id integer);
SELECT * FROM relusage_demo;
```

### Operational Boundaries

All settings require superuser: `pg_relusage.enabled` defaults on, `pg_relusage.log_level` defaults log and `pg_relusage.rel_kinds` defaults riSvmfp. `pg_relusage.nested_statements` accepts `all` (default: each statement), `top` (ignore nested statements) or `fold` (merge nested relations into the outer log). Version 1.0.1 fixes fold handling after a failed subtransaction.

Use `session_preload_libraries` for new sessions without a restart, or `shared_preload_libraries` for all sessions with a restart; preserve existing preload entries. LOAD normally requires superuser. Log messages omit SQL text, so configure `log_line_prefix` to identify the client.

System catalogs are excluded. Relations eliminated by planning or startup pruning are excluded, but partitions skipped later during execution can still appear. INSERT into a partitioned table reports only its parent. Deferred triggers at COMMIT, COPY FROM triggers and planner-invoked functions are treated as top-level statements. DDL and maintenance commands are excluded. A quiet log alone does not prove an object can be dropped.

### 1.0.2 and Partitioned Tables

Version 1.0.2 reduces relation-membership and partition-matching overhead for statements involving many partitions. Log format and settings are unchanged. Replace the library and restart if using shared preload, or reconnect when using session preload or LOAD. If installation was recorded with CREATE EXTENSION, also run `ALTER EXTENSION pg_relusage UPDATE`.
