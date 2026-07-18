## 用法

来源：

- [官方扩展控制文件](https://github.com/cahutton/text_domains/blob/a4e1aa7463e07ffdf1ce9f2d380f054bf1af1b5c/text_domains.control)

`text_domains` — 纯 SQL 的 PostgreSQL 文本域集合，为非空、字母、字母数字、大小写与固定长度文本提供约束校验。

已复核目录快照记录的版本为 `1.0`、类型为 `puresql`、实现语言为 `SQL`。
整理后的兼容版本集合为 `10`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "text_domains";
SELECT extversion
FROM pg_extension
WHERE extname = 'text_domains';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
