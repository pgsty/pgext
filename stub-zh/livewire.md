## 用法

来源：

- [官方扩展控制文件](https://github.com/b4oshany/livewire/blob/547e59a951fe5d088d89039a2d078186e73c7862/livewire.control)

`livewire` — livewire：通用类 PostgreSQL 扩展；上游说明为“电力输送建模”。

已复核目录快照记录的版本为 `0.2.0`、类型为 `puresql`、实现语言为 `SQL`。
应先安装并验证声明的扩展依赖：`pgrouting`, `plpgsql`, `postgis`, `postgis_topology`。
整理后的兼容版本集合为 `10`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "livewire";
SELECT extversion
FROM pg_extension
WHERE extname = 'livewire';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
