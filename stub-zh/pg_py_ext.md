## 用法

来源：

- [官方上游 README](https://api.pgxn.org/src/pg_py_ext/pg_py_ext-1.0.0/README.md)
- [官方扩展控制文件 (pg_py_ext.control)](https://api.pgxn.org/src/pg_py_ext/pg_py_ext-1.0.0/pg_py_ext.control)
- [官方扩展 SQL (pg_py_ext--1.0.0.sql)](https://api.pgxn.org/src/pg_py_ext/pg_py_ext-1.0.0/pg_py_ext--1.0.0.sql)

`pg_py_ext` — **一个使用 PL/Python3U 扩展来添加数字的 PostgreSQL 扩展**。当 SQL 需要这些特殊函数或聚合时，请使用此扩展。在目标 PostgreSQL 构建中测试链接的上游修订版本作为 API 边界。

### 核心工作流

```sql
CREATE EXTENSION pg_py_ext;
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行该示例，并在将其集成到应用程序 SQL 中之前验证安装的版本和返回值。

### 重要对象

- `add_numbers(a integer, b integer)` 是一个扩展函数，返回 `integer`。

### 要求与注意事项

- 审查后的控制文件声明默认版本为 `1.0.0`。
- 控制文件将扩展标记为可重定位。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况。
