## Usage

Sources:

- [Official upstream README](https://github.com/replicase/pgcapture/blob/e9f4d88d4be1f12bddd72c86bbecfeddc9ea2c62/README.md)
- [Official extension control file (pgcapture.control)](https://github.com/replicase/pgcapture/blob/e9f4d88d4be1f12bddd72c86bbecfeddc9ea2c62/hack/postgres/extension/pgcapture.control)
- [Official extension SQL (pgcapture--0.1.sql)](https://github.com/replicase/pgcapture/blob/e9f4d88d4be1f12bddd72c86bbecfeddc9ea2c62/hack/postgres/extension/pgcapture--0.1.sql)

`pgcapture` — A scalable Netflix DBLog implementation for PostgreSQL. Use it when moving, transforming, or integrating the corresponding data from PostgreSQL. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION pgcapture;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `pgcapture.current_query()` is an extension function and returns `TEXT`.
- `pgcapture.log_ddl()` is an extension function and returns `event_trigger`.
- `pgcapture.sql_command_tags(p_sql TEXT)` is an extension function and returns `TEXT[]`.
- `pgcapture.ddl_logs` is a table installed or managed by the extension.
- `pgcapture.sources` is a table installed or managed by the extension.

### Requirements and Caveats

- The reviewed control file declares default version `0.1`.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
