## Usage

Sources:

- [Official upstream README](https://github.com/acoustid/pg_acoustid/blob/102f3c870c6157704694c2ddbad3ae8ab2c7de91/README.md)
- [Official extension control file (acoustid.control)](https://github.com/acoustid/pg_acoustid/blob/102f3c870c6157704694c2ddbad3ae8ab2c7de91/acoustid.control)
- [Official extension SQL (acoustid--1.0.sql)](https://github.com/acoustid/pg_acoustid/blob/102f3c870c6157704694c2ddbad3ae8ab2c7de91/acoustid--1.0.sql)

`acoustid` — AcoustID utility functions for PostgreSQL =========================================. Use it when SQL needs these specialized functions or aggregates. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION acoustid;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `acoustid_compare2(int4[], int4[], int DEFAULT 0)` is an extension function and returns `float4`.
- `acoustid_compare3(int4[], int4[], int DEFAULT -1)` is an extension function and returns `float4`.
- `acoustid_extract_query(int4[])` is an extension function and returns `int4[]`.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
