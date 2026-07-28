## 用法

来源：

- [官方扩展控制文件（extension_drop.control）](https://api.pgxn.org/src/extension_drop/extension_drop-0.1.1/extension_drop.control)
- [官方扩展 SQL（extension_drop.sql）](https://api.pgxn.org/src/extension_drop/extension_drop-0.1.1/sql/extension_drop.sql)

`extension_drop` — 在卸载扩展时运行自定义命令。在管理或自动化上述数据库行为时使用它。在卸载扩展之前，必须先安装并验证其依赖项。

### 核心工作流

```sql
CREATE EXTENSION extension_drop;
```

在目标数据库中安装扩展，在可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `__extension_drop.create_function(function_name text , args text , options text , body text , comment text , grants text DEFAULT NULL)` 是一个扩展函数，返回 `void`。
- `__extension_drop.exec(sql text)` 是一个扩展函数，返回 `void`。
- `__extension_drop.safe_dump(relation regclass , filter text DEFAULT '')` 是一个扩展函数，返回 `void`。
- `__extension_drop.messages` 是由扩展安装或管理的表。
- `extension_drop__commands` 是由扩展安装或管理的表。

### 要求与注意事项

- 审核的控制文件声明默认版本为 `0.1.1`。
- 先安装并验证确认的扩展依赖项：`cat_tools`。
- 控制文件将扩展标记为不可重定位。
- 在生产使用前，确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源的一致性。
