## 用法

来源：

- [官方扩展控制文件](https://github.com/aquameta/meta_triggers/blob/3bc68b67afea896300adf828b52d5dbbf1d63549/meta_triggers.control)

`meta_triggers` — meta_triggers：通用类 PostgreSQL 扩展；上游说明为“为 meta 系统目录提供插入、更新和删除触发器”。

已复核目录快照记录的版本为 `0.5.0`、类型为 `puresql`、实现语言为 `SQL`。
应先安装并验证声明的扩展依赖：`hstore`, `meta`, `plpgsql`。

```sql
CREATE EXTENSION "meta_triggers";
SELECT extversion
FROM pg_extension
WHERE extname = 'meta_triggers';
```

整理后的生命周期为 `active`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
