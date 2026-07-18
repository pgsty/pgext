## Usage

Sources:

- [Official extension control file](https://github.com/eva-ics/eva4/blob/f0e6801d2bb79d6651487766e3a3843fc50e2c2b/eva_pg/eva_pg.control)
- [Official Rust package manifest](https://github.com/eva-ics/eva4/blob/f0e6801d2bb79d6651487766e3a3843fc50e2c2b/eva_pg/Cargo.toml)
- [Official upstream README](https://github.com/eva-ics/eva4/blob/f0e6801d2bb79d6651487766e3a3843fc50e2c2b/README.md)

`eva_pg` — EVA ICS pgrx extension for the EVA industrial IoT platform.

The reviewed catalog snapshot records version `0.0.1`, kind `standard`, and implementation language `Rust`.
The curated compatibility set is `13,14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "eva_pg";
SELECT extversion
FROM pg_extension
WHERE extname = 'eva_pg';
```

The curated lifecycle is `active`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
