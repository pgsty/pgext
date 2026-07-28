## 用法

来源：

- [官方扩展控制文件（tg_sanity.control）](https://api.pgxn.org/src/tg_sanity/tg_sanity-0.1.0/tg_sanity.control)
- [官方扩展 SQL（tg_sanity.sql）](https://api.pgxn.org/src/tg_sanity/tg_sanity-0.1.0/sql/tg_sanity.sql)

`tg_sanity` — 触发函数用于确保数据质量。当应用程序需要此特定数据库功能时，请使用它。在目标 PostgreSQL 构建中测试上游链接的固定版本作为 API 边界。

### 核心工作流

```sql
CREATE EXTENSION tg_sanity;
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行它，并在将其集成到应用程序 SQL 中之前验证安装版本和返回值。

### 重要对象

- `tg_sanity()` 是一个扩展函数，返回 `trigger`。

### 要求与注意事项

- 审查后的控制文件声明默认版本为 `0.1.0`。
- 控制文件将扩展标记为可重定位。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况，以匹配固定源。
