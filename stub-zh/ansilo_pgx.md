## 用法

来源：

- [官方上游 README](https://github.com/ansilo-data/ansilo/blob/819a32f1782c4d8d4c97e01fe908e2694a546f35/README.md)
- [官方扩展控制文件 (ansilo_pgx.control)](https://github.com/ansilo-data/ansilo/blob/819a32f1782c4d8d4c97e01fe908e2694a546f35/ansilo-pgx/ansilo_pgx.control)
- [官方实现源代码](https://github.com/ansilo-data/ansilo/blob/819a32f1782c4d8d4c97e01fe908e2694a546f35/ansilo-pgx/src/lib.rs)

`ansilo_pgx` — 一个 Postgres 接口，用于访问任何数据库。当 PostgreSQL 需要通过外部数据接口访问相应的外部数据源时，请使用它。请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION ansilo_pgx;
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行该示例，并在将其集成到应用程序 SQL 之前验证已安装的版本和返回值。

### 重要对象

- `hello_ansilo()` 是一个扩展函数。

### 要求与注意事项

- 元组记录版本 `0.1.0`。
- 控制文件将扩展标记为不可重定位。
- 控制文件不要求超级用户安装。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源代码中的信息一致。
