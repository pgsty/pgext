## Usage

Sources:

- [Official source](https://github.com/qsliu2017/pg_jevplanner/blob/5e888343e7b3d3730d13b0d473c8dea9af542377/README.md)
- [Official source](https://github.com/qsliu2017/pg_jevplanner/blob/5e888343e7b3d3730d13b0d473c8dea9af542377/benchmark/README.md)
- [Official source](https://github.com/qsliu2017/pg_jevplanner/blob/5e888343e7b3d3730d13b0d473c8dea9af542377/src/pg_jevplanner.c)
- [Official source](https://github.com/qsliu2017/pg_jevplanner/blob/5e888343e7b3d3730d13b0d473c8dea9af542377/sql/pg_jevplanner--0.1.0.sql)

`pg_jevplanner` 0.1.0 is an experimental PostgreSQL 18 join-ordering extension. It asks the external TypeSafe JEV service to choose legal inner-join merges, while PostgreSQL retains physical planning. Server headers for other majors are rejected by the source.

### Enable and Compare

Configure the server process with a valid `TYPESAFE_API_KEY` before starting it. Install as a superuser and load the library in each backend; creating the extension loads only the creating session. Administrator-configured session preloading can cover later sessions.

```sql
CREATE EXTENSION pg_jevplanner;
LOAD 'pg_jevplanner';
SELECT jevplanner_version();
SET jev.enabled = on;
SET max_parallel_workers_per_gather = 0;
EXPLAIN SELECT a.id FROM a JOIN b ON a.id=b.id JOIN c ON b.id=c.id;
SET jev.enabled = off;
```

The query example requires existing joinable tables. Compare actual plans and results with the extension disabled; benchmark execution time alone excludes planning and external API latency. This is an opt-in research planner, and unsupported SQL barriers are preserved.

### Configuration and Data Exposure

Planner controls require superuser authority. `jev.timeout_ms` defaults to 2000 milliseconds per request; `jev.planning_timeout_ms` defaults to 30000 for planning as a whole. `jev.max_decisions` and `jev.max_candidates` bound decision work; exceeding the candidate bound errors rather than silently shortlisting candidates.

Requests disclose query SQL, names, literals, relation statistics and join metadata to the external provider. `jev.log_http` defaults off; turning it on exposes request and response bodies in notices. Keep production secrets out of SQL text and logs, restrict credentials to the server environment, and assess provider access before enabling this hook. Replanning prepared statements can make further model requests.
