## 用法

来源：

- [官方上游 README](https://github.com/airpage-app/pg-utils/blob/2d56c14862dcf60d83cb79f1ebe0a80273d9e58d/readme.md)
- [官方扩展控制文件 (skitch-extension-defaults.control)](https://github.com/airpage-app/pg-utils/blob/2d56c14862dcf60d83cb79f1ebe0a80273d9e58d/packages/defaults/skitch-extension-defaults.control)
- [官方扩展 SQL (skitch-extension-defaults--0.0.7.sql)](https://github.com/airpage-app/pg-utils/blob/2d56c14862dcf60d83cb79f1ebe0a80273d9e58d/packages/defaults/sql/skitch-extension-defaults--0.0.7.sql)

`skitch-extension-defaults` — 默认角色。在实现相应的安全、审计或访问控制工作流时使用。必须首先安装并验证其扩展依赖项。

### 核心工作流

```sql
CREATE EXTENSION "skitch-extension-defaults";
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行该示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 要求与注意事项

- 审查过的控制文件声明默认版本 `0.0.7`。
- 首先安装并验证确认的扩展依赖项：`plpgsql`。
- 控制文件将扩展标记为不可重定位。
- 控制文件不要求超级用户安装。
- 在生产使用前，确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源保持一致。
