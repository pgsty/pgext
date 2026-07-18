## 用法

来源：

- [官方上游文档](https://github.com/MasahikoSawada/pg_keeper/blob/4bfc5773fae35f55daf6c6015a947456209c51a8/README.md)

`pg_keeper` — 基于后台工作进程的集群模块，用于监控流复制、切换同步复制模式并在故障时提升备用服务器。

已复核目录快照记录的版本为 `2.0`、类型为 `preload`、实现语言为 `C`。

```sql
CREATE EXTENSION "pg_keeper";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_keeper';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
