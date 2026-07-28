## Usage

Sources:

- [Official upstream README](https://github.com/halgari/level-pivot/blob/ca600015cac235c84cae804e5ef7a19dcdad9be0/README.md)
- [Official extension control file (level_pivot.control)](https://github.com/halgari/level-pivot/blob/ca600015cac235c84cae804e5ef7a19dcdad9be0/sql/level_pivot.control)
- [Official extension SQL (level_pivot--1.0.sql)](https://github.com/halgari/level-pivot/blob/ca600015cac235c84cae804e5ef7a19dcdad9be0/sql/level_pivot--1.0.sql)

`level_pivot` — LevelDB FDW with support for pivoting and key templating. Use it when PostgreSQL must access the corresponding external data source through a foreign-data interface. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION level_pivot;

SELECT * FROM users WHERE group_name = 'admins';
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `level_pivot_fdw_handler()` is an extension function and returns `fdw_handler`.
- `level_pivot_fdw_validator(text[], oid)` is an extension function and returns `void`.
- `level_pivot` is an extension-defined foreign data wrapper.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
