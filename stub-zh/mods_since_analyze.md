## 用法

来源：

- [官方上游 README](https://api.pgxn.org/src/mods_since_analyze/mods_since_analyze-1.0.0/README.md)
- [官方扩展 SQL (mods_since_analyze.sql)](https://api.pgxn.org/src/mods_since_analyze/mods_since_analyze-1.0.0/mods_since_analyze.sql)

`mods_since_analyze` — mods_since_analyze 是一个 PostgreSQL 扩展，通过函数 pg_stat_get_mod_since_analyze() 暴露了自上次分析以来元组数量的估计值。在收集或解释相应的 PostgreSQL 统计信息时使用它。使用上方链接的固定上游版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

已审核的分发包使用了过时的 SQL 或非控制文件安装布局，因此不会建立现代独立的 `CREATE EXTENSION` 和升级工作流。遵循固定上游的安装机制，并在隔离数据库中验证安装对象。

### 重要对象

- `pg_stat_get_mod_since_analyze(oid)` 是一个扩展函数，返回 `bigint`。

### 要求与注意事项

- 目录记录了 `1.0.0` 的版本。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况，与固定源代码进行比对。
