## 用法

来源：

- [官方上游 README](https://github.com/fljdin/mysql_migrator/blob/3c15bd7d7b9a2183d3c9753b3ec8fd238b5ca2cc/README.md)
- [官方扩展控制文件 (mysql_migrator.control)](https://github.com/fljdin/mysql_migrator/blob/3c15bd7d7b9a2183d3c9753b3ec8fd238b5ca2cc/mysql_migrator.control)
- [官方扩展 SQL (mysql_migrator--0.3.0.sql)](https://github.com/fljdin/mysql_migrator/blob/3c15bd7d7b9a2183d3c9753b3ec8fd238b5ca2cc/mysql_migrator--0.3.0.sql)

`mysql_migrator` — MySQL/MariaDB 到 PostgreSQL 迁移工具
使用此工具进行数据迁移、转换或集成。在应用之前，必须先安装并验证其扩展依赖。

### 核心工作流

```sql
CREATE EXTENSION mysql_migrator;
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行该示例，并在将其集成到应用程序 SQL 之前验证已安装的版本和返回值。

### 重要对象

- `db_migrator_callback(OUT create_metadata_views_fun regprocedure, OUT translate_datatype_fun regprocedure, OUT translate_identifier_fun regprocedure, OUT translate_expression_fun regprocedure, OUT create_foreign_table_fun regprocedure)` 是一个扩展函数，返回 `record`。
- `mysql_create_catalog(server name, schema name DEFAULT NAME 'public', options jsonb DEFAULT NULL)` 是一个扩展函数，返回 `void`。
- `mysql_migrate_identity(pgstage_schema name DEFAULT NAME 'pgsql_stage')` 是一个扩展函数，返回 `integer`。
- `mysql_mkforeign(server name, schema name, table_name name, orig_schema text, orig_table text, column_names name[], column_options jsonb[], orig_columns text[], data_types text[], nullable boolean[], options jsonb)` 是一个扩展函数，返回 `text`。
- `mysql_translate_datatype(v_type text, v_length integer, v_precision integer, v_scale integer)` 是一个扩展函数，返回 `text`。
- `mysql_translate_expression(s text)` 是一个扩展函数，返回 `text`。
- `mysql_translate_identifier_noop(text)` 是一个扩展函数，返回 `name`。

### 要求与注意事项

- 控制文件声明默认版本为 `0.3.0`。
- 首先安装并验证确认的扩展依赖：`mysql_fdw`, `db_migrator`。
- 控制文件将扩展标记为可重定位。
- 控制文件不要求超级用户安装。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源进行比对。
