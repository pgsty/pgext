## Usage

Sources:

- [Official extension control file (talib_ext.control)](https://github.com/sachaarbonel/pg_talib-rs/blob/fcc0ce3cb1475d8f12be7f2421f19b2e4e89384b/talib_ext/talib_ext.control)
- [Official implementation source](https://github.com/sachaarbonel/pg_talib-rs/blob/fcc0ce3cb1475d8f12be7f2421f19b2e4e89384b/talib_ext/src/lib.rs)
- [Official Rust package manifest](https://github.com/sachaarbonel/pg_talib-rs/blob/fcc0ce3cb1475d8f12be7f2421f19b2e4e89384b/talib_ext/Cargo.toml)

`talib_ext` — postgres extension to do technical analysis on postgres. Use it when SQL needs these specialized functions or aggregates. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION talib_ext;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `hello_talib_ext()` is an extension function.
- `macd` is an extension function.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- The control file marks the extension as non-relocatable.
- The control file does not require superuser-only installation.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
