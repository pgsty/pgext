## 用法

来源：

- [官方 PGXN 分发页](https://pgxn.org/dist/shard_manager/)

`shard_manager` — 基于 SQL/PLpgSQL 的模式分片工具，提供分片映射、模板表克隆与 64 位分布式 ID 生成。

已复核目录快照记录的版本为 `0.0.1`、类型为 `puresql`、实现语言为 `SQL`。

```sql
CREATE EXTENSION "shard_manager";
SELECT extversion
FROM pg_extension
WHERE extname = 'shard_manager';
```

整理后的生命周期为 `abandoned`。采用前应固定已复核构建并确认维护状态。
官方材料包含实验性、弃用、不支持或明确警告边界；用于非实验环境前必须阅读全文并测试失败场景。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
