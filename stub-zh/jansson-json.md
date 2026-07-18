## 用法

来源：

- [官方上游文档](https://github.com/claesjac/pg-json/blob/551e0067fe7731778362aa3eb265b810002ba9de/README.md)

`jansson-json` — 基于 Jansson JSON 库的 JSON 数据类型、操作符与函数扩展。

已复核目录快照记录的版本为 `0.0.2`、类型为 `standard`、实现语言为 `C`。
整理后的兼容版本集合为 `10,11,12,13,14,15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "jansson-json";
SELECT extversion
FROM pg_extension
WHERE extname = 'jansson-json';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
