## 用法

来源：

- [官方上游 README](https://github.com/constructive-io/pgpm-modules/blob/45d62108f508af564b72bd6512e86a5ec7cdc923/packages/types/README.md)
- [官方扩展控制文件 (pgpm-types.control)](https://github.com/constructive-io/pgpm-modules/blob/45d62108f508af564b72bd6512e86a5ec7cdc923/packages/types/pgpm-types.control)
- [官方扩展 SQL (pgpm-types--0.15.5.sql)](https://github.com/constructive-io/pgpm-modules/blob/45d62108f508af564b72bd6512e86a5ec7cdc923/packages/types/sql/pgpm-types--0.15.5.sql)

`pgpm-types` — 基础 PostgreSQL 数据类型与 SQL 脚本。当应用程序需要此类型、域或其操作符时使用它。在安装扩展依赖项并验证它们之前，请勿集成到应用程序 SQL 中。

### 核心工作流

```sql
CREATE EXTENSION "pgpm-types";
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行该示例，并在将其集成到应用程序 SQL 之前验证已安装的版本和返回值。

### 重要对象

- `attachment` 是扩展定义的域。
- `email` 是扩展定义的域。
- `hostname` 是扩展定义的域。
- `image` 是扩展定义的域。
- `origin` 是扩展定义的域。
- `upload` 是扩展定义的域。
- `url` 是扩展定义的域。

### 要求与注意事项

- 控制文件声明默认版本为 `0.15.5`。
- 请首先安装并验证确认的扩展依赖项：`plpgsql`, `citext`, `pgpm-verify`。
- 控制文件将扩展标记为不可重定位。
- 控制文件不要求超级用户安装。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源的一致性。
