## 用法

来源：

- [官方上游 README](https://github.com/x4m/pg_uuid_next/blob/21d2df4203448e8d4149859dce77ce420bc0d1d8/README.md)
- [官方扩展控制文件 (pg_uuid_next.control)](https://github.com/x4m/pg_uuid_next/blob/21d2df4203448e8d4149859dce77ce420bc0d1d8/pg_uuid_next.control)
- [官方扩展 SQL (pg_uuid_next--1.0.sql)](https://github.com/x4m/pg_uuid_next/blob/21d2df4203448e8d4149859dce77ce420bc0d1d8/pg_uuid_next--1.0.sql)

`pg_uuid_next` — 扩展用于生成 UUID 版本 7 和 8。当 SQL 需要这些特殊函数或聚合时，请使用此扩展。请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建中进行测试。

### 核心工作流

```sql
CREATE EXTENSION pg_uuid_next;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证安装的版本和返回值。

### 重要对象

- `gen_uuid_v7()` 是一个扩展函数，返回 `uuid`。
- `gen_uuid_v8()` 是一个扩展函数，返回 `uuid`。

### 要求与注意事项

- 审查后的控制文件声明默认版本为 `1.0`。
- 控制文件将扩展标记为可重定位。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源进行比对。
