## Usage

Sources:

- [README](https://github.com/s-hironobu/pg_query_plan/blob/d0e7683069e23630d2e543674469bea4b1d16cde/README.md)
- [C implementation](https://github.com/s-hironobu/pg_query_plan/blob/d0e7683069e23630d2e543674469bea4b1d16cde/pg_query_plan.c)
- [SQL installation](https://github.com/s-hironobu/pg_query_plan/blob/d0e7683069e23630d2e543674469bea4b1d16cde/pg_query_plan--1.0.sql)
- [Makefile](https://github.com/s-hironobu/pg_query_plan/blob/d0e7683069e23630d2e543674469bea4b1d16cde/Makefile)

`pg_query_plan` is a PostgreSQL 18 module for inspecting another backend's running plan and actual row counts. The reviewed source ships SQL version 1.0, with no control file or CREATE EXTENSION step. Its README labels a future release date; this entry describes the current source preview.

### Activate and Inspect

Set shared_preload_libraries to include pg_query_plan and restart PostgreSQL. Then run the supplied pg_query_plan--1.0.sql once in each database where the inspection function is needed. Preserve existing preload entries. The SQL script drops and recreates public.pg_query_plan(integer), so review any existing function and dependencies before running it.

```sql
SELECT pid FROM pg_stat_activity WHERE pid <> pg_backend_pid();
SELECT * FROM public.pg_query_plan(pg_backend_pid());
```

Use a PID selected from pg_stat_activity as the argument when inspecting a different backend. The caller must be a superuser or a member of the target role. Results include database, leader/worker identity, query ID, SQL text and the live plan. An idle target returns no rows.

### Operational Limits

The module enables row instrumentation on every query and adds overhead. It uses SIGUSR2 requests serviced at executor safe points; a long operation inside one plan node can delay a response, with a five-second timeout. The optional query_progress.py client needs Python and psycopg2, and its progress estimate depends on planner estimates. There is no background worker or automatic retention. Current license terms were not established from the repository.
