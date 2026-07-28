## 用法

来源：

- [官方扩展控制文件 (datafly_anon.control)](https://api.pgxn.org/src/datafly_anon/datafly_anon-1.0.30/datafly_anon.control)
- [官方扩展 SQL (datafly_anon--1.25.sql)](https://api.pgxn.org/src/datafly_anon/datafly_anon-1.0.30/sql/datafly_anon--1.25.sql)

`datafly_anon` — Datafly匿名化器。在实现相应的安全、审计或访问控制工作流时使用它。请使用上述链接的上游版本作为 API 边界，并在目标 PostgreSQL 构建中进行测试。

### 核心工作流

```sql
CREATE EXTENSION datafly_anon;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `add_level_generalization(sch_name varchar, attribute_name varchar,tbl_name varchar,generalization_rule varchar, new_level integer,function varchar, target_sch_name varchar, target_view varchar, re_init_anon bool)` 是一个扩展函数，返回 `void`。
- `check_if_generalization_rule_exists(attribute_name varchar, schema_name varchar, table_name varchar, rule varchar, target_schema_name varchar, target_table_name varchar)` 是一个扩展函数，返回 `boolean`。
- `check_if_level_exists(attribute_name varchar, schema_name varchar,table_name varchar, new_level integer, target_schema_name varchar, target_table_name varchar)` 是一个扩展函数，返回 `boolean`。
- `configure_plugin(json_config json)` 是一个扩展函数，返回 `text`。
- `does_column_exist_in_table(attribute_name varchar, sch_name varchar,tbl_name varchar)` 是一个扩展函数，返回 `boolean`。
- `does_table_exist(schema_name varchar, tbl_name varchar)` 是一个扩展函数，返回 `boolean`。
- `generalize(attribute_name varchar, target_sch_name varchar, target_view varchar, sch_name varchar, tbl_name varchar)` 是一个扩展函数，返回 `void`。
- `generalize_daterange(val DATE, step TEXT)` 是一个扩展函数，返回 `DATERANGE`。
- `generalize_numrange(val NUMERIC, step VARCHAR)` 是一个扩展函数，返回 `NUMRANGE`。
- `generate_init_view(sch_name varchar, tbl_name varchar, target_sch_name varchar, target_view varchar, test_mode bool, is_triggered bool)` 是一个扩展函数，返回 `varchar`。
- `generate_triggers(k_param integer,schema_name varchar, table_name varchar, target_schema_name varchar, target_table_name varchar)` 是一个扩展函数，返回 `void`。
- `init_datafly(k integer, sch_name varchar, tbl_name varchar, target_sch_name varchar, target_view varchar, test_mode bool, is_triggered bool default false)` 是一个扩展函数，返回 `text`。
- `init_datafly_tg()` 是一个扩展函数，返回 `TRIGGER`。
- `remove_level_generalization(attribute_name varchar, tbl_name varchar, sch_name varchar, target_sch_name varchar, target_view varchar, generalization_lvl integer,re_init_anon bool)` 是一个扩展函数，返回 `void`。

### 要求与注意事项

- 审查后的控制文件声明默认版本为 `1.25`。
- 控制文件将扩展标记为可重定位。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况。
