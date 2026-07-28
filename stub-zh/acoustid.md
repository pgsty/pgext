## 用法

来源：

- [官方上游 README](https://github.com/acoustid/pg_acoustid/blob/102f3c870c6157704694c2ddbad3ae8ab2c7de91/README.md)
- [官方扩展控制文件 (acoustid.control)](https://github.com/acoustid/pg_acoustid/blob/102f3c870c6157704694c2ddbad3ae8ab2c7de91/acoustid.control)
- [官方扩展 SQL (acoustid--1.0.sql)](https://github.com/acoustid/pg_acoustid/blob/102f3c870c6157704694c2ddbad3ae8ab2c7de91/acoustid--1.0.sql)

`acoustid` — AcoustID 工具函数 for PostgreSQL =========================================。当 SQL 需要这些特殊函数或聚合时使用它。使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建中进行测试。

### 核心工作流

```sql
CREATE EXTENSION acoustid;
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行它，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `acoustid_compare2(int4[], int4[], int DEFAULT 0)` 是一个扩展函数，返回 `float4`。
- `acoustid_compare3(int4[], int4[], int DEFAULT -1)` 是一个扩展函数，返回 `float4`。
- `acoustid_extract_query(int4[])` 是一个扩展函数，返回 `int4[]`。

### 要求与注意事项

- 审查后的控制文件声明默认版本为 `1.0`。
- 控制文件将扩展标记为可重定位。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源进行比对。
