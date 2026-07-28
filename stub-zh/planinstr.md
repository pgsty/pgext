## 用法

来源：

- [官方上游 README](https://api.pgxn.org/src/planinstr/planinstr-0.0.1/README.md)

`planinstr` — 确保你已经安装了 pg_config 并将其添加到路径中。如果你使用的是 RPM 包管理器安装 PostgreSQL，请确保也安装了 -devel 包。如果需要，可以在构建过程中指定其位置：使用它来收集或解释相应的 PostgreSQL 统计信息。使用链接的锁定上游版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

此组件在审查源代码中没有确认的独立 `CREATE EXTENSION` 工作流。仅通过上游机制构建、加载或启用它，然后在隔离数据库中验证其结果行为。

### 要求与注意事项

- 该目录记录了版本信息 `0.0.1`。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与锁定的源代码一致。
