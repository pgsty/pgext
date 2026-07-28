## 用法

来源：

- [官方上游 README](https://github.com/constructive-io/lib-count/blob/c04e5084d5fbdee6c452976d9e69532c85dcbe96/packages/npm/README.md)
- [官方扩展控制文件 (npm.control)](https://github.com/constructive-io/lib-count/blob/c04e5084d5fbdee6c452976d9e69532c85dcbe96/packages/npm/npm.control)
- [官方扩展 SQL (npm--0.0.1.sql)](https://github.com/constructive-io/lib-count/blob/c04e5084d5fbdee6c452976d9e69532c85dcbe96/packages/npm/sql/npm--0.0.1.sql)

`npm` — 用于跟踪 npm 包元数据和每日下载次数的模式和表。当应用程序需要此特定数据库功能时使用它。在安装扩展及其依赖项并验证它们之前，请勿将其集成到应用程序 SQL 中。

### 核心工作流

```sql
CREATE EXTENSION npm;
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行该示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `npm_count.set_id_from_pkg_date()` 是一个扩展函数，返回 `trigger`。
- `npm_count.update_updated_at()` 是一个扩展函数，返回 `trigger`。
- `npm_count.validate_download_date()` 是一个扩展函数，返回 `trigger`。
- `npm_count.missing_download_dates` 是一个由扩展定义的视图。
- `npm_count.category` 是一个由扩展安装或管理的表。
- `npm_count.daily_downloads` 是一个由扩展安装或管理的表。
- `npm_count.npm_package` 是一个由扩展安装或管理的表。
- `npm_count.package_category` 是一个由扩展安装或管理的表。
- `npm_count` 是一个由扩展创建的模式。

### 要求与注意事项

- 控制文件声明默认版本为 `0.0.1`。
- 请首先安装并验证确认的扩展依赖项：`btree_gist`, `plpgsql`, `uuid-ossp`。
- 控制文件将扩展标记为不可重定位。
- 控制文件不要求超级用户安装。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源中的信息一致。
