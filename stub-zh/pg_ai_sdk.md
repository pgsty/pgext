## 用法

来源：

- [官方上游 README](https://github.com/deeox/pg_ai_sdk/blob/24498354cb766ac274157d42127ff6702fa66649/README.md)
- [官方扩展控制文件 (pg_ai_sdk.control)](https://github.com/deeox/pg_ai_sdk/blob/24498354cb766ac274157d42127ff6702fa66649/pg_ai_sdk.control)
- [官方扩展 SQL (pg_ai_sdk--1.0.sql)](https://github.com/deeox/pg_ai_sdk/blob/24498354cb766ac274157d42127ff6702fa66649/pg_ai_sdk--1.0.sql)

`pg_ai_sdk` — 这个项目是一个 PostgreSQL 扩展，使用 ClickHouse/ai-sdk-cpp 将自然语言查询转换为 SQL。它允许用户通过用英语提问来查询其数据库。请在相应的向量、模型或检索工作流中使用它。使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION pg_ai_sdk;

-- Generate SQL query from natural language
SELECT generate_sql_from_text('how many matches were won by Royal Challengers Bangalore is 2023');
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `generate_sql_from_text(natural_language_query TEXT)` 是一个扩展函数，返回 `TEXT`。
- `generate_sql_from_text(natural_language_query TEXT, model_name TEXT)` 是一个扩展函数，返回 `TEXT`。
- `pg_ai_sdk_execute_json(natural_language_query text)` 是一个扩展函数，返回 `text`。
- `pg_ai_sdk_execute_json(natural_language_query text, model_name TEXT)` 是一个扩展函数，返回 `text`。

### 要求与注意事项

- 审查后的控制文件声明默认版本为 `1.0`。
- 控制文件将扩展标记为可重定位。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况，与固定源进行比对。
