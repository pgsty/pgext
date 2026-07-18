## 用法

来源：

- [官方 PGXN 分发页](https://pgxn.org/dist/pg_idx_advisor/)

`pg_idx_advisor` — 分析查询计划并建议候选索引的 PostgreSQL 扩展，现已停止更新。

已复核目录快照记录的版本为 `0.1.2`、类型为 `preload`、实现语言为 `C`。
应先安装并验证声明的扩展依赖：`plpgsql`。
整理后的兼容版本集合为 `10,11,12,13,14,15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pg_idx_advisor";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_idx_advisor';
```

整理后的生命周期为 `deprecated`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
