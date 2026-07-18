## 用法

来源：

- [官方扩展控制文件](https://github.com/autodatadirect/uniqueidentifier/blob/85976b564e5cd3362c30ba388b95c3df19bc4b34/uniqueidentifier.control)

`uniqueidentifier` — 为 PostgreSQL 提供旧式 16 字节 uniqueidentifier 类型、比较运算符、类型转换和 newid() 生成函数。

已复核目录快照记录的版本为 `1.0`、类型为 `standard`、实现语言为 `C`。
整理后的兼容版本集合为 `16`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "uniqueidentifier";
SELECT extversion
FROM pg_extension
WHERE extname = 'uniqueidentifier';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
