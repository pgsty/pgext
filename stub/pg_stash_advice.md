## Usage

Sources:

- [PostgreSQL development documentation](https://www.postgresql.org/docs/devel/pgstashadvice.html)
- [Extension control file at the reviewed PostgreSQL commit](https://github.com/postgres/postgres/blob/3cf5264557bee2ba848e5276beecc10571d468a6/contrib/pg_stash_advice/pg_stash_advice.control)
- [SQL API and default privileges](https://github.com/postgres/postgres/blob/3cf5264557bee2ba848e5276beecc10571d468a6/contrib/pg_stash_advice/pg_stash_advice--1.0.sql)

`pg_stash_advice` stores plan-advice strings in named, dynamically allocated shared-memory stashes keyed by PostgreSQL query ID. A session selects a stash with `pg_stash_advice.stash_name`; when a matching query is planned, the saved advice is applied automatically.

### Create and Select a Stash

Preloading is the normal way to make advice application and automatic persistence available in every session.

```conf
shared_preload_libraries = 'pg_stash_advice'
pg_stash_advice.persist = on
pg_stash_advice.persist_interval = 30
```

After restarting, obtain the query ID with `EXPLAIN (VERBOSE)` and the advice string with `EXPLAIN (PLAN_ADVICE)`, then store them:

```sql
CREATE EXTENSION pg_stash_advice;

SELECT pg_create_advice_stash('critical_queries');
SELECT pg_set_stashed_advice(
  'critical_queries',
  :'query_id'::bigint,
  :'plan_advice'
);

SET pg_stash_advice.stash_name = 'critical_queries';
```

Inspection functions are `pg_get_advice_stashes()` and `pg_get_advice_stash_contents(text)`. `pg_drop_advice_stash(text)` removes a stash. The install script revokes public execution of all management and inspection functions.

### Caveats

- This page describes PostgreSQL development-branch code, and the official page labels that branch unsupported. Confirm that the exact server build includes version `1.0`; do not assume availability in a released major version.
- Stashes consume dynamic shared memory. Applying advice adds planning overhead and can prevent the planner from adapting as data distribution changes, so constrain both stash size and scope.
- Persistence is enabled by default and writes `pg_stash_advice.tsv` in the data directory through a background worker. Automatic worker startup requires preloading; `pg_start_stash_advice_worker()` can start persistence manually but does not load an existing file.
- The security model is intentionally limited. Although management functions are not public by default, any user may select a known stash name, which may reveal its advice strings. Do not store secrets in plan advice.
- Query IDs require `compute_query_id`; loading the module enables computation when that setting is `auto`. Advice tied to obsolete query IDs must be maintained as queries and schemas change.
