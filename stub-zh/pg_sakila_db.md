## 用法

来源：

- [官方上游 README](https://github.com/benizar/pg_sakila_db/blob/85f43e570893f27577ba273b8b9853a2de7438b5/README.md)
- [官方扩展控制文件 (pg_sakila_db.control)](https://github.com/benizar/pg_sakila_db/blob/85f43e570893f27577ba273b8b9853a2de7438b5/pg_sakila_db.control)

`pg_sakila_db` — 引言 Postgres xtensions 贡献者指南 开始使用 安装扩展 依赖关系 相关项目 待办事项。在需要此特定数据库功能的应用程序中使用它。在集成到应用程序 SQL 之前，必须先安装并验证其扩展依赖关系。

### 核心工作流

```sql
CREATE EXTENSION pg_sakila_db;
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行该示例，并在将其集成到应用程序 SQL 之前验证已安装的版本和返回值。

### 要求与注意事项

- 审核的控制文件声明默认版本为 `0.0.1`。
- 首先安装并验证确认的扩展依赖关系：`plpgsql`。
- 控制文件将该扩展标记为可重定位。
- 在生产使用之前，确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源保持一致。
