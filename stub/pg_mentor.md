## Usage

Sources:

- [Official upstream README](https://github.com/danolivo/pg_mentor/blob/cd7177d756dbfa71f0ccf9a043e3b99685f4a887/README.md)
- [Official extension control file (pg_mentor.control)](https://github.com/danolivo/pg_mentor/blob/cd7177d756dbfa71f0ccf9a043e3b99685f4a887/pg_mentor.control)
- [Official extension SQL (pg_mentor--0.1.sql)](https://github.com/danolivo/pg_mentor/blob/cd7177d756dbfa71f0ccf9a043e3b99685f4a887/pg_mentor--0.1.sql)

`pg_mentor` — Lightweight extension that employs query statistics stored in the pg_stat_statements extension to decide which type of plan mode (custom, generic or auto) to use. Use it when administering or automating the database behavior described above. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION pg_mentor;

SELECT reconsider_ps_modes();
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `pg_mentor_nail_long_planned()` is an extension function and returns `integer`.
- `pg_mentor_reload_conf(void)` is an extension function.
- `pg_mentor_reset()` is an extension function and returns `integer`.
- `pg_mentor_set_plan_mode(queryId bigint, status integer, ref_total_time float8 DEFAULT NULL, ref_nblocks float8 DEFAULT NULL, fixed bool DEFAULT false)` is an extension function.
- `pg_mentor_show_prepared_statements(IN status integer, OUT queryid bigint, OUT refcounter integer, OUT plan_cache_mode int, OUT since TimestampTz, OUT fixed boolean, OUT statnum integer, OUT nblocks bigint[], OUT exec_times float8[], OUT avg_nblocks float8, OUT avg_exec_time float8, OUT ref_nblo…)` is an extension function and returns `SETOF`.
- `reconsider_ps_modes(OUT to_generic bigint, OUT to_custom bigint, OUT unchanged bigint)` is an extension function and returns `record`.

### Requirements and Caveats

- The reviewed control file declares default version `0.1`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
