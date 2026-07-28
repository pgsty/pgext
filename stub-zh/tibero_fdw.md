## 用法

来源：

- [官方上游 README](https://github.com/tmaxopensql/tibero-fdw/blob/49d2260e9f9228471ca318ada0f7630628dca974/README.md)
- [官方扩展控制文件 (tibero_fdw.control)](https://github.com/tmaxopensql/tibero-fdw/blob/49d2260e9f9228471ca318ada0f7630628dca974/tibero_fdw.control)
- [官方扩展 SQL (tibero_fdw--1.0.sql)](https://github.com/tmaxopensql/tibero-fdw/blob/49d2260e9f9228471ca318ada0f7630628dca974/tibero_fdw--1.0.sql)

`tibero_fdw` — 这个 PostgreSQL 扩展实现了一个外部数据源适配器（Foreign Data Wrapper，FDW），用于 Tibero。当 PostgreSQL 需要通过外部数据接口访问相应的外部数据源时，请使用此扩展。请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建中进行测试。

### 核心工作流

```sql
CREATE EXTENSION tibero_fdw;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `tibero_fdw_handler()` 是一个扩展函数，返回 `fdw_handler`。
- `tibero_fdw_validator(text[], oid)` 是一个扩展函数，返回 `void`。
- `tibero_fdw` 是一个扩展定义的外部数据源适配器。

### 要求与注意事项

- 控制文件声明默认版本为 `1.0`。
- 控制文件将扩展标记为可重定位。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为以及失败情况，以与固定源进行比对。
