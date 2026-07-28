## 用法

来源：

- [官方上游 README](https://github.com/okbob/simple/blob/e22ac05dbc0a78cf491df0d8fe1b08d66dbab28c/README.md)
- [官方扩展控制文件 (simple.control)](https://github.com/okbob/simple/blob/e22ac05dbc0a78cf491df0d8fe1b08d66dbab28c/simple.control)
- [官方扩展 SQL (simple--1.0.sql)](https://github.com/okbob/simple/blob/e22ac05dbc0a78cf491df0d8fe1b08d66dbab28c/sql/simple--1.0.sql)

`simple` — 这是一个用于 Postgres 开发培训的小型 PostgreSQL 扩展。当应用程序需要特定的数据库功能时，请使用它。请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION simple;
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行该示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `int_func(int)` 是一个扩展函数，返回 `int`。
- `text_func(text)` 是一个扩展函数，返回 `text`。

### 要求与注意事项

- 控制文件声明默认版本为 `1.0`。
- 控制文件将扩展标记为可重定位。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为以及失败情况，以与固定源进行比对。
