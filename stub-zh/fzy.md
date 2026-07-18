## 用法

来源：

- [官方上游文档](https://github.com/romgrk/pg_fzy/blob/0c0aab00004c738e86b031d618c215d900f1f5f4/README.md)

`fzy` — PostgreSQL 的 FZY 模糊匹配评分函数。

已复核目录快照记录的版本为 `0.0.1`、类型为 `standard`、实现语言为 `C`。
整理后的兼容版本集合为 `11,12,13`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "fzy";
SELECT extversion
FROM pg_extension
WHERE extname = 'fzy';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
