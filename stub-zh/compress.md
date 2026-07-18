## 用法

来源：

- [官方上游文档](https://github.com/snaiffer/pg_compress/blob/4d254039d5bd36f4a8c8bc9add225cf4f0c9b92e/README.md)

`compress` — 使用 C/C++ 将文本压缩为 Base64 编码的 ZIP 归档，或从中解压文本。

已复核目录快照记录的版本为 `0.0.2`、类型为 `standard`、实现语言为 `C++`。

```sql
CREATE EXTENSION "compress";
SELECT extversion
FROM pg_extension
WHERE extname = 'compress';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
