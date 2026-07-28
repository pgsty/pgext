## Usage

Sources:

- [Official upstream README](https://github.com/timescale/pg_traceam/blob/a8aafe6448e6fa110f283dafb0b201c846857e57/README.md)
- [Official extension control file (traceam.control)](https://github.com/timescale/pg_traceam/blob/a8aafe6448e6fa110f283dafb0b201c846857e57/traceam.control)
- [Official extension SQL (traceam--0.1.sql)](https://github.com/timescale/pg_traceam/blob/a8aafe6448e6fa110f283dafb0b201c846857e57/traceam--0.1.sql)

`traceam` — Once the extension is installed it is necessary to install it in the database as well, which is done using CREATE EXTENSION:. Use it when an application needs this specific database capability. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION traceam;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `traceam_handler(internal)` is an extension function and returns `table_am_handler`.
- `traceam` is a schema created by the extension.
- `traceam` is an extension-defined access method.

### Requirements and Caveats

- The reviewed control file declares default version `0.1`.
- The control file marks the extension as non-relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
