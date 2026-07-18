## 用法

来源：

- [官方扩展控制文件](https://github.com/sean-/pg_consul/blob/37a6c84ec26f981a048cf5c7fc070660472dd773/pg_consul.control)

`pg_consul` — C++ PostgreSQL 扩展，通过 SQL 函数访问 Consul 状态、服务目录与键值 API。

已复核目录快照记录的版本为 `0.1.0`、类型为 `standard`、实现语言为 `C++`。
整理后的兼容版本集合为 `10,11,12,13,14,15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pg_consul";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_consul';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
