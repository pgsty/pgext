## 用法

来源：

- [官方上游 README](https://github.com/constructive-io/utils/blob/949bad998f1ac6eaf873463ce2f788765068bc1a/packages/stamps/README.md)
- [官方扩展控制文件 (launchql-stamps.control)](https://github.com/constructive-io/utils/blob/949bad998f1ac6eaf873463ce2f788765068bc1a/packages/stamps/launchql-stamps.control)
- [官方扩展 SQL (launchql-stamps--0.4.5.sql)](https://github.com/constructive-io/utils/blob/949bad998f1ac6eaf873463ce2f788765068bc1a/packages/stamps/sql/launchql-stamps--0.4.5.sql)

`launchql-stamps` — PostgreSQL 扩展，提供触发函数以自动为数据库表添加时间戳和用户跟踪。此扩展简化了审计跟踪的实现，通过自动记录创建和更新时间戳以及执行这些操作的用户。当 SQL 需要这些特殊功能或聚合时，请使用此扩展。在安装此扩展之前，必须先安装并验证其扩展依赖项。

### 核心工作流

```sql
CREATE EXTENSION "launchql-stamps";
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行该示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `stamps.peoplestamps()` 是一个扩展函数，返回 `trigger`。
- `stamps.timestamps()` 是一个扩展函数，返回 `trigger`。
- `stamps` 是由扩展创建的模式。

### 要求与注意事项

- 控制文件声明默认版本为 `0.4.5`。
- 请先安装确认的扩展依赖项：`plpgsql`，`launchql-jwt-claims`。
- 控制文件将此扩展标记为不可重定位。
- 控制文件不要求超级用户安装。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源中的信息一致。
