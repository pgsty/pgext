## Usage

Sources:

- [Official upstream README](https://github.com/florents-tselai/pg_fts_greek/blob/c4700c43945f5980308028824b3d13d762960f2c/README.md)
- [Official extension control file (pg_fts_greek.control)](https://github.com/florents-tselai/pg_fts_greek/blob/c4700c43945f5980308028824b3d13d762960f2c/pg_fts_greek.control)
- [Official extension SQL (pg_fts_greek--0.1.sql)](https://github.com/florents-tselai/pg_fts_greek/blob/c4700c43945f5980308028824b3d13d762960f2c/sql/pg_fts_greek--0.1.sql)

`pg_fts_greek` — Postgres FTS Improvements for Greek. Use it for the corresponding text-search, parsing, or linguistic workflow. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION pg_fts_greek;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Requirements and Caveats

- The reviewed control file declares default version `0.1`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
