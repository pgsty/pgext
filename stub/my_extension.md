## Usage

Sources:

- [Official upstream README](https://api.pgxn.org/src/my_extension/my_extension-1.1.0/README.md)
- [Official extension control file (my_extension.control)](https://api.pgxn.org/src/my_extension/my_extension-1.1.0/my_extension.control)
- [Official extension SQL (my_extension--1.0.1.sql)](https://api.pgxn.org/src/my_extension/my_extension-1.1.0/my_extension--1.0.1.sql)

`my_extension` — my_extension is a basic PostgreSQL extension that provides additional functionality for efficient data manipulation and calculations. Use it when an application needs this specific database capability. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION my_extension;

SELECT add(1, 2);
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `add(a integer, b integer)` is an extension function and returns `integer`.
- `complex_add(integer[])` is an extension function and returns `integer`.
- `multiply(a integer, b integer)` is an extension function and returns `integer`.
- `my_table` is a table installed or managed by the extension.

### Requirements and Caveats

- The reviewed control file declares default version `1.1.0`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
