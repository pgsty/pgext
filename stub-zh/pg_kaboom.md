## 用法

来源：

- [官方上游 README](https://api.pgxn.org/src/pg_kaboom/pg_kaboom-0.0.1/README.md)
- [官方扩展控制文件 (pg_kaboom.control)](https://api.pgxn.org/src/pg_kaboom/pg_kaboom-0.0.1/pg_kaboom.control)
- [官方扩展 SQL (pg_kaboom--0.0.1.sql)](https://api.pgxn.org/src/pg_kaboom/pg_kaboom-0.0.1/pg_kaboom--0.0.1.sql)

`pg_kaboom` — 该扩展用于以多种破坏性方式使 PostgreSQL 崩溃。当应用程序需要此特定数据库功能时，请使用它。请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION pg_kaboom;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证安装的版本和返回值。

### 重要对象

- `pg_kaboom(method text, payload jsonb default NULL)` 是一个扩展函数，返回 `boolean`。
- `pg_kaboom_arsenal()` 是一个扩展函数，返回 `TABLE`。

### 要求与注意事项

- 控制文件声明默认版本为 `0.0.1`。
- 控制文件将扩展标记为可重定位。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源进行比对。
