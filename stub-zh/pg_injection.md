## 用法

来源：

- [官方上游 README](https://api.pgxn.org/src/pg_injection/pg_injection-1.0.0/README.md)

`pg_injection` — PostgreSQL 注入。在实现相应的安全、审计或访问控制工作流时使用它。使用上述链接的固定上游版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

经过审核的分发使用了过时的 SQL 或非控制文件安装布局，因此没有建立现代独立的 `CREATE EXTENSION` 和升级工作流。遵循固定上游的安装机制，并在隔离数据库中验证安装的对象。

### 要求与注意事项

- 该目录记录了版本 `1.0.0`。
- 在生产使用前，确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源进行对比。
