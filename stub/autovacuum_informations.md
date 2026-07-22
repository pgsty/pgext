## Usage

Sources:

- [Official README](https://github.com/gleu/autovacuum_informations/blob/0dde21773a3929e0b0bacca018607d911fa71f00/README.md)
- [Official extension SQL](https://github.com/gleu/autovacuum_informations/blob/0dde21773a3929e0b0bacca018607d911fa71f00/autovacuum_informations--1.0.sql)
- [Official implementation](https://github.com/gleu/autovacuum_informations/blob/0dde21773a3929e0b0bacca018607d911fa71f00/autovacuum_informations.c)

`autovacuum_informations` 1.0 exposes PostgreSQL's in-memory autovacuum launcher and worker state through two C functions. It is a diagnostic prototype tied to PostgreSQL internals; upstream explicitly says that it is work in progress and must not be used in production.

### Core Workflow

Install the extension, then read the launcher PID or expand the record returned for each active worker:

```sql
CREATE EXTENSION autovacuum_informations;

SELECT get_autovacuum_launcher_pid();

SELECT *
FROM get_autovacuum_workers_infos() AS w(
  idx integer,
  datid oid,
  relid oid,
  pid integer,
  launchtime timestamptz,
  dobalance boolean,
  cost_delay integer,
  cost_limit integer,
  cost_limit_base integer
);
```

The launcher function returns a `bigint` PID, or null when no launcher PID is available. The worker function is declared as `SETOF record`, so callers must supply the nine-column record definition shown above.

### Important Objects

- `get_autovacuum_launcher_pid()` reads the launcher process ID from autovacuum shared memory.
- `get_autovacuum_workers_infos()` reports the current worker index, database and relation OIDs, PID, launch time, balancing flag, and cost settings.

### Operational Notes

The control file fixes the extension in `pg_catalog` and marks it non-relocatable. It has no GUC or preload requirement, but its C implementation uses PostgreSQL private autovacuum structures and locks, so server-version compatibility must be validated against the exact build. Treat all output as transient diagnostic state and retain upstream's explicit non-production warning.
