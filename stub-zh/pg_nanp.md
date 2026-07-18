## 用法

来源：

- [官方扩展控制文件](https://github.com/rlichtenwalter/pg_nanp/blob/8cfb5213faaeef92a19997978f0b6801f0f61aa0/pg_nanp.control)

`pg_nanp` — 用于校验并统一格式化北美号码规划电话号码的 PostgreSQL 数据类型

已复核目录快照记录的版本为 `1.0`、类型为 `standard`、实现语言为 `C`。

```sql
CREATE EXTENSION "pg_nanp";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_nanp';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
