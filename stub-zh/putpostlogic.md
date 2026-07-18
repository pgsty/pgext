## 用法

来源：

- [官方扩展控制文件](https://github.com/rektide/putpostlogic/blob/0d71733aa77085b4f62c49c0bba972c241d0fcbc/putpostlogic.control)

`putpostlogic` — 未发布的逻辑解码输出插件原型，可将 JSON 变更通过 PostgreSQL、Kafka 或 Nanomsg 输出。

已复核目录快照记录的版本为 `0.1.0`、类型为 `headless`、实现语言为 `C`。

整理记录表明该组件没有独立扩展 DDL；只能按上游说明的加载或服务商流程集成。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
