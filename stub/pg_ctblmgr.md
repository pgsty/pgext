## Usage

Sources:

- [Official upstream README](https://github.com/spd010273/pg_ctblmgr/blob/e15de246e28b50a9f319f8dcffd34a1c0adae751/README.md)
- [Official extension control file (pg_ctblmgr.control)](https://github.com/spd010273/pg_ctblmgr/blob/e15de246e28b50a9f319f8dcffd34a1c0adae751/server/pg_ctblmgr.control)
- [Official extension SQL (pg_ctblmgr--0.1.sql)](https://github.com/spd010273/pg_ctblmgr/blob/e15de246e28b50a9f319f8dcffd34a1c0adae751/server/sql/pg_ctblmgr--0.1.sql)

`pg_ctblmgr` — Logical Replication Based, Asynchronous Eager Materialized Views for PostgreSQL. Use it when an application needs this specific database capability. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION pg_ctblmgr;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Requirements and Caveats

- The reviewed control file declares default version `0.1`.
- The control file marks the extension as relocatable.
- The control file does not require superuser-only installation.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
