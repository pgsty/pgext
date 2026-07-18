## Usage

Sources:

- [Official extension control file](https://github.com/postgrespro/aqo/blob/4a7fcd7291a8c4209c9102d1736f9f754d311cf2/aqo.control)
- [Official upstream documentation](https://github.com/postgrespro/aqo/blob/4a7fcd7291a8c4209c9102d1736f9f754d311cf2/README.md)

`aqo` — Adaptive query optimization extension using execution statistics to improve cardinality estimation.

The reviewed catalog snapshot records version `1.6`, kind `preload`, and implementation language `C`.
The curated compatibility set is `10,11,12,13,14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "aqo";
SELECT extversion
FROM pg_extension
WHERE extname = 'aqo';
```

The curated lifecycle is `active`. Pin the reviewed build and verify maintenance status before adoption.
The official material contains an experimental, deprecated, unsupported, or explicit warning boundary; read it in full and test failure cases before non-lab use.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
