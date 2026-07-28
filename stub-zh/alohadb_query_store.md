## 用法

来源：

- [官方上游 README](https://github.com/vern-allworks-llc/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/contrib/README)
- [官方扩展控制文件 (alohadb_query_store.control)](https://github.com/vern-allworks-llc/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/contrib/alohadb_query_store/alohadb_query_store.control)
- [官方扩展 SQL (alohadb_query_store--1.0.sql)](https://github.com/vern-allworks-llc/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/contrib/alohadb_query_store/alohadb_query_store--1.0.sql)

`alohadb_query_store` — AlohaDB 查询存储和索引顾问。在收集或解释相应的 PostgreSQL 统计信息时使用它。上游将此功能描述为实验性功能。

### 核心工作流

```sql
CREATE EXTENSION alohadb_query_store;
```

在目标数据库中安装扩展，当可用时运行上游提供的最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `autovacuum_suggestions()` 是一个扩展函数，返回 `TABLE`。
- `index_advisor_recommend()` 是一个扩展函数，返回 `TABLE`。
- `index_advisor_unused_indexes()` 是一个扩展函数，返回 `TABLE`。
- `query_store_entries()` 是一个扩展函数，返回 `TABLE`。
- `query_store_reset()` 是一个扩展函数，返回 `void`。
- `query_store_stats()` 是一个扩展函数，返回 `TABLE`。

### 要求与注意事项

- 审查过的控制文件声明默认版本为 `1.0`。
- 控制文件标记该扩展为不可重定位。
- 上游将项目的一部分或全部标记为实验性。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况，以匹配固定源。
