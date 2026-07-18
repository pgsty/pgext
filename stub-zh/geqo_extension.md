## 用法

来源：

- [官方上游文档](https://github.com/parkag/geqo/blob/0f5bea3a679c9719bd8d48d97e1ea84f1d77bbed/README.md)

`geqo_extension` — 从 PostgreSQL 中提取的仅加载式遗传查询优化器，通过规划器钩子提供可配置的 GEQO 行为。

已复核目录快照记录的版本为 `0.0.1`、类型为 `headless`、实现语言为 `C`。

整理记录表明该组件没有独立扩展 DDL；只能按上游说明的加载或服务商流程集成。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
