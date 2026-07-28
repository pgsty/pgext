## 用法

来源：

- [官方上游 README](https://github.com/turbot/steampipe-postgres-fdw/blob/f8098e5e79cda44af01fcb9cd77b9ac684e70690/fdw/README.md)
- [官方扩展控制文件 (steampipe_postgres_fdw.control)](https://github.com/turbot/steampipe-postgres-fdw/blob/f8098e5e79cda44af01fcb9cd77b9ac684e70690/fdw/steampipe_postgres_fdw.control)
- [官方扩展 SQL (steampipe_postgres_fdw--1.0.sql)](https://github.com/turbot/steampipe-postgres-fdw/blob/f8098e5e79cda44af01fcb9cd77b9ac684e70690/fdw/steampipe_postgres_fdw--1.0.sql)

`steampipe_postgres_fdw` — Fdw 是一个用 Go 编写的 Postgres 外部数据封装接口。动态外部表通过 gRPC 插件定义，使其安全、高效且易于构建。当 PostgreSQL 需要通过外部数据接口访问相应的外部数据源时，请使用此扩展。在目标 PostgreSQL 构建中测试链接的上游修订版本作为 API 边界。

### 核心工作流

```sql
CREATE EXTENSION steampipe_postgres_fdw;

create server
  fdw_aws
foreign data wrapper
  fdw
options (
  wrapper 'aws'
);
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `fdw_handler()` 是一个扩展函数，返回 `fdw_handler`。
- `fdw_validator(text[], oid)` 是一个扩展函数，返回 `void`。
- `steampipe_postgres_fdw` 是一个扩展定义的外部数据封装。

### 要求与注意事项

- 审查后的控制文件声明默认版本为 `1.0`。
- 控制文件将扩展标记为可重定位。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源进行验证。
