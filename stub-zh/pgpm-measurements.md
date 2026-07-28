## 用法

来源：

- [官方上游 README](https://github.com/constructive-io/pgpm-modules/blob/45d62108f508af564b72bd6512e86a5ec7cdc923/packages/measurements/README.md)
- [官方扩展控制文件 (pgpm-measurements.control)](https://github.com/constructive-io/pgpm-modules/blob/45d62108f508af564b72bd6512e86a5ec7cdc923/packages/measurements/pgpm-measurements.control)
- [官方扩展 SQL (pgpm-measurements--0.15.5.sql)](https://github.com/constructive-io/pgpm-modules/blob/45d62108f508af564b72bd6512e86a5ec7cdc923/packages/measurements/sql/pgpm-measurements--0.15.5.sql)

`pgpm-measurements` — @pgpm/measurements 提供了一套标准化的系统，用于在 PostgreSQL 应用程序中跟踪测量和数量。此扩展包定义了一个存储测量类型及其单位和描述的模式，从而在应用程序中实现一致的度量跟踪。当应用程序需要此特定数据库功能时，请使用它。上游将此功能描述为实验性。

### 核心工作流

```sql
CREATE EXTENSION "pgpm-measurements";
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行该示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `measurements.quantities` 是由扩展安装或管理的表。
- `measurements` 是由扩展创建的模式。

### 要求与注意事项

- 审查后的控制文件声明默认版本为 `0.15.5`。
- 请首先安装确认的扩展依赖项：`plpgsql`，`pgpm-verify`。
- 控制文件将扩展标记为不可重定位。
- 控制文件不要求超级用户安装。
- 上游将项目的一部分或全部标记为实验性。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源的一致性。
