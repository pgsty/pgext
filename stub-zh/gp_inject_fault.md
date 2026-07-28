## 用法

来源：

- [官方上游 README](https://github.com/apache/cloudberry/blob/83022b58b2f4e62fb8fd2ebf9137b5c29e4fcb5b/gpcontrib/gp_inject_fault/README)
- [官方扩展控制文件 (gp_inject_fault.control)](https://github.com/apache/cloudberry/blob/83022b58b2f4e62fb8fd2ebf9137b5c29e4fcb5b/gpcontrib/gp_inject_fault/gp_inject_fault.control)
- [官方扩展 SQL (gp_inject_fault--1.0.sql)](https://github.com/apache/cloudberry/blob/83022b58b2f4e62fb8fd2ebf9137b5c29e4fcb5b/gpcontrib/gp_inject_fault/gp_inject_fault--1.0.sql)

`gp_inject_fault` — infinite_loop 循环直到查询取消或接收到终止信号。当应用程序需要此特定数据库功能时使用它。请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建中进行测试。

### 核心工作流

```sql
CREATE EXTENSION gp_inject_fault;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证安装的版本和返回值。

### 重要对象

- `force_mirrors_to_catch_up()` 是一个扩展函数，返回 `VOID`。
- `gp_inject_fault(faultname text, type text, db_id int4)` 是一个扩展函数，返回 `text`。
- `gp_inject_fault(faultname text, type text, db_id int4, gp_session_id int4)` 是一个扩展函数，返回 `text`。
- `gp_inject_fault(faultname text, type text, ddl text, database text, tablename text, start_occurrence int4, end_occurrence int4, extra_arg int4, db_id int4)` 是一个扩展函数，返回 `text`。
- `gp_inject_fault(faultname text, type text, ddl text, database text, tablename text, start_occurrence int4, end_occurrence int4, extra_arg int4, db_id int4, gp_session_id int4)` 是一个扩展函数，返回 `text`。
- `gp_inject_fault_infinite(faultname text, type text, db_id int4)` 是一个扩展函数，返回 `text`。
- `gp_wait_until_triggered_fault(faultname text, numtimestriggered int4, db_id int4)` 是一个扩展函数，返回 `text`。
- `insert_noop_xlog_record()` 是一个扩展函数，返回 `VOID`。

### 要求与注意事项

- 控制文件声明默认版本为 `1.0`。
- 控制文件将扩展标记为可重定位。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况，与固定源进行比对。
