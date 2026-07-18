## 用法

来源：

- [官方上游文档](https://www.alibabacloud.com/help/en/polardb/polardb-for-postgresql/read-and-write-external-data-files-by-using-oss-fdw)

`oss_fdw` — 用于在阿里云 PostgreSQL 兼容托管服务中读写 OSS 对象的外部数据包装器扩展。

已复核目录快照记录的版本为 `unknown`、类型为 `standard`、实现语言为 `C`。
整理后的兼容版本集合为 `11,14,16`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "oss_fdw";
SELECT extversion
FROM pg_extension
WHERE extname = 'oss_fdw';
```

这是 `Alibaba Cloud` 的服务商专用组件；可用性、启用、权限与升级遵循该服务，而不是可移植的社区软件包。

整理后的生命周期为 `active`。采用前应固定已复核构建并确认维护状态。
官方材料包含实验性、弃用、不支持或明确警告边界；用于非实验环境前必须阅读全文并测试失败场景。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
