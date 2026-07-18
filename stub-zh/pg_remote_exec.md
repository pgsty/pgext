## 用法

来源：

- [官方扩展控制文件](https://github.com/cybertec-postgresql/pg_remote_exec/blob/1ac5d03725b86a1bebcf754b2e7f0212a46218c5/pg_remote_exec.control)
- [官方上游文档](https://github.com/cybertec-postgresql/pg_remote_exec/blob/1ac5d03725b86a1bebcf754b2e7f0212a46218c5/README.md)

`pg_remote_exec` — 从 SQL 执行 shell 命令，并可选择返回标准输出。

已复核目录快照记录的版本为 `1.0`、类型为 `standard`、实现语言为 `C`。
整理后的兼容版本集合为 `12,13,14,15,16`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pg_remote_exec";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_remote_exec';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
