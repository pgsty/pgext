## 用法

来源：

- [官方上游 README](https://api.pgxn.org/src/omnidb_plugin/omnidb_plugin-0.0.1/README.md)
- [官方扩展控制文件 (omnidb_plugin.control)](https://api.pgxn.org/src/omnidb_plugin/omnidb_plugin-0.0.1/omnidb_plugin.control)
- [官方扩展 SQL (omnidb_plugin--0.0.1.sql)](https://api.pgxn.org/src/omnidb_plugin/omnidb_plugin-0.0.1/omnidb_plugin--0.0.1.sql)

`omnidb_plugin` — nano /etc/postgresql/X.Y/main/postgresql.conf shared_preload_libraries = '/opt/omnidb-plugin/omnidb_plugin_XY'. 使用它当数据库代码必须在或与该过程语言进行交互运行时。使用上方链接的上游固定版本作为API边界，并在目标PostgreSQL构建中进行测试。

### 核心工作流

```sql
CREATE EXTENSION omnidb_plugin;
```

在目标数据库中安装扩展，当可用时运行上方的最小上游示例，并在将其集成到应用程序SQL之前验证安装的版本和返回值。

### 重要对象

- `omnidb.omnidb_enable_debugger(character varying)` 是一个扩展函数并返回 `void`。
- `omnidb.contexts` 是由扩展安装或管理的表。
- `omnidb.statistics` 是由扩展安装或管理的表。
- `omnidb.variables` 是由扩展安装或管理的表。
- `omnidb` 是由扩展创建的模式。

### 要求与注意事项

- 审查后的控制文件声明默认版本 `0.0.1`。
- 控制文件将扩展标记为不可重定位。
- 在生产使用前，确认权限、支持的PostgreSQL版本、升级行为和失败情况与固定源进行对比。
