## Usage

Sources:

- [Official upstream README](https://github.com/shaifimran/pg_embed/blob/0934d0b68caeb32813a3b661859f5552f6595287/readme.md)
- [Official extension control file (pg_embed.control)](https://github.com/shaifimran/pg_embed/blob/0934d0b68caeb32813a3b661859f5552f6595287/pg_embed.control)
- [Official extension SQL (pg_embed--1.0.sql)](https://github.com/shaifimran/pg_embed/blob/0934d0b68caeb32813a3b661859f5552f6595287/pg_embed--1.0.sql)

`pg_embed` — pg_embed is a PostgreSQL extension written in PL/Python that enables you to generate and store text embeddings directly in your database using HuggingFace's Inference API. This extension is ideal for adding semantic search, similarity, and AI-powered features to your PostgreSQL-backed applications. Use it for the corresponding vector, model, or retrieval workflow. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION pg_embed;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `generate_and_store_embeddings(tbl_name TEXT, txt_col TEXT, hf_api_key TEXT)` is an extension function and returns `VOID`.
- `get_embedding(input_text TEXT, hf_api_key TEXT)` is an extension function and returns `FLOAT8[]`.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
