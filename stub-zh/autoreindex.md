## 用法

来源：

- [官方扩展控制文件](https://github.com/okbob/autoreindex/blob/363e5299bda99f9e5166c228fd82c8845bf77500/autoreindex.control)

`autoreindex` — 需要 shared_preload_libraries 加载的后台 worker 扩展，用于扫描膨胀索引并在后台执行并发重建索引维护。

已复核目录快照记录的版本为 `1.0`、类型为 `headless`、实现语言为 `C`。

整理记录表明该组件没有独立扩展 DDL；只能按上游说明的加载或服务商流程集成。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
