## 用法

来源：

- [官方扩展控制文件](https://github.com/kotsachin/pg_shardman/blob/17820427be98e931aab1ef0f2813d526dcb4ac58/pg_shardman.control)
- [官方上游文档](https://github.com/kotsachin/pg_shardman/blob/17820427be98e931aab1ef0f2813d526dcb4ac58/readme.md)

`pg_shardman` — 在 PostgreSQL 10 中使用 pg_pathman、postgres_fdw 和逻辑复制管理跨节点哈希分片表。

已复核目录快照记录的版本为 `1.0`、类型为 `standard`、实现语言为 `C`。
应先安装并验证声明的扩展依赖：`pg_pathman`, `plpgsql`, `postgres_fdw`。
整理后的兼容版本集合为 `10`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pg_shardman";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_shardman';
```
官方材料包含实验性、弃用、不支持或明确警告边界；用于非实验环境前必须阅读全文并测试失败场景。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
