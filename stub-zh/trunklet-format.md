## 用法

来源：

- [官方上游 README](https://api.pgxn.org/src/trunklet-format/trunklet-format-0.2.0/README.md)
- [官方扩展控制文件 (trunklet-format.control)](https://api.pgxn.org/src/trunklet-format/trunklet-format-0.2.0/trunklet-format.control)
- [官方扩展 SQL (trunklet-format.sql)](https://api.pgxn.org/src/trunklet-format/trunklet-format-0.2.0/sql/trunklet-format.sql)

`trunklet-format` — 确保你已经安装了 pg_config 并将其添加到路径中。如果你使用的是 RPM 包管理器安装 PostgreSQL，请确保也安装了 -devel 包。如果需要，可以告诉构建过程 pg_config 的位置：使用它来对应 SQL 或数据库实用工具工作流。扩展的依赖项必须首先被安装并验证。

### 核心工作流

```sql
CREATE EXTENSION "trunklet-format";
```

在目标数据库中安装扩展，在可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 之前验证已安装的版本和返回值。

### 要求与注意事项

- 审查过的控制文件声明默认版本为 `0.2.0`。
- 在安装确认的扩展依赖项之前，请先安装它们：`trunklet`。
- 控制文件将扩展标记为可重定位。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源进行比对。
