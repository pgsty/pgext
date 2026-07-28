## Usage

Sources:

- [Official upstream README](https://github.com/emielm/gofdw/blob/1f3b8256a0b09f49d017cf580b4ccb6f301f409b/README.md)
- [Official extension control file (gofdw.control)](https://github.com/emielm/gofdw/blob/1f3b8256a0b09f49d017cf580b4ccb6f301f409b/gofdw.control)
- [Official extension SQL (gofdw.sql)](https://github.com/emielm/gofdw/blob/1f3b8256a0b09f49d017cf580b4ccb6f301f409b/gofdw.sql)

`gofdw` — A Postgres Foreign Data Wrapper implemented in go using cgo. Very experimental. Use it when PostgreSQL must access the corresponding external data source through a foreign-data interface. Upstream describes this capability as experimental.

### Core Workflow

```sql
CREATE EXTENSION gofdw;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `gofdw_handler()` is an extension function and returns `fdw_handler`.
- `gofdw_validator(text[], oid)` is an extension function.
- `gofdw` is an extension-defined foreign data wrapper.

### Requirements and Caveats

- The reviewed control file declares default version `0.1`.
- The control file marks the extension as relocatable.
- Upstream labels part or all of the project experimental.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
