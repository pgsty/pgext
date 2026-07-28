## Usage

Sources:

- [Official upstream README](https://github.com/chanukyasds/pg_running_stats/blob/20d42698a0c7b594d930a55bfe81ed5c7163a058/README.md)
- [Official extension control file (pg_running_stats.control)](https://github.com/chanukyasds/pg_running_stats/blob/20d42698a0c7b594d930a55bfe81ed5c7163a058/pg_running_stats.control)
- [Official extension SQL (pg_running_stats--1.0.sql)](https://github.com/chanukyasds/pg_running_stats/blob/20d42698a0c7b594d930a55bfe81ed5c7163a058/pg_running_stats--1.0.sql)

`pg_running_stats` — Mergeable Running Statistics (Welford/Chan) for PostgreSQL. Use it when SQL needs these specialized functions or aggregates. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION pg_running_stats;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `rstat_combine(a bytea, b bytea)` is an extension function and returns `bytea`.
- `rstat_final(state bytea)` is an extension function and returns `rstat_result_t`.
- `rstat_sfunc(state bytea, x double precision)` is an extension function and returns `bytea`.
- `rstat_state_merge(a bytea, b bytea)` is an extension function and returns `bytea`.
- `rstat_state_result(state bytea)` is an extension function and returns `rstat_result_t`.
- `rstat_state` is an aggregate exposed by the extension.
- `running_stats` is an aggregate exposed by the extension.
- `rstat_result_t` is an extension-defined type.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
