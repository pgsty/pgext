## Usage

Sources:

- [Official upstream README](https://github.com/unknowntpo/playground-2022/blob/c2e4935c7a575bec01d6d5301d027f6adf801a62/README.md)
- [Official extension control file (exploring_aggregates.control)](https://github.com/unknowntpo/playground-2022/blob/c2e4935c7a575bec01d6d5301d027f6adf801a62/rust/pgx/exploring_aggregates/exploring_aggregates.control)
- [Official implementation source](https://github.com/unknowntpo/playground-2022/blob/c2e4935c7a575bec01d6d5301d027f6adf801a62/rust/pgx/exploring_aggregates/src/lib.rs)

`exploring_aggregates` — pgrx example defining a custom integer-sum aggregate with serialized state. Use it when SQL needs these specialized functions or aggregates. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION exploring_aggregates;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Requirements and Caveats

- The catalog records version `0.0.0`.
- The control file marks the extension as non-relocatable.
- The control file does not require superuser-only installation.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
