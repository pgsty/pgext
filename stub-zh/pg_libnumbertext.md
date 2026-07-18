## 用法

来源：

- [官方扩展控制文件](https://github.com/blm768/pg-libnumbertext/blob/67262d721c47c052b4a24d8ad3dd0623b8cbf9a5/pg_libnumbertext.control)
- [官方上游文档](https://github.com/blm768/pg-libnumbertext/blob/67262d721c47c052b4a24d8ad3dd0623b8cbf9a5/README.md)

`pg_libnumbertext` — 通过 libnumbertext 将数字转换为多种语言的文字表达。

已复核目录快照记录的版本为 `0.1.0`、类型为 `standard`、实现语言为 `C++`。

```sql
CREATE EXTENSION "pg_libnumbertext";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_libnumbertext';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
