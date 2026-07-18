## 用法

来源：

- [官方扩展控制文件](https://github.com/glynastill/pg_jsonb_opx/blob/adb985a8c3192e2577bf7ef863560baf4e566c5e/jsonb_opx.control)
- [官方上游 README](https://github.com/glynastill/pg_jsonb_opx/blob/adb985a8c3192e2577bf7ef863560baf4e566c5e/README.md)

`jsonb_opx` — 为 jsonb 提供 hstore 风格的函数和操作符。

已复核目录快照记录的版本为 `1.0`、类型为 `standard`、实现语言为 `C`。

```sql
CREATE EXTENSION "jsonb_opx";
SELECT extversion
FROM pg_extension
WHERE extname = 'jsonb_opx';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
