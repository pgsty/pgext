## 用法

来源：

- [官方扩展控制文件（capitalize.control）](https://github.com/darshan744/postgres-extension/blob/2a22caeecee3505a327ff37b3752f51bbb1c017a/CapitalizeExtension/capitalize.control)
- [官方扩展 SQL（capitalize--1.0.sql）](https://github.com/darshan744/postgres-extension/blob/2a22caeecee3505a327ff37b3752f51bbb1c017a/CapitalizeExtension/capitalize--1.0.sql)

`capitalize` — 一个用于将所有给定字符大写的函数。在 SQL 需要这些特殊函数或聚合时使用它。使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建中进行测试。

### 核心工作流

```sql
CREATE EXTENSION capitalize;
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行它，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `capitalize(text)` 是一个扩展函数，返回 `TEXT`。

### 要求与注意事项

- 审查后的控制文件声明默认版本为 `1.0`。
- 控制文件将扩展标记为可重定位。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况，以匹配固定源代码。
