## Usage

Sources:

- [Official extension control file](https://github.com/rixlhq/snowid-postgres/blob/9be325a349805a3c10ebcb37cd00f1e3a1a150d2/pg_snowid.control)
- [Official upstream documentation](https://github.com/rixlhq/snowid-postgres/blob/9be325a349805a3c10ebcb37cd00f1e3a1a150d2/README.md)
- [Official Rust package manifest](https://github.com/rixlhq/snowid-postgres/blob/9be325a349805a3c10ebcb37cd00f1e3a1a150d2/Cargo.toml)

`pg_snowid` — Generate thread-safe, time-sorted Snowflake-like IDs in PostgreSQL

The reviewed catalog snapshot records version `3.1.0`, kind `preload`, and implementation language `Rust`.
The curated compatibility set is `18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pg_snowid";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_snowid';
```

The upstream project is associated with `Rixl`; verify its current support, license, packaging, and deployment boundary from the linked source.

The curated lifecycle is `active`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
