## 用法

来源：

- [官方上游 README](https://github.com/li3ds/pg-li3ds/blob/60a45f4e291aa2f14ef702bcb8d6ffc5811d0576/README.rst)
- [官方扩展控制文件 (li3ds.control)](https://github.com/li3ds/pg-li3ds/blob/60a45f4e291aa2f14ef702bcb8d6ffc5811d0576/extension/li3ds.control)
- [官方扩展 SQL (li3ds--1.0.0.sql)](https://github.com/li3ds/pg-li3ds/blob/60a45f4e291aa2f14ef702bcb8d6ffc5811d0576/extension/li3ds--1.0.0.sql)

`li3ds` — PostgreSQL 扩展，用于管理 3D 传感器数据。适用于相应的空间数据或地理空间工作流。在安装此扩展之前，必须先安装并验证其扩展依赖项。

### 核心工作流

```sql
CREATE EXTENSION li3ds;
```

在目标数据库中安装扩展，当可用时运行上游最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `check_datasource_uri(uri text)` 是一个扩展函数，返回 `boolean`。
- `check_pcpatch_column(schema_table_column varchar)` 是一个扩展函数，返回 `boolean`。
- `check_timezone_name(timezone varchar)` 是一个扩展函数，返回 `boolean`。
- `check_transfo_args(parameters jsonb, transfo_type_id int)` 是一个扩展函数，返回 `boolean`。
- `check_transfotree_istree(transfo_trees integer[])` 是一个扩展函数，返回 `boolean`。
- `dijkstra(config integer, source integer, target integer, stoptosensor varchar default '')` 是一个扩展函数，返回 `integer[]`。
- `foreign_key_array(arr integer[], foreign_table regclass)` 是一个扩展函数，返回 `boolean`。
- `isconnected(transfos integer[], doubletransfo boolean default False)` 是一个扩展函数，返回 `boolean`。
- `postgres_version()` 是一个扩展函数，返回 `text`。
- `transform(box4d libox4d, config integer, source integer, target integer, ttime float8 default 0.0)` 是一个扩展函数，返回 `libox4d`。
- `transform(box4d libox4d, config integer, source integer, target integer, ttime text)` 是一个扩展函数，返回 `libox4d`。
- `transform(box4d libox4d, func_name text, func_sign text[], params text)` 是一个扩展函数，返回 `libox4d`。
- `transform(box4d libox4d, transfo integer, ttime float8 default 0.0)` 是一个扩展函数，返回 `libox4d`。
- `transform(box4d libox4d, transfo integer, ttime text)` 是一个扩展函数，返回 `libox4d`。

### 要求与注意事项

- 审核的控制文件声明默认版本为 `1.0.0`。
- 先安装并验证确认的扩展依赖项：`postgis`, `plpython2u`, `pointcloud`, `pointcloud_postgis`。
- 控制文件将此扩展标记为不可重定位。
- 控制文件不要求超级用户安装。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况。
