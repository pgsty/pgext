## 用法

来源：

- [官方上游 README](https://github.com/eatonphil/pgleak/blob/543555544763d5fe746bdf7869c6d9a90581527d/README.md)
- [官方扩展控制文件 (pgleak.control)](https://github.com/eatonphil/pgleak/blob/543555544763d5fe746bdf7869c6d9a90581527d/pgleak.control)
- [官方扩展 SQL (pgleak--0.0.1.sql)](https://github.com/eatonphil/pgleak/blob/543555544763d5fe746bdf7869c6d9a90581527d/pgleak--0.0.1.sql)

`pgleak` — 创建并启动一个新的 Postgres 数据库：当应用程序需要此特定数据库功能时使用。使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION pgleak;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `pgleak.leak_sum(INT, INT)` 是一个扩展函数，返回 `INT`。
- `pgleak.leak_sum_malloc(INT, INT)` 是一个扩展函数，返回 `INT`。
- `pgleak` 是由扩展创建的一个模式。

### 要求与注意事项

- 控制文件声明默认版本为 `0.0.1`。
- 控制文件将扩展标记为可重定位。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为以及失败情况。
