## 用法

来源：

- [官方扩展控制文件](https://github.com/pgstuff/base32_4b/blob/0518fe2e17a7e634405c79eb51ca5b36f0e1275c/base32_4b.control)
- [官方上游 README](https://github.com/pgstuff/base32_4b/blob/0518fe2e17a7e634405c79eb51ca5b36f0e1275c/README.md)

`base32_4b` — 提供以 4 字节存储的 Base32 数据类型及比较运算符。

已复核目录快照记录的版本为 `0.0.1`、类型为 `standard`、实现语言为 `C`。
整理后的兼容版本集合为 `10,11,12,13,14,15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "base32_4b";
SELECT extversion
FROM pg_extension
WHERE extname = 'base32_4b';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
