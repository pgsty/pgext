## 用法

来源：

- [官方上游 README](https://github.com/emielm/gofdw/blob/1f3b8256a0b09f49d017cf580b4ccb6f301f409b/README.md)
- [官方扩展控制文件 (gofdw.control)](https://github.com/emielm/gofdw/blob/1f3b8256a0b09f49d017cf580b4ccb6f301f409b/gofdw.control)
- [官方扩展 SQL (gofdw.sql)](https://github.com/emielm/gofdw/blob/1f3b8256a0b09f49d017cf580b4ccb6f301f409b/gofdw.sql)

`gofdw` — 一个使用 cgo 实现的 Postgres 外部数据封装器，非常实验性。当 PostgreSQL 需要通过外部数据接口访问相应的外部数据源时，请使用此扩展。上游将此功能描述为实验性。

### 核心工作流

```sql
CREATE EXTENSION gofdw;
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行该示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `gofdw_handler()` 是一个扩展函数，返回 `fdw_handler`。
- `gofdw_validator(text[], oid)` 是一个扩展函数。
- `gofdw` 是一个扩展定义的外部数据封装器。

### 要求与注意事项

- 审查后的控制文件声明默认版本为 `0.1`。
- 控制文件将扩展标记为可重定位。
- 上游将项目的一部分或全部标记为实验性。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况，以与固定源进行比对。
