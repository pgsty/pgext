## 用法

来源：

- [官方上游文档](https://www.alibabacloud.com/help/en/rds/apsaradb-rds-for-postgresql/vector-pyramid)
- [官方主文档](https://www.alibabacloud.com/help/en/rds/apsaradb-rds-for-postgresql/extensions-supported-by-apsaradb-rds-for-postgresql)

`ganos_geometry_pyramid` — 为大规模二维几何数据构建稀疏矢量瓦片金字塔，以实现快速渲染。

已复核目录快照记录的版本为 `7.0`、类型为 `standard`、实现语言为 `C`。
整理后的兼容版本集合为 `10,11,12,13,14,15,16,17`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "ganos_geometry_pyramid";
SELECT extversion
FROM pg_extension
WHERE extname = 'ganos_geometry_pyramid';
```

这是 `Alibaba Cloud` 的服务商专用组件；可用性、启用、权限与升级遵循该服务，而不是可移植的社区软件包。

整理后的生命周期为 `active`。采用前应固定已复核构建并确认维护状态。
官方材料包含实验性、弃用、不支持或明确警告边界；用于非实验环境前必须阅读全文并测试失败场景。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
