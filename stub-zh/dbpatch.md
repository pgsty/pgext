## 用法

来源：

- [官方上游 README](https://api.pgxn.org/src/dbpatch/dbpatch-1.3.0/README.md)
- [官方扩展 SQL (dbpatch--unpackaged--1.0.0.sql)](https://api.pgxn.org/src/dbpatch/dbpatch-1.3.0/sql/dbpatch--unpackaged--1.0.0.sql)

`dbpatch` — postgresql-dbpatch ==================

在管理或自动化上述数据库行为时使用它。在使用之前，必须先安装并验证其扩展依赖项。

### 核心工作流

经过审核的分发包使用了过时的 SQL 或非控制文件安装布局，因此没有建立现代独立的 `CREATE EXTENSION` 和升级工作流。请遵循固定在上游的安装机制，并在隔离数据库中验证安装的对象。

### 重要对象

- `IF` 是由扩展安装或管理的表。

### 要求与注意事项

- 该目录记录了版本信息 `1.3.0`。
- 在生产使用之前，先安装并验证确认的扩展依赖项：`plpgsql`。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为以及失败情况。
