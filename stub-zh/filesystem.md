## 用法

来源：

- [官方上游 README](https://github.com/aquametalabs/aquameta/blob/ec5dbe6b84ea61b5e8b4b10263a0911b848b0025/extensions/filesystem/README.md)
- [官方扩展控制文件 (filesystem.control)](https://github.com/aquametalabs/aquameta/blob/ec5dbe6b84ea61b5e8b4b10263a0911b848b0025/extensions/filesystem/filesystem.control)

`filesystem` — 文件系统外部数据封装器 =============================== 将文件系统暴露给 PostgreSQL，允许通过 SQL 命令读取文件和目录。当 PostgreSQL 需要通过外部数据接口访问相应的外部数据源时，请使用此扩展。在安装并验证其扩展依赖项之前，请勿使用此扩展。

### 核心工作流

```sql
CREATE EXTENSION filesystem;
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行该示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 要求与注意事项

- 审核后的控制文件声明默认版本为 `0.4.0`。
- 请首先安装并验证确认的扩展依赖项：`meta`。
- 控制文件标记该扩展为不可重定位。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源保持一致。
