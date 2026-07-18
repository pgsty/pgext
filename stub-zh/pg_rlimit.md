## 用法

来源：

- [官方扩展控制文件](https://github.com/harukat/pg_rlimit/blob/51aec672a347f92fe3e08ed2be75c8ee6b3427e2/pg_rlimit.control)
- [官方上游文档](https://github.com/harukat/pg_rlimit/blob/51aec672a347f92fe3e08ed2be75c8ee6b3427e2/README.md)

`pg_rlimit` — 按会话通过 SQL 函数或 GUC 设置、读取进程地址空间资源限制。

已复核目录快照记录的版本为 `1.0`、类型为 `preload`、实现语言为 `C`。
整理后的兼容版本集合为 `10,11,12,13,14,15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pg_rlimit";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_rlimit';
```

整理后的生命周期为 `active`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
