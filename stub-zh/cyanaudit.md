## 用法

来源：

- [官方上游 README](https://api.pgxn.org/src/cyanaudit/cyanaudit-2.2.0/README.md)
- [官方扩展 SQL (cyanaudit--2.0.sql)](https://api.pgxn.org/src/cyanaudit/cyanaudit-2.2.0/sql/cyanaudit--2.0.sql)

`cyanaudit` — Cyan Audit 是一个 PostgreSQL 工具，提供全面且易于搜索的 DML（INSERT/UPDATE/DELETE）活动日志。在实施相应的安全、审计或访问控制工作流时，请使用它。请使用上述链接的固定上游版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

此组件在审查源代码中没有确认的独立 `CREATE EXTENSION` 工作流。仅通过上游机制构建、加载或启用它，然后在隔离数据库中验证其结果服务器行为。

### 重要对象

- `cyanaudit.fn_activate_partition(in_partition_name varchar)` 是一个扩展函数，返回 `void`。
- `cyanaudit.fn_add_trigger_to_extension(in_table_schema varchar, in_table_name varchar)` 是一个扩展函数，返回 `void`。
- `cyanaudit.fn_after_audit_field_change()` 是一个扩展函数，返回 `trigger`。
- `cyanaudit.fn_archive_partition(in_partition_name varchar)` 是一个扩展函数，返回 `void`。
- `cyanaudit.fn_before_audit_field_change()` 是一个扩展函数，返回 `trigger`。
- `cyanaudit.fn_create_event_trigger()` 是一个扩展函数，返回 `void`。
- `cyanaudit.fn_create_new_partition(in_new_table_name varchar default 'tb_audit_event_' || to_char(now(), 'YYYYMMDD_HH24MI'))` 是一个扩展函数，返回 `varchar`。
- `cyanaudit.fn_create_partition_indexes(in_table_name varchar)` 是一个扩展函数，返回 `void`。
- `cyanaudit.fn_get_active_partition_name()` 是一个扩展函数，返回 `varchar`。
- `cyanaudit.fn_get_current_uid()` 是一个扩展函数，返回 `integer`。
- `cyanaudit.fn_get_email_by_uid(in_uid integer)` 是一个扩展函数，返回 `varchar`。
- `cyanaudit.fn_get_last_txid()` 是一个扩展函数，返回 `bigint`。
- `cyanaudit.fn_get_or_create_audit_field(in_table_schema varchar, in_table_name varchar, in_column_name varchar)` 是一个扩展函数，返回 `integer`。
- `cyanaudit.fn_get_or_create_audit_transaction_type(in_label varchar)` 是一个扩展函数，返回 `integer`。

### 要求与注意事项

- 请确认版本记录 `2.2.0`。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源代码一致。
