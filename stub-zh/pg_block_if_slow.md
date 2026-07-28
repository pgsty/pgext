## 用法

来源：

- [官方上游 README](https://github.com/eshwar-333/pg_block_if_slow/blob/2313c00bb439332fe3d54321037c13d9522f53ee/README.md)
- [官方扩展控制文件 (pg_block_if_slow.control)](https://github.com/eshwar-333/pg_block_if_slow/blob/2313c00bb439332fe3d54321037c13d9522f53ee/pg_block_if_slow.control)
- [官方扩展 SQL (pg_block_if_slow--1.0.sql)](https://github.com/eshwar-333/pg_block_if_slow/blob/2313c00bb439332fe3d54321037c13d9522f53ee/pg_block_if_slow--1.0.sql)

`pg_block_if_slow` — 当启用时，pg_block_if_slow 扩展可以防止任何查询在执行前，如果其 **估计成本超过定义的阈值**。在管理或自动化上述数据库行为时使用它。请使用链接中的固定上游版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION pg_block_if_slow;
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行它，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 要求与注意事项

- 审查过的控制文件声明默认版本为 `1.0`。
- 控制文件将扩展标记为不可重定位。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为以及失败情况，以匹配固定源代码。
