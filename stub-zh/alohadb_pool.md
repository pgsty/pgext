## 用法

来源：

- [官方上游 README](https://github.com/vern-allworks-llc/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/contrib/README)
- [官方扩展控制文件 (alohadb_pool.control)](https://github.com/vern-allworks-llc/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/contrib/alohadb_pool/alohadb_pool.control)
- [官方扩展 SQL (alohadb_pool--1.0.sql)](https://github.com/vern-allworks-llc/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/contrib/alohadb_pool/alohadb_pool--1.0.sql)

`alohadb_pool` — AlohaDB 内置连接池。在管理或自动化上述数据库行为时使用它。上游将此功能描述为实验性功能。

### 核心工作流

```sql
CREATE EXTENSION alohadb_pool;
```

在目标数据库中安装扩展，当可用时运行上游提供的最小示例，并在将其集成到应用程序 SQL 之前验证已安装的版本和返回值。

### 重要对象

- `pool_reset(pool_name text DEFAULT NULL)` 是一个扩展函数，返回 `void`。
- `pool_settings()` 是一个扩展函数，返回 `TABLE`。
- `pool_status()` 是一个扩展函数，返回 `TABLE`。

### 要求与注意事项

- 审查过的控制文件声明默认版本为 `1.0`。
- 控制文件标记该扩展为不可重定位。
- 上游将项目的一部分或全部标记为实验性。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况，以匹配固定源。
