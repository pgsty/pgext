## 用法

来源：

- [官方上游 README](https://api.pgxn.org/src/log_functions/log_functions-1.0.0/README)

`log_functions` — log_functions 是一个 PostgreSQL 模块，记录每个执行的函数。在收集或解释相应的 PostgreSQL 统计信息时使用它。使用上方链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

审核过的分发包使用了过时的 SQL 或非控制文件安装布局，因此不会建立现代独立的 `CREATE EXTENSION` 和升级工作流。遵循上方链接的上游安装机制，并在隔离数据库中验证安装的对象。

### 要求与注意事项

- 目录记录 `1.0.0` 的版本。
- 在生产使用前，确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与上方链接的源代码一致。
