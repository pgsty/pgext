## Usage

Sources:

- [Official upstream README](https://github.com/valehdba/pgx_warnings/blob/723431d3557d7b341ec663cf3aa16f23dc2f5973/README.md)
- [Official extension control file (pgx_warnings.control)](https://github.com/valehdba/pgx_warnings/blob/723431d3557d7b341ec663cf3aa16f23dc2f5973/pgx_warnings.control)
- [Official extension SQL (pgx_warnings--1.0.sql)](https://github.com/valehdba/pgx_warnings/blob/723431d3557d7b341ec663cf3aa16f23dc2f5973/pgx_warnings--1.0.sql)

`pgx_warnings` — A PostgreSQL extension that **captures all WARNING, ERROR, FATAL, and PANIC messages** from the server log in real time and **sends instant notifications to a Telegram channel**. Use it when collecting or interpreting the corresponding PostgreSQL statistics. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION pgx_warnings;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `pgx_warnings_clear()` is an extension function and returns `void`.
- `pgx_warnings_list(IN max_entries integer DEFAULT 100, OUT "timestamp" timestamptz, OUT level text, OUT database text, OUT message text, OUT pid integer, OUT sent boolean)` is an extension function and returns `SETOF`.
- `pgx_warnings_stats(OUT current_entries integer, OUT buffer_size integer, OUT total_captured bigint, OUT total_sent bigint, OUT total_failed bigint, OUT enabled boolean)` is an extension function and returns `record`.
- `pgx_warnings_test()` is an extension function and returns `text`.
- `pgx_warnings` is an extension-defined view.
- `pgx_warnings_info` is an extension-defined view.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- The control file marks the extension as non-relocatable.
- The control file requires a superuser for installation.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
