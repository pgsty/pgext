## Usage

Sources:

- [Official upstream README](https://github.com/animeshs34/timeql/blob/a3032739628758757e223d49fc67b276c6e14f8c/README.md)
- [Official extension control file (timeql.control)](https://github.com/animeshs34/timeql/blob/a3032739628758757e223d49fc67b276c6e14f8c/timeql.control)
- [Official extension SQL (timeql--1.0.sql)](https://github.com/animeshs34/timeql/blob/a3032739628758757e223d49fc67b276c6e14f8c/timeql--1.0.sql)

`timeql` — TimeQL is a PostgreSQL extension that provides native, high-performance temporal data tracking and "Point-in-Time" querying capabilities. It uses PostgreSQL's native Range Types (tsrange) and GiST indexing to ensure that tracking your data's history doesn't kill your performance. Use it for the corresponding scheduling, temporal, or time-series workflow. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION timeql;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `timeql.create_history_table(p_schema TEXT, p_table TEXT)` is an extension function and returns `void`.
- `timeql.create_history_table(p_schema TEXT, p_table TEXT, p_retention_period INTERVAL DEFAULT NULL)` is an extension function and returns `void`.
- `timeql.disable_temporal(p_schema TEXT, p_table TEXT)` is an extension function and returns `void`.
- `timeql.enable_temporal(p_schema TEXT, p_table TEXT)` is an extension function and returns `void`.
- `timeql.purge_all_history()` is an extension function and returns `void`.
- `timeql.purge_history(p_schema TEXT, p_table TEXT)` is an extension function and returns `void`.
- `timeql.restore_at(p_table_reg regclass, p_pk_value ANYELEMENT, p_timestamp TIMESTAMP WITHOUT TIME ZONE)` is an extension function and returns `void`.
- `timeql.set_retention_policy(p_schema TEXT, p_table TEXT, p_retention_period INTERVAL)` is an extension function and returns `void`.
- `timeql.tql_at(p_table_reg regclass, p_timestamp TIMESTAMP WITHOUT TIME ZONE)` is an extension function and returns `SETOF`.
- `timeql.tql_between(p_table_reg regclass, p_start TIMESTAMP WITHOUT TIME ZONE, p_end TIMESTAMP WITHOUT TIME ZONE)` is an extension function and returns `SETOF`.
- `timeql.track_delete()` is an extension function and returns `TRIGGER`.
- `timeql.track_insert()` is an extension function and returns `TRIGGER`.
- `timeql.track_update()` is an extension function and returns `TRIGGER`.
- `timeql_version()` is an extension function and returns `text`.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- The control file marks the extension as non-relocatable.
- The control file requires a superuser for installation.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
