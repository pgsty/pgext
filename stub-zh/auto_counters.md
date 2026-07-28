## 用法

来源：

- [官方上游 README](https://github.com/mbcheikh/auto_counters/blob/918ef7b428697d563572b51e271a8593f74847f6/README.md)
- [官方扩展控制文件 (auto_counters.control)](https://github.com/mbcheikh/auto_counters/blob/918ef7b428697d563572b51e271a8593f74847f6/auto_counters.control)
- [官方扩展 SQL (auto_counters--1.0.sql)](https://github.com/mbcheikh/auto_counters/blob/918ef7b428697d563572b51e271a8593f74847f6/auto_counters--1.0.sql)

`auto_counters` — 一个基于多个字段组合的 **自动上下文编号** 的强大且灵活的 PostgreSQL 扩展。适用于生成文档编号、发票编号或任何依赖于上下文信息（如年份、部门或类别）的编号系统。当 SQL 需要这些特殊功能或聚合函数时，请使用此扩展。请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建中进行测试。

### 核心工作流

```sql
CREATE EXTENSION auto_counters;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `auto_counters_install()` 是一个扩展函数，返回 `void`。
- `create_counter_def(p_counter_id VARCHAR, p_table_name VARCHAR, p_fields TEXT [], p_description TEXT DEFAULT NULL, p_is_active BOOLEAN DEFAULT TRUE)` 是一个扩展函数，返回 `VOID`。
- `create_counter_trigger_on_def_insert()` 是一个扩展函数，返回 `TRIGGER`。
- `delete_counter_def(p_counter_id VARCHAR, p_cascade BOOLEAN DEFAULT FALSE)` 是一个扩展函数，返回 `VOID`。
- `generic_counter_trigger()` 是一个扩展函数，返回 `TRIGGER`。
- `get_counter_def(p_counter_id VARCHAR DEFAULT NULL)` 是一个扩展函数，返回 `TABLE`。
- `get_counter_status(p_counter_id VARCHAR DEFAULT NULL)` 是一个扩展函数，返回 `TABLE`。
- `get_next_counter_value(p_counter_id VARCHAR, p_key_values TEXT [])` 是一个扩展函数，返回 `INTEGER`。
- `set_field(record_data ANYELEMENT, field_name TEXT, field_value ANYELEMENT)` 是一个扩展函数，返回 `ANYELEMENT`。
- `sync_all_counter_triggers()` 是一个扩展函数，返回 `void`。
- `toggle_counter_def(p_counter_id VARCHAR, p_is_active BOOLEAN)` 是一个扩展函数，返回 `VOID`。
- `update_counter_def(p_counter_id VARCHAR, p_table_name VARCHAR DEFAULT NULL, p_fields TEXT [] DEFAULT NULL, p_description TEXT DEFAULT NULL, p_is_active BOOLEAN DEFAULT NULL)` 是一个扩展函数，返回 `VOID`。
- `update_counter_trigger_on_def_change()` 是一个扩展函数，返回 `TRIGGER`。
- `vw_counter_status` 是一个扩展定义视图。

### 要求与注意事项

- 控制文件声明默认版本为 `1.0`。
- 控制文件标记该扩展为不可重定位。
- 控制文件不要求超级用户安装。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况。
