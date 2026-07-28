## Usage

Sources:

- [Official upstream README](https://github.com/mayflower/pgturbohybrid/blob/cd68155df20545b163dd610fb95e8aa7e62fc108/extensions/pg_colbert_llama/README.md)
- [Official extension control file (llama_embed.control)](https://github.com/mayflower/pgturbohybrid/blob/cd68155df20545b163dd610fb95e8aa7e62fc108/extensions/pg_colbert_llama/llama_embed.control)
- [Official extension SQL (llama_embed--0.1.0.sql)](https://github.com/mayflower/pgturbohybrid/blob/cd68155df20545b163dd610fb95e8aa7e62fc108/extensions/pg_colbert_llama/sql/llama_embed--0.1.0.sql)

`llama_embed` — llama_embed is the SQL-facing companion extension for running llama.cpp embedding models inside PostgreSQL. It can return dense vector embeddings, token-level vector[] embeddings, or pgturbohybrid multivectors for ColBERT and other late-interaction models. Use it for the corresponding vector, model, or retrieval workflow. Upstream describes this capability as experimental.

### Core Workflow

```sql
CREATE EXTENSION llama_embed;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `llama_embed` is an extension function.
- `llama_embed_model_info(model pg_catalog.text)` is an extension function and returns `pg_catalog`.
- `llama_embed_mv` is an extension function.
- `llama_embed_mv_batch` is an extension function.
- `llama_embed_sparse` is an extension function.
- `llama_embed_sparse_batch` is an extension function.
- `llama_embed_sparse_model_info` is an extension function.
- `llama_embed_tokens` is an extension function.
- `llama_embed_vector` is an extension function.
- `llama_embed_vector_batch` is an extension function.

### Requirements and Caveats

- The reviewed control file declares default version `0.1.0`.
- Install the confirmed extension dependencies first: `pgturbohybrid_experimental`.
- The control file marks the extension as relocatable.
- Upstream labels part or all of the project experimental.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
