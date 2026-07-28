## 用法

来源：

- [官方上游 README](https://api.pgxn.org/src/pg_global_catalog/pg_global_catalog-0.0.1/README.md)
- [官方扩展控制文件 (pg_global_catalog.control)](https://api.pgxn.org/src/pg_global_catalog/pg_global_catalog-0.0.1/pg_global_catalog.control)
- [官方扩展 SQL (pg_global_catalog--0.0.1.sql)](https://api.pgxn.org/src/pg_global_catalog/pg_global_catalog-0.0.1/pg_global_catalog--0.0.1.sql)

`pg_global_catalog` — PostgreSQL 扩展，用于在每个数据库中将 pg_catalog 合并到一个名为 global_catalog 的单一模式中。当您需要管理或自动化上述描述的数据库行为时，请使用此扩展。请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION pg_global_catalog;

select
datname, count(*)
from global_catalog.pg_class c
join pg_database d
on c.dbid = d.oid
group by datname
order by datname;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `pggc_create_fdws()` 是一个扩展函数，返回 `void`。
- `pggc_create_global_views()` 是一个扩展函数，返回 `void`。

### 要求与注意事项

- 审查后的控制文件声明默认版本为 `0.0.1`。
- 控制文件将扩展标记为不可重定位。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源进行比对。
