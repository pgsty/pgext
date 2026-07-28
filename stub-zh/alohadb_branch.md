## 用法

来源：

- [官方上游 README](https://github.com/vern-allworks-llc/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/contrib/README)
- [官方扩展控制文件 (alohadb_branch.control)](https://github.com/vern-allworks-llc/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/contrib/alohadb_branch/alohadb_branch.control)
- [官方扩展 SQL (alohadb_branch--1.0.sql)](https://github.com/vern-allworks-llc/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/contrib/alohadb_branch/alohadb_branch--1.0.sql)

`alohadb_branch` — 轻量级数据库分支，用于测试迁移和实验。当应用程序需要此特定数据库功能时使用。上游将此功能描述为实验性。

### 核心工作流

```sql
CREATE EXTENSION alohadb_branch;
```

在目标数据库中安装扩展，运行可用的最小上游示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `alohadb_create_branch(name text, from_lsn pg_lsn DEFAULT NULL, OUT branch_name text, OUT port int, OUT data_dir text)` 是一个扩展函数，返回 `record`。
- `alohadb_drop_branch(name text)` 是一个扩展函数，返回 `void`。
- `alohadb_list_branches()` 是一个扩展函数，返回 `TABLE`。
- `alohadb_branches` 是由扩展安装或管理的表。

### 要求与注意事项

- 控制文件声明默认版本为 `1.0`。
- 控制文件标记该扩展为不可重定位。
- 控制文件要求超级用户进行安装。
- 上游将项目的一部分或全部标记为实验性。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源的一致性。
