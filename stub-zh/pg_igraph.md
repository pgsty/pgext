## 用法

来源：

- [官方上游 README](https://api.pgxn.org/src/pg_igraph/pg_igraph-1.1.0/README.md)
- [官方扩展控制文件 (pg_igraph.control)](https://api.pgxn.org/src/pg_igraph/pg_igraph-1.1.0/pg_igraph.control)
- [官方扩展 SQL (pg_igraph--1.0.sql)](https://api.pgxn.org/src/pg_igraph/pg_igraph-1.1.0/pg_igraph--1.0.sql)

`pg_igraph` — **PostgreSQL 高性能图遍历引擎**。当应用程序需要此特定数据库功能时使用它。请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION pg_igraph;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `graph_add_complex_field(type_id SMALLINT, pos SMALLINT, field_name TEXT)` 是一个扩展函数，返回 `VOID`。
- `graph_add_complex_type(type_name TEXT)` 是一个扩展函数，返回 `SMALLINT`。
- `graph_add_edge(from_id BIGINT, to_id BIGINT, rel_name TEXT)` 是一个扩展函数，返回 `VOID`。
- `graph_add_edge(from_id BIGINT, to_id BIGINT, rel_name TEXT, table_prefix TEXT DEFAULT '')` 是一个扩展函数，返回 `VOID`。
- `graph_add_edge(from_id INT, to_id INT, rel_name TEXT)` 是一个扩展函数，返回 `VOID`。
- `graph_add_edge(from_id INT, to_id INT, rel_name TEXT, table_prefix TEXT DEFAULT '')` 是一个扩展函数，返回 `VOID`。
- `graph_add_node(label_name TEXT)` 是一个扩展函数，返回 `BIGINT`。
- `graph_add_node(label_name TEXT, table_prefix TEXT DEFAULT '')` 是一个扩展函数，返回 `BIGINT`。
- `graph_delete_node(node_id BIGINT)` 是一个扩展函数，返回 `VOID`。
- `graph_delete_node(node_id BIGINT, table_prefix TEXT DEFAULT '')` 是一个扩展函数，返回 `VOID`。
- `graph_delete_node(node_id INT)` 是一个扩展函数，返回 `VOID`。
- `graph_delete_node(node_id INT, table_prefix TEXT DEFAULT '')` 是一个扩展函数，返回 `VOID`。
- `graph_delete_property(node_id BIGINT, prop_name TEXT)` 是一个扩展函数，返回 `VOID`。
- `graph_delete_property(node_id INT, prop_name TEXT)` 是一个扩展函数，返回 `VOID`。

### 要求与注意事项

- 控制文件声明默认版本为 `1.1`。
- 控制文件标记该扩展为可重定位。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况。
