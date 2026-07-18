## Usage

Sources:

- [Official PGXN distribution page](https://pgxn.org/dist/roleman/)

`roleman` — Parameterized role and privilege management functions for PostgreSQL.

The reviewed catalog snapshot records version `0.3.0`, kind `puresql`, and implementation language `SQL`.
Install and validate the declared extension dependencies first: `plpgsql`.
The curated compatibility set is `10,11,12,13,14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "roleman";
SELECT extversion
FROM pg_extension
WHERE extname = 'roleman';
```

The curated lifecycle is `archived`. Pin the reviewed build and verify maintenance status before adoption.
The official material contains an experimental, deprecated, unsupported, or explicit warning boundary; read it in full and test failure cases before non-lab use.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
