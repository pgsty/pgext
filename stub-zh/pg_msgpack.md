## 用法

来源：

- [官方扩展控制文件](https://github.com/choplin/pg_msgpack/blob/7cf18817053f4a42df80c34d816003505de8b1c1/pg_msgpack.control)

`pg_msgpack` — pg_msgpack：提供 MessagePack 数据类型、JSON/bytea 转换和访问操作符的 PostgreSQL 扩展。

已复核目录快照记录的版本为 `0.0.1`、类型为 `standard`、实现语言为 `C`。

```sql
CREATE EXTENSION "pg_msgpack";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_msgpack';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
