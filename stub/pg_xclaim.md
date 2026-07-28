## Usage

Sources:

- [Official upstream README](https://github.com/fenoman/pg_xclaim/blob/e26926e06a9442b8554e6032e22a8080360d5d1d/README.md)
- [Official extension control file (pg_xclaim.control)](https://github.com/fenoman/pg_xclaim/blob/e26926e06a9442b8554e6032e22a8080360d5d1d/pg_xclaim.control)
- [Official extension SQL (pg_xclaim--1.0.0-rc1.sql)](https://github.com/fenoman/pg_xclaim/blob/e26926e06a9442b8554e6032e22a8080360d5d1d/sql/pg_xclaim--1.0.0-rc1.sql)

`pg_xclaim` — **Экспериментальное** PostgreSQL-расширение: альтернативный механизм хранения транзакционных claim'ов в собственной партиционированной shared-memory hashtable. Use it when an application needs this specific database capability. Upstream explicitly says it is not production-ready.

### Core Workflow

```sql
CREATE EXTENSION pg_xclaim;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `xclaim.count()` is an extension function and returns `int8`.
- `xclaim.debug()` is an extension function and returns `TABLE`.
- `xclaim.debug_inject_stale(scope int4, key int4)` is an extension function and returns `void`.
- `xclaim.debug_snapshot()` is an extension function and returns `TABLE`.
- `xclaim.session_reset()` is an extension function and returns `void`.
- `xclaim.stats()` is an extension function and returns `TABLE`.
- `xclaim.try(classid int4, objid int4)` is an extension function and returns `boolean`.
- `xclaim.try(key int8)` is an extension function and returns `boolean`.
- `xclaim.try_many(classid int4, objids int4[])` is an extension function and returns `boolean[]`.
- `xclaim.try_many(keys int8[])` is an extension function and returns `boolean[]`.

### Requirements and Caveats

- The reviewed control file declares default version `1.0.0-rc1`.
- The control file marks the extension as non-relocatable.
- The control file requires a superuser for installation.
- The control file marks the extension as not trusted.
- Upstream explicitly says the project is not production-ready.
- Upstream labels part or all of the project experimental.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
