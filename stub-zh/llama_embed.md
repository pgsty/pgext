## 用法

来源：

- [官方上游 README](https://github.com/mayflower/pgturbohybrid/blob/cd68155df20545b163dd610fb95e8aa7e62fc108/extensions/pg_colbert_llama/README.md)
- [官方扩展控制文件 (llama_embed.control)](https://github.com/mayflower/pgturbohybrid/blob/cd68155df20545b163dd610fb95e8aa7e62fc108/extensions/pg_colbert_llama/llama_embed.control)
- [官方扩展 SQL (llama_embed--0.1.0.sql)](https://github.com/mayflower/pgturbohybrid/blob/cd68155df20545b163dd610fb95e8aa7e62fc108/extensions/pg_colbert_llama/sql/llama_embed--0.1.0.sql)

`llama_embed` — llama_embed 是在 PostgreSQL 内运行 llama.cpp 嵌入模型的 SQL 面向伴侣扩展。它可以返回密集向量嵌入、标记级向量[]嵌入或 ColBERT 等晚期交互模型的 pgturbohybrid 多向量。使用它来进行相应的向量、模型或检索工作流。上游将此功能描述为实验性。

### 核心工作流

```sql
CREATE EXTENSION llama_embed;
```

在目标数据库中安装扩展，当可用时运行上游提供的最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `llama_embed` 是一个扩展函数。
- `llama_embed_model_info(model pg_catalog.text)` 是一个扩展函数并返回 `pg_catalog`。
- `llama_embed_mv` 是一个扩展函数。
- `llama_embed_mv_batch` 是一个扩展函数。
- `llama_embed_sparse` 是一个扩展函数。
- `llama_embed_sparse_batch` 是一个扩展函数。
- `llama_embed_sparse_model_info` 是一个扩展函数。
- `llama_embed_tokens` 是一个扩展函数。
- `llama_embed_vector` 是一个扩展函数。
- `llama_embed_vector_batch` 是一个扩展函数。

### 要求与注意事项

- 审查后的控制文件声明默认版本为 `0.1.0`。
- 先安装确认的扩展依赖项：`pgturbohybrid_experimental`。
- 控制文件标记该扩展为可重定位。
- 上游将项目的一部分或全部标记为实验性。
- 在生产使用前，确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源进行比对。
