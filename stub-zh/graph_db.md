## 用法

来源：

- [官方上游 README](https://github.com/dmitriy-m1319/graphextension/blob/e4446ba491752b7ac280f52a32c1533d9e652721/README.md)
- [官方扩展控制文件 (graph_db.control)](https://github.com/dmitriy-m1319/graphextension/blob/e4446ba491752b7ac280f52a32c1533d9e652721/extension/graph_db.control)
- [官方扩展 SQL (graph_db--1.0.sql)](https://github.com/dmitriy-m1319/graphextension/blob/e4446ba491752b7ac280f52a32c1533d9e652721/extension/graph_db--1.0.sql)

`graph_db` — 一个用于图数据库使用的 PostgreSQL 扩展。当应用程序需要这种特定的数据库功能时，请使用它。请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION graph_db;
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行该示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `cypher(cstring, cstring)` 是一个扩展函数，返回 `cstring`。
- `graph_node_in(cstring)` 是一个扩展函数，返回 `graph_node`。
- `graph_node_out(graph_node)` 是一个扩展函数，返回 `cstring`。
- `graph_nodes()` 是一个扩展函数，返回 `graph_node`。
- `key_value(cstring, cstring)` 是一个扩展函数，返回 `cstring`。
- `graph_node` 是一个扩展定义的类型。

### 要求与注意事项

- 审查过的控制文件声明默认版本为 `1.0`。
- 控制文件标记该扩展为不可重定位。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为以及失败情况与固定源代码中的信息一致。
