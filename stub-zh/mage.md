## 用法

来源：

- [官方上游 README](https://github.com/yugabyte/yugabyte-db/blob/0f345ba55edef263a2f67c44438dbdffcbb44754/src/postgres/third-party-extensions/mage/README.md)
- [官方扩展控制文件 (mage.control)](https://github.com/yugabyte/yugabyte-db/blob/0f345ba55edef263a2f67c44438dbdffcbb44754/src/postgres/third-party-extensions/mage/mage.control)
- [官方扩展 SQL (mage--1.5.0--1.6.0.sql)](https://github.com/yugabyte/yugabyte-db/blob/0f345ba55edef263a2f67c44438dbdffcbb44754/src/postgres/third-party-extensions/mage/mage--1.5.0--1.6.0.sql)

`mage` — 由于 AGE 是基于强大的 PostgreSQL 关系型数据库系统构建的，因此它具有强大的功能和全面的特性。当应用程序需要这种特定的数据库能力时，请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION mage;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `mag_catalog.age_graph_stats(agtype)` 是一个扩展函数，返回 `agtype`。
- `mag_catalog.age_is_valid_label_name(agtype)` 是一个扩展函数，返回 `boolean`。
- `mag_catalog.age_tostring("any")` 是一个扩展函数，返回 `agtype`。
- `mag_catalog.agtype_array_to_agtype(agtype[])` 是一个扩展函数，返回 `agtype`。
- `mag_catalog.agtype_contained_by_top_level(agtype, agtype)` 是一个扩展函数，返回 `boolean`。
- `mag_catalog.agtype_contains_top_level(agtype, agtype)` 是一个扩展函数，返回 `boolean`。
- `mag_catalog.agtype_to_json(agtype)` 是一个扩展函数，返回 `json`。
- `mag_catalog.create_elabel(graph_name cstring, label_name cstring)` 是一个扩展函数，返回 `void`。
- `mag_catalog.create_vlabel(graph_name cstring, label_name cstring)` 是一个扩展函数，返回 `void`。
- `mag_catalog.graph_exists(graph_name name)` 是一个扩展函数，返回 `agtype`。
- `mag_catalog.load_edges_from_file(graph_name name, label_name name, file_path text, load_as_agtype bool default false)` 是一个扩展函数，返回 `void`。
- `mag_catalog.load_labels_from_file(graph_name name, label_name name, file_path text, id_field_exists bool default true, load_as_agtype bool default false)` 是一个扩展函数，返回 `void`。
- `mag_catalog.gin_agtype_ops` 是一个扩展定义的运算符类。

### 要求与注意事项

- 审查后的控制文件声明默认版本为 `1.6.0`。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为以及失败情况。
