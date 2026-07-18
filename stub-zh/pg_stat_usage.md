## 用法

来源：

- [官方扩展控制文件](https://github.com/mpihlak/pg_stat_usage/blob/8e8173b3a8fd61a36491e86b4c9d8c021ad77a30/pg_stat_usage.control)
- [官方上游文档](https://github.com/mpihlak/pg_stat_usage/blob/8e8173b3a8fd61a36491e86b4c9d8c021ad77a30/README.md)

`pg_stat_usage` — pg_stat_usage：统计观测类 PostgreSQL 扩展；上游说明为“追踪存储过程调用的使用统计信息”。

已复核目录快照记录的版本为 `1.0`、类型为 `preload`、实现语言为 `C`。

```sql
CREATE EXTENSION "pg_stat_usage";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_stat_usage';
```

整理后的生命周期为 `abandoned`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
