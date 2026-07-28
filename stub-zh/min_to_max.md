## 用法

来源：

- [官方上游 README](https://github.com/chanukyasds/pg_extensions/blob/d732560f24a2741225182a8b74a36837cc4abc3f/README.md)
- [官方扩展控制文件 (min_to_max.control)](https://github.com/chanukyasds/pg_extensions/blob/d732560f24a2741225182a8b74a36837cc4abc3f/aggregates/min_to_max/min_to_max.control)
- [官方扩展 SQL (min_to_max--1.0.sql)](https://github.com/chanukyasds/pg_extensions/blob/d732560f24a2741225182a8b74a36837cc4abc3f/aggregates/min_to_max/min_to_max--1.0.sql)

`min_to_max` — 为 PostgreSQL 服务器提供聚合函数和函数。当 SQL 需要这些特殊函数或聚合时使用它。请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建中进行测试。

### 核心工作流

```sql
CREATE EXTENSION min_to_max;
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行该示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `min_to_max_ffunc(internal,anynonarray)` 是一个扩展函数，返回 `text`。
- `min_to_max_sfunc(internal,anynonarray)` 是一个扩展函数，返回 `internal`。
- `min_to_max` 是由扩展公开的聚合函数。

### 要求与注意事项

- 控制文件声明默认版本为 `1.0`。
- 控制文件标记该扩展为可重定位。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源进行比对。
