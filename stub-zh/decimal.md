## 用法

来源：

- [官方扩展控制文件](https://github.com/okbob/pgDecimal/blob/9a49dbb0ec73508af1f6ace9184ec6a1f93f0dc7/decimal.control)
- [官方上游文档](https://github.com/okbob/pgDecimal/blob/9a49dbb0ec73508af1f6ace9184ec6a1f93f0dc7/README.md)
- [官方 PGXN 分发页](https://pgxn.org/dist/pgdecimal/)

`decimal` — 提供基于 _Decimal32 与 _Decimal64 的十进制浮点数据类型，支持算术运算和类型转换。

已复核目录快照记录的版本为 `1.0`、类型为 `standard`、实现语言为 `C`。

```sql
CREATE EXTENSION "decimal";
SELECT extversion
FROM pg_extension
WHERE extname = 'decimal';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
