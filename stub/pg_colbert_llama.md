## Usage

Sources:

- [Official upstream README](https://github.com/mayflower/pgturbohybrid/blob/cd68155df20545b163dd610fb95e8aa7e62fc108/extensions/pg_colbert_llama/README.md)
- [Official extension control file (pg_colbert_llama.control)](https://github.com/mayflower/pgturbohybrid/blob/cd68155df20545b163dd610fb95e8aa7e62fc108/extensions/pg_colbert_llama/pg_colbert_llama.control)
- [Official extension SQL (pg_colbert_llama--0.1.0.sql)](https://github.com/mayflower/pgturbohybrid/blob/cd68155df20545b163dd610fb95e8aa7e62fc108/extensions/pg_colbert_llama/sql/pg_colbert_llama--0.1.0.sql)

`pg_colbert_llama` — llama_embed is the SQL-facing companion extension for running llama.cpp embedding models inside PostgreSQL. It can return dense vector embeddings, token-level vector[] embeddings, or pgturbohybrid multivectors for ColBERT and other late-interaction models. Use it for the corresponding vector, model, or retrieval workflow. Upstream describes this capability as experimental.

### Core Workflow

```sql
CREATE EXTENSION pg_colbert_llama;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `colbert(model pg_catalog.text, input pg_catalog.text)` is an extension function and returns `pg_catalog`.
- `colbert_debug(model pg_catalog.text, input pg_catalog.text)` is an extension function and returns `pg_catalog`.
- `colbert_dim(model pg_catalog.text, input pg_catalog.text)` is an extension function and returns `pg_catalog`.
- `colbert_float4(model pg_catalog.text, input pg_catalog.text)` is an extension function and returns `pg_catalog`.
- `colbert_model_info(model pg_catalog.text)` is an extension function and returns `pg_catalog`.
- `colbert_mv(model pg_catalog.text, input pg_catalog.text)` is an extension function and returns `turbohybrid_multivector`.
- `colbert_mv_batch(model pg_catalog.text, inputs pg_catalog.text[])` is an extension function and returns `turbohybrid_multivector[]`.
- `colbert_vectors(model pg_catalog.text, input pg_catalog.text)` is an extension function and returns `vector[]`.

### Requirements and Caveats

- The reviewed control file declares default version `0.1.0`.
- Install the confirmed extension dependencies first: `pgturbohybrid_experimental`.
- The control file marks the extension as relocatable.
- Upstream labels part or all of the project experimental.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
