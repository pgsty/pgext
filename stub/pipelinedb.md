## Usage

Sources:

- [Official project or provider page](https://www.pipelinedb.com)

`pipelinedb` — Run continuous SQL queries for high-throughput incremental time-series aggregation in PostgreSQL

The reviewed catalog snapshot records version `1.1.0`, kind `preload`, and implementation language `C`.
The curated compatibility set is `10,11`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pipelinedb";
SELECT extversion
FROM pg_extension
WHERE extname = 'pipelinedb';
```

The curated lifecycle is `deprecated`. Pin the reviewed build and verify maintenance status before adoption.
The official material contains an experimental, deprecated, unsupported, or explicit warning boundary; read it in full and test failure cases before non-lab use.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
