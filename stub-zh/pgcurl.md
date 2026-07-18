## 用法

来源：

- [官方上游文档](https://github.com/pandrewhk/pgcurl/blob/97d1e7e7447adae345fd90695916031237e1b06b/README.md)

`pgcurl` — pgcurl：数据集成类 PostgreSQL 扩展；上游说明为“从 PostgreSQL 调用 curl”。

已复核目录快照记录的版本为 `0.0.1`、类型为 `standard`、实现语言为 `C`。

```sql
CREATE EXTENSION "pgcurl";
SELECT extversion
FROM pg_extension
WHERE extname = 'pgcurl';
```

整理后的生命周期为 `abandoned`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
