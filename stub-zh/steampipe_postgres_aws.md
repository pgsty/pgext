## 用法

来源：

- [官方扩展控制文件](https://database.dev/toliver38/steampipe_postgres_aws)
- [官方上游文档](https://hub.steampipe.io/plugins/turbot/aws#postgres-fdw)
- [官方项目或服务商页面](https://hub.steampipe.io/plugins/turbot/aws)

`steampipe_postgres_aws` — 将实时 AWS API 映射为 PostgreSQL 外部表的 Steampipe AWS 原生 FDW。

已复核目录快照记录的版本为 `0.125.0`、类型为 `standard`、实现语言为 `Go`。
整理后的兼容版本集合为 `14,15`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "steampipe_postgres_aws";
SELECT extversion
FROM pg_extension
WHERE extname = 'steampipe_postgres_aws';
```

该上游项目与 `Turbot` 相关；应根据所链接来源核实当前支持、许可证、打包方式与部署边界。
官方材料包含实验性、弃用、不支持或明确警告边界；用于非实验环境前必须阅读全文并测试失败场景。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
