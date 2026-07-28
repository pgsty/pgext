## Usage

Sources:

- [Official upstream README](https://github.com/x4m/pg_uuid_next/blob/21d2df4203448e8d4149859dce77ce420bc0d1d8/README.md)
- [Official extension control file (pg_uuid_next.control)](https://github.com/x4m/pg_uuid_next/blob/21d2df4203448e8d4149859dce77ce420bc0d1d8/pg_uuid_next.control)
- [Official extension SQL (pg_uuid_next--1.0.sql)](https://github.com/x4m/pg_uuid_next/blob/21d2df4203448e8d4149859dce77ce420bc0d1d8/pg_uuid_next--1.0.sql)

`pg_uuid_next` — Extension to generate UUID version 7 and 8. Use it when SQL needs these specialized functions or aggregates. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION pg_uuid_next;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `gen_uuid_v7()` is an extension function and returns `uuid`.
- `gen_uuid_v8()` is an extension function and returns `uuid`.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
