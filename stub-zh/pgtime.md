## 用法

来源：

- [官方上游 README](https://github.com/sanidhyavijay24/pgtime/blob/d5815a9780dc3f943e67ffbc9cc97d92c9db53e7/README.md)
- [官方扩展控制文件 (pgtime.control)](https://github.com/sanidhyavijay24/pgtime/blob/d5815a9780dc3f943e67ffbc9cc97d92c9db53e7/extension/pgtime.control)
- [官方扩展 SQL (pgtime--0.1.sql)](https://github.com/sanidhyavijay24/pgtime/blob/d5815a9780dc3f943e67ffbc9cc97d92c9db53e7/extension/pgtime--0.1.sql)

`pgtime` — PostgreSQL 的时间表扩展。使用范围类型和高性能的 C 基础 AFTER ROW 触发器自动跟踪任何表的事务时间（系统时间）。适用于相应的调度、时间或时间序列工作流。请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建中进行测试。

### 核心工作流

```sql
CREATE EXTENSION pgtime;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `pgtime.as_of(target_table TEXT, ts TIMESTAMPTZ)` 是一个扩展函数，返回 `SETOF`。
- `pgtime.attach(target_table TEXT)` 是一个扩展函数，返回 `VOID`。
- `pgtime.detach(target_table TEXT)` 是一个扩展函数，返回 `VOID`。
- `pgtime.diff(target_table TEXT, t1 TIMESTAMPTZ, t2 TIMESTAMPTZ)` 是一个扩展函数，返回 `SETOF`。
- `pgtime.history(target_table TEXT, row_id anyelement)` 是一个扩展函数，返回 `SETOF`。
- `pgtime.pgtime_trigger_fn()` 是一个扩展函数，返回 `TRIGGER`。
- `pgtime.versions(target_table TEXT, row_id anyelement)` 是一个扩展函数，返回 `BIGINT`。
- `pgtime` 是由扩展创建的一个模式。

### 要求与注意事项

- 审查后的控制文件声明默认版本为 `0.1`。
- 控制文件将扩展标记为可重定位。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况。
