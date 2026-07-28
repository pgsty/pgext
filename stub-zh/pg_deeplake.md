## 用法

来源：

- [官方上游 README](https://github.com/activeloopai/deeplake/blob/f432041fabbf4a1fc1d342aeb550e0a7de41b0da/postgres/README.md)
- [官方扩展控制文件 (pg_deeplake.control)](https://github.com/activeloopai/deeplake/blob/f432041fabbf4a1fc1d342aeb550e0a7de41b0da/postgres/pg_deeplake.control)
- [官方扩展 SQL (pg_deeplake--1.0.sql)](https://github.com/activeloopai/deeplake/blob/f432041fabbf4a1fc1d342aeb550e0a7de41b0da/postgres/pg_deeplake--1.0.sql)

`pg_deeplake` — PostgreSQL 扩展，用于向量相似度搜索、全文搜索和使用 DeepLake 的混合搜索。请在相应的向量、模型或检索工作流中使用它。请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION pg_deeplake;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `array_ndims(anyarray)` 是一个扩展函数，返回 `int`。
- `contains(col text, search text)` 是一个扩展函数，返回 `boolean`。
- `create_deeplake_table(tablename TEXT, path TEXT)` 是一个扩展函数，返回 `void`。
- `deeplake_bm25_cmp(text, text)` 是一个扩展函数，返回 `int4`。
- `deeplake_bm25_similarity(text, text)` 是一个扩展函数，返回 `float4`。
- `deeplake_cosine_similarity(float4[], float4[])` 是一个扩展函数，返回 `float4`。
- `deeplake_hybrid_cmp(deeplake_hybrid_record, deeplake_hybrid_record_weighted)` 是一个扩展函数，返回 `int4`。
- `deeplake_hybrid_record(embedding float4[], text_value text)` 是一个扩展函数，返回 `deeplake_hybrid_record_weighted`。
- `deeplake_hybrid_record(embedding float4[], text_value text, embedding_weight float8, text_weight float8)` 是一个扩展函数，返回 `deeplake_hybrid_record_weighted`。
- `deeplake_hybrid_record_to_weighted(deeplake_hybrid_record)` 是一个扩展函数，返回 `deeplake_hybrid_record_weighted`。
- `deeplake_hybrid_record_weighted_to_simple(deeplake_hybrid_record_weighted)` 是一个扩展函数，返回 `deeplake_hybrid_record`。
- `deeplake_hybrid_search(deeplake_hybrid_record, deeplake_hybrid_record_weighted)` 是一个扩展函数，返回 `float4`。
- `deeplake_index_handler(INTERNAL)` 是一个扩展函数，返回 `index_am_handler`。
- `deeplake_maxsim(float4[][], float4[][])` 是一个扩展函数，返回 `float4`。

### 要求与注意事项

- 审核的控制文件声明默认版本为 `1.0`。
- 控制文件将扩展标记为可重定位。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况。
