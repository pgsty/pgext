## 用法

来源：

- [官方上游 README](https://github.com/aquametalabs/aquameta/blob/ec5dbe6b84ea61b5e8b4b10263a0911b848b0025/extensions/ide/README.md)
- [官方扩展控制文件 (ide.control)](https://github.com/aquametalabs/aquameta/blob/ec5dbe6b84ea61b5e8b4b10263a0911b848b0025/extensions/ide/ide.control)

`ide` — Aquameta IDE。当应用程序需要此特定数据库功能时使用它。在安装扩展及其依赖项并验证之前，请勿将其集成到应用程序 SQL 中。

### 核心工作流

```sql
CREATE EXTENSION ide;
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行该示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 要求与注意事项

- 控制文件声明默认版本为 `0.5.0`。
- 请首先安装并验证扩展依赖项：`endpoint`、`widget`、`meta`。
- 控制文件标记该扩展为不可重定位。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源保持一致。
