## Usage

Sources:

- [Official upstream README](https://github.com/takanoriyanagitani/dir_fdw/blob/c6edbbb4a9928a687a576ccef058e496e2aaffc9/README.md)
- [Official extension control file (dir_fdw.control)](https://github.com/takanoriyanagitani/dir_fdw/blob/c6edbbb4a9928a687a576ccef058e496e2aaffc9/dir_fdw.control)
- [Official extension SQL (dir_fdw--1.0.sql)](https://github.com/takanoriyanagitani/dir_fdw/blob/c6edbbb4a9928a687a576ccef058e496e2aaffc9/dir_fdw--1.0.sql)

`dir_fdw` — foreign-data wrapper for readdir. Use it when PostgreSQL must access the corresponding external data source through a foreign-data interface. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION dir_fdw;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `dir_fdw_handler()` is an extension function and returns `fdw_handler`.
- `dir_fdw_validator(text[], oid)` is an extension function and returns `void`.
- `dir_fdw` is an extension-defined foreign data wrapper.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
