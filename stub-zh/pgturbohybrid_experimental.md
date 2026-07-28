## 用法

来源：

- [官方上游 README](https://github.com/mayflower/pgturbohybrid/blob/cd68155df20545b163dd610fb95e8aa7e62fc108/README.md)
- [官方扩展控制文件 (pgturbohybrid_experimental.control)](https://github.com/mayflower/pgturbohybrid/blob/cd68155df20545b163dd610fb95e8aa7e62fc108/pgturbohybrid_experimental.control)
- [官方扩展 SQL (pgturbohybrid_experimental--0.2.0.sql)](https://github.com/mayflower/pgturbohybrid/blob/cd68155df20545b163dd610fb95e8aa7e62fc108/sql/pgturbohybrid_experimental--0.2.0.sql)

`pgturbohybrid_experimental` — 本 README 帮助您了解 pgturbohybrid 的功能、混合搜索何时有用、如何安装、如何创建第一个索引以及如何检查快速路径是否正常工作。请根据相应的向量、模型或检索工作流使用它。上游将此功能描述为实验性。

### 核心工作流

```sql
CREATE EXTENSION pgturbohybrid_experimental;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `turbohybrid_experimental_compact_code_score(query_codes pg_catalog.int2[], doc_codes pg_catalog.int2[], experimental pg_catalog.bool DEFAULT false, force_kernel pg_catalog.text DEFAULT 'auto')` 是一个扩展函数，返回 `pg_catalog`。
- `turbohybrid_multivector(vector[])` 是一个扩展函数，返回 `turbohybrid_multivector`。
- `turbohybrid_multivector_context_count(turbohybrid_multivector)` 是一个扩展函数，返回 `pg_catalog`。
- `turbohybrid_multivector_context_maxsim(query turbohybrid_multivector, doc turbohybrid_multivector)` 是一个扩展函数，返回 `pg_catalog`。
- `turbohybrid_multivector_context_offsets(turbohybrid_multivector)` 是一个扩展函数，返回 `pg_catalog`。
- `turbohybrid_multivector_count(turbohybrid_multivector)` 是一个扩展函数，返回 `pg_catalog`。
- `turbohybrid_multivector_dims(turbohybrid_multivector)` 是一个扩展函数，返回 `pg_catalog`。
- `turbohybrid_multivector_distance(turbohybrid_multivector, turbohybrid_query)` 是一个扩展函数，返回 `pg_catalog`。
- `turbohybrid_multivector_field_ids(turbohybrid_multivector)` 是一个扩展函数，返回 `pg_catalog`。
- `turbohybrid_multivector_field_weighted_maxsim(query turbohybrid_multivector, doc turbohybrid_multivector, field_ids pg_catalog.int4[], weights pg_catalog.float4[])` 是一个扩展函数，返回 `pg_catalog`。
- `turbohybrid_multivector_from_contexts(raw_values pg_catalog.float4[], dim pg_catalog.int4, context_offsets pg_catalog.int4[])` 是一个扩展函数，返回 `turbohybrid_multivector`。
- `turbohybrid_multivector_from_contexts_and_fields(raw_values pg_catalog.float4[], dim pg_catalog.int4, context_offsets pg_catalog.int4[], field_ids pg_catalog.int4[])` 是一个扩展函数，返回 `turbohybrid_multivector`。
- `turbohybrid_multivector_from_float4(raw_values pg_catalog.float4[], dim pg_catalog.int4)` 是一个扩展函数，返回 `turbohybrid_multivector`。
- `turbohybrid_multivector_in(pg_catalog.cstring)` 是一个扩展函数，返回 `turbohybrid_multivector`。

### 要求与注意事项

- 审核的控制文件声明默认版本为 `0.2.0`。
- 先安装确认的扩展依赖项：`pgturbohybrid`。
- 控制文件标记该扩展为可重定位。
- 上游将项目的一部分或全部标记为实验性。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况，以匹配固定源。
