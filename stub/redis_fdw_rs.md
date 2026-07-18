## Usage

Sources:

- [Official extension control file](https://github.com/isdaniel/redis_fdw_rs/blob/ecd495ac4db9ba6ae48ff4d778592885883f11d1/redis_fdw_rs.control)
- [Official upstream documentation](https://github.com/isdaniel/redis_fdw_rs/blob/ecd495ac4db9ba6ae48ff4d778592885883f11d1/README.md)
- [Official Rust package manifest](https://github.com/isdaniel/redis_fdw_rs/blob/ecd495ac4db9ba6ae48ff4d778592885883f11d1/Cargo.toml)

`redis_fdw_rs` — This extension allows PostgreSQL to directly query and manipulate Redis data as if it were regular PostgreSQL tables.

The reviewed catalog snapshot records version `0.6.0`, kind `standard`, and implementation language `Rust`.
The curated compatibility set is `14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "redis_fdw_rs";
SELECT extversion
FROM pg_extension
WHERE extname = 'redis_fdw_rs';
```

The curated lifecycle is `active`. Pin the reviewed build and verify maintenance status before adoption.
The official material contains an experimental, deprecated, unsupported, or explicit warning boundary; read it in full and test failure cases before non-lab use.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
