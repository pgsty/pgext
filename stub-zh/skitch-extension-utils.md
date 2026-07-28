## 用法

来源：

- [官方上游 README](https://github.com/airpage-app/pg-utils/blob/2d56c14862dcf60d83cb79f1ebe0a80273d9e58d/readme.md)
- [官方扩展控制文件 (skitch-extension-utils.control)](https://github.com/airpage-app/pg-utils/blob/2d56c14862dcf60d83cb79f1ebe0a80273d9e58d/packages/utils/skitch-extension-utils.control)

`skitch-extension-utils` — PostgreSQL 工具。使用它来进行相应的 SQL 或数据库工具工作流。在将其集成到应用程序 SQL 中之前，必须先安装并验证其扩展依赖项。

### 核心工作流

```sql
CREATE EXTENSION "skitch-extension-utils";
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行该示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 要求与注意事项

- 审查过的控制文件声明默认版本为 `0.0.7`。
- 请先安装确认的扩展依赖项：`plpgsql`。
- 控制文件将该扩展标记为不可重定位。
- 控制文件不要求超级用户安装。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源保持一致。
