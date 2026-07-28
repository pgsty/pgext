## 用法

来源：

- [官方上游 README](https://github.com/misachi/pg_wal_recovery/blob/61ac3633d26bf414d9b145b153bc96d6b0e75a77/README.md)
- [官方扩展控制文件 (pg_wal_recovery.control)](https://github.com/misachi/pg_wal_recovery/blob/61ac3633d26bf414d9b145b153bc96d6b0e75a77/pg_wal_recovery.control)
- [官方扩展 SQL (pg_wal_recovery--1.0.sql)](https://github.com/misachi/pg_wal_recovery/blob/61ac3633d26bf414d9b145b153bc96d6b0e75a77/pg_wal_recovery--1.0.sql)

`pg_wal_recovery` — pg_wal_recovery 是一个用于数据库恢复的教育性 PostgreSQL 扩展，专注于从 Write-Ahead Logs (WAL) 恢复数据库并支持点时间恢复。请参阅我对此的帖子。在管理或自动化上述数据库行为时使用它。上游将其描述为一个正在进行中的项目。

### 核心工作流

```sql
CREATE EXTENSION pg_wal_recovery;
```

在目标数据库中安装扩展，当可用时运行上游提供的最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `wal_list_records` 是一个扩展函数。

### 要求与注意事项

- 审查的控制文件声明默认版本为 `1.0`。
- 控制文件标记该扩展为可重定位。
- 上游将该项目描述为一个正在进行中的项目。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况。
