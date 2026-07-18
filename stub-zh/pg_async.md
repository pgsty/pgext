## 用法

来源：

- [官方扩展控制文件](https://github.com/RekGRpth/pg_async/blob/20c2984087d89a6d9d30ea727325c808051147bb/pg_async.control)
- [官方上游文档](https://github.com/RekGRpth/pg_async/blob/20c2984087d89a6d9d30ea727325c808051147bb/README.md)

`pg_async` — C 扩展，为热备只读副本提供异步 LISTEN/NOTIFY 队列支持。

已复核目录快照记录的版本为 `1.0`、类型为 `preload`、实现语言为 `C`。
整理后的兼容版本集合为 `13,14`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pg_async";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_async';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
