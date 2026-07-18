## Usage

Sources:

- [Official upstream documentation](https://github.com/citusdata/pg_shard/blob/0014e199c5723020f3e0c982882a10ca53694cad/README.md)
- [Official PGXN distribution page](https://pgxn.org/dist/pg_shard/)
- [Official project or provider page](https://github.com/citusdata/citus)

`pg_shard` — Shard and replicate PostgreSQL tables across remote servers

The reviewed catalog snapshot records version `1.2`, kind `preload`, and implementation language `C`.

```sql
CREATE EXTENSION "pg_shard";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_shard';
```

The upstream project is associated with `Citus Data`; verify its current support, license, packaging, and deployment boundary from the linked source.

The curated lifecycle is `deprecated`. Pin the reviewed build and verify maintenance status before adoption.
The official material contains an experimental, deprecated, unsupported, or explicit warning boundary; read it in full and test failure cases before non-lab use.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
