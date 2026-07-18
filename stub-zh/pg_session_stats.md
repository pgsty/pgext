## 用法

来源：

- [官方扩展控制文件](https://github.com/RyanMarcus/pg_session_stats/blob/8d114d5b653304a3039741c5cad8fee751bea1b3/pg_session_stats.control)
- [官方上游文档](https://github.com/RyanMarcus/pg_session_stats/blob/8d114d5b653304a3039741c5cad8fee751bea1b3/README.md)

`pg_session_stats` — 将会话及并行 worker 的 CPU、进程状态和 I/O 数据记录到外部 SQLite 数据库。

已复核目录快照记录的版本为 `0.0.1`、类型为 `headless`、实现语言为 `C`。
整理后的兼容版本集合为 `12`；仍需针对目标服务器确认实际构建。

整理记录表明该组件没有独立扩展 DDL；只能按上游说明的加载或服务商流程集成。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
