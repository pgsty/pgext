## 用法

来源：

- [官方上游 README](https://api.pgxn.org/src/pg_procfs/pg_procfs-0.0.2/README.md)
- [官方扩展控制文件 (pg_procfs.control)](https://api.pgxn.org/src/pg_procfs/pg_procfs-0.0.2/pg_procfs.control)
- [官方扩展 SQL (pg_procfs--0.0.1.sql)](https://api.pgxn.org/src/pg_procfs/pg_procfs-0.0.2/pg_procfs--0.0.1.sql)

`pg_procfs` — 一个用于从 SQL 显示 /proc 文件系统数据的 PostgreSQL 扩展。在收集或解释相应的 PostgreSQL 统计信息时使用它。请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION pg_procfs;

select * from pg_procfs('/proc/version');
 line |                                                                                     message

------+--------------------------------------------------------------------------------------------------------------------------------------
--------------------------------------------
    0 | Linux version 4.18.0-372.19.1.el8_6.x86_64 (mockbuild@49c5e54ed716424c9ae8c1a3d1fef96f) (gcc version 8.5.0 20210514 (Red Hat 8.5.0-10
) (GCC)) #1 SMP Tue Aug 2 13:42:59 EDT 2022
(1 row)
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `pg_procfs(IN filename cstring, OUT line integer, OUT data text)` 是一个扩展函数，返回 `SETOF record`。

### 要求与注意事项

- 审查过的控制文件声明默认版本为 `0.0.1`。
- 控制文件标记该扩展为不可重定位。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况，以匹配固定源代码。
