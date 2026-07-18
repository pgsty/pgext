## 用法

来源：

- [官方扩展控制文件](https://github.com/extio1/pg_isolate/blob/503ab1fd3c0844adc332c75626f96647e8b96ef7/pg_isolate.control)
- [官方上游文档](https://github.com/extio1/pg_isolate/blob/503ab1fd3c0844adc332c75626f96647e8b96ef7/README.md)

`pg_isolate` — 通过 Linux cgroups v2 隔离 PostgreSQL 后端与后台工作进程的扩展

已复核目录快照记录的版本为 `1.0`、类型为 `headless`、实现语言为 `C`。

整理记录表明该组件没有独立扩展 DDL；只能按上游说明的加载或服务商流程集成。

整理后的生命周期为 `active`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
