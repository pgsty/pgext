## 用法

来源：

- [官方扩展控制文件](https://github.com/melanieplageman/format_x/blob/e65255bcb23aa17a4e00fab66d9e77908fc2b69b/format_x.control)
- [官方 PGXN 分发页](https://pgxn.org/dist/format_x/)
- [官方上游 README](https://github.com/melanieplageman/format_x/blob/e65255bcb23aa17a4e00fab66d9e77908fc2b69b/README.md)

`format_x` — 支持在复合类型、JSONB 与 HSTORE 值中按名称查找的多态可变参数字符串格式化函数。

已复核目录快照记录的版本为 `0.0.1`、类型为 `standard`、实现语言为 `C`。
整理后的兼容版本集合为 `10,11,12,13,14,15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "format_x";
SELECT extversion
FROM pg_extension
WHERE extname = 'format_x';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
