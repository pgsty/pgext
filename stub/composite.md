## Usage

Sources:

- [Official upstream README](https://github.com/chanukyasds/pg_extensions/blob/d732560f24a2741225182a8b74a36837cc4abc3f/README.md)
- [Official extension control file (composite.control)](https://github.com/chanukyasds/pg_extensions/blob/d732560f24a2741225182a8b74a36837cc4abc3f/functions/composite/composite.control)
- [Official extension SQL (composite--1.0.sql)](https://github.com/chanukyasds/pg_extensions/blob/d732560f24a2741225182a8b74a36837cc4abc3f/functions/composite/composite--1.0.sql)

`composite` — aggregates and functions for PostgreSQL Server. Use it when SQL needs these specialized functions or aggregates. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION composite;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `composite(custom_type[])` is an extension function and returns `custom_type[]`.
- `custom_type` is an extension-defined type.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
