## Usage

Sources:

- [Official upstream README](https://api.pgxn.org/src/pg_monetdb/pg_monetdb-1.4.0/README.md)
- [Official extension control file (pg_monetdb.control)](https://api.pgxn.org/src/pg_monetdb/pg_monetdb-1.4.0/pg_monetdb.control)
- [Official extension SQL (pg_monetdb--1.3.sql)](https://api.pgxn.org/src/pg_monetdb/pg_monetdb-1.4.0/pg_monetdb--1.3.sql)

`pg_monetdb` — pg_monetdb is a fork of monetdb_fdw focused on stronger pushdown for analytical query shapes derived from TPC-H and TPC-DS-style workloads. Use it when PostgreSQL must access the corresponding external data source through a foreign-data interface. Upstream describes this capability as experimental.

### Core Workflow

```sql
CREATE EXTENSION pg_monetdb;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `monet_query(server name, statement text)` is an extension function and returns `SETOF`.
- `monet_query_to_array(server name, statement text)` is an extension function and returns `SETOF`.
- `monet_query_to_jsonb(server name, statement text, column_names text[])` is an extension function and returns `SETOF`.
- `monetdb_execute(server name, statement text)` is an extension function and returns `void`.
- `monetdb_fdw_handler()` is an extension function and returns `fdw_handler`.
- `pg_monetdb_execute(server name, statement text)` is an extension function and returns `void`.
- `pg_monetdb_handler()` is an extension function and returns `fdw_handler`.
- `pg_monetdb_query(server name, statement text)` is an extension function and returns `SETOF`.
- `pg_monetdb_query_to_array(server name, statement text)` is an extension function and returns `SETOF`.
- `pg_monetdb_query_to_jsonb(server name, statement text, column_names text[])` is an extension function and returns `SETOF`.
- `BLOB` is an extension-defined domain.
- `CLOB` is an extension-defined domain.
- `HUGEINT` is an extension-defined domain.
- `STRING` is an extension-defined domain.

### Requirements and Caveats

- The reviewed control file declares default version `1.4`.
- The control file marks the extension as relocatable.
- Upstream labels part or all of the project experimental.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
