## 用法

来源：

- [官方上游 README](https://github.com/takanoriyanagitani/dir_fdw/blob/c6edbbb4a9928a687a576ccef058e496e2aaffc9/README.md)
- [官方扩展控制文件 (dir_fdw.control)](https://github.com/takanoriyanagitani/dir_fdw/blob/c6edbbb4a9928a687a576ccef058e496e2aaffc9/dir_fdw.control)
- [官方扩展 SQL (dir_fdw--1.0.sql)](https://github.com/takanoriyanagitani/dir_fdw/blob/c6edbbb4a9928a687a576ccef058e496e2aaffc9/dir_fdw--1.0.sql)

`dir_fdw` — 外部数据封装器，用于 readdir。当 PostgreSQL 需要通过外部数据接口访问相应的外部数据源时使用它。请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建中进行测试。

### 核心工作流

```sql
CREATE EXTENSION dir_fdw;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `dir_fdw_handler()` 是一个扩展函数，返回 `fdw_handler`。
- `dir_fdw_validator(text[], oid)` 是一个扩展函数，返回 `void`。
- `dir_fdw` 是一个扩展定义的外部数据封装器。

### 要求与注意事项

- 控制文件声明默认版本为 `1.0`。
- 控制文件将扩展标记为可重定位。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源进行比对。
