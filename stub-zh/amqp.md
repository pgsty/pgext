## 用法

来源：

- [官方扩展控制文件](https://github.com/omniti-labs/pg_amqp/blob/240d477d40c5e7a579b931c98eb29cef4edda164/amqp.control)
- [官方上游文档](https://github.com/omniti-labs/pg_amqp/blob/240d477d40c5e7a579b931c98eb29cef4edda164/doc/amqp.md)
- [官方 PGXN 分发页](https://pgxn.org/dist/pg_amqp/)

`amqp` — 允许 PostgreSQL 语句向 AMQP 消息代理发布消息。

已复核目录快照记录的版本为 `0.4.2`、类型为 `preload`、实现语言为 `C`。
整理后的兼容版本集合为 `10,11,12,13,14,15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "amqp";
SELECT extversion
FROM pg_extension
WHERE extname = 'amqp';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
