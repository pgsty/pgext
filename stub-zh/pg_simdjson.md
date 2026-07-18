## 用法

来源：

- [官方扩展控制文件](https://github.com/erthalion/pg_simdjson/blob/4e5a20eab6296d48e41da6f255c70c16bbb3d413/pg_simdjson.control)
- [官方上游文档](https://github.com/erthalion/pg_simdjson/blob/4e5a20eab6296d48e41da6f255c70c16bbb3d413/README.md)

`pg_simdjson` — 将 JSON 文本转换为 jsonb 的 SIMD 加速解析器原型

已复核目录快照记录的版本为 `0.1`、类型为 `standard`、实现语言为 `C++`。

```sql
CREATE EXTENSION "pg_simdjson";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_simdjson';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
