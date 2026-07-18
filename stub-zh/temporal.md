## 用法

来源：

- [官方扩展控制文件](https://api.pgxn.org/src/temporal/temporal-0.7.1/temporal.control)
- [官方上游文档](https://pgxn.org/dist/temporal/doc/html/reference.html)
- [官方 PGXN 分发页](https://pgxn.org/dist/temporal/)

`temporal` — C 扩展：提供基于带时区时间戳的 PERIOD 类型、时间区间关系运算符与函数，以及 GiST 索引支持。

已复核目录快照记录的版本为 `0.7.1`、类型为 `standard`、实现语言为 `C`。

```sql
CREATE EXTENSION "temporal";
SELECT extversion
FROM pg_extension
WHERE extname = 'temporal';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
