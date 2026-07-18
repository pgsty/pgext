## 用法

来源：

- [官方扩展控制文件](https://github.com/pgEdge/pgedge-safesession/blob/eae953e2d2e90bf8a4c912073f4b6193001028c1/pgedge_safesession.control)
- [官方上游文档](https://github.com/pgEdge/pgedge-safesession/blob/eae953e2d2e90bf8a4c912073f4b6193001028c1/docs/index.md)
- [官方上游 README](https://github.com/pgEdge/pgedge-safesession/blob/eae953e2d2e90bf8a4c912073f4b6193001028c1/README.md)

`pgedge_safesession` — 为指定 PostgreSQL 角色强制只读会话，阻止 DML、DDL 等写入操作。

已复核目录快照记录的版本为 `1.0-alpha1`、类型为 `preload`、实现语言为 `C`。
整理后的兼容版本集合为 `14,15,16,17,18`；仍需针对目标服务器确认实际构建。

整理记录表明该组件没有独立扩展 DDL；只能按上游说明的加载或服务商流程集成。

该上游项目与 `pgEdge` 相关；应根据所链接来源核实当前支持、许可证、打包方式与部署边界。

整理后的生命周期为 `preview`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
