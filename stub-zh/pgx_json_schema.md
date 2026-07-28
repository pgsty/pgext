## 用法

来源：

- [官方上游 README](https://github.com/jefbarn/pgx_json_schema/blob/00a573b165139eae9ff78149b22a3b4dcc2a9c69/README.md)
- [官方扩展控制文件 (pgx_json_schema.control)](https://github.com/jefbarn/pgx_json_schema/blob/00a573b165139eae9ff78149b22a3b4dcc2a9c69/pgx_json_schema.control)
- [官方实现源代码](https://github.com/jefbarn/pgx_json_schema/blob/00a573b165139eae9ff78149b22a3b4dcc2a9c69/src/lib.rs)

`pgx_json_schema` — 一个用 Rust 实现的 Postgres JSON Schema 验证器。当应用程序需要此特定数据库功能时，请使用它。使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建中进行测试。

### 核心工作流

```sql
CREATE EXTENSION pgx_json_schema;
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行该示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 要求与注意事项

- 审查过的控制文件声明默认版本为 `1.0`。
- 控制文件将扩展标记为不可重定位。
- 控制文件不要求超级用户安装。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源代码中的信息一致。
