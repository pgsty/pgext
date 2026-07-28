## 用法

来源：

- [官方上游 README](https://github.com/apache/cloudberry/blob/83022b58b2f4e62fb8fd2ebf9137b5c29e4fcb5b/gpcontrib/gp_debug_numsegments/README.md)
- [官方扩展控制文件 (gp_debug_numsegments.control)](https://github.com/apache/cloudberry/blob/83022b58b2f4e62fb8fd2ebf9137b5c29e4fcb5b/gpcontrib/gp_debug_numsegments/gp_debug_numsegments.control)
- [官方扩展 SQL (gp_debug_numsegments--1.0.sql)](https://github.com/apache/cloudberry/blob/83022b58b2f4e62fb8fd2ebf9137b5c29e4fcb5b/gpcontrib/gp_debug_numsegments/gp_debug_numsegments--1.0.sql)

`gp_debug_numsegments` — 默认情况下，所有表都会创建在所有段上。使用此扩展可以检查或更改默认行为。在管理或自动化上述数据库行为时，请使用上述链接的固定上游版本作为 API 边界，并在目标 PostgreSQL 构建中进行测试。

### 核心工作流

```sql
CREATE EXTENSION gp_debug_numsegments;
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行该示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `gp_debug_get_create_table_default_numsegments()` 是一个扩展函数，返回 `text`。
- `gp_debug_reset_create_table_default_numsegments()` 是一个扩展函数，返回 `void`。
- `gp_debug_reset_create_table_default_numsegments(integer)` 是一个扩展函数，返回 `void`。
- `gp_debug_reset_create_table_default_numsegments(text)` 是一个扩展函数，返回 `void`。
- `gp_debug_set_create_table_default_numsegments(integer)` 是一个扩展函数，返回 `text`。
- `gp_debug_set_create_table_default_numsegments(text)` 是一个扩展函数，返回 `text`。

### 要求与注意事项

- 审查过的控制文件声明默认版本为 `1.0`。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为以及失败情况。
