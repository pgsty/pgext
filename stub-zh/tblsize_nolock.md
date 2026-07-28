## 用法

来源：

- [官方上游 README](https://api.pgxn.org/src/tblsize_nolock/tblsize_nolock-1.0.0/README)

`tblsize_nolock` — 函数用于在不获取目标关系的 AccessShareLock 的情况下计算关系大小。当进行数据库管理或自动化上述行为时，请使用此功能。请使用上述链接的可重定位上游版本作为 API 边界，并在目标 PostgreSQL 构建中进行测试。

### 核心工作流

此组件在审查源代码中没有确认的独立 `CREATE EXTENSION` 工作流。仅通过精确的上游机制构建、加载或启用它，然后在隔离数据库中验证最终的服务器行为。

### 要求与注意事项

- 请确认目录项记录的版本 `1.0.0`。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源代码的一致性。
