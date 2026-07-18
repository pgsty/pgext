## 用法

来源：

- [官方项目或服务商页面](https://pgexporter.github.io/)

`pgexporter_ext` — 为 pgexporter 暴露额外的操作系统、CPU、内存、网络、负载、磁盘和 FIPS 指标。

已复核目录快照记录的版本为 `0.3.1`、类型为 `preload`、实现语言为 `C`。
整理后的兼容版本集合为 `13,14,15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pgexporter_ext";
SELECT extversion
FROM pg_extension
WHERE extname = 'pgexporter_ext';
```

整理后的生命周期为 `active`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
