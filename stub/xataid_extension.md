## Usage

Sources:

- [Official upstream README](https://github.com/tsirysndr/rocksky/blob/9260ce846604f72431f0efb8ee04657174b49249/README.md)
- [Official extension control file (xataid_extension.control)](https://github.com/tsirysndr/rocksky/blob/9260ce846604f72431f0efb8ee04657174b49249/crates/xataid-extension/xataid_extension.control)
- [Official implementation source](https://github.com/tsirysndr/rocksky/blob/9260ce846604f72431f0efb8ee04657174b49249/crates/xataid-extension/src/lib.rs)

`xataid_extension` — Generate Xata-style rec_ identifiers with a 20-character lowercase alphanumeric suffix. Use it when SQL needs these specialized functions or aggregates. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION xataid_extension;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `xata_id()` is an extension function.

### Requirements and Caveats

- The catalog records version `0.0.0`.
- The control file marks the extension as non-relocatable.
- The control file requires a superuser for installation.
- The control file marks the extension as not trusted.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
