## 用法

来源：

- [官方上游 README](https://github.com/mahipv/pg_ai/blob/890dcd622528c00f710214ec966987ab2d04c2d4/README.md)
- [官方扩展控制文件 (pg_ai.control)](https://github.com/mahipv/pg_ai/blob/890dcd622528c00f710214ec966987ab2d04c2d4/pg_ai.control)
- [官方扩展 SQL (pg_ai--0.0.1.sql)](https://github.com/mahipv/pg_ai/blob/890dcd622528c00f710214ec966987ab2d04c2d4/sql/pg_ai--0.0.1.sql)

`pg_ai` — PostgreSQL 扩展，内置 RAG 能力，允许通过自然语言和 SQL 函数来解释和查询数据。使用它来进行相应的向量、模型或检索工作流。请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION pg_ai;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `pg_ai_create_vector_store(store NAME, sql_query TEXT, notes NAME = NULL)` 是一个扩展函数，返回 `TEXT`。
- `pg_ai_generate_image(column_name TEXT, prompt TEXT = NULL)` 是一个扩展函数，返回 `TEXT`。
- `pg_ai_help()` 是一个扩展函数，返回 `TEXT`。
- `pg_ai_insight(column_name TEXT, prompt TEXT = NULL)` 是一个扩展函数，返回 `TEXT`。
- `pg_ai_moderation(column_name TEXT, prompt TEXT = NULL)` 是一个扩展函数，返回 `TEXT`。
- `pg_ai_query_vector_store(store NAME, nl_query TEXT, count INT = 2)` 是一个扩展函数，返回 `SETOF`。
- `pg_ai_generate_image_agg` 是由扩展公开的聚合函数。
- `pg_ai_insight_agg` 是由扩展公开的聚合函数。
- `pg_ai_moderation_agg` 是由扩展公开的聚合函数。

### 要求与注意事项

- 审查后的控制文件声明默认版本为 `0.0.1`。
- 控制文件标记该扩展为不可重定位。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为以及失败情况，以匹配固定源代码。
