## 用法

来源：

- [官方上游 README](https://github.com/danolivo/pg_index_stats/blob/36be5041b5a5173a6076153c65bfc8437a41eaf2/README.md)
- [官方扩展控制文件 (pg_index_stats.control)](https://github.com/danolivo/pg_index_stats/blob/36be5041b5a5173a6076153c65bfc8437a41eaf2/pg_index_stats.control)
- [官方扩展 SQL (pg_index_stats--0.2.sql)](https://github.com/danolivo/pg_index_stats/blob/36be5041b5a5173a6076153c65bfc8437a41eaf2/pg_index_stats--0.2.sql)

`pg_index_stats` — 一个轻量级的 PostgreSQL 扩展，基于索引定义生成扩展统计信息。它引入了统计信息对相应索引的依赖关系。在管理或自动化上述数据库行为时使用它。请使用链接中的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE TABLE test(x integer, y integer);
CREATE INDEX ON test (x,y);
CREATE EXTENSION pg_index_stats;
SELECT pg_index_stats_rebuild();
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `pg_index_stats_build(idxname text, mode text DEFAULT 'mcv, ndistinct')` 是一个扩展函数，返回 `boolean`。
- `pg_index_stats_rebuild()` 是一个扩展函数，返回 `integer`。
- `pg_index_stats_remove()` 是一个扩展函数，返回 `integer`。

### 要求与注意事项

- 审查过的控制文件声明默认版本为 `0.2`。
- 控制文件将扩展标记为可重定位。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源进行比对。
