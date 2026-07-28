## 用法

来源：

- [官方上游 README](https://github.com/percona-lab/pg_coredump/blob/e63295452fc6b379302f10193d1edfeee7fa8a94/README.md)
- [官方扩展控制文件 (pg_coredump.control)](https://github.com/percona-lab/pg_coredump/blob/e63295452fc6b379302f10193d1edfeee7fa8a94/pg_coredump.control)
- [官方扩展 SQL (pg_coredump--1.0.sql)](https://github.com/percona-lab/pg_coredump/blob/e63295452fc6b379302f10193d1edfeee7fa8a94/pg_coredump--1.0.sql)

`pg_coredump` — 该 PostgreSQL 扩展使得在发生崩溃时生成核心转储文件更加容易。在管理或自动化上述数据库行为时，请使用它。请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION pg_coredump;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `pg_coredump(dumpdir text)` 是一个扩展函数，返回 `boolean`。

### 要求与注意事项

- 审查过的控制文件声明默认版本为 `1.0`。
- 控制文件将扩展标记为可重定位。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况。
