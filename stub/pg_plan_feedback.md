## Usage

Sources:

- [Official source](https://github.com/pganalyze/pg_plan_feedback/blob/26769778708c9e49cef1ea3f4c3e48cdc40b6f3f/README.md)
- [Official source](https://github.com/pganalyze/pg_plan_feedback/blob/26769778708c9e49cef1ea3f4c3e48cdc40b6f3f/pg_plan_feedback.control)
- [Official source](https://github.com/pganalyze/pg_plan_feedback/blob/26769778708c9e49cef1ea3f4c3e48cdc40b6f3f/pg_plan_feedback--1.0.sql)
- [Official source](https://github.com/lfittl/postgres/pull/82)

`pg_plan_feedback` learns per-query, per-relation row-count corrections from execution and can feed them into later planning. It requires an upstream PostgreSQL core patch; stock PostgreSQL compatibility is not established. Upstream marks this experimental and unsuitable for production.

### Observe Before Applying

On a server with the required core hooks and installed library:

```sql
LOAD 'pg_plan_feedback';
CREATE EXTENSION pg_plan_feedback;
SET pg_plan_feedback.mode = 'observe';
SELECT count(*) FROM pg_class;
SELECT pg_plan_feedback_flush();
SELECT * FROM pg_plan_feedback;
```

The viewing role must have `pg_read_all_stats` or equivalent explicitly granted privileges. `observe` records feedback without changing estimates; `apply` uses learned corrections; `off` disables the feature. Exercise representative queries and inspect feedback before switching modes.

### Objects and Limits

`pg_plan_feedback_reset(queryid)` clears corrections; omitting its argument resets all. `pg_plan_feedback_set(queryid, relation, correction)` installs a manual factor. Both are restricted by default because they affect planning for other users. `pg_plan_feedback_stats()` reports counts, and the view exposes estimates, observed rows and correction factors across databases.

The library may be loaded through shared or session preload, or explicit `LOAD`; shared preload is not mandatory for its dynamic shared memory. The SQL extension is relocatable and requires a superuser to install. It learns base-relation row counts, not a forced join order or full plan. Incomplete scans do not supply reliable feedback; cached plans require replanning before estimates change. Do not assume stable storage or interfaces across experimental revisions.
