## 用法

来源：

- [官方扩展控制文件](https://github.com/citusdata/pg_octopus/blob/a6e892480f007a6fe00c956ad43fc814b9520928/pg_octopus.control)

`pg_octopus` — 运行后台工作进程，周期性连接已配置的 PostgreSQL 节点并记录健康状态。

已复核目录快照记录的版本为 `0.1`、类型为 `preload`、实现语言为 `C`。
整理后的兼容版本集合为 `10,11,12,13,14,15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pg_octopus";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_octopus';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
