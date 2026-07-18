## Usage

Sources:

- [Official extension control file](https://github.com/foxflow/pg_zen_engine/blob/e4924af4f5a0ad7da3d0031ee5e54ef6698db0b8/pg_zen_engine.control)

`pg_zen_engine` — Evaluate JSON Decision Model graphs against PostgreSQL JSONB data with Zen Engine

The reviewed catalog snapshot records version `0.0.1`, kind `standard`, and implementation language `Rust`.
The curated compatibility set is `12,13,14,15,16`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pg_zen_engine";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_zen_engine';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
