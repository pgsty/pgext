## 用法

来源：

- [官方上游 README](https://api.pgxn.org/src/my_extension/my_extension-1.1.0/README.md)
- [官方扩展控制文件 (my_extension.control)](https://api.pgxn.org/src/my_extension/my_extension-1.1.0/my_extension.control)
- [官方扩展 SQL (my_extension--1.0.1.sql)](https://api.pgxn.org/src/my_extension/my_extension-1.1.0/my_extension--1.0.1.sql)

`my_extension` — my_extension 是一个基本的 PostgreSQL 扩展，提供了高效数据操作和计算的额外功能。当应用程序需要这种特定的数据库能力时，请使用它。请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION my_extension;

SELECT add(1, 2);
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行该示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `add(a integer, b integer)` 是一个扩展函数，返回 `integer`。
- `complex_add(integer[])` 是一个扩展函数，返回 `integer`。
- `multiply(a integer, b integer)` 是一个扩展函数，返回 `integer`。
- `my_table` 是一个由扩展安装或管理的表。

### 要求与注意事项

- 控制文件声明默认版本为 `1.1.0`。
- 控制文件将扩展标记为可重定位。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为以及失败情况与固定源进行比对。
