## 用法

来源：

- [官方上游 README](https://github.com/vern-allworks-llc/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/contrib/README)
- [官方扩展控制文件 (alohadb_approx.control)](https://github.com/vern-allworks-llc/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/contrib/alohadb_approx/alohadb_approx.control)
- [官方扩展 SQL (alohadb_approx--1.0--1.1.sql)](https://github.com/vern-allworks-llc/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/contrib/alohadb_approx/alohadb_approx--1.0--1.1.sql)

`alohadb_approx` — 近似查询处理：HLL、Count-Min Sketch、Top-K。当 SQL 需要这些特殊函数或聚合时使用它。上游将此功能描述为实验性。

### 核心工作流

```sql
CREATE EXTENSION alohadb_approx;
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行该示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `approx_count_distinct_finalfn(internal)` 是一个扩展函数，返回 `int8`。
- `approx_count_distinct_transfn(internal, anyelement)` 是一个扩展函数，返回 `internal`。
- `approx_percentile_finalfn(internal)` 是一个扩展函数，返回 `float8`。
- `approx_percentile_transfn(internal, float8, float8, float8)` 是一个扩展函数，返回 `internal`。
- `bloom_add(bf bloom_filter, item text)` 是一个扩展函数，返回 `bloom_filter`。
- `bloom_agg_finalfn(internal)` 是一个扩展函数，返回 `bloom_filter`。
- `bloom_agg_transfn(internal, text, int, float8)` 是一个扩展函数，返回 `internal`。
- `bloom_contains(bf bloom_filter, item text)` 是一个扩展函数，返回 `boolean`。
- `bloom_create(expected_items int, fpr float8 DEFAULT 0.01)` 是一个扩展函数，返回 `bloom_filter`。
- `bloom_in(cstring)` 是一个扩展函数，返回 `bloom_filter`。
- `bloom_merge(bf1 bloom_filter, bf2 bloom_filter)` 是一个扩展函数，返回 `bloom_filter`。
- `bloom_out(bloom_filter)` 是一个扩展函数，返回 `cstring`。
- `bloom_stats(bf bloom_filter)` 是一个扩展函数，返回 `TABLE`。
- `cms_add(cms, text)` 是一个扩展函数，返回 `cms`。

### 要求与注意事项

- 审查的控制文件声明默认版本为 `1.1`。
- 控制文件标记该扩展为可重定位。
- 上游将项目的一部分或全部标记为实验性。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况，以匹配固定源。
