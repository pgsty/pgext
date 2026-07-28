## 用法

来源：

- [官方上游 README](https://github.com/darthunix/pg_vx/blob/5916aafcfde5b7a5c5a90b6ca7ea1879faf85a4d/README.md)
- [官方扩展控制文件 (pg_vx.control)](https://github.com/darthunix/pg_vx/blob/5916aafcfde5b7a5c5a90b6ca7ea1879faf85a4d/pg_vx.control)
- [官方扩展 SQL (pg_vx--0.1.sql)](https://github.com/darthunix/pg_vx/blob/5916aafcfde5b7a5c5a90b6ca7ea1879faf85a4d/pg_vx--0.1.sql)

`pg_vx` — PG_VX 是 PostgreSQL 向量化执行器。它是一个测试项目，旨在了解 PostgreSQL CustomScanAPI 的工作原理，并验证其是否可以用于对 OLAP 工作负载进行向量化加速的扫描、聚合和连接操作。当应用程序需要此特定数据库功能时，请使用它。在将扩展集成到应用程序 SQL 之前，请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION pg_vx;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 之前验证已安装的版本和返回值。

### 要求与注意事项

- 审查后的控制文件声明默认版本为 `0.1`。
- 控制文件将扩展标记为可重定位。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为以及失败情况，以匹配固定源代码。
