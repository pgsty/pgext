## 用法

来源：

- [官方上游文档](https://github.com/koctep/jsonb-extend/blob/fb704ad2b19a8d61d7cc0bcd50ecf507c629dba1/README.md)

`jsonb_extend` — 合并两个或多个 jsonb 值的 C 扩展函数。

已复核目录快照记录的版本为 `1.0`、类型为 `standard`、实现语言为 `C`。
整理后的兼容版本集合为 `10,11,12,13,14,15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "jsonb_extend";
SELECT extversion
FROM pg_extension
WHERE extname = 'jsonb_extend';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
