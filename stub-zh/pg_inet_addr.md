## 用法

来源：

- [官方扩展控制文件](https://github.com/eulerto/pg_inet_addr/blob/ac8bee1ed392e4a07ed75f24cb87414b297fe40c/pg_inet_addr.control)
- [官方上游文档](https://github.com/eulerto/pg_inet_addr/blob/ac8bee1ed392e4a07ed75f24cb87414b297fe40c/README.md)

`pg_inet_addr` — 列出 PostgreSQL 服务器主机网络接口所分配的 IPv4、IPv6 地址及网络掩码。

已复核目录快照记录的版本为 `1.0`、类型为 `standard`、实现语言为 `C`。
整理后的兼容版本集合为 `10,11,12,13,14,15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pg_inet_addr";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_inet_addr';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
