## 用法

来源：

- [官方扩展控制文件](https://github.com/ycui1984/pg_thrift/blob/1de265a661ab3c61aa593d7e99d3a313024170fc/pg_thrift.control)
- [官方上游文档](https://github.com/ycui1984/pg_thrift/blob/1de265a661ab3c61aa593d7e99d3a313024170fc/README.md)

`pg_thrift` — 在 PostgreSQL 中编码、解码并查询 Apache Thrift 二进制与紧凑协议数据。

已复核目录快照记录的版本为 `1.0`、类型为 `standard`、实现语言为 `C`。
整理后的兼容版本集合为 `10,11`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pg_thrift";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_thrift';
```

整理后的生命周期为 `abandoned`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
