## 用法

来源：

- [官方上游 README](https://github.com/vsradkevich/pg_meritrank/blob/6157eac1d22b52357bd84c5c6b0dac76eb33160a/README.md)
- [官方扩展控制文件 (pg_meritrank.control)](https://github.com/vsradkevich/pg_meritrank/blob/6157eac1d22b52357bd84c5c6b0dac76eb33160a/pg_meritrank.control)
- [官方实现源码](https://github.com/vsradkevich/pg_meritrank/blob/6157eac1d22b52357bd84c5c6b0dac76eb33160a/src/lib.rs)

`pg_meritrank` — Postgres Merit Rank 是一个为 PostgreSQL 提供计算和排名功能的扩展。此 README 提供了使用 cargo pgx test 测试扩展以及在 PostgreSQL 数据库中安装它的说明。当需要这些特殊函数或聚合时使用它。使用上方链接的锁定上游版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION pg_meritrank;
```

在目标数据库中安装扩展，在可用时运行上方的最小上游示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 要求与注意事项

- 元数据记录版本 `0.0.1`。
- 控制文件将扩展标记为不可重定位。
- 控制文件要求超级用户进行安装。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与锁定源代码进行验证。
