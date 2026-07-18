## 用法

来源：

- [官方扩展控制文件](https://github.com/adjust/all_substrings_tokenizer/blob/f459a88981981659b7f21fab955ca1cbc802c219/all_substrings_tokenizer.control)
- [官方上游文档](https://github.com/adjust/all_substrings_tokenizer/blob/f459a88981981659b7f21fab955ca1cbc802c219/README.md)

`all_substrings_tokenizer` — 从字符串中提取长度超过固定字符阈值的所有子串。

已复核目录快照记录的版本为 `0.1`、类型为 `standard`、实现语言为 `C`。

```sql
CREATE EXTENSION "all_substrings_tokenizer";
SELECT extversion
FROM pg_extension
WHERE extname = 'all_substrings_tokenizer';
```

整理后的生命周期为 `archived`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
