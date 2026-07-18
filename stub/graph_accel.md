## Usage

Sources:

- [Official extension control file](https://github.com/aaronsb/knowledge-graph-system/blob/b6ebe037b5d6d89921fa62e09bc8c0a54394aa18/graph-accel/ext/graph_accel.control)
- [Official upstream documentation](https://github.com/aaronsb/knowledge-graph-system/blob/b6ebe037b5d6d89921fa62e09bc8c0a54394aa18/graph-accel/README.md)
- [Official Rust package manifest](https://github.com/aaronsb/knowledge-graph-system/blob/b6ebe037b5d6d89921fa62e09bc8c0a54394aa18/graph-accel/ext/Cargo.toml)

`graph_accel` — In-memory graph traversal acceleration extension for Apache AGE.

The reviewed catalog snapshot records version `0.5.0`, kind `standard`, and implementation language `Rust`.
Install and validate the declared extension dependencies first: `age`.
The curated compatibility set is `13,14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "graph_accel";
SELECT extversion
FROM pg_extension
WHERE extname = 'graph_accel';
```

The curated lifecycle is `active`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
