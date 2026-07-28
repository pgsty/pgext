## 用法

来源：

- [官方上游 README](https://github.com/apache/cloudberry/blob/83022b58b2f4e62fb8fd2ebf9137b5c29e4fcb5b/README.md)
- [官方扩展控制文件 (gp_exttable_fdw.control)](https://github.com/apache/cloudberry/blob/83022b58b2f4e62fb8fd2ebf9137b5c29e4fcb5b/gpcontrib/gp_exttable_fdw/gp_exttable_fdw.control)
- [官方扩展 SQL (gp_exttable_fdw--1.0.sql)](https://github.com/apache/cloudberry/blob/83022b58b2f4e62fb8fd2ebf9137b5c29e4fcb5b/gpcontrib/gp_exttable_fdw/gp_exttable_fdw--1.0.sql)

`gp_exttable_fdw` — 外部表外模式数据封装器，用于 Greenplum 家族数据库。当 PostgreSQL 需要通过外模式接口访问相应的外部数据源时，请使用此扩展。请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建中进行测试。

### 核心工作流

```sql
CREATE EXTENSION gp_exttable_fdw;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `gp_exttable_fdw_handler()` 是一个扩展函数，返回 `fdw_handler`。
- `gp_exttable_permission_check(text[], oid)` 是一个扩展函数，返回 `void`。
- `pg_exttable` 是一个扩展定义视图。
- `gp_exttable_fdw` 是一个扩展定义的外模式数据封装器。

### 要求与注意事项

- 控制文件声明默认版本为 `1.0`。
- 控制文件将扩展标记为不可重定位。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源中的信息一致。
