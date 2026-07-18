## 用法

来源：

- [官方项目或服务商页面](https://fasttransfer.arpe.io/)

`pg_fasttransfer` — 从 SQL 函数调用 FastTransfer 工具，在 PostgreSQL 数据库之间高速传输数据。

已复核目录快照记录的版本为 `1.0`、类型为 `standard`、实现语言为 `C`。
应先安装并验证声明的扩展依赖：`pgcrypto`。
整理后的兼容版本集合为 `15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pg_fasttransfer";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_fasttransfer';
```

整理后的生命周期为 `active`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
