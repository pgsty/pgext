## 用法

来源：

- [官方上游 README](https://github.com/bigsmoke/pg_safer_settings/blob/6ef74b74e6ad2799cae9d29c87990d353786e821/README.md)
- [官方扩展控制文件 (pg_safer_settings_table_dependent_extension.control)](https://github.com/bigsmoke/pg_safer_settings/blob/6ef74b74e6ad2799cae9d29c87990d353786e821/pg_safer_settings_table_dependent_extension/pg_safer_settings_table_dependent_extension.control)
- [官方扩展 SQL (pg_safer_settings_table_dependent_extension--forever.sql)](https://github.com/bigsmoke/pg_safer_settings/blob/6ef74b74e6ad2799cae9d29c87990d353786e821/pg_safer_settings_table_dependent_extension/pg_safer_settings_table_dependent_extension--forever.sql)

`pg_safer_settings_table_dependent_extension` — 一组函数和机制，旨在使在 PostgreSQL 中处理配置设置更加安全。当需要管理或自动化上述数据库行为时，请使用此扩展。在集成到应用程序 SQL 之前，必须先安装并验证其扩展依赖项。

### 核心工作流

```sql
CREATE EXTENSION pg_safer_settings_table_dependent_extension;

-- To change for the duration of the session:
SET app.settings.bla = 'blegherrerbypass';  -- or:
SELECT set_config('app.settings.bla', 'blegherrerbypass', false);

-- To change for the duration of the transaction:
SET LOCAL app.settings.bla = 'blegherrerbypass';  -- or:
SELECT set_config('app.settings.bla', 'blegherrerbypass', true);
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行该示例，并在将其集成到应用程序 SQL 之前验证已安装的版本和返回值。

### 要求与注意事项

- 控制文件声明默认版本为 `forever`。
- 请先安装并验证确认的扩展依赖项：`pg_safer_settings`。
- 控制文件将该扩展标记为可重定位。
- 控制文件不要求超级用户安装。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源中的信息一致。
