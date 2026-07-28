## 用法

来源：

- [官方上游 README](https://api.pgxn.org/src/pg_monetdb/pg_monetdb-1.4.0/README.md)
- [官方扩展控制文件 (pg_monetdb.control)](https://api.pgxn.org/src/pg_monetdb/pg_monetdb-1.4.0/pg_monetdb.control)
- [官方扩展 SQL (pg_monetdb--1.3.sql)](https://api.pgxn.org/src/pg_monetdb/pg_monetdb-1.4.0/pg_monetdb--1.3.sql)

`pg_monetdb` — pg_monetdb 是 monetdb_fdw 的分支，专注于为从 TPC-H 和 TPC-DS 样式的负载中派生的分析查询提供更强的下推能力。当 PostgreSQL 需要通过外部数据接口访问相应的外部数据源时，请使用此扩展。上游将此功能描述为实验性。

### 核心工作流

```sql
CREATE EXTENSION pg_monetdb;
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行该示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `monet_query(server name, statement text)` 是一个扩展函数，返回 `SETOF`。
- `monet_query_to_array(server name, statement text)` 是一个扩展函数，返回 `SETOF`。
- `monet_query_to_jsonb(server name, statement text, column_names text[])` 是一个扩展函数，返回 `SETOF`。
- `monetdb_execute(server name, statement text)` 是一个扩展函数，返回 `void`。
- `monetdb_fdw_handler()` 是一个扩展函数，返回 `fdw_handler`。
- `pg_monetdb_execute(server name, statement text)` 是一个扩展函数，返回 `void`。
- `pg_monetdb_handler()` 是一个扩展函数，返回 `fdw_handler`。
- `pg_monetdb_query(server name, statement text)` 是一个扩展函数，返回 `SETOF`。
- `pg_monetdb_query_to_array(server name, statement text)` 是一个扩展函数，返回 `SETOF`。
- `pg_monetdb_query_to_jsonb(server name, statement text, column_names text[])` 是一个扩展函数，返回 `SETOF`。
- `BLOB` 是一个扩展定义的域。
- `CLOB` 是一个扩展定义的域。
- `HUGEINT` 是一个扩展定义的域。
- `STRING` 是一个扩展定义的域。

### 要求与注意事项

- 审查后的控制文件声明默认版本为 `1.4`。
- 控制文件将扩展标记为可重定位。
- 上游将项目的一部分或全部标记为实验性。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况，以匹配固定源。
