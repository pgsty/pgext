
## Usage

> [pg_explain_ui: send EXPLAIN plans to explain.dalibo.com](https://github.com/davidgomes/pg-explain-ui)
> Blog: [Writing a Postgres Extension With Pgrx for Visual Query Plans](https://davidgomes.com/writing-postgres-extension-with-pgrx-query-plans/)

This function will explain query and post it to the [`explain.dalibo.com`](https://explain.dalibo.com/) to visualize the plan

> WARNING: http post will be sent to the internet service!

```sql
SELECT explain_ui($$query$$);
```


### Example

Let's generate pgbench sample tables

```bash
pgbench -is10
```

Then explain the query with:

```sql
CREATE EXTENSION explain_ui;
SELECT explain_ui($$SELECT * FROM pgbench_accounts;$$);
```

You'll get an URL for the visual plan:

```
postgres@u22:5432/postgres=# SELECT explain_ui($$SELECT * FROM pgbench_accounts;$$);
                    explain_ui
--------------------------------------------------
 https://explain.dalibo.com/plan/05377227a29f0418
(1 row)

Time: 2284.667 ms (00:02.285)
```

Click or open that [URL]( https://explain.dalibo.com/plan/05377227a29f0418) with your browser. 
