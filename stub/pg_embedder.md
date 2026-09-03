## Usage

Sources:

- [Official documentation](https://github.com/riclab/pg_embedder/blob/230cfc5f1785975d5e2fe78c2f2646be953c649a/README.md)
- [Extension control file](https://github.com/riclab/pg_embedder/blob/230cfc5f1785975d5e2fe78c2f2646be953c649a/pg_embedder.control)
- [Build manifest](https://github.com/riclab/pg_embedder/blob/230cfc5f1785975d5e2fe78c2f2646be953c649a/Cargo.toml)

`pg_embedder` In-backend text embedding inference using Rust, Candle, and packaged models.

### Enablement

Install the files for the intended server, then create `pg_embedder` in the target database:

```sql
CREATE EXTENSION pg_embedder;
```

### Core Workflow

The following example is taken from the reviewed upstream documentation. Adapt object names, paths, credentials, and workload values before use.

```sql
-- index side
UPDATE docs SET embedding = embed_encode(content);

-- search side, correct for every model
SELECT id, content, cosine_similarity(embedding, embed_query('how does vector search work')) AS score
FROM docs
ORDER BY score DESC
LIMIT 5;
```

### Main Objects

The reviewed install surface includes these prominent objects:

| Object | Kind | Operational role |
|---|---|---|
| `embed_encode` | FUNCTION | Callable function from the reviewed install surface. |
| `embed_text` | FUNCTION | Callable function from the reviewed install surface. |
| `docs` | TABLE | Extension-owned table; account for its data in backup and upgrades. |
| `cosine_similarity` | FUNCTION | Callable function from the reviewed install surface. |
| `embed_init` | FUNCTION | Callable function from the reviewed install surface. |
| `embed_model` | FUNCTION | Callable function from the reviewed install surface. |
| `embed_query` | FUNCTION | Callable function from the reviewed install surface. |
| `embed_info` | FUNCTION | Callable function from the reviewed install surface. |

### Operations and Boundaries

- The verified PostgreSQL-major evidence covers 13, 14, 15, 16, 17, 18; do not infer unlisted majors.
