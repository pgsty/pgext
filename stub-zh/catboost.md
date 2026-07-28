## 用法

来源：

- [官方上游 README](https://github.com/akalend/pg_catboost/blob/c5f10185700ed82b6a35a5c327f644c8bf7259fd/README.md)
- [官方扩展控制文件 (catboost.control)](https://github.com/akalend/pg_catboost/blob/c5f10185700ed82b6a35a5c327f644c8bf7259fd/catboost.control)
- [官方扩展 SQL (catboost--0.1.sql)](https://github.com/akalend/pg_catboost/blob/c5f10185700ed82b6a35a5c327f644c8bf7259fd/catboost--0.1.sql)

`catboost` — 一个使用 CatBoost 的机器学习模块。请使用上述链接的上游修订版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION catboost;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `ml_learn(name text, model_type int, options text, table_name text, filename text)` 是一个扩展函数，返回 `float`。
- `ml_meta(OUT name text, OUT loss_function text, OUT model_type char(1), OUT acc real, OUT args text, OUT classes text)` 是一个扩展函数，返回 `setof`。
- `ml_predict(model text, tablename text, join_field text DEFAULT 'row', OUT index text, OUT predict float, OUT class text)` 是一个扩展函数，返回 `setof`。
- `ml_predict_internal(model text, tablename text, join_field text DEFAULT 'row', isQuery bool DEFAULT FALSE, OUT index text, OUT predict float, OUT class text)` 是一个扩展函数，返回 `setof`。
- `ml_predict_query(model text, query text, join_field text DEFAULT 'row', OUT index text, OUT predict float, OUT class text)` 是一个扩展函数，返回 `setof`。
- `ml_test(name Name)` 是一个扩展函数，返回 `text`。
- `ml_model` 是一个由扩展安装或管理的表。

### 要求与注意事项

- 控制文件声明默认版本为 `0.1`。
- 控制文件标记该扩展为可重定位。
- 控制文件标记该扩展为可信。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源中的信息一致。
