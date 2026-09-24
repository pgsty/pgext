## Usage

Sources:

- [PostgreSQL 19 documentation](https://www.postgresql.org/docs/19/pgplanadvice.html)
- [Module definition](https://github.com/postgres/postgres/blob/REL_19_BETA3/contrib/pg_plan_advice/Makefile)
- [Module implementation](https://github.com/postgres/postgres/blob/REL_19_BETA3/contrib/pg_plan_advice/pg_plan_advice.c)
- [Versioned manual](https://github.com/postgres/postgres/blob/REL_19_BETA3/doc/src/sgml/pgplanadvice.sgml)

`pg_plan_advice` is a supplied PostgreSQL 19 module that records planner choices as advice and constrains later planning. This page uses PostgreSQL 19 beta3 as its version boundary; the module has no independent SQL extension version.

### Session Workflow

```sql
LOAD 'pg_plan_advice';
CREATE TEMP TABLE advice_demo (id integer PRIMARY KEY, body text);
EXPLAIN (COSTS OFF, PLAN_ADVICE)
  SELECT * FROM advice_demo d WHERE id = 1;
SET pg_plan_advice.advice = 'SEQ_SCAN(d)';
EXPLAIN (COSTS OFF, PLAN_ADVICE)
  SELECT * FROM advice_demo d WHERE id = 1;
RESET pg_plan_advice.advice;
```

### Enablement and Feedback

Use `LOAD` with sufficient privileges for one session, `session_preload_libraries` for new sessions, or `shared_preload_libraries` followed by a server restart. There is no extension DDL step. `EXPLAIN (PLAN_ADVICE)` generates advice using relation aliases as targets. `pg_plan_advice.advice` supplies constraints on scans, join order, join methods, or parallelism.

Inspect the returned feedback for unmatched, conflicting, inapplicable, or failed advice. `pg_plan_advice.feedback_warnings` can emit warnings, and `pg_plan_advice.always_store_advice_details` preserves details useful when explaining prepared statements at additional planning cost.

### Limits

Advice only constrains plans the core planner considers viable. It cannot force a semantically invalid plan, control aggregation strategy, or control set operations. Fixed advice can age poorly as data changes. `pg_stash_advice` optionally stores advice by query ID; `pg_plan_guard` monitors drift from approved baselines.
