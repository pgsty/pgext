## 用法

来源：

- [官方上游 README](https://github.com/mcadariu/pg_changepoint/blob/1133b4f9c51e42d4f9866fd055b9b6279a6fa19e/README.md)
- [官方扩展控制文件 (pg_change_point_detection.control)](https://github.com/mcadariu/pg_changepoint/blob/1133b4f9c51e42d4f9866fd055b9b6279a6fa19e/pg_change_point_detection.control)
- [官方扩展 SQL (pg_change_point_detection--1.0.sql)](https://github.com/mcadariu/pg_changepoint/blob/1133b4f9c51e42d4f9866fd055b9b6279a6fa19e/pg_change_point_detection--1.0.sql)

`pg_change_point_detection` — pg_changepoint 是一个用于检测表数据中变化点的 PostgreSQL 扩展。它是 Andrey Akinshin 的 ED-PELT 算法实现的移植版本。使用它来进行相应的调度、时间序列或时间工作流。请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION pg_change_point_detection;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证安装的版本和返回值。

### 重要对象

- `pg_change_point_detection(data double precision[])` 是一个扩展函数，返回 `integer[]`。
- `pg_change_point_detection_in_column(table_name text, column_name text, order_column text DEFAULT NULL)` 是一个扩展函数，返回 `integer[]`。

### 要求与注意事项

- 审查后的控制文件声明默认版本为 `1.0`。
- 控制文件将扩展标记为可重定位。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源进行比对。
