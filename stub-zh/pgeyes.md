## 用法

来源：

- [官方上游文档](https://github.com/ycku/pgeyes/blob/e09e243faaf4837a05366c3fc9d573909dbd053d/README.md)

`pgeyes` — pgeyes：PL/pgSQL 管理辅助扩展，提供 pgeyes() 环境检查函数与 foreign_db 元数据表。

已复核目录快照记录的版本为 `0.0.2`、类型为 `puresql`、实现语言为 `SQL`。
应先安装并验证声明的扩展依赖：`plpgsql`。
整理后的兼容版本集合为 `10,11,12,13,14,15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pgeyes";
SELECT extversion
FROM pg_extension
WHERE extname = 'pgeyes';
```

整理后的生命周期为 `abandoned`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
