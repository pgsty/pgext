## Usage

Sources:

- [Official upstream README](https://github.com/project-tsurugi/tsurugi_fdw/blob/e777ab56b5cbff4df43f608299cd73739e92e7aa/README.md)
- [Official extension control file (tsurugi_fdw.control)](https://github.com/project-tsurugi/tsurugi_fdw/blob/e777ab56b5cbff4df43f608299cd73739e92e7aa/tsurugi_fdw.control)
- [Official extension SQL (tsurugi_fdw--1.4.0--1.5.0.sql)](https://github.com/project-tsurugi/tsurugi_fdw/blob/e777ab56b5cbff4df43f608299cd73739e92e7aa/tsurugi_fdw--1.4.0--1.5.0.sql)

`tsurugi_fdw` — tsurugi_fdw is a PostgreSQL extension that provides a Foreign Data Wrapper for access to Tsurugi. Use it when PostgreSQL must access the corresponding external data source through a foreign-data interface. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION tsurugi_fdw;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `tg_execute_ddl(text DEFAULT null, text DEFAULT null)` is an extension function and returns `TEXT`.
- `tg_set_exclusive_read_areas(variadic text[])` is an extension function and returns `cstring`.
- `tg_set_inclusive_read_areas(variadic text[])` is an extension function and returns `cstring`.
- `tg_set_transaction(text)` is an extension function and returns `cstring`.
- `tg_set_transaction(text, text)` is an extension function and returns `cstring`.
- `tg_set_transaction(text, text, text)` is an extension function and returns `cstring`.
- `tg_set_write_preserve(variadic text[])` is an extension function and returns `cstring`.
- `tg_show_tables(text DEFAULT null, text DEFAULT null, text DEFAULT 'detail', boolean DEFAULT true)` is an extension function and returns `JSON`.
- `tg_show_transaction()` is an extension function and returns `cstring`.
- `tg_verify_tables(text DEFAULT null, text DEFAULT null, text DEFAULT null, text DEFAULT 'summary', boolean DEFAULT true)` is an extension function and returns `JSON`.
- `tsurugi_fdw_handler()` is an extension function and returns `fdw_handler`.
- `tsurugi_fdw_validator(options text[], catalog oid)` is an extension function and returns `void`.
- `tsurugi_fdw` is an extension-defined foreign data wrapper.

### Requirements and Caveats

- The reviewed control file declares default version `1.5.0`.
- The control file marks the extension as non-relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
