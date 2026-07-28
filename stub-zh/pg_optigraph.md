## 用法

来源：

- [官方上游 README](https://github.com/cloverdb/pg_optigraph/blob/83e362f68774c0c87fc2de1508795011c4170dfe/README.md)
- [官方扩展控制文件 (pg_optigraph.control)](https://github.com/cloverdb/pg_optigraph/blob/83e362f68774c0c87fc2de1508795011c4170dfe/extension/pg_optigraph.control)
- [官方扩展 SQL (pg_optigraph--0.1.0.sql)](https://github.com/cloverdb/pg_optigraph/blob/83e362f68774c0c87fc2de1508795011c4170dfe/extension/sql/pg_optigraph--0.1.0.sql)

`pg_optigraph` — pg Optigraph 是一个使用 OptiGraph ML 模型的 PostgreSQL 扩展。请使用它进行相应的向量、模型或检索工作流。请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION pg_optigraph;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `optigraph.extract_plans(query_text TEXT)` 是一个扩展函数，返回 `TABLE`。
- `optigraph.health_check()` 是一个扩展函数，返回 `TABLE`。
- `optigraph.reset_stats()` 是一个扩展函数，返回 `VOID`。
- `optigraph.stats()` 是一个扩展函数，返回 `TABLE`。
- `optigraph.status()` 是一个扩展函数，返回 `TABLE`。
- `optigraph.test_optimize()` 是一个扩展函数，返回 `TABLE`。
- `optigraph.configuration` 是一个扩展定义视图。
- `optigraph` 是由扩展创建的一个模式。

### 要求与注意事项

- 控制文件声明默认版本为 `0.1.0`。
- 控制文件标记该扩展为不可重定位。
- 控制文件要求超级用户进行安装。
- 控制文件标记该扩展为不受信任。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源中的内容一致。
