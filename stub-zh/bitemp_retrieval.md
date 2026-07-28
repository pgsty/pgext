## 用法

来源：

- [官方上游 README](https://github.com/aditditto/bitemp_retrieval/blob/8bc2fc4f9c7a84f92f0c1224ea557bd7f53a76b6/README.md)
- [官方扩展控制文件 (bitemp_retrieval.control)](https://github.com/aditditto/bitemp_retrieval/blob/8bc2fc4f9c7a84f92f0c1224ea557bd7f53a76b6/bitemp_retrieval.control)
- [官方扩展 SQL (bitemp_retrieval--0.0.1.sql)](https://github.com/aditditto/bitemp_retrieval/blob/8bc2fc4f9c7a84f92f0c1224ea557bd7f53a76b6/bitemp_retrieval--0.0.1.sql)

`bitemp_retrieval` — 用于在使用 pg_bitemporal 扩展创建的二时态表上进行检索功能的 Postgresql 扩展，由 Henrietta Dombrovskaya 开发。请使用上述链接的上游修订版本作为 API 边界，并在目标 PostgreSQL 构建中进行测试。

### 核心工作流

```sql
CREATE EXTENSION bitemp_retrieval;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `bitemp_contains_now(effective_range temporal_relationships.timeperiod, asserted_range temporal_relationships.timeperiod)` 是一个扩展函数，返回 `BOOLEAN`。
- `bitemp_contains_ts(effective_range temporal_relationships.timeperiod, asserted_range temporal_relationships.timeperiod, effective_ts timestamptz, asserted_ts timestamptz)` 是一个扩展函数，返回 `BOOLEAN`。
- `bitemporal_internal.ll_register_temporal_attribute_property(p_schema TEXT, p_table TEXT, p_attr_name TEXT, p_attr_property bitemporal_internal.temporal_attribute_property_enum)` 是一个扩展函数，返回 `INTEGER`。
- `get_interval_overlap(a temporal_relationships.timeperiod, b temporal_relationships.timeperiod)` 是一个扩展函数，返回 `temporal_relationships`。
- `get_sum(int, int)` 是一个扩展函数，返回 `int`。
- `interval_contains_now(interv temporal_relationships.timeperiod)` 是一个扩展函数，返回 `BOOLEAN`。
- `interval_contains_ts(interv temporal_relationships.timeperiod, ts timestamptz)` 是一个扩展函数，返回 `BOOLEAN`。
- `interval_join(a temporal_relationships.timeperiod, b temporal_relationships.timeperiod)` 是一个扩展函数，返回 `temporal_relationships`。
- `interval_joinable(a temporal_relationships.timeperiod, b temporal_relationships.timeperiod)` 是一个扩展函数，返回 `BOOLEAN`。
- `interval_len(interv temporal_relationships.timeperiod)` 是一个扩展函数，返回 `INTERVAL`。
- `intervals_contains_now(intervs temporal_relationships.timeperiod[])` 是一个扩展函数，返回 `BOOLEAN`。
- `intervals_contains_ts(intervs temporal_relationships.timeperiod[], ts timestamptz)` 是一个扩展函数，返回 `BOOLEAN`。
- `ita_now(p_schema TEXT, p_table TEXT, p_group_by TEXT[], p_aggr_funcs TEXT[], p_aggr_target TEXT[], p_aggr_fieldnames TEXT[])` 是一个扩展函数，返回 `SETOF`。
- `mwta_now(p_schema TEXT, p_table TEXT, p_group_by TEXT[], p_aggr_funcs TEXT[], p_aggr_target TEXT[], p_aggr_fieldnames TEXT[], p_window_size INTERVAL)` 是一个扩展函数，返回 `SETOF`。

### 要求与注意事项

- 控制文件声明默认版本为 `0.0.1`。
- 控制文件标记该扩展为不可重定位。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况。
