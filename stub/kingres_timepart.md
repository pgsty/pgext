## Usage

Sources:

- [Official source](https://github.com/KingresDB/Kingres/blob/3c98d1cd4a8d81e2fee4d859da0951b65ac44e72/contrib/kingres_timepart/README.md)
- [Official source](https://github.com/KingresDB/Kingres/blob/3c98d1cd4a8d81e2fee4d859da0951b65ac44e72/contrib/kingres_timepart/kingres_timepart.control)
- [Official source](https://github.com/KingresDB/Kingres/blob/3c98d1cd4a8d81e2fee4d859da0951b65ac44e72/contrib/kingres_timepart/kingres_timepart--0.1.sql)
- [Official source](https://github.com/KingresDB/Kingres/blob/3c98d1cd4a8d81e2fee4d859da0951b65ac44e72/kingres/README.md)

`kingres_timepart` 0.1 manages creation and retention of ordinary time-range partitions using SQL and PL/pgSQL. It has no shared library, preload requirement or background scheduler. The same installation SQL targets PostgreSQL 18/19.

### Define and inspect a policy

Installation needs a superuser because it creates event triggers. The example is an administrative workflow; a maintenance run may drop expired partitions. Inspect its plan first, and issue the procedure outside a transaction block.

```sql
CREATE EXTENSION kingres_timepart;
CREATE TABLE public.events (ts timestamptz, payload text) PARTITION BY RANGE (ts);
SELECT kingres.timepart_set_policy(
    'public.events', '1 day', create_ahead => '7 days', retention => '30 days'
);
SELECT * FROM kingres.timepart_plan('public.events');
CALL kingres.timepart_run('public.events');
```

### Maintenance and privileges

Tables must use a single plain `date`, `timestamp` or `timestamptz` RANGE key and have no DEFAULT partition. Owners with explicit execution grants can configure policies and use `kingres.create_time_partitions`, `kingres.drop_old_time_partitions`, `kingres.timepart_plan` and `kingres.timepart_check`. The maintenance procedure is superuser-only and commits between actions; a partially failed run does not roll back already committed work.

`kingres.timepart_policies` and run-history views expose owner-filtered state. Policies do nothing until an external caller runs them. Time grids default to UTC and distinguish calendar days from fixed hours. Misaligned partitions are reported rather than retired automatically. Distributed tables, automatic scheduling and columnar-conversion procedures are not implemented features of this version. Leave the reserved columnar policy fields unused.
