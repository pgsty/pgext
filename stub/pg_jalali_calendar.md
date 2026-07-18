## Usage

Sources:

- [Official Rust package manifest](https://github.com/alisol911/pg_jalali_calendar/blob/968267cb54de153c24169a2b9fb0ec5a3dac72f1/Cargo.toml)

`pg_jalali_calendar` — Jalali/Persian calendar date conversion and helper functions for PostgreSQL.

The reviewed catalog snapshot records version `0.1.0`, kind `standard`, and implementation language `Rust`.
The curated compatibility set is `18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pg_jalali_calendar";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_jalali_calendar';
```

The curated lifecycle is `active`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
