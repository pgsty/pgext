## Usage

Sources:

- [Official extension control file (tg_sanity.control)](https://api.pgxn.org/src/tg_sanity/tg_sanity-0.1.0/tg_sanity.control)
- [Official extension SQL (tg_sanity.sql)](https://api.pgxn.org/src/tg_sanity/tg_sanity-0.1.0/sql/tg_sanity.sql)

`tg_sanity` — Trigger function for enforcing data quality. Use it when an application needs this specific database capability. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION tg_sanity;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `tg_sanity()` is an extension function and returns `trigger`.

### Requirements and Caveats

- The reviewed control file declares default version `0.1.0`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
