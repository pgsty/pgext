## Usage

Sources:

- [Official upstream README](https://github.com/solidcoredata/pgwire4/blob/13f90441a022cce962f6e6bf1b710703c21af19a/README.md)
- [Official extension control file (pgwire4.control)](https://github.com/solidcoredata/pgwire4/blob/13f90441a022cce962f6e6bf1b710703c21af19a/ext/pgwire4.control)
- [Official extension SQL (pgwire4--1.0.sql)](https://github.com/solidcoredata/pgwire4/blob/13f90441a022cce962f6e6bf1b710703c21af19a/ext/sql/pgwire4--1.0.sql)

`pgwire4` — PostgreSQL extension to provide a new wire protocol. Use it when an application needs this specific database capability. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION pgwire4;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `pgwire4.bulk_int8(stream_name text)` is an extension function and returns `SETOF`.
- `pgwire4.bulk_scan(stream_name text)` is an extension function and returns `SETOF`.
- `pgwire4.stats(OUT accepted bigint, OUT dispatch_failures bigint, OUT workers integer, OUT databases integer, OUT listener_pid integer)` is an extension function and returns `record`.
- `pgwire4.status(OUT slot integer, OUT pid integer, OUT database text, OUT state text, OUT sessions bigint, OUT queries bigint, OUT rows_streamed bigint, OUT cancels bigint, OUT errors bigint, OUT cache_hits bigint, OUT cache_misses bigint, OUT session_user_name text, OUT sess…)` is an extension function and returns `SETOF`.
- `pgwire4.version()` is an extension function and returns `text`.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- The control file marks the extension as non-relocatable.
- The control file requires a superuser for installation.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
