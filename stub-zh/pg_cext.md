## 用法

来源：

- [官方上游 README](https://api.pgxn.org/src/pg_cext/pg_cext-1.0.0/README.md)
- [官方扩展控制文件 (pg_cext.control)](https://api.pgxn.org/src/pg_cext/pg_cext-1.0.0/pg_cext.control)
- [官方扩展 SQL (pg_cext--1.0.0.sql)](https://api.pgxn.org/src/pg_cext/pg_cext-1.0.0/pg_cext--1.0.0.sql)

`pg_cext` — 这是一个用 C 语言实现的 PostgreSQL 扩展，用于添加两个数字。它展示了如何使用 C 语言为 PostgreSQL 创建一个简单的扩展。当 SQL 需要这些特殊函数或聚合时，请使用它。使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建中进行测试。

### 核心工作流

```sql
CREATE EXTENSION pg_cext;
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行该示例，并在将其集成到应用程序 SQL 中之前验证安装版本和返回值。

### 重要对象

- `add_nums(int,int)` 是一个扩展函数，返回 `int`。

### 要求与注意事项

- 审查后的控制文件声明默认版本为 `1.0.0`。
- 控制文件将扩展标记为可重定位。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况，以匹配固定源代码。
