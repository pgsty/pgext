## Usage

Sources:

- [Official upstream README](https://github.com/babelfish-for-postgresql/babelfish_extensions/blob/46617a93de0eb666ce98591cfcbae4554f6f6ea0/contrib/babelfishpg_unit/README.md)
- [Official extension control file (babelfishpg_unit.control)](https://github.com/babelfish-for-postgresql/babelfish_extensions/blob/46617a93de0eb666ce98591cfcbae4554f6f6ea0/contrib/babelfishpg_unit/babelfishpg_unit.control)
- [Official extension SQL (babelfishpg_unit--1.0.0.sql)](https://github.com/babelfish-for-postgresql/babelfish_extensions/blob/46617a93de0eb666ce98591cfcbae4554f6f6ea0/contrib/babelfishpg_unit/babelfishpg_unit--1.0.0.sql)

`babelfishpg_unit` — Babelfish has introduced a new extension named babelfishpg_unit which enables us to run unit tests. Please follow the build instructions to build and install the babelfishpg_unit extension. Use it when an application needs this specific database capability. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION babelfishpg_unit;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `babelfishpg_unit.babelfishpg_unit_run_tests()` is an extension function and returns `TABLE`.
- `babelfishpg_unit.babelfishpg_unit_run_tests(VARIADIC name text[])` is an extension function and returns `TABLE`.

### Requirements and Caveats

- The reviewed control file declares default version `1.0.0`.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
