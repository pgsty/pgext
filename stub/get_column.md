## Usage

Sources:

- [Official extension control file (get_column.control)](https://api.pgxn.org/src/get_column/get_column-1.0.0/get_column.control)
- [Official extension SQL (get_column--1.0.sql)](https://api.pgxn.org/src/get_column/get_column-1.0.0/get_column--1.0.sql)

`get_column` — Fetch a column value from a record by name. Use it when SQL needs these specialized functions or aggregates. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION get_column;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `get_column(record, text)` is an extension function and returns `anyelement`.

### Requirements and Caveats

- The reviewed control file declares default version `1.0.0`.
- The control file marks the extension as non-relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
