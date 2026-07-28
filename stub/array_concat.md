## Usage

Sources:

- [Official upstream README](https://github.com/chanukyasds/pg_extensions/blob/d732560f24a2741225182a8b74a36837cc4abc3f/README.md)
- [Official extension control file (array_concat.control)](https://github.com/chanukyasds/pg_extensions/blob/d732560f24a2741225182a8b74a36837cc4abc3f/functions/array_concat/array_concat.control)
- [Official extension SQL (array_concat--1.0.sql)](https://github.com/chanukyasds/pg_extensions/blob/d732560f24a2741225182a8b74a36837cc4abc3f/functions/array_concat/array_concat--1.0.sql)

`array_concat` — aggregates and functions for PostgreSQL Server. Use it when SQL needs these specialized functions or aggregates. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION array_concat;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `array_concat(text[])` is an extension function and returns `text`.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
