## Usage

Sources:

- [Official extension control file (tg_sanity_tap.control)](https://api.pgxn.org/src/tg_sanity/tg_sanity-0.1.0/tg_sanity_tap.control)
- [Official extension SQL (tg_sanity_tap.sql)](https://api.pgxn.org/src/tg_sanity/tg_sanity-0.1.0/sql/tg_sanity_tap.sql)

`tg_sanity_tap` — pgtap testing function for tg sanity triggers. Use it when an application needs this specific database capability. Its extension dependencies must be installed and validated first.

### Core Workflow

```sql
CREATE EXTENSION tg_sanity_tap;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `tg_sanity_tap(trigger_table regclass , trigger_name text , timing text , events text , trigger_arguments text)` is an extension function and returns `SETOF`.

### Requirements and Caveats

- The reviewed control file declares default version `0.1.0`.
- Install the confirmed extension dependencies first: `tg_sanity`, `pgtap`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
