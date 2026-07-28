## 用法

来源：

- [官方上游 README](https://github.com/scokmen/pg_mask/blob/e614f6ebe6ac98f89c977694cd4624c978c1caf1/README.md)
- [官方扩展控制文件 (pg_mask.control)](https://github.com/scokmen/pg_mask/blob/e614f6ebe6ac98f89c977694cd4624c978c1caf1/pg_mask.control)
- [官方扩展 SQL (pg_mask--1.0.0.sql)](https://github.com/scokmen/pg_mask/blob/e614f6ebe6ac98f89c977694cd4624c978c1caf1/pg_mask--1.0.0.sql)

`pg_mask` — 该项目使用 C 语言开发，并利用 PostgreSQL 的 C 语言函数来提供一个扩展 makefile，该 makefile 管理库目录和安装目标。为了包含 PostgreSQL 扩展 makefile，二进制文件必须可用。在实现相应的安全、审计或访问控制工作流时，请使用上述链接的固定上游版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION pg_mask;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证安装的版本和返回值。

### 重要对象

- `pg_mask()` 是一个扩展函数，返回 `text`。

### 要求与注意事项

- 控制文件声明默认版本为 `1.0.0`。
- 控制文件将扩展标记为不可重定位。
- 控制文件不要求超级用户安装。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源代码中的信息一致。
