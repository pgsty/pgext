## 用法

来源：

- [官方上游 README](https://github.com/accarniel/festival/blob/b4ae7e2a3eff43e1be65806dc84242df8be88e3b/README.md)
- [官方扩展控制文件 (festival.control)](https://github.com/accarniel/festival/blob/b4ae7e2a3eff43e1be65806dc84242df8be88e3b/festival.control)
- [官方扩展 SQL (festival--1.1.1.sql)](https://github.com/accarniel/festival/blob/b4ae7e2a3eff43e1be65806dc84242df8be88e3b/festival--1.1.1.sql)

`festival` — FESTIval 是一个框架，作为 PostgreSQL 扩展实现，用于进行空间索引结构的实验性评估。完整的 FESTIval 文档可在此处找到。使用它来进行相应的空间数据或地理空间工作流。上游将此功能描述为实验性。

### 核心工作流

```sql
CREATE EXTENSION festival;
```

在目标数据库中安装扩展，当可用时运行上游提供的最小示例，并在将其集成到应用程序 SQL 中之前验证安装版本和返回值。

### 重要对象

- `FT_ADelete(absolute_path text, p int4, geom geometry, statistic_options int4 default 1, location_statistics int4 default 1, file_statistics text default NULL)` 是一个扩展函数，返回 `int4`。
- `FT_ADelete(index_name text, index_path text, p int4, geom geometry, statistic_options int4 default 1, location_statistics int4 default 1, file_statistics text default NULL)` 是一个扩展函数，返回 `int4`。
- `FT_AInsert(absolute_path text, p int4, geom geometry, statistic_options int4 default 1, location_statistics int4 default 1, file_statistics text default NULL)` 是一个扩展函数，返回 `int4`。
- `FT_AInsert(index_name text, index_path text, p int4, geom geometry, statistic_options int4 default 1, location_statistics int4 default 1, file_statistics text default NULL)` 是一个扩展函数，返回 `int4`。
- `FT_ApplyAllModificationsForFAI(absolute_path text)` 是一个扩展函数。
- `FT_ApplyAllModificationsForFAI(index_name text, index_path text)` 是一个扩展函数。
- `FT_ApplyAllModificationsFromBuffer(absolute_path text)` 是一个扩展函数。
- `FT_ApplyAllModificationsFromBuffer(index_name text, index_path text)` 是一个扩展函数。
- `FT_AQuerySpatialIndex(absolute_path text, type_query int4, obj geometry, predicate int4, processing_option int4 default 1, statistic_options int4 default 1, location_statistics int4 default 1, file_statistics text default NULL)` 是一个扩展函数，返回 `SETOF`。
- `FT_AQuerySpatialIndex(index_name text, index_path text, type_query int4, obj geometry, predicate int4, processing_option int4 default 1, statistic_options int4 default 1, location_statistics int4 default 1, file_statistics text default NULL)` 是一个扩展函数，返回 `SETOF`。
- `FT_AUpdate(absolute_path text, old_p int4, old_geom geometry, new_p int4, new_geom geometry, statistic_options int4 default 1, location_statistics int4 default 1, file_statistics text default NULL)` 是一个扩展函数，返回 `int4`。
- `FT_AUpdate(index_name text, index_path text, old_p int4, old_geom geometry, new_p int4, new_geom geometry, statistic_options int4 default 1, location_statistics int4 default 1, file_statistics text default NULL)` 是一个扩展函数，返回 `int4`。
- `FT_CollectOrderOfReadWrite()` 是一个扩展函数。
- `FT_CreateEmptySpatialIndex(index_id int4, absolute_path text, src_id int4, bc_id int4, sc_id int4, buf_id int4)` 是一个扩展函数。

### 要求与注意事项

- 审查后的控制文件声明默认版本为 `1.1.1`。
- 控制文件标记该扩展为可重定位。
- 上游将项目的一部分或全部标记为实验性。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况，以匹配固定源。
