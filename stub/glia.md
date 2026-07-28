## Usage

Sources:

- [Official upstream README](https://api.pgxn.org/src/glia/glia-0.0.1/README.md)
- [Official extension control file (glia.control)](https://api.pgxn.org/src/glia/glia-0.0.1/glia.control)
- [Official extension SQL (glia--0.0.1.sql)](https://api.pgxn.org/src/glia/glia-0.0.1/glia--0.0.1.sql)

`glia` — *__A PostgreSQL extension for data mining.__*. Use it for the corresponding vector, model, or retrieval workflow. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION glia;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Requirements and Caveats

- The reviewed control file declares default version `0.0.1`.
- The control file marks the extension as relocatable.
- The former GitHub repository URL returned 404 during the 2026-07-28 review; treat the pinned PGXN distribution above as the available source boundary.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
