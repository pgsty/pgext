## Usage

Sources:

- [Official upstream README](https://api.pgxn.org/src/pg_kaboom/pg_kaboom-0.0.1/README.md)
- [Official extension control file (pg_kaboom.control)](https://api.pgxn.org/src/pg_kaboom/pg_kaboom-0.0.1/pg_kaboom.control)
- [Official extension SQL (pg_kaboom--0.0.1.sql)](https://api.pgxn.org/src/pg_kaboom/pg_kaboom-0.0.1/pg_kaboom--0.0.1.sql)

`pg_kaboom` — This extension serves to crash postgresql in multiple varied and destructive ways. Use it when an application needs this specific database capability. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION pg_kaboom;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `pg_kaboom(method text, payload jsonb default NULL)` is an extension function and returns `boolean`.
- `pg_kaboom_arsenal()` is an extension function and returns `TABLE`.

### Requirements and Caveats

- The reviewed control file declares default version `0.0.1`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
