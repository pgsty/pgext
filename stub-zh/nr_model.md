## 用法

来源：

- [官方上游 README](https://github.com/neurdb/neurdb/blob/d9aefea3681a8c754ca5836354a314f0dfea23a9/aiengine/pgext/nr_modelmanager/readme.md)
- [官方扩展控制文件 (nr_model.control)](https://github.com/neurdb/neurdb/blob/d9aefea3681a8c754ca5836354a314f0dfea23a9/aiengine/pgext/nr_modelmanager/nr_model.control)
- [官方扩展 SQL (nr_model--1.0.0.sql)](https://github.com/neurdb/neurdb/blob/d9aefea3681a8c754ca5836354a314f0dfea23a9/aiengine/pgext/nr_modelmanager/sql/nr_model--1.0.0.sql)

`nr_model` — pg-model 是一个简单的 Postgres 扩展，允许模型管理和数据库内推理。使用它来进行相应的向量、模型或检索工作流。请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION nr_model;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `pgm_get_model_id_by_name(model_name text)` 是一个扩展函数，返回 `INT`。
- `pgm_predict_float4(model_name text, input anyarray)` 是一个扩展函数，返回 `SETOF`。
- `pgm_predict_table(model_name text, batch_size int, table_name text, column_names text[])` 是一个扩展函数，返回 `SETOF`。
- `pgm_register_model(model_name text, model_path text)` 是一个扩展函数，返回 `BOOL`。
- `pgm_store_model(model_name text, model_path text)` 是一个扩展函数，返回 `BOOL`。
- `pgm_unregister_model(model_name text)` 是一个扩展函数，返回 `BOOL`。
- `layer` 是一个由扩展安装或管理的表。
- `model` 是一个由扩展安装或管理的表。
- `router` 是一个由扩展安装或管理的表。

### 要求与注意事项

- 审查后的控制文件声明默认版本为 `1.0.0`。
- 控制文件将扩展标记为可重定位。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源进行比对。
