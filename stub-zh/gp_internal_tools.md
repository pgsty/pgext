## 用法

来源：

- [官方上游 README](https://github.com/gridgentoo/gpdb/blob/f3dc101a7b4fa3d392f79cc5146b20c83894eb19/contrib/gp_internal_tools/README)
- [官方扩展控制文件 (gp_internal_tools.control)](https://github.com/gridgentoo/gpdb/blob/f3dc101a7b4fa3d392f79cc5146b20c83894eb19/contrib/gp_internal_tools/gp_internal_tools.control)
- [官方扩展 SQL (gp_internal_tools--1.0.0.sql)](https://github.com/gridgentoo/gpdb/blob/f3dc101a7b4fa3d392f79cc5146b20c83894eb19/contrib/gp_internal_tools/gp_internal_tools--1.0.0.sql)

`gp_internal_tools` — 不同的 Greenplum 内部工具。在进行数据库管理或自动化上述数据库行为时使用此扩展。使用链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION gp_internal_tools;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `session_state_memory_entries_f_on_master()` 是一个扩展函数，返回 `SETOF`。
- `session_state_memory_entries_f_on_segments()` 是一个扩展函数，返回 `SETOF`。
- `session_level_memory_consumption` 是一个扩展定义视图。
- `session_state` 是由扩展创建的一个模式。

### 要求与注意事项

- 控制文件声明默认版本为 `1.0.0`。
- 控制文件将扩展标记为可重定位。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况。
