## 用法

来源：

- [官方扩展控制文件](https://github.com/instructure/jsoncdc/blob/46a04f733b554d792c2609f50831a5edd5a256af/jsoncdc.control)
- [官方上游文档](https://github.com/instructure/jsoncdc/blob/46a04f733b554d792c2609f50831a5edd5a256af/README.md)
- [官方 Rust 包清单](https://github.com/instructure/jsoncdc/blob/46a04f733b554d792c2609f50831a5edd5a256af/Cargo.toml)

`jsoncdc` — 将 PostgreSQL 逻辑解码变更流转换为逐行 JSON 的 Rust 输出插件。

已复核目录快照记录的版本为 `0.1.0`、类型为 `headless`、实现语言为 `Rust`。

整理记录表明该组件没有独立扩展 DDL；只能按上游说明的加载或服务商流程集成。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
