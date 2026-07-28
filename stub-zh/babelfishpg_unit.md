## 用法

来源：

- [官方上游 README](https://github.com/babelfish-for-postgresql/babelfish_extensions/blob/46617a93de0eb666ce98591cfcbae4554f6f6ea0/contrib/babelfishpg_unit/README.md)
- [官方扩展控制文件 (babelfishpg_unit.control)](https://github.com/babelfish-for-postgresql/babelfish_extensions/blob/46617a93de0eb666ce98591cfcbae4554f6f6ea0/contrib/babelfishpg_unit/babelfishpg_unit.control)
- [官方扩展 SQL (babelfishpg_unit--1.0.0.sql)](https://github.com/babelfish-for-postgresql/babelfish_extensions/blob/46617a93de0eb666ce98591cfcbae4554f6f6ea0/contrib/babelfishpg_unit/babelfishpg_unit--1.0.0.sql)

`babelfishpg_unit` — Babelfish 引入了一个名为 babelfishpg_unit 的新扩展，使我们能够运行单元测试。请遵循构建说明来构建并安装 babelfishpg_unit 扩展。当应用程序需要此特定数据库功能时，请使用它。在目标 PostgreSQL 构建上使用链接的上游版本作为 API 边界，并进行测试。

### 核心工作流

```sql
CREATE EXTENSION babelfishpg_unit;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证安装版本和返回值。

### 重要对象

- `babelfishpg_unit.babelfishpg_unit_run_tests()` 是一个扩展函数，返回 `TABLE`。
- `babelfishpg_unit.babelfishpg_unit_run_tests(VARIADIC name text[])` 是一个扩展函数，返回 `TABLE`。

### 要求与注意事项

- 审查后的控制文件声明默认版本为 `1.0.0`。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为以及失败情况。
