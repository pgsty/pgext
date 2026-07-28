## Usage

Sources:

- [Official upstream README](https://github.com/matdehaast/pg_tigerbeetle/blob/c8e8b581e4b344a6cc084db55d1d19f4aa1ff442/README.md)
- [Official extension control file (pg_tigerbeetle.control)](https://github.com/matdehaast/pg_tigerbeetle/blob/c8e8b581e4b344a6cc084db55d1d19f4aa1ff442/extension/pg_tigerbeetle.control)
- [Official extension SQL (pg_tigerbeetle--0.1.sql)](https://github.com/matdehaast/pg_tigerbeetle/blob/c8e8b581e4b344a6cc084db55d1d19f4aa1ff442/extension/pg_tigerbeetle--0.1.sql)

`pg_tigerbeetle` — Experimental Zig extension for looking up TigerBeetle accounts from PostgreSQL. Use it for the corresponding SQL or database utility workflow. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION pg_tigerbeetle;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `lookup_account()` is an extension function and returns `TEXT`.
- `query_by_id(int4)` is an extension function and returns `TEXT`.

### Requirements and Caveats

- The reviewed control file declares default version `0.1`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
