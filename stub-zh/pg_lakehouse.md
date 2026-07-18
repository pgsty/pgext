## 用法

来源：

- [官方上游 README](https://github.com/paradedb/paradedb/blob/59a9224db5b22893cd308993bcddfbfdbc14cfc1/README.md)

`pg_lakehouse` — ParadeDB 已弃用的 PostgreSQL 数据湖分析查询引擎

已复核目录快照记录的版本为 `0.7.6`、类型为 `preload`、实现语言为 `Rust`。
整理后的兼容版本集合为 `14,15,16`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pg_lakehouse";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_lakehouse';
```

该上游项目与 `ParadeDB` 相关；应根据所链接来源核实当前支持、许可证、打包方式与部署边界。

整理后的生命周期为 `deprecated`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
