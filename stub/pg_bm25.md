## Usage

Sources:

- [Official extension control file](https://github.com/devflowinc/pg_bm25_fork/blob/be44998446d1095e67970d42141cb2874c945289/pg_bm25/pg_bm25.control)
- [Official upstream documentation](https://github.com/devflowinc/pg_bm25_fork/blob/be44998446d1095e67970d42141cb2874c945289/README.md)
- [Official Rust package manifest](https://github.com/devflowinc/pg_bm25_fork/blob/be44998446d1095e67970d42141cb2874c945289/pg_bm25/Cargo.toml)

`pg_bm25` — Full text search for PostgreSQL using BM25

The reviewed catalog snapshot records version `0.0.0`, kind `standard`, and implementation language `Rust`.
Install and validate the declared extension dependencies first: `plpgsql`.
The curated compatibility set is `12,13,14,15,16`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pg_bm25";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_bm25';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
