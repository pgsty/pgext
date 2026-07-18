## Usage

Sources:

- [Official extension control file](https://github.com/vee-huyvunguyen/postgres_rust_extension/blob/667aaa420473f1b3b89507152dac5e7f4240163e/date_string_converter/date_string_converter.control)
- [Official upstream documentation](https://github.com/vee-huyvunguyen/postgres_rust_extension/blob/667aaa420473f1b3b89507152dac5e7f4240163e/README.md)
- [Official Rust package manifest](https://github.com/vee-huyvunguyen/postgres_rust_extension/blob/667aaa420473f1b3b89507152dac5e7f4240163e/date_string_converter/Cargo.toml)

`date_string_converter` — pgrx extension skeleton for date-string conversion; current source exposes hello_date_string_converter().

The reviewed catalog snapshot records version `0.0.0`, kind `standard`, and implementation language `Rust`.
The curated compatibility set is `11,12,13,14,15,16`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "date_string_converter";
SELECT extversion
FROM pg_extension
WHERE extname = 'date_string_converter';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
