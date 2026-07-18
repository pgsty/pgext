## 用法

来源：

- [官方扩展控制文件](https://github.com/filiprem/pg-tools/blob/3ad4b21874bed6b629f984d046650de354e95d92/ddl_versioning/ddl_versioning.control)
- [官方上游文档](https://github.com/filiprem/pg-tools/blob/3ad4b21874bed6b629f984d046650de354e95d92/ddl_versioning/README.ddl_versioning)
- [官方上游 README](https://github.com/filiprem/pg-tools/blob/3ad4b21874bed6b629f984d046650de354e95d92/README.md)

`ddl_versioning` — 通过事件触发器记录表、索引、函数和视图等 DDL 对象版本。

已复核目录快照记录的版本为 `1.0`、类型为 `puresql`、实现语言为 `SQL`。
整理后的兼容版本集合为 `13,14,15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "ddl_versioning";
SELECT extversion
FROM pg_extension
WHERE extname = 'ddl_versioning';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
