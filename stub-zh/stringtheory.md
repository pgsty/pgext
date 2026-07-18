## 用法

来源：

- [官方上游文档](https://github.com/hydradatabase/pg_stringtheory/blob/249f4cd83bdb6557ffb0143ed4f3407f71598224/README.md)

`stringtheory` — 为 PostgreSQL 文本值提供 SIMD 优化精确相等与子串搜索函数的 C++ 扩展。

已复核目录快照记录的版本为 `1.0.2`、类型为 `standard`、实现语言为 `C++`。
整理后的兼容版本集合为 `14,15,16`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "stringtheory";
SELECT extversion
FROM pg_extension
WHERE extname = 'stringtheory';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
