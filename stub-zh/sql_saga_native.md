## 用法

来源：

- [官方上游 README](https://github.com/veridit/sql_saga/blob/9e39fcc7fc862fcb793f59d996eb01268cf9e822/README.md)

`sql_saga_native` — 为 sql_saga 的 temporal_merge 规划与执行器内省提供加速的 Rust/pgrx 伴生库。

已复核目录快照记录的版本为 `0.1.0`、类型为 `headless`、实现语言为 `Rust`。
应先安装并验证声明的扩展依赖：`sql_saga`。
整理后的兼容版本集合为 `18`；仍需针对目标服务器确认实际构建。

整理记录表明该组件没有独立扩展 DDL；只能按上游说明的加载或服务商流程集成。
官方材料包含实验性、弃用、不支持或明确警告边界；用于非实验环境前必须阅读全文并测试失败场景。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
