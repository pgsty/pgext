## Usage

Sources:

- [Official upstream README](https://github.com/yuuch/pg_yaap/blob/41eda4a6020c9e0c0185586ffdd5ed4c55e94585/README.md)
- [Official extension control file (pg_yaap.control)](https://github.com/yuuch/pg_yaap/blob/41eda4a6020c9e0c0185586ffdd5ed4c55e94585/pg_yaap.control)
- [Official extension SQL (pg_yaap--1.0.sql)](https://github.com/yuuch/pg_yaap/blob/41eda4a6020c9e0c0185586ffdd5ed4c55e94585/pg_yaap--1.0.sql)

`pg_yaap` — pg_yaap is a PostgreSQL extension that replaces PostgreSQL's planner and executor path for supported analytical queries with a YAAP-owned optimizer and a C++ columnar execution engine. Use it for the corresponding analytical or storage workflow. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION pg_yaap;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
