## 用法

来源：

- [官方上游 README](https://github.com/gabbasb/pg_gnuplot/blob/1716363eb3138f63b353eb749bee82782307e181/README.md)
- [官方扩展控制文件 (pg_gnuplot.control)](https://github.com/gabbasb/pg_gnuplot/blob/1716363eb3138f63b353eb749bee82782307e181/pg_gnuplot.control)
- [官方扩展 SQL (pg_gnuplot--1.0.sql)](https://github.com/gabbasb/pg_gnuplot/blob/1716363eb3138f63b353eb749bee82782307e181/pg_gnuplot--1.0.sql)

`pg_gnuplot` — PostgreSQL 扩展，用于使用 GNUPlot 绘制图形。在相应的 SQL 或数据库实用程序工作流中使用它。使用上述链接的上游修订版本作为 API 边界，并在目标 PostgreSQL 构建中进行测试。

### 核心工作流

```sql
CREATE EXTENSION pg_gnuplot;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证安装的版本和返回值。

### 重要对象

- `gnuplot_version()` 是一个扩展函数，返回 `cstring`。
- `pg_gnuplot_version()` 是一个扩展函数，返回 `pg_catalog`。
- `pg_plot(db_query pg_catalog.text, plot_cmd pg_catalog.text)` 是一个扩展函数，返回 `pg_catalog`。

### 要求与注意事项

- 审查后的控制文件声明默认版本为 `1.0`。
- 控制文件将扩展标记为可重定位。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况，以匹配固定源代码。
