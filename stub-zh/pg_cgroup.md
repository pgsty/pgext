## 用法

来源：

- [官方扩展控制文件](https://github.com/MasahikoSawada/pg_cgroup/blob/09214a99f5acfced8697ff3c72b0b05dc6bc7bb2/pg_cgroup.control)
- [官方上游文档](https://github.com/MasahikoSawada/pg_cgroup/blob/09214a99f5acfced8697ff3c72b0b05dc6bc7bb2/README.md)

`pg_cgroup` — 提供 SQL 可调用函数，用于创建、配置 Linux cgroup 并将进程加入资源组。

已复核目录快照记录的版本为 `1.0`、类型为 `preload`、实现语言为 `C`。
整理后的兼容版本集合为 `10`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pg_cgroup";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_cgroup';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
