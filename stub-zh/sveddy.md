## 用法

来源：

- [官方上游 README](https://github.com/robbiegm/sveddy/blob/aee5495b2e4cfc0207a415a1fcd479c63808ed8e/README.md)
- [官方扩展控制文件 (sveddy.control)](https://github.com/robbiegm/sveddy/blob/aee5495b2e4cfc0207a415a1fcd479c63808ed8e/sveddy.control)
- [官方扩展 SQL (sveddy--0.1.0.sql)](https://github.com/robbiegm/sveddy/blob/aee5495b2e4cfc0207a415a1fcd479c63808ed8e/sveddy--0.1.0.sql)

`sveddy` — Sveddy 是一个在数据库中进行协同过滤的 PostgreSQL 扩展。使用它来进行相应的向量、模型或检索工作流。请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION sveddy;

SELECT predict_uv(
    (SELECT weights FROM ratings_sveddy_model_u WHERE id = 3),
    (SELECT weights FROM ratings_sveddy_model_v WHERE id = 5)
);
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证安装的版本和返回值。

### 重要对象

- `get_initial_weights_uv(integer)` 是一个扩展函数，返回 `real[]`。
- `predict_uv(real[], real[])` 是一个扩展函数，返回 `real`。
- `update_model_uv()` 是一个扩展函数，返回 `TRIGGER`。
- `garbage_collect_uv` 是一个扩展过程。
- `initialize_model_uv` 是一个扩展过程。
- `train_uv` 是一个扩展过程。
- `sveddy_models_uv` 是一个由扩展安装或管理的表。

### 要求与注意事项

- 审查过的控制文件声明默认版本为 `0.1.0`。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为以及失败情况。
