## 用法

来源：

- [官方上游 README](https://github.com/19pine-ai/transparent-offload/blob/b81a4282dde314bd7fe86c9529dddcc6130116d8/README.md)
- [官方扩展控制文件 (accel_ext.control)](https://github.com/19pine-ai/transparent-offload/blob/b81a4282dde314bd7fe86c9529dddcc6130116d8/transparent-runtime/apps/pg_accel/accel_ext.control)
- [官方扩展 SQL (accel_ext--1.0.sql)](https://github.com/19pine-ai/transparent-offload/blob/b81a4282dde314bd7fe86c9529dddcc6130116d8/transparent-runtime/apps/pg_accel/accel_ext--1.0.sql)

`accel_ext` — 细粒度加速器卸载运行时，可以在 GPU、HSM 或推理工作与其他 PostgreSQL 请求之间重叠执行。当应用程序需要此特定数据库功能时，请使用它。上游将此功能描述为实验性。

### 核心工作流

```sql
CREATE EXTENSION accel_ext;
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行该示例，并在将其集成到应用程序 SQL 之前验证已安装的版本和返回值。

### 重要对象

- `accel_async(integer)` 是一个扩展函数，返回 `integer`。
- `accel_sync(integer)` 是一个扩展函数，返回 `integer`。

### 要求与注意事项

- 审查后的控制文件声明默认版本为 `1.0`。
- 控制文件将扩展标记为可重定位。
- 上游将项目的一部分或全部标记为实验性。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况，以匹配固定源。
