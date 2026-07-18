## 用法

来源：

- [官方扩展控制文件](https://github.com/idrawone/get_sum/blob/43f8ed114da79321c4125c190044ab4e06e1322e/get_sum.control)

`get_sum` — get_sum：通用类 PostgreSQL 扩展；上游说明为“使用 C 编写、用于 PostgreSQL 的两个整数简单求和扩展”。

已复核目录快照记录的版本为 `0.0.1`、类型为 `standard`、实现语言为 `C`。

```sql
CREATE EXTENSION "get_sum";
SELECT extversion
FROM pg_extension
WHERE extname = 'get_sum';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
