## Usage

Sources:

- [Official upstream README](https://github.com/robertmu/pg-lakebase/blob/a5baec33934b069b0832644ff9cee64b429c14cb/README.md)
- [Official extension control file (pg_lakebase_runtime.control)](https://github.com/robertmu/pg-lakebase/blob/a5baec33934b069b0832644ff9cee64b429c14cb/pg-lakebase-runtime/pg_lakebase_runtime.control)
- [Official implementation source](https://github.com/robertmu/pg-lakebase/blob/a5baec33934b069b0832644ff9cee64b429c14cb/pg-lakebase-runtime/src/lib.rs)

`pg_lakebase_runtime` — Shared runtime for Lakebase table access method extensions. Use it for the corresponding analytical or storage workflow. Upstream describes this capability as experimental.

### Core Workflow

```sql
CREATE EXTENSION pg_lakebase_runtime;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `deregister_worker` is an extension function.
- `maintenance_worker` is an extension function.
- `observe_object_tree` is an extension function.
- `register_worker_impl` is an extension function.
- `request_worker_wakeup` is an extension function.
- `retry_maintenance_item` is an extension function.

### Requirements and Caveats

- The catalog records version `0.1.0`.
- The control file marks the extension as non-relocatable.
- The control file requires a superuser for installation.
- Upstream labels part or all of the project experimental.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
