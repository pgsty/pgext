## 用法

来源：

- [官方扩展控制文件](https://github.com/anayrat/pg_sampletolog/blob/478aa18638c987efba6403e619f2ff41ad0e1d48/pg_sampletolog.control)
- [官方上游文档](https://github.com/anayrat/pg_sampletolog/blob/478aa18638c987efba6403e619f2ff41ad0e1d48/README.md)
- [官方 PGXN 分发页](https://pgxn.org/dist/pg_sampletolog/)

`pg_sampletolog` — 通过执行器与实用语句 hook，将语句或完整事务按比例采样写入 PostgreSQL 日志。

已复核目录快照记录的版本为 `2.0.0`、类型为 `headless`、实现语言为 `C`。
整理后的兼容版本集合为 `10,11`；仍需针对目标服务器确认实际构建。

整理记录表明该组件没有独立扩展 DDL；只能按上游说明的加载或服务商流程集成。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
