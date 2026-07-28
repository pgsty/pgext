## Usage

Sources:

- [Official upstream README](https://github.com/kurtbuilds/sid/blob/cfd88358d24675b49cf6907e42397f06264a3809/README.md)
- [Official extension control file (sid.control)](https://github.com/kurtbuilds/sid/blob/cfd88358d24675b49cf6907e42397f06264a3809/pg/sid.control)
- [Official implementation source](https://github.com/kurtbuilds/sid/blob/cfd88358d24675b49cf6907e42397f06264a3809/pg/src/lib.rs)

`sid` — Sortable labeled 128-bit identifier type with compact human-readable encoding. Use it when application data needs this type, domain, or its operators. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION sid;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `sid_from_uuid` is an extension function.
- `sid_new` is an extension function.
- `sid_null` is an extension function.

### Requirements and Caveats

- The catalog records version `0.0.0`.
- The control file marks the extension as non-relocatable.
- The control file requires a superuser for installation.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
