## 用法

来源：

- [官方扩展控制文件](https://github.com/zilder/pg_udp/blob/ab0f84f8820996bf58eb597eeb5c299476f9137d/pg_udp.control)

`pg_udp` — 通过 PostgreSQL 函数向指定主机和端口发送 UDP 数据报。

已复核目录快照记录的版本为 `0.1`、类型为 `standard`、实现语言为 `C`。

```sql
CREATE EXTENSION "pg_udp";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_udp';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
