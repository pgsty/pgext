## 用法

来源：

- [官方上游 README](https://api.pgxn.org/src/pg_start_sql/pg_start_sql-0.0.2/README.md)
- [官方扩展控制文件 (pg_start_sql.control)](https://api.pgxn.org/src/pg_start_sql/pg_start_sql-0.0.2/pg_start_sql.control)

`pg_start_sql` — PostgreSQL 扩展，用于在实例启动时执行 SQL 语句。当需要管理或自动化上述描述的数据库行为时，请使用此扩展。请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION pg_start_sql;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 要求与注意事项

- 审查过的控制文件声明默认版本为 `1.0`。
- 控制文件将扩展标记为可重定位。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为以及失败情况。
