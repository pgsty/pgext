## 用法

来源：

- [官方扩展控制文件](https://github.com/vidardb/pgrocks-fdw/blob/8f89fcad7e9c471e8dacfd167457716b7cf3d6e5/kv_fdw.control)

`kv_fdw` — kv_fdw：用于访问由 RocksDB/VidarDB 支持的键值存储的 PostgreSQL 外部数据包装器。

已复核目录快照记录的版本为 `0.0.1`、类型为 `preload`、实现语言为 `C++`。
整理后的兼容版本集合为 `13`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "kv_fdw";
SELECT extversion
FROM pg_extension
WHERE extname = 'kv_fdw';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
