## 用法

来源：

- [官方扩展控制文件（cat_tools.control）](https://api.pgxn.org/src/cat_tools/cat_tools-0.2.1/cat_tools.control)
- [官方扩展 SQL（cat_tools--0.1.0--0.1.3.sql）](https://api.pgxn.org/src/cat_tools/cat_tools-0.2.1/sql/cat_tools--0.1.0--0.1.3.sql)

`cat_tools` — 用于与目录进行交互的工具。在管理或自动化上述数据库行为时使用它。请使用链接中的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建中进行测试。

### 核心工作流

```sql
CREATE EXTENSION cat_tools;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `__cat_tools.create_function(function_name text , args text , options text , body text , grants text DEFAULT NULL)` 是一个扩展函数，返回 `void`。
- `__cat_tools.exec(sql text)` 是一个扩展函数，返回 `void`。
- `pg_temp.create_function(function_name text , args text , options text , body text , grants text DEFAULT NULL)` 是一个扩展函数，返回 `void`。
- `pg_temp.exec(sql text)` 是一个扩展函数，返回 `void`。

### 要求与注意事项

- 审查过的控制文件声明默认版本为 `0.2.1`。
- 控制文件将扩展标记为不可重定位。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源进行比对。
