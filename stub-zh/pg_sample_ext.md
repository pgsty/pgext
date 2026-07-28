## 用法

来源：

- [官方上游 README](https://api.pgxn.org/src/pg_sample_ext/pg_sample_ext-1.0.2/README.md)
- [官方扩展控制文件 (pg_sample_ext.control)](https://api.pgxn.org/src/pg_sample_ext/pg_sample_ext-1.0.2/pg_sample_ext.control)
- [官方扩展 SQL (pg_sample_ext--1.0.0.sql)](https://api.pgxn.org/src/pg_sample_ext/pg_sample_ext-1.0.2/pg_sample_ext--1.0.0.sql)

`pg_sample_ext` — pg_sample_ext 是一个 PostgreSQL 扩展，提供了一个示例函数来演示如何扩展 PostgreSQL 的功能。当需要这些特殊函数或聚合时，请使用它。请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION pg_sample_ext;

SELECT square(5);  -- Returns 25
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行该示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `square(num integer)` 是一个扩展函数，返回 `integer`。
- `person_type` 是一个扩展定义的类型。
- `status_type` 是一个扩展定义的类型。
- `positive_integer` 是一个扩展定义的域。

### 要求与注意事项

- 审查后的控制文件声明默认版本为 `1.0.2`。
- 控制文件将扩展标记为可重定位。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为以及失败情况，以与固定源进行比对。
