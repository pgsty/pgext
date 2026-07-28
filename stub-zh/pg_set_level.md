## 用法

来源：

- [官方上游 README](https://api.pgxn.org/src/pg_set_level/pg_set_level-0.0.1/README.md)
- [官方扩展控制文件 (pg_set_level.control)](https://api.pgxn.org/src/pg_set_level/pg_set_level-0.0.1/pg_set_level.control)
- [官方扩展 SQL (pg_set_level--0.0.1.sql)](https://api.pgxn.org/src/pg_set_level/pg_set_level-0.0.1/pg_set_level--0.0.1.sql)

`pg_set_level` — pg_set_level 是一个 PostgreSQL 扩展，允许自定义 SET 语句。在管理或自动化上述数据库行为时使用它。请使用链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION pg_set_level;
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行它，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 要求与注意事项

- 审查过的控制文件声明默认版本为 `0.0.1`。
- 控制文件标记该扩展为不可重定位。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为以及失败情况，以匹配固定源代码。
