## 用法

来源：

- [官方上游 README](https://github.com/postgres-artificialintelligence/pgai/blob/5c5f01d833ab4ee80779e0958c7ca0f6022e5c94/README.md)
- [官方扩展控制文件 (pgai.control)](https://github.com/postgres-artificialintelligence/pgai/blob/5c5f01d833ab4ee80779e0958c7ca0f6022e5c94/pgai.control)
- [官方扩展 SQL (pgai--1.0.sql)](https://github.com/postgres-artificialintelligence/pgai/blob/5c5f01d833ab4ee80779e0958c7ca0f6022e5c94/pgai--1.0.sql)

`pgai` — 本指南提供了在各种操作系统上安装 PostgreSQL 的逐步说明。请根据相应的向量、模型或检索工作流使用它。在目标 PostgreSQL 构建上测试链接的上游修订版本作为 API 边界。

### 核心工作流

```sql
CREATE EXTENSION pgai;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证安装的版本和返回值。

### 重要对象

- `fetch_apple_stock(input_date DATE)` 是一个扩展函数，返回 `TABLE`。
- `fetch_google_stock(input_date DATE)` 是一个扩展函数，返回 `TABLE`。
- `fetch_tesla_stock(input_date DATE)` 是一个扩展函数，返回 `TABLE`。
- `get_arima_prediction(target_date TEXT)` 是一个扩展函数，返回 `FLOAT`。
- `get_db_config()` 是一个扩展函数，返回 `TABLE`。
- `get_prediction_for_date(prediction_date DATE)` 是一个扩展函数，返回 `TABLE`。
- `predict_stock_close_value_apple(input_date_str TEXT)` 是一个扩展函数，返回 `FLOAT`。
- `predict_stock_close_value_tesla(input_date_str TEXT)` 是一个扩展函数，返回 `FLOAT`。
- `apple_stocks` 是一个扩展定义的物化视图。
- `google_stocks` 是一个扩展定义的物化视图。
- `tesla_stocks` 是一个扩展定义的物化视图。

### 要求与注意事项

- 控制文件声明默认版本为 `1.0`。
- 控制文件将扩展标记为可重定位。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况。
