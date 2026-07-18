## Usage

Sources:

- [Official extension control file](https://github.com/neurondb/neurondb/blob/66284d0476a864bbe074017b517a34daef454fba/neurondb.control)
- [Official upstream documentation](https://www.neurondb.ai/docs)
- [Official project or provider page](https://www.neurondb.ai)

`neurondb` — AI database extension for vector search, ML inference, hybrid search, and RAG in PostgreSQL.

The reviewed catalog snapshot records version `4.0.0-devel`, kind `preload`, and implementation language `C`.
The curated compatibility set is `16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "neurondb";
SELECT extversion
FROM pg_extension
WHERE extname = 'neurondb';
```

The upstream project is associated with `NeuronDB`; verify its current support, license, packaging, and deployment boundary from the linked source.

The curated lifecycle is `active`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
