## 用法

来源：

- [官方扩展控制文件](https://github.com/paroga/pg-cbor/blob/23afc682e6168e59c4d27d64fb10e8c5e624722a/cbor.control)
- [官方上游文档](https://github.com/paroga/pg-cbor/blob/23afc682e6168e59c4d27d64fb10e8c5e624722a/README.md)

`cbor` — cbor：为 PostgreSQL 提供 Concise Binary Object Representation 数据类型。

已复核目录快照记录的版本为 `0.1.0`、类型为 `standard`、实现语言为 `C`。
整理后的兼容版本集合为 `10,11,12,13,14,15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "cbor";
SELECT extversion
FROM pg_extension
WHERE extname = 'cbor';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
