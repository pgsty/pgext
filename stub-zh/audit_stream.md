## 用法

来源：

- [官方上游 README](https://github.com/mizcausevic-dev/pg-audit-stream-extension/blob/73a456241b0c6fa0ef4f46ba84c97bc0e39cc126/README.md)
- [官方扩展控制文件 (audit_stream.control)](https://github.com/mizcausevic-dev/pg-audit-stream-extension/blob/73a456241b0c6fa0ef4f46ba84c97bc0e39cc126/audit_stream.control)
- [官方扩展 SQL (audit_stream--0.1.0.sql)](https://github.com/mizcausevic-dev/pg-audit-stream-extension/blob/73a456241b0c6fa0ef4f46ba84c97bc0e39cc126/audit_stream--0.1.0.sql)

`audit_stream` — 这是一个 Postgres 扩展，无需编写一行应用程序代码即可将任何表级 CRUD 转换为 audit-stream-py 兼容的治理事件。在实现相应的安全、审计或访问控制工作流时使用它。请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION audit_stream;
SELECT audit_stream.watch('decisions', 'decision_card_status_changed', 'procurement-api');
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证安装的版本和返回值。

### 重要对象

- `list_watches()` 是一个扩展函数，返回 `TABLE`。
- `unwatch(p_table TEXT)` 是一个扩展函数，返回 `BOOLEAN`。
- `watch(p_table TEXT, p_event_kind TEXT, p_source TEXT DEFAULT NULL)` 是一个扩展函数，返回 `TEXT`。
- `watches` 是由扩展安装或管理的表。

### 要求与注意事项

- 审查后的控制文件声明默认版本为 `0.1.0`。
- 控制文件将扩展标记为不可重定位。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源进行比对。
