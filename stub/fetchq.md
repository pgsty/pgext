## Usage

Sources:

- [Official upstream README](https://github.com/fetchq/pg-extension/blob/b9c3f62e226401c94635709fec32e48fd85f754a/README.md)
- [Official extension control file (fetchq.control)](https://github.com/fetchq/pg-extension/blob/b9c3f62e226401c94635709fec32e48fd85f754a/src/fetchq.control)

`fetchq` — Postgres extension that enables FetchQ capabilities. Use it when an application needs this specific database capability. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION fetchq;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Requirements and Caveats

- The reviewed control file declares default version `4.0.2`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
