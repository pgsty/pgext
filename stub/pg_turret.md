## Usage

Sources:

- [Official extension control file](https://github.com/lasect/pg_turret/blob/5c7407afd7e3268a7d8583c643a4c9554a6a2c42/pg_turret.control)
- [Official upstream documentation](https://github.com/lasect/pg_turret/blob/5c7407afd7e3268a7d8583c643a4c9554a6a2c42/README.md)
- [Official Rust package manifest](https://github.com/lasect/pg_turret/blob/5c7407afd7e3268a7d8583c643a4c9554a6a2c42/Cargo.toml)

`pg_turret` — Capture PostgreSQL logs and export structured events to HTTP, Kafka, or Sentry

The reviewed catalog snapshot records version `0.0.0`, kind `preload`, and implementation language `Rust`.
The curated compatibility set is `13,14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pg_turret";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_turret';
```

The curated lifecycle is `active`. Pin the reviewed build and verify maintenance status before adoption.
The official material contains an experimental, deprecated, unsupported, or explicit warning boundary; read it in full and test failure cases before non-lab use.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
