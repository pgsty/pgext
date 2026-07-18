## 用法

来源：

- [官方扩展控制文件](https://github.com/vibhorkum/efm_extension/blob/ecb305cc185b05eaf0e766d6a9c5d23b3072e2fb/efm_extension.control)
- [官方上游文档](https://github.com/vibhorkum/efm_extension/blob/ecb305cc185b05eaf0e766d6a9c5d23b3072e2fb/README.md)

`efm_extension` — 通过 SQL 管理和监控 EDB Failover Manager 集群。

已复核目录快照记录的版本为 `1.1`、类型为 `preload`、实现语言为 `C`。
应先安装并验证声明的扩展依赖：`dblink`, `pgcrypto`。
整理后的兼容版本集合为 `14,15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "efm_extension";
SELECT extversion
FROM pg_extension
WHERE extname = 'efm_extension';
```

整理后的生命周期为 `active`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
