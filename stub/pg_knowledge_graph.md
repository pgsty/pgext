## Usage

Sources:

- [Official extension control file](https://github.com/hiyenwong/pg_knowledge_graph/blob/904e15c6500e23897dc626338b6469d639270d76/pg_knowledge_graph.control)
- [Official upstream documentation](https://github.com/hiyenwong/pg_knowledge_graph/blob/904e15c6500e23897dc626338b6469d639270d76/README.md)
- [Official Rust package manifest](https://github.com/hiyenwong/pg_knowledge_graph/blob/904e15c6500e23897dc626338b6469d639270d76/Cargo.toml)

`pg_knowledge_graph` — Knowledge graph extension for PostgreSQL with graph algorithms and pgvector integration.

The reviewed catalog snapshot records version `0.1.0`, kind `standard`, and implementation language `Rust`.
Install and validate the declared extension dependencies first: `vector`.
The curated compatibility set is `16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pg_knowledge_graph";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_knowledge_graph';
```

The curated lifecycle is `active`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
