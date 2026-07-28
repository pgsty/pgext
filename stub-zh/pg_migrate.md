## 用法

来源：

- [官方上游 README](https://api.pgxn.org/src/pg_migrate/pg_migrate-0.1.1/README.md)

`pg_migrate` — pg_migrate 是一个 PostgreSQL 扩展和命令行界面，允许你对表和索引进行模式更改。与 ALTER TABLE 不同，它在线工作，在迁移过程中不会持有长时间的独占锁。它会构建目标表的副本并进行替换。在从 PostgreSQL 移动、转换或集成相应数据时使用它。使用上方链接的固定上游版本作为 API 边界，并在目标 PostgreSQL 构建中进行测试。

### 核心工作流

经过审核的分发使用了过时的 SQL 或非控制文件安装布局，因此不会建立现代独立的 `CREATE EXTENSION` 和升级工作流。遵循固定上游的安装机制，并在隔离数据库中验证安装的对象。

### 要求与注意事项

- 目录记录 `0.1.1` 的版本。
- 在生产使用前，确认权限、支持的 PostgreSQL 版本、升级行为和失败情况，与固定源进行验证。
