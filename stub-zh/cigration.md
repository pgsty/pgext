## 用法

来源：

- [官方扩展控制文件](https://github.com/cloud-sn2/cigration/blob/7da7778375544eb0fee789f51e7b92cf3019b94c/cigration.control)

`cigration` — 在线迁移与再平衡 Citus 共置分片，用于集群扩容、缩容及 Worker 节点替换。

已复核目录快照记录的版本为 `1.1`、类型为 `puresql`、实现语言为 `SQL`。
应先安装并验证声明的扩展依赖：`citus`, `dblink`。
整理后的兼容版本集合为 `10,11,12,13`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "cigration";
SELECT extversion
FROM pg_extension
WHERE extname = 'cigration';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
