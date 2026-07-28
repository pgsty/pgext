## 用法

来源：

- [官方上游 README](https://api.pgxn.org/src/json_enhancements/json_enhancements-1.0.2/README.md)
- [官方扩展控制文件 (json_enhancements_with_hstore.control)](https://api.pgxn.org/src/json_enhancements/json_enhancements-1.0.2/json_enhancements_with_hstore.control)

`json_enhancements_with_hstore` — Json 扩展增强功能 for PostgreSQL 9.2 ====================================。当应用程序需要此特定数据库功能时使用它。所审核的上游项目已归档或不再维护。

### 核心工作流

```sql
CREATE EXTENSION json_enhancements_with_hstore;
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行该示例，并在将其集成到应用程序 SQL 中之前验证安装版本和返回值。

### 要求与注意事项

- 所审核的控制文件声明默认版本 `1.0.0`。
- 首先安装确认的扩展依赖项：`hstore`。
- 控制文件标记该扩展为可重定位。
- 在生产使用前，确认权限、支持的 PostgreSQL 版本、升级行为和失败情况。
