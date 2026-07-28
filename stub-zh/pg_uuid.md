## 用法

来源：

- [官方上游 README](https://github.com/tvondra/pg_uuid/blob/5abbd6a5ed12b65674f6ef1c6a6d07c16f4b6f68/README.md)
- [官方扩展控制文件 (pg_uuid.control)](https://github.com/tvondra/pg_uuid/blob/5abbd6a5ed12b65674f6ef1c6a6d07c16f4b6f68/pg_uuid.control)
- [官方扩展 SQL (pg_uuid--1.0.0.sql)](https://github.com/tvondra/pg_uuid/blob/5abbd6a5ed12b65674f6ef1c6a6d07c16f4b6f68/pg_uuid--1.0.0.sql)

`pg_uuid` — 用于生成新的 UUID 版本。当 SQL 需要这些特殊函数或聚合时使用它。使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建中进行测试。

### 核心工作流

```sql
CREATE EXTENSION pg_uuid;
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行它，并在将其集成到应用程序 SQL 中之前验证安装的版本和返回值。

### 重要对象

- `uuid_generate_v6()` 是一个扩展函数，返回 `uuid`。
- `uuid_generate_v7()` 是一个扩展函数，返回 `uuid`。

### 要求与注意事项

- 审查后的控制文件声明默认版本为 `1.0.0`。
- 控制文件将扩展标记为可重定位。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为以及失败情况与固定源进行比对。
