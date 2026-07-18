## Usage

Sources:

- [Official extension control file](https://github.com/messagebird/clickhouse-postgres-fdw/blob/7ff697b2142fa3c52d7aa5ceb170f4b8d9e974b3/ch_fdw.control)
- [Official upstream README](https://github.com/messagebird/clickhouse-postgres-fdw/blob/7ff697b2142fa3c52d7aa5ceb170f4b8d9e974b3/README.md)

`ch_fdw` — Foreign data wrapper for federated access from PostgreSQL to ClickHouse.

The reviewed catalog snapshot records version `1.0`, kind `standard`, and implementation language `C`.
The curated compatibility set is `13`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "ch_fdw";
SELECT extversion
FROM pg_extension
WHERE extname = 'ch_fdw';
```

The curated lifecycle is `archived`. Pin the reviewed build and verify maintenance status before adoption.
The official material contains an experimental, deprecated, unsupported, or explicit warning boundary; read it in full and test failure cases before non-lab use.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
