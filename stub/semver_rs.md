## Usage

Sources:

- [Official upstream README](https://github.com/fabmation-gmbh/pg_semver-rs/blob/0bc1aa00db74b824852027cbc6a369a5ccbd3f10/README.md)
- [Official extension control file (semver_rs.control)](https://github.com/fabmation-gmbh/pg_semver-rs/blob/0bc1aa00db74b824852027cbc6a369a5ccbd3f10/semver_rs.control)
- [Official implementation source](https://github.com/fabmation-gmbh/pg_semver-rs/blob/0bc1aa00db74b824852027cbc6a369a5ccbd3f10/src/lib.rs)

`semver_rs` — The awesome [pg-semver][] extension but implemented in Rust. It supports all the operations [pg-semver][] does. Use it when application data needs this type, domain, or its operators. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION semver_rs;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Requirements and Caveats

- The catalog records version `0.1.0`.
- The control file marks the extension as non-relocatable.
- The control file does not require superuser-only installation.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
