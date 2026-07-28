## 用法

来源：

- [官方扩展控制文件（object_reference.control）](https://api.pgxn.org/src/object_reference/object_reference-0.1.0/object_reference.control)
- [官方扩展 SQL（object_reference.sql）](https://api.pgxn.org/src/object_reference/object_reference-0.1.0/sql/object_reference.sql)

`object_reference` — 提供 Postgres 对象的不可变引用。当应用程序需要此特定数据库功能时，请使用它。在集成到应用程序 SQL 之前，必须先安装并验证其扩展依赖项。

### 核心工作流

```sql
CREATE EXTENSION object_reference;
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行该示例，并在将其集成到应用程序 SQL 之前验证已安装的版本和返回值。

### 重要对象

- `__object_reference.create_function(function_name text , args text , options text , body text , comment text , grants text DEFAULT NULL)` 是一个扩展函数，返回 `void`。
- `__object_reference.exec(sql text)` 是一个扩展函数，返回 `void`。
- `__object_reference.safe_dump(relation regclass , filter text DEFAULT '')` 是一个扩展函数，返回 `void`。
- `snitch()` 是一个扩展函数，返回 `event_trigger`。
- `_object_reference.object` 是由扩展安装或管理的表。
- `_object_reference.object_group` 是由扩展安装或管理的表。
- `_object_reference.object_group__object` 是由扩展安装或管理的表。

### 要求与注意事项

- 审核的控制文件声明默认版本为 `0.1.0`。
- 请先安装并验证确认的扩展依赖项：`cat_tools`，`count_nulls`。
- 控制文件将该扩展标记为不可重定位。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源保持一致。
