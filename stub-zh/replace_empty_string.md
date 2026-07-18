## 用法

来源：

- [官方上游文档](https://github.com/okbob/replace_empty_string/blob/ee225b0b52a97ade45e52f6e71d2281c45c01b11/README.md)

`replace_empty_string` — 将空字符串替换为 NULL 的 C 语言触发器函数

已复核目录快照记录的版本为 `1.0`、类型为 `standard`、实现语言为 `C`。

```sql
CREATE EXTENSION "replace_empty_string";
SELECT extversion
FROM pg_extension
WHERE extname = 'replace_empty_string';
```

整理后的生命周期为 `deprecated`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
