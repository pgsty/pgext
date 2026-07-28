## 用法

来源：

- [官方上游 README](https://github.com/michaelpq/pg_plugins/blob/626fb56b0a0b833d4f23ca55359ce56d38162864/pg_rusage/README)
- [官方扩展控制文件 (pg_rusage.control)](https://github.com/michaelpq/pg_plugins/blob/626fb56b0a0b833d4f23ca55359ce56d38162864/pg_rusage/pg_rusage.control)
- [官方扩展 SQL (pg_rusage--1.0.sql)](https://github.com/michaelpq/pg_plugins/blob/626fb56b0a0b833d4f23ca55359ce56d38162864/pg_rusage/pg_rusage--1.0.sql)

`pg_rusage` — 这个模块是一个 PostgreSQL 扩展，可以启用 CPU 测量，包含一个用于启用测量的 SQL 函数和一个用于禁用测量的 SQL 函数。在禁用时，现有的累积结果将会显示出来。在收集或解释相应的 PostgreSQL 统计信息时使用它。请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建中进行测试。

### 核心工作流

```sql
CREATE EXTENSION pg_rusage;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证安装的版本和返回值。

### 重要对象

- `pg_rusage_print()` 是一个扩展函数，返回 `void`。
- `pg_rusage_reset()` 是一个扩展函数，返回 `void`。

### 要求与注意事项

- 审查后的控制文件声明默认版本为 `1.0`。
- 控制文件将扩展标记为可重定位。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况，与固定源进行比对。
