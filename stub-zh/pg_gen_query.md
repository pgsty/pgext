## 用法

来源：

- [官方上游 README](https://github.com/nimit/whatsinmypg/blob/6728d32cfa680120edac83597d7c319dba4f00f8/README.md)
- [官方扩展控制文件 (pg_gen_query.control)](https://github.com/nimit/whatsinmypg/blob/6728d32cfa680120edac83597d7c319dba4f00f8/pg_gen_query.control)
- [官方扩展 SQL (pg_gen_query--1.0.sql)](https://github.com/nimit/whatsinmypg/blob/6728d32cfa680120edac83597d7c319dba4f00f8/sql/pg_gen_query--1.0.sql)

`pg_gen_query` — 此 PostgreSQL 扩展提供了一个函数 pg_gen_query，它可以将自然语言输入转换为等效的 SQL 命令。在相应的向量、模型或检索工作流中使用它。请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建中进行测试。

### 核心工作流

```sql
CREATE EXTENSION pg_gen_query;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `pg_gen_query(query text)` 是一个扩展函数，返回 `TEXT`。
- `regen_schema_cache()` 是一个扩展函数，返回 `void`。
- `regen_schema_cache_trigger()` 是一个扩展函数，返回 `event_trigger`。

### 要求与注意事项

- 审查过的控制文件声明默认版本为 `1.0`。
- 控制文件将扩展标记为可重定位。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源进行比对。
