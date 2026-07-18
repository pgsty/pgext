## 用法

来源：

- [官方扩展控制文件](https://github.com/frost242/ssi_failures/blob/8b6b38b1f688632c3520d1b6aaac24feb7734c84/pg_stat_ssi_failures.control)
- [官方上游文档](https://github.com/frost242/ssi_failures/blob/8b6b38b1f688632c3520d1b6aaac24feb7734c84/README.md)

`pg_stat_ssi_failures` — pg_stat_ssi_failures：统计 PostgreSQL 集群级 SSI 序列化失败次数的监控扩展。

已复核目录快照记录的版本为 `0.1`、类型为 `preload`、实现语言为 `C`。

```sql
CREATE EXTENSION "pg_stat_ssi_failures";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_stat_ssi_failures';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
