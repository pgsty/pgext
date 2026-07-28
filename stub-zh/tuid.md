## 用法

来源：

- [官方上游 README](https://github.com/tanglebones/pg_tuid/blob/eca0bbd95cd17dda004d1adbe2546048de5650b9/README.md)
- [官方扩展控制文件 (tuid.control)](https://github.com/tanglebones/pg_tuid/blob/eca0bbd95cd17dda004d1adbe2546048de5650b9/pg_c/tuid.control)
- [官方扩展 SQL (tuid--0.3.0.sql)](https://github.com/tanglebones/pg_tuid/blob/eca0bbd95cd17dda004d1adbe2546048de5650b9/pg_c/tuid--0.3.0.sql)

`tuid` — tuid 数据类型。当 SQL 需要这些特殊函数或聚合时使用它。请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建中进行测试。

### 核心工作流

```sql
CREATE EXTENSION tuid;

select uuidv7();
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行它，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `stuid_generate()` 是一个扩展函数，返回 `bytea`。
- `tuid_generate()` 是一个扩展函数，返回 `uuid`。

### 要求与注意事项

- 审查后的控制文件声明默认版本为 `0.3.0`。
- 控制文件将扩展标记为可重定位。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源进行比对。
