## 用法

来源：

- [官方上游 README](https://github.com/mayflower/pgturbohybrid/blob/cd68155df20545b163dd610fb95e8aa7e62fc108/README.md)
- [官方扩展控制文件 (pgturbohybrid.control)](https://github.com/mayflower/pgturbohybrid/blob/cd68155df20545b163dd610fb95e8aa7e62fc108/pgturbohybrid.control)
- [官方扩展 SQL (pgturbohybrid--0.1.2--0.2.0.sql)](https://github.com/mayflower/pgturbohybrid/blob/cd68155df20545b163dd610fb95e8aa7e62fc108/sql/pgturbohybrid--0.1.2--0.2.0.sql)

`pgturbohybrid` — 本 README 帮助您了解 pgturbohybrid 的功能、混合搜索何时有用、如何安装、如何创建第一个索引以及如何检查快速路径是否正常工作。请根据相应的向量、模型或检索工作流使用它。上游将此功能描述为实验性。

### 核心工作流

```sql
CREATE EXTENSION pgturbohybrid;
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行该示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `turbohybrid_cosine_distance(vector, turbohybrid_query)` 是一个扩展函数，返回 `pg_catalog`。
- `turbohybrid_dense_query(vector_query vector, final_k pg_catalog.int4 DEFAULT NULL, dense_k pg_catalog.int4 DEFAULT NULL)` 是一个扩展函数，返回 `turbohybrid_query`。
- `turbohybrid_distance(vector, turbohybrid_query)` 是一个扩展函数，返回 `pg_catalog`。
- `turbohybrid_estimate_memory(pg_catalog.regclass)` 是一个扩展函数，返回 `pg_catalog`。
- `turbohybrid_graph_repair_dry_run(index pg_catalog.regclass, sample_nodes pg_catalog.int4 DEFAULT 1000, search_ef pg_catalog.int4 DEFAULT 400, candidate_limit pg_catalog.int4 DEFAULT 200)` 是一个扩展函数，返回 `pg_catalog`。
- `turbohybrid_handler(pg_catalog.internal)` 是一个扩展函数，返回 `pg_catalog`。
- `turbohybrid_hybrid_query(vector_query vector, text_query pg_catalog.tsquery, final_k pg_catalog.int4 DEFAULT NULL, dense_k pg_catalog.int4 DEFAULT NULL, bm25_k pg_catalog.int4 DEFAULT NULL)` 是一个扩展函数，返回 `turbohybrid_query`。
- `turbohybrid_index_stats(pg_catalog.regclass)` 是一个扩展函数，返回 `pg_catalog`。
- `turbohybrid_l2_distance(vector, turbohybrid_query)` 是一个扩展函数，返回 `pg_catalog`。
- `turbohybrid_last_build_stats()` 是一个扩展函数，返回 `pg_catalog`。
- `turbohybrid_last_scan_diagnosis()` 是一个扩展函数，返回 `pg_catalog`。
- `turbohybrid_last_scan_stats()` 是一个扩展函数，返回 `pg_catalog`。
- `turbohybrid_negative_inner_product(vector, turbohybrid_query)` 是一个扩展函数，返回 `pg_catalog`。
- `turbohybrid_prewarm(pg_catalog.regclass)` 是一个扩展函数，返回 `pg_catalog`。

### 要求与注意事项

- 审核的控制文件声明默认版本为 `0.2.0`。
- 先安装确认的扩展依赖项：`vector`。
- 控制文件标记该扩展为可重定位。
- 上游将项目的一部分或全部标记为实验性。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况，以匹配固定源。
