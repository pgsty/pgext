## 用法

来源：

- [官方上游 README](https://github.com/timescale/pg_traceam/blob/a8aafe6448e6fa110f283dafb0b201c846857e57/README.md)
- [官方扩展控制文件 (traceam.control)](https://github.com/timescale/pg_traceam/blob/a8aafe6448e6fa110f283dafb0b201c846857e57/traceam.control)
- [官方扩展 SQL (traceam--0.1.sql)](https://github.com/timescale/pg_traceam/blob/a8aafe6448e6fa110f283dafb0b201c846857e57/traceam--0.1.sql)

`traceam` — 安装扩展后，还需要在数据库中安装它，这可以通过 CREATE EXTENSION: 完成。当应用程序需要此特定数据库功能时使用它。使用上方链接的固定上游版本作为 API 边界，并在目标 PostgreSQL 构建中进行测试。

### 核心工作流

```sql
CREATE EXTENSION traceam;
```

在目标数据库中安装扩展，当可用时运行上方的最小上游示例，并在将其集成到应用程序 SQL 之前验证已安装的版本和返回值。

### 重要对象

- `traceam_handler(internal)` 是一个扩展函数并返回 `table_am_handler`。
- `traceam` 是由扩展创建的一个模式。
- `traceam` 是一个由扩展定义的访问方法。

### 要求与注意事项

- 审核后的控制文件声明默认版本 `0.1`。
- 控制文件将扩展标记为不可重定位。
- 在生产使用前，确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源代码一致。
