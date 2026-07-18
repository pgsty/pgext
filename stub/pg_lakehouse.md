## Usage

Sources:

- [Official upstream README](https://github.com/paradedb/paradedb/blob/59a9224db5b22893cd308993bcddfbfdbc14cfc1/README.md)

`pg_lakehouse` — An analytical query engine for PostgreSQL

The reviewed catalog snapshot records version `0.7.6`, kind `preload`, and implementation language `Rust`.
The curated compatibility set is `14,15,16`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pg_lakehouse";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_lakehouse';
```

The upstream project is associated with `ParadeDB`; verify its current support, license, packaging, and deployment boundary from the linked source.

The curated lifecycle is `deprecated`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
