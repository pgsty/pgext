## Usage

Sources:

- [PGXN 1.0.0 README](https://pgxn.org/dist/pg_plan_guard/1.0.0/README.html)
- [pg_plan_guard control file](https://api.pgxn.org/src/pg_plan_guard/pg_plan_guard-1.0.0/pg_plan_guard.control)
- [pg_plan_guard 1.0 SQL definitions](https://api.pgxn.org/src/pg_plan_guard/pg_plan_guard-1.0.0/pg_plan_guard--1.0.sql)
- [PostgreSQL license](https://api.pgxn.org/src/pg_plan_guard/pg_plan_guard-1.0.0/LICENSE)

`pg_plan_guard` captures an approved PostgreSQL 19 plan shape as plan-advice text and reports when replanning produces different advice. It is a monitoring layer for critical queries, not an optimizer or automatic baseline selector.

### Core Workflow

```sql
CREATE EXTENSION pg_plan_guard;

SELECT plan_guard.capture(
    'semantic_search',
    'SELECT id FROM docs ORDER BY embedding <=> ''[1,2,3]'' LIMIT 10',
    'must use the vector index'
);

SELECT * FROM plan_guard.verify();
SELECT * FROM plan_guard.status WHERE state <> 'ok';
```

`plan_guard.capture(...)` uses `EXPLAIN (PLAN_ADVICE)` without executing the query. `plan_guard.verify(...)` replans stored SQL, compares the generated advice, records a transition to drift or error, and continues when one baseline can no longer be planned.

### Objects and Optional Stash

- `plan_guard.baselines` stores approved query text, advice, human-readable plans, and current state.
- `plan_guard.drift_log` is an append-only transition history.
- `plan_guard.status` is the monitoring view.
- `plan_guard.advice_for(...)` and `plan_guard.query_id_for(...)` expose lower-level planning information.
- `plan_guard.sync_stash(...)` copies approved advice into a `pg_stash_advice` stash.

The extension revokes public execution of functions that plan stored SQL. Grant capture and verification only to a role allowed to plan every stored query and to update the baseline tables.

### Requirements and Boundaries

`pg_plan_guard` requires PostgreSQL 19 and the `pg_plan_advice` module; it does not work on PostgreSQL 18. Applying advice through `plan_guard.sync_stash(...)` additionally needs `pg_stash_advice` in `shared_preload_libraries` and a nonempty `pg_stash_advice.stash_name` setting. Verification detects drift but does not force a plan unless that optional stash path is configured.

Queries are stored and replanned as executable SQL text, so parameterized application queries need representative literal values. A baseline can bless a poor plan; review captured advice before approval. Although `EXPLAIN` does not execute the query, planning can still acquire catalog locks, invoke planning hooks, and fail after schema or privilege changes.
