## 用法

来源：

- [官方上游 README](https://api.pgxn.org/src/pgds/pgds-0.0.3/README.md)
- [官方扩展控制文件 (pgds.control)](https://api.pgxn.org/src/pgds/pgds-0.0.3/pgds.control)
- [官方扩展 SQL (pgds--0.0.3.sql)](https://api.pgxn.org/src/pgds/pgds-0.0.3/pgds--0.0.3.sql)

`pgds` — PostgreSQL 扩展，用于收集动态统计信息。在管理或自动化上述数据库行为时使用此扩展。请使用链接中的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION pgds;
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行该示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `find_tables(p_oid oid)` 是一个扩展函数，返回 `TABLE`。

### 要求与注意事项

- 审查过的控制文件声明默认版本为 `0.0.1`。
- 控制文件标记该扩展为不可重定位。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况。
