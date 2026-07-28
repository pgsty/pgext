## 用法

来源：

- [官方上游 README](https://github.com/aquametalabs/aquameta/blob/ec5dbe6b84ea61b5e8b4b10263a0911b848b0025/extensions/bundle/README.md)
- [官方扩展控制文件 (bundle.control)](https://github.com/aquametalabs/aquameta/blob/ec5dbe6b84ea61b5e8b4b10263a0911b848b0025/extensions/bundle/bundle.control)

`bundle` — 此扩展为 PostgreSQL 提供行级快照数据版本控制功能，类似于 git。当应用程序需要此特定数据库功能时，请使用此扩展。在安装此扩展之前，必须先安装并验证其扩展依赖项。

### 核心工作流

```sql
CREATE EXTENSION bundle;
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行该示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 要求与注意事项

- 控制文件声明默认版本为 `0.4.0`。
- 请先安装确认的扩展依赖项：`meta`, `"uuid-ossp"`, `pgcrypto`。
- 控制文件将此扩展标记为不可重定位。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源保持一致。
