## 用法

来源：

- [官方上游 README](https://github.com/mayflower/pgturbohybrid/blob/cd68155df20545b163dd610fb95e8aa7e62fc108/extensions/pg_colbert_llama/README.md)
- [官方扩展控制文件 (pg_colbert_llama.control)](https://github.com/mayflower/pgturbohybrid/blob/cd68155df20545b163dd610fb95e8aa7e62fc108/extensions/pg_colbert_llama/pg_colbert_llama.control)
- [官方扩展 SQL (pg_colbert_llama--0.1.0.sql)](https://github.com/mayflower/pgturbohybrid/blob/cd68155df20545b163dd610fb95e8aa7e62fc108/extensions/pg_colbert_llama/sql/pg_colbert_llama--0.1.0.sql)

`pg_colbert_llama` — llama_embed 是在 PostgreSQL 内运行 llama.cpp 嵌入模型的 SQL 面向伴侣扩展。它可以返回密集向量嵌入、标记级向量[]嵌入或 ColBERT 及其他晚期交互模型的 pgturbohybrid 多向量。使用它来进行相应的向量、模型或检索工作流。上游将此功能描述为实验性。

### 核心工作流

```sql
CREATE EXTENSION pg_colbert_llama;
```

在目标数据库中安装扩展，当可用时运行上游提供的最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `colbert(model pg_catalog.text, input pg_catalog.text)` 是一个扩展函数，返回 `pg_catalog`。
- `colbert_debug(model pg_catalog.text, input pg_catalog.text)` 是一个扩展函数，返回 `pg_catalog`。
- `colbert_dim(model pg_catalog.text, input pg_catalog.text)` 是一个扩展函数，返回 `pg_catalog`。
- `colbert_float4(model pg_catalog.text, input pg_catalog.text)` 是一个扩展函数，返回 `pg_catalog`。
- `colbert_model_info(model pg_catalog.text)` 是一个扩展函数，返回 `pg_catalog`。
- `colbert_mv(model pg_catalog.text, input pg_catalog.text)` 是一个扩展函数，返回 `turbohybrid_multivector`。
- `colbert_mv_batch(model pg_catalog.text, inputs pg_catalog.text[])` 是一个扩展函数，返回 `turbohybrid_multivector[]`。
- `colbert_vectors(model pg_catalog.text, input pg_catalog.text)` 是一个扩展函数，返回 `vector[]`。

### 要求与注意事项

- 审查后的控制文件声明默认版本为 `0.1.0`。
- 先安装确认的扩展依赖项：`pgturbohybrid_experimental`。
- 控制文件标记该扩展为可重定位。
- 上游将项目的一部分或全部标记为实验性。
- 在生产使用前，确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源进行比对。
