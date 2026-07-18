## 用法

来源：

- [官方扩展控制文件](https://github.com/RekGRpth/pg_grpc/blob/adbdf3a00c68581b6691c48bb45e81f8295e442e/pg_grpc.control)

`pg_grpc` — 从 PostgreSQL 调用 gRPC C Core API 的扩展

已复核目录快照记录的版本为 `1.0`、类型为 `standard`、实现语言为 `C`。

```sql
CREATE EXTENSION "pg_grpc";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_grpc';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
