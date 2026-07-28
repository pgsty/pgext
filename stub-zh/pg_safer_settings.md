## 用法

来源：

- [官方上游 README](https://github.com/bigsmoke/pg_safer_settings/blob/6ef74b74e6ad2799cae9d29c87990d353786e821/README.md)
- [官方扩展控制文件 (pg_safer_settings.control)](https://github.com/bigsmoke/pg_safer_settings/blob/6ef74b74e6ad2799cae9d29c87990d353786e821/pg_safer_settings.control)

`pg_safer_settings` — pg_safer_settings 提供了一些函数和机制，使在 Postgres 中处理配置项变得更加安全。当进行数据库管理或自动化上述数据库行为时，请使用此扩展。请使用链接中的锁定上游版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION pg_safer_settings;

-- To change for the duration of the session:
SET app.settings.bla = 'blegherrerbypass';  -- or:
SELECT set_config('app.settings.bla', 'blegherrerbypass', false);

-- To change for the duration of the transaction:
SET LOCAL app.settings.bla = 'blegherrerbypass';  -- or:
SELECT set_config('app.settings.bla', 'blegherrerbypass', true);
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行该示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 要求与注意事项

- 审查过的控制文件声明默认版本为 `1.0.0`。
- 控制文件将扩展标记为可重定位。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为以及失败情况。
