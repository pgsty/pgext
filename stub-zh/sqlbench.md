## 用法

来源：

- [官方上游 README](https://github.com/swida/sqlbench/blob/db8d31dc7e02517c61e555c03daac1ee4d1b82a4/README)
- [官方扩展控制文件 (sqlbench.control)](https://github.com/swida/sqlbench/blob/db8d31dc7e02517c61e555c03daac1ee4d1b82a4/src/storeproc/pgsql/c/sqlbench.control)
- [官方扩展 SQL (sqlbench--1.0.0.sql)](https://github.com/swida/sqlbench/blob/db8d31dc7e02517c61e555c03daac1ee4d1b82a4/src/storeproc/pgsql/c/sqlbench--1.0.0.sql)

`sqlbench` — 该项目源自 dbt2，可以执行标准 TPC-C 测试。当应用程序需要此特定数据库功能时，请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION sqlbench;
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行该示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `delivery(INTEGER, INTEGER)` 是一个扩展函数，返回 `INTEGER`。
- `new_order(INTEGER, INTEGER, INTEGER, INTEGER, INTEGER, INTEGER, INTEGER, INTEGER, INTEGER, INTEGER, INTEGER, INTEGER, INTEGER, INTEGER, INTEGER, INTEGER, INTEGER, INTEGER, INTEGER, INTEGER, INTEGER, INTEGER, INTEGER, INTEGER, INTEGER, INTEGER, INTEGER, INTEGER, INTEGER,…)` 是一个扩展函数，返回 `INTEGER`。
- `order_status(INTEGER, INTEGER, INTEGER, TEXT)` 是一个扩展函数，返回 `SETOF`。
- `payment(INTEGER, INTEGER, INTEGER, INTEGER, INTEGER, TEXT, REAL)` 是一个扩展函数，返回 `INTEGER`。
- `stock_level(INTEGER, INTEGER, INTEGER)` 是一个扩展函数，返回 `INTEGER`。

### 要求与注意事项

- 审查后的控制文件声明默认版本为 `1.0.0`。
- 控制文件将扩展标记为可重定位。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为以及失败情况。
