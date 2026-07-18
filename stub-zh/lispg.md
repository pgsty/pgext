## 用法

来源：

- [官方扩展控制文件](https://github.com/niquola/lispg/blob/e43d7961477259e910761218a5541211d2cf8860/lispg.control)

`lispg` — 为 PostgreSQL 提供 Lisp 风格表达式类型与求值器的 C 扩展。

已复核目录快照记录的版本为 `1.0`、类型为 `standard`、实现语言为 `C`。

```sql
CREATE EXTENSION "lispg";
SELECT extversion
FROM pg_extension
WHERE extname = 'lispg';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
