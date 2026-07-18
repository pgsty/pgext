## 用法

来源：

- [官方扩展控制文件](https://github.com/MasaoFujii/postgres_security/blob/8a5d11461b5381b4cd3a7cae70d1d97bb73f2711/postgres_security.control)

`postgres_security` — 定义一个复制 PostgreSQL text 存储与输入输出行为的模式内文本类型

已复核目录快照记录的版本为 `1.0.0`、类型为 `standard`、实现语言为 `C`。

```sql
CREATE EXTENSION "postgres_security";
SELECT extversion
FROM pg_extension
WHERE extname = 'postgres_security';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
