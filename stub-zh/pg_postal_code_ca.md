## 用法

来源：

- [官方扩展控制文件](https://github.com/rlichtenwalter/pg_postal_code_ca/blob/30db999d5f5e89007882afd3c2ca2a499ac755c7/pg_postal_code_ca.control)

`pg_postal_code_ca` — 高效校验、规范格式化并支持 B-tree 排序的加拿大邮政编码类型。

已复核目录快照记录的版本为 `1.0`、类型为 `standard`、实现语言为 `C`。

```sql
CREATE EXTENSION "pg_postal_code_ca";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_postal_code_ca';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
