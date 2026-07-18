## Usage

Sources:

- [Official extension control file](https://github.com/fabianlindfors/pgfdb/blob/598aca8937e5c61c805592c90719baaa5e1a4d73/pgfdb.control)

`pgfdb` — Experimental pgrx extension that stores PostgreSQL table and index data in FoundationDB for distributed, fault-tolerant, horizontally scalable operation.

The reviewed catalog snapshot records version `0.0.2`, kind `standard`, and implementation language `Rust`.
The curated compatibility set is `13,14,15,16,17`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pgfdb";
SELECT extversion
FROM pg_extension
WHERE extname = 'pgfdb';
```

The curated lifecycle is `active`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
