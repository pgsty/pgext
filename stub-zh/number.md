## 用法

来源：

- [官方扩展控制文件](https://github.com/df7cb/postgresql-number/blob/e11fec180898e8c323bcc0bff4f22af6118628d4/number.control)

`number` — 为 PostgreSQL 提供可变宽度整数数据类型，用紧凑格式存储 64 位整数。

已复核目录快照记录的版本为 `1`、类型为 `standard`、实现语言为 `C`。

```sql
CREATE EXTENSION "number";
SELECT extversion
FROM pg_extension
WHERE extname = 'number';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
