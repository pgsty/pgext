## Usage

Sources:

- [Official upstream README](https://github.com/unknowntpo/playground-2022/blob/c2e4935c7a575bec01d6d5301d027f6adf801a62/README.md)
- [Official extension control file (plannerlog.control)](https://github.com/unknowntpo/playground-2022/blob/c2e4935c7a575bec01d6d5301d027f6adf801a62/rust/pgx/plannerlog/plannerlog.control)
- [Official implementation source](https://github.com/unknowntpo/playground-2022/blob/c2e4935c7a575bec01d6d5301d027f6adf801a62/rust/pgx/plannerlog/src/lib.rs)

`plannerlog` — pgrx planner-hook example that logs each planned query. Use it when collecting or interpreting the corresponding PostgreSQL statistics. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION plannerlog;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `hello_hello_world()` is an extension function.

### Requirements and Caveats

- The catalog records version `0.0.0`.
- The control file marks the extension as non-relocatable.
- The control file does not require superuser-only installation.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
