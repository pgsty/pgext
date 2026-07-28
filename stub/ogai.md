## Usage

Sources:

- [Official upstream README](https://github.com/opengauss-mirror/opengauss-server/blob/7508e682a7348c515c63244da8be4800b54b90cf/contrib/README)
- [Official extension control file (ogai.control)](https://github.com/opengauss-mirror/opengauss-server/blob/7508e682a7348c515c63244da8be4800b54b90cf/contrib/ogai/ogai.control)
- [Official extension SQL (ogai--1.0.sql)](https://github.com/opengauss-mirror/opengauss-server/blob/7508e682a7348c515c63244da8be4800b54b90cf/contrib/ogai/ogai--1.0.sql)

`ogai` — OGAI - AI vectorization and search framework for openGauss. Use it for the corresponding vector, model, or retrieval workflow. Upstream describes this capability as experimental.

### Core Workflow

```sql
CREATE EXTENSION ogai;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `ogai.ai_unvectorize(p_task_name TEXT)` is an extension function and returns `TABLE`.
- `ogai.ai_vectorize(p_task_name TEXT, p_task_type TEXT, p_index_type TEXT, p_embed_model TEXT, p_src_schema TEXT, p_src_table TEXT, p_src_col TEXT, p_primary_key TEXT, p_table_method TEXT, p_dim INTEGER, p_max_chunk_size INTEGER DEFAULT 1000, p_max_chunk_overlap INTEGER DEFAULT 2…)` is an extension function and returns `TABLE`.
- `ogai.encrypt_api_key_trigger()` is an extension function and returns `TRIGGER`.
- `ogai.hybrid_search(p_task_name TEXT, p_query TEXT, p_return_cols TEXT DEFAULT '', p_limit INTEGER DEFAULT 10, p_where_clause TEXT DEFAULT '')` is an extension function and returns `TABLE`.
- `ogai.rag(p_user_question TEXT, p_task_name TEXT, p_reranker_model TEXT, p_chat_model TEXT, p_rerank_limit INTEGER DEFAULT 5, p_search_limit INTEGER DEFAULT 20)` is an extension function and returns `TEXT`.
- `ogai.search(p_task_name TEXT, p_query TEXT, p_return_cols TEXT DEFAULT '', p_limit INTEGER DEFAULT 10, p_where_clause TEXT DEFAULT '')` is an extension function and returns `TABLE`.
- `ogai.vectorize_param_trigger()` is an extension function and returns `TRIGGER`.
- `ogai.vectorize_trigger_handle(p_content TEXT, p_embed_model TEXT, p_dim INTEGER, p_task_name TEXT, p_task_type TEXT, p_src_schema TEXT, p_src_table TEXT, p_primary_key TEXT, p_table_method TEXT, p_operation TEXT, p_pk_value TEXT, p_max_chunk_size INTEGER DEFAULT 1000, p_max_chunk_overlap I…)` is an extension function and returns `VOID`.
- `ogai.model_provider_type` is an extension-defined type.
- `ogai.queue_status` is an extension-defined type.
- `ogai.table_method` is an extension-defined type.
- `ogai.task_type` is an extension-defined type.
- `ogai.model_sources` is a table installed or managed by the extension.
- `ogai.vectorize_queue` is a table installed or managed by the extension.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- The control file marks the extension as non-relocatable.
- Upstream labels part or all of the project experimental.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
