## 用法

来源：

- [官方上游 README](https://api.pgxn.org/src/pg_rand_ext/pg_rand_ext-1.0.1/README.md)
- [官方扩展控制文件 (pg_rand_ext.control)](https://api.pgxn.org/src/pg_rand_ext/pg_rand_ext-1.0.1/pg_rand_ext.control)
- [官方扩展 SQL (pg_rand_ext--1.0.sql)](https://api.pgxn.org/src/pg_rand_ext/pg_rand_ext-1.0.1/pg_rand_ext--1.0.sql)

`pg_rand_ext` — PostgreSQL 中构建扩展模块的常见方式。已在 PostgreSQL v14 及以上版本中进行测试。当需要使用这些特殊函数或聚合时，请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建中进行测试。

### 核心工作流

```sql
CREATE EXTENSION pg_rand_ext;
```

在目标数据库中安装扩展，如果有可用的上游最小示例，请运行该示例，并在将其集成到应用程序 SQL 中之前验证安装版本和返回值。

### 重要对象

- `rand_ext.random_exponential(bigint, bigint, double precision)` 是一个扩展函数，返回 `bigint`。
- `rand_ext.random_gaussian(bigint, bigint, double precision)` 是一个扩展函数，返回 `bigint`。
- `rand_ext.random_zipfian(bigint, bigint, double precision)` 是一个扩展函数，返回 `bigint`。
- `rand_ext` 是由扩展创建的一个模式。

### 要求与注意事项

- 控制文件声明默认版本为 `1.0`。
- 控制文件将扩展标记为可重定位。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为以及失败情况，与固定源进行比对。
