## Usage

Sources:

- [Official upstream README](https://github.com/stephen-hilton/ugly.network/blob/e8297ceb5be82a8f662a4c8cb2fc993901f30aaa/old/README.md)
- [Official extension control file (helloworldtable.control)](https://github.com/stephen-hilton/ugly.network/blob/e8297ceb5be82a8f662a4c8cb2fc993901f30aaa/helloworldtable/helloworldtable.control)
- [Official implementation source](https://github.com/stephen-hilton/ugly.network/blob/e8297ceb5be82a8f662a4c8cb2fc993901f30aaa/helloworldtable/src/lib.rs)

`helloworldtable` — pgrx example returning a one-row table containing a Hello World greeting. Use it when SQL needs these specialized functions or aggregates. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION helloworldtable;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `helloworldtable()` is an extension function.

### Requirements and Caveats

- The catalog records version `0.0.0`.
- The control file marks the extension as non-relocatable.
- The control file requires a superuser for installation.
- The control file marks the extension as not trusted.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
