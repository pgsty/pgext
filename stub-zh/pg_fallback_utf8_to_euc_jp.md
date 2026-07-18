## 用法

来源：

- [官方扩展控制文件](https://github.com/MasaoFujii/pg_fallback_utf8_to_euc_jp/blob/101376656db7619d2f12e51b5b90a67b53966ff1/pg_fallback_utf8_to_euc_jp.control)
- [官方上游 README](https://github.com/MasaoFujii/pg_fallback_utf8_to_euc_jp/blob/101376656db7619d2f12e51b5b90a67b53966ff1/README.md)

`pg_fallback_utf8_to_euc_jp` — 提供 UTF-8 到 EUC_JP 的后备编码转换，将无法映射的字符替换为问号。

已复核目录快照记录的版本为 `1.0`、类型为 `standard`、实现语言为 `C`。

```sql
CREATE EXTENSION "pg_fallback_utf8_to_euc_jp";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_fallback_utf8_to_euc_jp';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
