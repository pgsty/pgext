## 用法

来源：

- [官方扩展控制文件](https://github.com/logplex/pg_logfebe/blob/bac15ef2386fc70197ba889e23027bffe00d0f5b/pg_logfebe.control)

`pg_logfebe` — 通过 Unix 套接字以 PostgreSQL FE/BE 帧格式发送结构化日志。

已复核目录快照记录的版本为 `1.0`、类型为 `preload`、实现语言为 `C`。

整理记录表明该组件没有独立扩展 DDL；只能按上游说明的加载或服务商流程集成。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
