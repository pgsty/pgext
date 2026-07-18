## 用法

来源：

- [官方扩展控制文件](https://github.com/microsoft/pg_paxos/blob/7d4f656b1cfa0943626b658fbc23a375eab41c3a/pg_paxos.control)
- [官方上游文档](https://github.com/microsoft/pg_paxos/blob/7d4f656b1cfa0943626b658fbc23a375eab41c3a/README.md)

`pg_paxos` — 为 PostgreSQL 提供 Paxos 分布式共识函数和基于 Paxos 的表复制。

已复核目录快照记录的版本为 `0.2`、类型为 `preload`、实现语言为 `C`。
应先安装并验证声明的扩展依赖：`dblink`, `plpgsql`。

```sql
CREATE EXTENSION "pg_paxos";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_paxos';
```

整理后的生命周期为 `archived`。采用前应固定已复核构建并确认维护状态。
官方材料包含实验性、弃用、不支持或明确警告边界；用于非实验环境前必须阅读全文并测试失败场景。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
