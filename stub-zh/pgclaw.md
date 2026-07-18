## 用法

来源：

- [官方扩展控制文件](https://github.com/calebwin/pgclaw/blob/ea6c3abe27602724f01de199b87eac8cfcd47212/pgclaw.control)
- [官方上游文档](https://github.com/calebwin/pgclaw/blob/ea6c3abe27602724f01de199b87eac8cfcd47212/README.md)
- [官方 Rust 包清单](https://github.com/calebwin/pgclaw/blob/ea6c3abe27602724f01de199b87eac8cfcd47212/Cargo.toml)

`pgclaw` — 将 LLM 或 OpenClaw 智能体作为 PostgreSQL 值存储，并通过后台工作进程处理受监控的数据行

已复核目录快照记录的版本为 `0.1.0`、类型为 `preload`、实现语言为 `Rust`。
整理后的兼容版本集合为 `14,15,16,17`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pgclaw";
SELECT extversion
FROM pg_extension
WHERE extname = 'pgclaw';
```

整理后的生命周期为 `active`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
