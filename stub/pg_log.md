## Usage

Sources:

- [Official upstream README](https://api.pgxn.org/src/pg_log/pg_log-0.0.3/README.md)
- [Official extension control file (pg_log.control)](https://api.pgxn.org/src/pg_log/pg_log-0.0.3/pg_log.control)
- [Official extension SQL (pg_log--0.0.1.sql)](https://api.pgxn.org/src/pg_log/pg_log-0.0.3/pg_log--0.0.1.sql)

`pg_log` — PostgreSQL extension to display log from SQL. Use it when collecting or interpreting the corresponding PostgreSQL statistics. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION pg_log;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `pg_get_logname()` is an extension function and returns `cstring`.
- `pg_log(OUT line integer, OUT message text)` is an extension function and returns `SETOF`.
- `pg_log_refresh()` is an extension function and returns `void`.
- `pg_read(cstring)` is an extension function and returns `void`.
- `log` is an extension-defined view.
- `pglog` is a table installed or managed by the extension.

### Requirements and Caveats

- The catalog records version `1.0.0`, while the reviewed control file declares default version `0.0.1`.
- The control file marks the extension as non-relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
