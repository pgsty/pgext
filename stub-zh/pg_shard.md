## 用法

来源：

- [官方上游文档](https://github.com/citusdata/pg_shard/blob/0014e199c5723020f3e0c982882a10ca53694cad/README.md)
- [官方 PGXN 分发页](https://pgxn.org/dist/pg_shard/)
- [官方项目或服务商页面](https://github.com/citusdata/citus)

`pg_shard` — 在远程 PostgreSQL 服务器之间分片并复制数据表

已复核目录快照记录的版本为 `1.2`、类型为 `preload`、实现语言为 `C`。

```sql
CREATE EXTENSION "pg_shard";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_shard';
```

该上游项目与 `Citus Data` 相关；应根据所链接来源核实当前支持、许可证、打包方式与部署边界。

整理后的生命周期为 `deprecated`。采用前应固定已复核构建并确认维护状态。
官方材料包含实验性、弃用、不支持或明确警告边界；用于非实验环境前必须阅读全文并测试失败场景。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
