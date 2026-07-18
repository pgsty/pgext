## Usage

Sources:

- [Official upstream documentation](https://neon.com/docs/extensions/neon)
- [Official project or provider page](https://neon.com)
- [Official upstream README](https://github.com/neondatabase/neon/blob/8f60b04da47ffefe0e52bda2440134b42874eb75/README.md)

`neon` — Neon-specific statistics and Local File Cache metrics for Neon Postgres computes.

The reviewed catalog snapshot records version `1.9`, kind `preload`, and implementation language `C`.
The curated compatibility set is `14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "neon";
SELECT extversion
FROM pg_extension
WHERE extname = 'neon';
```

The upstream project is associated with `Neon`; verify its current support, license, packaging, and deployment boundary from the linked source.

The curated lifecycle is `active`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
