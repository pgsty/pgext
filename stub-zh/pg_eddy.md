## 用法

来源：

- [官方上游 README](https://github.com/trickle-labs/pg-eddy/blob/5afc17c9b630788dffeee5a6a603a6df99f22183/README.md)
- [官方扩展控制文件 (pg_eddy.control)](https://github.com/trickle-labs/pg-eddy/blob/5afc17c9b630788dffeee5a6a603a6df99f22183/pg_eddy/pg_eddy.control)
- [官方扩展 SQL (pg_eddy--0.1.0.sql)](https://github.com/trickle-labs/pg-eddy/blob/5afc17c9b630788dffeee5a6a603a6df99f22183/pg_eddy/sql/pg_eddy--0.1.0.sql)

`pg_eddy` — 一个用于带标签的属性图的 Postgres 扩展，具有无索引邻接性和内置的物化视图。当应用程序需要此特定数据库功能时，请使用它。请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION pg_eddy;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `pg_eddy_edge_handler(internal)` 是一个扩展函数，返回 `table_am_handler`。
- `pg_eddy_node_handler(internal)` 是一个扩展函数，返回 `table_am_handler`。
- `edges` 是一个由扩展定义的视图。
- `nodes` 是一个由扩展定义的视图。
- `_pg_eddy.edge_type_dst` 是一个由扩展安装或管理的表。
- `_pg_eddy.edge_type_src` 是一个由扩展安装或管理的表。
- `_pg_eddy.edges` 是一个由扩展安装或管理的表。
- `_pg_eddy.label_index` 是一个由扩展安装或管理的表。
- `_pg_eddy.label_registry` 是一个由扩展安装或管理的表。
- `_pg_eddy.nodes` 是一个由扩展安装或管理的表。
- `_pg_eddy.property_key_registry` 是一个由扩展安装或管理的表。
- `_pg_eddy.rel_type_registry` 是一个由扩展安装或管理的表。
- `edges` 是一个由扩展安装或管理的表。
- `nodes` 是一个由扩展安装或管理的表。

### 要求与注意事项

- 该目录记录版本 `0.6.0`。
- 控制文件将扩展标记为不可重定位。
- 控制文件要求超级用户进行安装。
- 控制文件将扩展标记为不受信任。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源中的信息一致。
