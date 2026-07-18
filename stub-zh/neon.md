## 用法

来源：

- [官方上游文档](https://neon.com/docs/extensions/neon)
- [官方项目或服务商页面](https://neon.com)
- [官方上游 README](https://github.com/neondatabase/neon/blob/8f60b04da47ffefe0e52bda2440134b42874eb75/README.md)

`neon` — 用于 Neon Postgres 计算节点的 Neon 专用统计信息与本地文件缓存指标扩展。

已复核目录快照记录的版本为 `1.9`、类型为 `preload`、实现语言为 `C`。
整理后的兼容版本集合为 `14,15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "neon";
SELECT extversion
FROM pg_extension
WHERE extname = 'neon';
```

该上游项目与 `Neon` 相关；应根据所链接来源核实当前支持、许可证、打包方式与部署边界。

整理后的生命周期为 `active`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
