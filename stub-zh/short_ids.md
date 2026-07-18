## 用法

来源：

- [官方 PGXN 分发页](https://pgxn.org/dist/short_ids/)
- [官方上游来源](https://gitlab.com/depesz/short_ids)

`short_ids` — 通过冲突检查与事务级咨询锁生成短随机文本 ID 的 PL/pgSQL 扩展。

已复核目录快照记录的版本为 `0.1.0`、类型为 `puresql`、实现语言为 `SQL`。
应先安装并验证声明的扩展依赖：`plpgsql`。
整理后的兼容版本集合为 `10,11,12,13,14,15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "short_ids";
SELECT extversion
FROM pg_extension
WHERE extname = 'short_ids';
```

整理后的生命周期为 `abandoned`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
