## 用法

来源：

- [官方扩展控制文件](https://github.com/XiaYingyin/ctidscan/blob/101a177a42ee94c0aac5bffa8ca63ae9f0790d5d/ctidscan.control)
- [官方上游文档](https://github.com/XiaYingyin/ctidscan/blob/101a177a42ee94c0aac5bffa8ca63ae9f0790d5d/README.md)

`ctidscan` — 使用 Custom Scan Provider API 加速 CTID 范围条件的示例模块。

已复核目录快照记录的版本为 `1.1`、类型为 `headless`、实现语言为 `C`。

整理记录表明该组件没有独立扩展 DDL；只能按上游说明的加载或服务商流程集成。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
