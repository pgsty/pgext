## 用法

来源：

- [官方上游 README](https://github.com/opengauss-mirror/opengauss-server/blob/7508e682a7348c515c63244da8be4800b54b90cf/contrib/README)
- [官方扩展控制文件 (ogai.control)](https://github.com/opengauss-mirror/opengauss-server/blob/7508e682a7348c515c63244da8be4800b54b90cf/contrib/ogai/ogai.control)
- [官方扩展 SQL (ogai--1.0.sql)](https://github.com/opengauss-mirror/opengauss-server/blob/7508e682a7348c515c63244da8be4800b54b90cf/contrib/ogai/ogai--1.0.sql)

`ogai` — OGAI - 用于 openGauss 的 AI 向量化和搜索框架。使用它来进行相应的向量、模型或检索工作流。上游将此功能描述为实验性。

### 核心工作流

```sql
CREATE EXTENSION ogai;
```

在目标数据库中安装扩展，当可用时运行上游提供的最小示例，并在将其集成到应用程序 SQL 之前验证安装的版本和返回值。

### 重要对象

- `ogai.ai_unvectorize(p_task_name TEXT)` 是一个扩展函数，返回 `TABLE`。
- `ogai.ai_vectorize(p_task_name TEXT, p_task_type TEXT, p_index_type TEXT, p_embed_model TEXT, p_src_schema TEXT, p_src_table TEXT, p_src_col TEXT, p_primary_key TEXT, p_table_method TEXT, p_dim INTEGER, p_max_chunk_size INTEGER DEFAULT 1000, p_max_chunk_overlap INTEGER DEFAULT 2…)` 是一个扩展函数，返回 `TABLE`。
- `ogai.encrypt_api_key_trigger()` 是一个扩展函数，返回 `TRIGGER`。
- `ogai.hybrid_search(p_task_name TEXT, p_query TEXT, p_return_cols TEXT DEFAULT '', p_limit INTEGER DEFAULT 10, p_where_clause TEXT DEFAULT '')` 是一个扩展函数，返回 `TABLE`。
- `ogai.rag(p_user_question TEXT, p_task_name TEXT, p_reranker_model TEXT, p_chat_model TEXT, p_rerank_limit INTEGER DEFAULT 5, p_search_limit INTEGER DEFAULT 20)` 是一个扩展函数，返回 `TEXT`。
- `ogai.search(p_task_name TEXT, p_query TEXT, p_return_cols TEXT DEFAULT '', p_limit INTEGER DEFAULT 10, p_where_clause TEXT DEFAULT '')` 是一个扩展函数，返回 `TABLE`。
- `ogai.vectorize_param_trigger()` 是一个扩展函数，返回 `TRIGGER`。
- `ogai.vectorize_trigger_handle(p_content TEXT, p_embed_model TEXT, p_dim INTEGER, p_task_name TEXT, p_task_type TEXT, p_src_schema TEXT, p_src_table TEXT, p_primary_key TEXT, p_table_method TEXT, p_operation TEXT, p_pk_value TEXT, p_max_chunk_size INTEGER DEFAULT 1000, p_max_chunk_overlap I…)` 是一个扩展函数，返回 `VOID`。
- `ogai.model_provider_type` 是一个扩展定义的类型。
- `ogai.queue_status` 是一个扩展定义的类型。
- `ogai.table_method` 是一个扩展定义的类型。
- `ogai.task_type` 是一个扩展定义的类型。
- `ogai.model_sources` 是由扩展安装或管理的表。
- `ogai.vectorize_queue` 是由扩展安装或管理的表。

### 要求与注意事项

- 审查后的控制文件声明默认版本为 `1.0`。
- 控制文件标记该扩展为不可重定位。
- 上游将项目的一部分或全部标记为实验性。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况，以匹配固定源。
