## Usage

Sources:

- [Official extension control file (hello_pg.control)](https://github.com/stevelauc/debug_macos_build_failure_pgrx/blob/98ffdcf0dc09d2f657c43f36f5d48778c4ae2665/hello_pg.control)
- [Official implementation source](https://github.com/stevelauc/debug_macos_build_failure_pgrx/blob/98ffdcf0dc09d2f657c43f36f5d48778c4ae2665/src/lib.rs)
- [Official Rust package manifest](https://github.com/stevelauc/debug_macos_build_failure_pgrx/blob/98ffdcf0dc09d2f657c43f36f5d48778c4ae2665/Cargo.toml)

`hello_pg` — Minimal pgrx hello-world extension used to reproduce a macOS build issue. Use it when SQL needs these specialized functions or aggregates. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION hello_pg;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `hello_hello_pg()` is an extension function.

### Requirements and Caveats

- The catalog records version `0.0.0`.
- The control file marks the extension as non-relocatable.
- The control file requires a superuser for installation.
- The control file marks the extension as not trusted.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
