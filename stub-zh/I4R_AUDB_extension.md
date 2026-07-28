## 用法

来源：

- [官方上游 README](https://github.com/asxvi/audb/blob/6018074c6e7e7e8416ccc7a04f1de9365d2f44dc/README.md)
- [官方扩展控制文件 (I4R_AUDB_extension.control)](https://github.com/asxvi/audb/blob/6018074c6e7e7e8416ccc7a04f1de9365d2f44dc/c_extension/i4r_audb_extension/I4R_AUDB_extension.control)
- [官方扩展 SQL (i4r_audb_extension--1.1.sql)](https://github.com/asxvi/audb/blob/6018074c6e7e7e8416ccc7a04f1de9365d2f44dc/c_extension/i4r_audb_extension/i4r_audb_extension--1.1.sql)

`I4R_AUDB_extension` — AUDB 操作对于 PostgreSQL int4range 值。当 SQL 需要这些特殊函数或聚合时使用它。请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION "I4R_AUDB_extension";
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行它，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `agg_avg_range_finalfunc(internal)` 是一个扩展函数，返回 `int4range`。
- `agg_avg_range_transfunc(state internal, range int4range, mult int4range)` 是一个扩展函数，返回 `internal`。
- `agg_count_transfunc(state int4range, input int4range)` 是一个扩展函数，返回 `int4range`。
- `agg_max_range_transfunc(state int4range, input int4range)` 是一个扩展函数，返回 `int4range`。
- `agg_max_set_transfunc(state int4range[], input int4range[])` 是一个扩展函数，返回 `int4range[]`。
- `agg_min_max_set_finalfunc(int4range[])` 是一个扩展函数，返回 `int4range[]`。
- `agg_min_range_transfunc(state int4range, input int4range)` 是一个扩展函数，返回 `int4range`。
- `agg_min_set_transfunc(state int4range[], input int4range[])` 是一个扩展函数，返回 `int4range[]`。
- `agg_sum_range_transfunc(int4range, int4range)` 是一个扩展函数，返回 `int4range`。
- `agg_sum_set_finalfunc(internal)` 是一个扩展函数，返回 `int4range[]`。
- `agg_sum_set_finalfunc_metrics(internal)` 是一个扩展函数，返回 `sum_set_metrics`。
- `agg_sum_set_transfunc(internal, set int4range[], resizeTrigger integer, reduceToSize integer)` 是一个扩展函数，返回 `internal`。
- `agg_sum_set_transfunc_metrics(internal, int4range[], integer, integer, bool)` 是一个扩展函数，返回 `internal`。
- `array_length(set int4range[])` 是一个扩展函数，返回 `int4`。

### 要求与注意事项

- 审查后的控制文件声明默认版本为 `1.1`。
- 控制文件将扩展标记为可重定位。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况。
