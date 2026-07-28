## Usage

Sources:

- [Official upstream README](https://github.com/tmaxopensql/tibero-fdw/blob/49d2260e9f9228471ca318ada0f7630628dca974/README.md)
- [Official extension control file (tibero_fdw.control)](https://github.com/tmaxopensql/tibero-fdw/blob/49d2260e9f9228471ca318ada0f7630628dca974/tibero_fdw.control)
- [Official extension SQL (tibero_fdw--1.0.sql)](https://github.com/tmaxopensql/tibero-fdw/blob/49d2260e9f9228471ca318ada0f7630628dca974/tibero_fdw--1.0.sql)

`tibero_fdw` — This PostgreSQL extension implements a Foreign Data Wrapper (FDW) for Tibero. Use it when PostgreSQL must access the corresponding external data source through a foreign-data interface. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION tibero_fdw;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `tibero_fdw_handler()` is an extension function and returns `fdw_handler`.
- `tibero_fdw_validator(text[], oid)` is an extension function and returns `void`.
- `tibero_fdw` is an extension-defined foreign data wrapper.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
