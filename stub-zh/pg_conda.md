## 用法

来源：

- [官方扩展控制文件](https://github.com/JeanChristopheMorinPerso/pg_conda/blob/bb2ee4f348de550a76599ab85bfcb65ebad2601c/pg_conda.control)

`pg_conda` — pg_conda：为 PostgreSQL 添加 Conda 生态版本类型、聚合与函数。

已复核目录快照记录的版本为 `1.0.1`、类型为 `standard`、实现语言为 `Rust`。
整理后的兼容版本集合为 `13,14,15,16,17`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pg_conda";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_conda';
```

整理后的生命周期为 `active`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
