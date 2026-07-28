## 用法

来源：

- [官方上游 README](https://github.com/aquametalabs/aquameta/blob/ec5dbe6b84ea61b5e8b4b10263a0911b848b0025/README.md)
- [官方扩展控制文件 (semantics.control)](https://github.com/aquametalabs/aquameta/blob/ec5dbe6b84ea61b5e8b4b10263a0911b848b0025/extensions/semantics/semantics.control)

`semantics` — Aquameta Semantics 扩展。当应用程序需要此特定数据库功能时使用它。上游将此功能描述为实验性。

### 核心工作流

```sql
CREATE EXTENSION semantics;
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行该示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 要求与注意事项

- 审查过的控制文件声明默认版本为 `0.5.0`。
- 首先安装确认的扩展依赖项：`meta`, `widget`。
- 控制文件标记该扩展为不可重定位。
- 上游将项目的一部分或全部标记为实验性。
- 在生产使用前，确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源保持一致。
