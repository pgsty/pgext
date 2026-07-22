## Usage

Sources:

- [Extension README](https://github.com/KBVE/kbve/blob/6d2eecf849454b8a6396bb2600d490da718a21c8/apps/kbve/kilobase/README.md)
- [Background-worker implementation](https://github.com/KBVE/kbve/blob/6d2eecf849454b8a6396bb2600d490da718a21c8/apps/kbve/kilobase/src/lib.rs)
- [SQL object definitions](https://github.com/KBVE/kbve/blob/6d2eecf849454b8a6396bb2600d490da718a21c8/apps/kbve/kilobase/src/sql.rs)
- [Extension control file](https://github.com/KBVE/kbve/blob/6d2eecf849454b8a6396bb2600d490da718a21c8/apps/kbve/kilobase/kilobase.control)

`kilobase` 17.4.1 is a pgrx background worker that schedules materialized-view refreshes. The current implementation is an experimental single-database scheduler: the worker is hard-coded to connect to the `postgres` database, so its extension objects and jobs must be created there.

### Preload and Database Setup

Add the library to `shared_preload_libraries` and restart PostgreSQL; loading it only through `CREATE EXTENSION` deliberately skips worker registration.

```conf
shared_preload_libraries = 'kilobase'
```

Connect to the `postgres` database and install the extension. Its SQL uses unqualified object names, so keep the extension in the default `public` schema unless you have audited and patched every worker query.

```sql
CREATE EXTENSION kilobase;
```

Installation has an unusual side effect: if the database has no materialized view, the install script creates and registers a demonstration view named `sample_stats`. Review and remove that example and its job if it is not wanted.

### Register a Refresh Job

Create a materialized view. A suitable unique index lets the worker attempt `REFRESH MATERIALIZED VIEW CONCURRENTLY`; without one, it falls back to a blocking regular refresh.

```sql
CREATE TABLE orders (
    id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    customer_id bigint NOT NULL,
    amount numeric NOT NULL
);

CREATE MATERIALIZED VIEW order_totals AS
SELECT customer_id, sum(amount) AS total
FROM orders
GROUP BY customer_id;

CREATE UNIQUE INDEX order_totals_customer_idx
    ON order_totals (customer_id);

SELECT register_matview_refresh('public', 'order_totals', 300, 'orders');

SELECT * FROM matview_refresh_status;
SELECT * FROM matview_refresh_history;
```

The optional fourth argument names one source table in the same schema. When supplied, the worker compares the cumulative insert/update/delete counters in `pg_stat_user_tables` and skips a refresh if the counter did not change. These statistics can reset and do not describe all dependencies, so omit the argument when correctness requires every scheduled refresh.

### Management Objects

- `register_matview_refresh(text, text, integer, text)`: create or reactivate a job and return its integer ID.
- `unregister_matview_refresh(text, text)`: mark a job inactive.
- `matview_refresh_jobs`: scheduler state; active due jobs are processed in batches of at most ten.
- `matview_refresh_log`: refresh outcomes and durations.
- `matview_refresh_status` and `matview_refresh_history`: operational views.
- `cleanup_matview_refresh_logs(integer)`: delete history older than the requested number of days; the worker also runs seven-day cleanup periodically.
- `kilobase_health_check()`: summarize worker visibility, due jobs, failures, durations, and views without a unique index.

### Operational Boundaries

One worker serializes the refresh workload in the `postgres` database. Concurrent refresh can still fail and then be retried as a blocking refresh; a slow refresh delays later jobs. Test lock impact, worker restart behavior, job permissions, and recovery after errors before relying on the schedule. The extension creates ordinary tables, views, functions, and a trigger without hardening their privileges, so restrict modification and execution rights to trusted administrators.
