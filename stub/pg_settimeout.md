## Usage

Sources:

- [README.md](https://github.com/zesoup/pg_settimeout/blob/60525e8b2a4076359a8546879f27e24ce561f686/README.md)
- [src/pg_settimeout--1.1.sql](https://github.com/zesoup/pg_settimeout/blob/60525e8b2a4076359a8546879f27e24ce561f686/src/pg_settimeout--1.1.sql)
- [src/pg_settimeout.c](https://github.com/zesoup/pg_settimeout/blob/60525e8b2a4076359a8546879f27e24ce561f686/src/pg_settimeout.c)
- [src/pg_settimeout.control](https://github.com/zesoup/pg_settimeout/blob/60525e8b2a4076359a8546879f27e24ce561f686/src/pg_settimeout.control)

`pg_settimeout` starts dynamic background workers to execute SQL after a delay. Upstream reports testing on PostgreSQL 9.6, 10 and 11; compatibility with later majors is unverified.

### Core Workflow

```sql
CREATE EXTENSION pg_settimeout;
SELECT pg_settimeout('SELECT 1', 1000);
```

### Operational Boundaries

`pg_settimeout(query, timeout)` runs once; `pg_setinterval(query, timeout)` repeats indefinitely. The timeout is in milliseconds and the result is the worker PID. Workers connect to the current database as the calling user and execute in separate transactions; rolling back the scheduling transaction does not undo an independently completed task. Available worker slots limit scheduling. The module has no startup routines and needs no preload. Installation uses the control-file superuser default.
