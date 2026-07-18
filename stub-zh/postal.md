## 用法

来源：

- [官方上游文档](https://github.com/pramsey/pgsql-postal/blob/105a6ca855dbf4d33b160ca50d00c3c76a7e2061/README.md)
- [官方 PGXN 分发页](https://pgxn.org/dist/postal/)

`postal` — postal 为 PostgreSQL 提供 libpostal 绑定，可标准化自由格式地址并将其解析为带标签的组成部分。

已复核目录快照记录的版本为 `1.0`、类型为 `standard`、实现语言为 `C`。
整理后的兼容版本集合为 `10,11,12,13,14,15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "postal";
SELECT extversion
FROM pg_extension
WHERE extname = 'postal';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
