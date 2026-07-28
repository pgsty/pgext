## Usage

Sources:

- [Official upstream README](https://github.com/tsirysndr/xata_id_extension/blob/bf2071f81c25d29e67b11d821a1f81300f78a6fc/README.md)
- [Official extension control file (xata_id_extension.control)](https://github.com/tsirysndr/xata_id_extension/blob/bf2071f81c25d29e67b11d821a1f81300f78a6fc/xata_id_extension.control)
- [Official implementation source](https://github.com/tsirysndr/xata_id_extension/blob/bf2071f81c25d29e67b11d821a1f81300f78a6fc/src/lib.rs)

`xata_id_extension` — Generates 24-character unique IDs with a rec_ prefix and 20 random characters (a-z0-9). Use it when SQL needs these specialized functions or aggregates. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION xata_id_extension;
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
