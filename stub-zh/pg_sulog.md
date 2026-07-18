## 用法

来源：

- [官方扩展控制文件](https://github.com/nuko-yokohama/pg_sulog/blob/e07a87f76261ccfd884da42bff84419b395cba81/pg_sulog.control)
- [官方上游文档](https://github.com/nuko-yokohama/pg_sulog/blob/e07a87f76261ccfd884da42bff84419b395cba81/README.md)

`pg_sulog` — pg_sulog：记录并可阻断 PostgreSQL 超级用户操作的安全类扩展。

已复核目录快照记录的版本为 `1.0`、类型为 `preload`、实现语言为 `C`。

整理记录表明该组件没有独立扩展 DDL；只能按上游说明的加载或服务商流程集成。

整理后的生命周期为 `abandoned`。采用前应固定已复核构建并确认维护状态。
官方材料包含实验性、弃用、不支持或明确警告边界；用于非实验环境前必须阅读全文并测试失败场景。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
