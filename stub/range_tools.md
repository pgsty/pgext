## Usage

Sources:

- [Official extension control file (range_tools.control)](https://api.pgxn.org/src/range_tools/range_tools-0.1.2/range_tools.control)
- [Official extension SQL (range_tools.sql)](https://api.pgxn.org/src/range_tools/range_tools-0.1.2/sql/range_tools.sql)

`range_tools` — Tools for use with range types. Use it when SQL needs these specialized functions or aggregates. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION range_tools;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `range_type` is an extension-defined view.

### Requirements and Caveats

- The reviewed control file declares default version `0.1.2`.
- The control file marks the extension as non-relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
