## 用法

来源：

- [官方上游文档](https://github.com/maciekgajewski/postgres-uints/blob/4bdc1d98729f77ab6dd32fa563a1c7b7c7d3559d/README.md)

`uints` — 提供无符号 16 位与 32 位整数类型及类型转换、运算符和数值支持函数的 C 扩展。

已复核目录快照记录的版本为 `0.9`、类型为 `standard`、实现语言为 `C`。

```sql
CREATE EXTENSION "uints";
SELECT extversion
FROM pg_extension
WHERE extname = 'uints';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
