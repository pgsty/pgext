## Usage

Sources:

- [Official upstream README](https://github.com/timvaillancourt/pg_mysql/blob/d79179caf3013967faf220af08207adbf9c790e5/README.md)
- [Official extension control file (pg_mysql.control)](https://github.com/timvaillancourt/pg_mysql/blob/d79179caf3013967faf220af08207adbf9c790e5/pg_mysql.control)
- [Official extension SQL (pg_mysql--1.0.sql)](https://github.com/timvaillancourt/pg_mysql/blob/d79179caf3013967faf220af08207adbf9c790e5/pg_mysql--1.0.sql)

`pg_mysql` — Production-grade replication for PostgreSQL - at last! Use it when porting or emulating the corresponding database API. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION pg_mysql;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `mysql.delete(port integer)` is an extension function and returns `text`.
- `mysql.query(sql text, port integer DEFAULT 0)` is an extension function and returns `text`.
- `mysql.start(port integer DEFAULT 0, semi_sync boolean DEFAULT false)` is an extension function and returns `mysql`.
- `mysql.start_replica(source_host text, source_port integer, port integer DEFAULT 0, semi_sync boolean DEFAULT false)` is an extension function and returns `mysql`.
- `mysql.status(port integer DEFAULT 0)` is an extension function and returns `SETOF`.
- `mysql.stop(port integer)` is an extension function and returns `text`.
- `mysql.start_result` is an extension-defined type.
- `mysql.status_info` is an extension-defined type.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- The control file marks the extension as non-relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
