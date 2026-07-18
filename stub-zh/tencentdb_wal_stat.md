## 用法

来源：

- [官方上游文档](https://cloud.tencent.com/document/product/409/130512)
- [官方项目或服务商页面](https://cloud.tencent.com/product/postgres)

`tencentdb_wal_stat` — 腾讯云数据库专有 WAL 诊断扩展：扫描近期 WAL 文件，按数据库、模式与关系归因记录数、字节数、资源管理器类型和全页镜像。

已复核目录快照记录的版本为 `1.0`、类型为 `standard`、实现语言为 `C`。
整理后的兼容版本集合为 `14,15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "tencentdb_wal_stat";
SELECT extversion
FROM pg_extension
WHERE extname = 'tencentdb_wal_stat';
```

这是 `Tencent Cloud` 的服务商专用组件；可用性、启用、权限与升级遵循该服务，而不是可移植的社区软件包。

整理后的生命周期为 `active`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
