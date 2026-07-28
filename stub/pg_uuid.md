## Usage

Sources:

- [Official upstream README](https://github.com/tvondra/pg_uuid/blob/5abbd6a5ed12b65674f6ef1c6a6d07c16f4b6f68/README.md)
- [Official extension control file (pg_uuid.control)](https://github.com/tvondra/pg_uuid/blob/5abbd6a5ed12b65674f6ef1c6a6d07c16f4b6f68/pg_uuid.control)
- [Official extension SQL (pg_uuid--1.0.0.sql)](https://github.com/tvondra/pg_uuid/blob/5abbd6a5ed12b65674f6ef1c6a6d07c16f4b6f68/pg_uuid--1.0.0.sql)

`pg_uuid` — generator of new UUID versions. Use it when SQL needs these specialized functions or aggregates. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION pg_uuid;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `uuid_generate_v6()` is an extension function and returns `uuid`.
- `uuid_generate_v7()` is an extension function and returns `uuid`.

### Requirements and Caveats

- The reviewed control file declares default version `1.0.0`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
