## 用法

来源：

- [官方上游 README](https://api.pgxn.org/src/pg_log_userqueries/pg_log_userqueries-1.0.0/README)

`pg_log_userqueries` — pg_log_userqueries 是一个 PostgreSQL 模块，用于记录超级用户执行的每个查询。它将每个查询记录在标准日志文件中。在收集或解释相应的 PostgreSQL 统计信息时使用它。请使用链接中的固定上游版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

经过审核的分发包使用了过时的 SQL 或非控制文件安装布局，因此没有建立现代独立的 `CREATE EXTENSION` 和升级工作流。请遵循固定上游的安装机制，并在隔离数据库中验证安装对象。

### 要求与注意事项

- 该目录记录版本 `1.0.0`。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为以及失败情况，与固定源进行比对。
