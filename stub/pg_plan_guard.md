## Usage

Sources:

- [README.md](https://api.pgxn.org/src/pg_plan_guard/pg_plan_guard-1.1.9/README.md)
- [pg_plan_guard.control](https://api.pgxn.org/src/pg_plan_guard/pg_plan_guard-1.1.9/pg_plan_guard.control)
- [CHANGELOG.md](https://api.pgxn.org/src/pg_plan_guard/pg_plan_guard-1.1.9/CHANGELOG.md)
- [pg_plan_guard--1.1.8--1.1.9.sql](https://api.pgxn.org/src/pg_plan_guard/pg_plan_guard-1.1.9/pg_plan_guard--1.1.8--1.1.9.sql)

`pg_plan_guard` 1.1.9 records approved PostgreSQL 19 plan-advice text and detects changes when stored queries are replanned. It requires the loaded `pg_plan_advice` module and does not support PostgreSQL 18.

### Core Workflow

```sql
LOAD 'pg_plan_advice';
CREATE EXTENSION pg_plan_guard;

SELECT plan_guard.capture(
    'semantic_search',
    'SELECT id FROM docs ORDER BY embedding <=> ''[1,2,3]'' LIMIT 10',
    'must use the vector index'
);

SELECT * FROM plan_guard.verify();
SELECT * FROM plan_guard.status WHERE state <> 'ok';
```

### Operational Boundaries

`plan_guard.capture()` records a query, its advice, author and search path. `plan_guard.verify()` replans it, and `plan_guard.status` exposes state; `plan_guard.drift_log` records transitions and rejects modification. Planning happens read-only in a rolled-back subtransaction as the captured author, so the checking role must be allowed to assume that role. A baseline from before 1.1.7 has no author and must be captured again; otherwise it returns an error. The installer rejects a pre-created schema owned by an unrelated unprivileged role. Planning can invoke functions, so stored SQL must still be trusted and the baseline reviewed. `plan_guard.watch()` optionally uses `pg_living_assertions`; ordinary capture/verification do not require it. `plan_guard.sync_stash()` additionally needs `pg_stash_advice` preload, restart and a configured stash. Current Apache-2.0 licensing starts at 1.1.3.

### Role Isolation in 1.1.9

Version 1.1.9 plans stored SQL inside a temporary author-owned `SECURITY DEFINER` function when role isolation is needed. A function evaluated by the planner can no longer regain the runner's identity by resetting role or session authorization. The function is created and rolled back inside the read-only planning subtransaction. The author needs database `TEMP` permission when this frame is required; the runner needs authority to act as the author. Missing privileges produce a baseline error. This fixes the role-reset escape in 1.1.7/1.1.8 without turning stored SQL into untrusted input.
