## Usage

Sources:

- [Official Rust package manifest](https://github.com/lmwnshn/hypostats/blob/0181b25d84430d7d7cfc1971b1c94c270cfebbb4/Cargo.toml)

`hypostats` — Dump and load pg_statistic rows as JSON for statistics experiments.

The reviewed catalog snapshot records version `0.0.0`, kind `standard`, and implementation language `Rust`.
The curated compatibility set is `12,13,14,15,16,17`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "hypostats";
SELECT extversion
FROM pg_extension
WHERE extname = 'hypostats';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
