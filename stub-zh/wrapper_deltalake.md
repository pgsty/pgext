## 用法

来源：

- [官方扩展控制文件](https://github.com/Giangblackk/wrapper_deltalake/blob/0459b9f308ee790df21dafdd9f366b5b54fb33b0/wrapper_deltalake.control)
- [官方上游文档](https://github.com/Giangblackk/wrapper_deltalake/blob/0459b9f308ee790df21dafdd9f366b5b54fb33b0/README.md)

`wrapper_deltalake` — 原计划实现 Delta Lake 外部数据包装器、但仅停留在 pgrx Hello World 的未完成原型。

已复核目录快照记录的版本为 `0.0.0`、类型为 `standard`、实现语言为 `Rust`。
整理后的兼容版本集合为 `14,15,16`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "wrapper_deltalake";
SELECT extversion
FROM pg_extension
WHERE extname = 'wrapper_deltalake';
```

整理后的生命周期为 `abandoned`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
