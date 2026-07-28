## 用法

来源：

- [官方上游 README](https://github.com/eto-ai/rikai-pg/blob/9a245266bb6190fda7e5cc9edc82c80a267f6a84/README.md)
- [官方扩展控制文件 (rikai.control)](https://github.com/eto-ai/rikai-pg/blob/9a245266bb6190fda7e5cc9edc82c80a267f6a84/rikai.control)
- [官方扩展 SQL (rikai--0.1.sql)](https://github.com/eto-ai/rikai-pg/blob/9a245266bb6190fda7e5cc9edc82c80a267f6a84/sql/rikai--0.1.sql)

`rikai` — rikai ML 扩展。请用于相应的向量、模型或检索工作流。上游明确表示该项目尚未准备好生产使用。

### 核心工作流

```sql
CREATE EXTENSION rikai;

Create a model via `INSERT INTO`
```

在目标数据库中安装扩展，当可用时运行上游提供的最小示例，并在将其集成到应用程序 SQL 之前验证安装的版本和返回值。

### 重要对象

- `iou(box1 box, box2 box)` 是一个扩展函数，返回 `real`。
- `ml.create_model_trigger()` 是一个扩展函数，返回 `TRIGGER`。
- `ml.cuda_info()` 是一个扩展函数，返回 `JSON`。
- `ml.delete_model_trigger()` 是一个扩展函数，返回 `TRIGGER`。
- `ml.is_cuda_available()` 是一个扩展函数，返回 `BOOL`。
- `ml.version()` 是一个扩展函数，返回 `table`。
- `detection` 是一个扩展定义的类型。
- `image` 是一个扩展定义的类型。
- `mask` 是一个扩展定义的类型。
- `mask_type` 是一个扩展定义的类型。
- `ml.models` 是一个由扩展安装或管理的表。
- `ml` 是一个由扩展创建的模式。

### 要求与注意事项

- 控制文件声明默认版本为 `0.1`。
- 控制文件标记该扩展为可重定位。
- 上游明确表示该项目尚未准备好生产使用。
- 上游将项目的一部分或全部标记为实验性。
- 上游将项目描述为概念验证。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况。
