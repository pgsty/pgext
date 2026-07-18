## 用法

来源：

- [官方上游文档](https://github.com/BrandwatchLtd/trunkfingerprint/blob/a4ed83cd533f233e70b23aac3ebb7d91248c7087/README.md)

`trunkfingerprint` — 计算可比较的 PostgreSQL 数据库结构指纹，并可选择包含数据指纹。

已复核目录快照记录的版本为 `1.3.0`、类型为 `puresql`、实现语言为 `SQL`。
整理后的兼容版本集合为 `10,11,12,13,14,15`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "trunkfingerprint";
SELECT extversion
FROM pg_extension
WHERE extname = 'trunkfingerprint';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
