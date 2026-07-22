## 用法

来源：

- [官方 README](https://github.com/benizar/pg_adm/blob/78314723ddfeb2128847018d5198d4156b527fb4/README.md)
- [扩展 control 文件](https://github.com/benizar/pg_adm/blob/78314723ddfeb2128847018d5198d4156b527fb4/pg_adm.control)
- [官方 SQL 工具集合](https://github.com/benizar/pg_adm/tree/78314723ddfeb2128847018d5198d4156b527fb4/sql)

`pg_adm` 0.0.1 将大量 PostgreSQL 管理查询封装为固定 `adm` 模式中的视图与函数。它适合交互式检查容量、膨胀、索引、缓冲区、依赖、权限和扩展对象；在允许日常用户执行会修改状态的辅助函数前，必须逐一审查。

### 核心流程

该扩展依赖 `pg_buffercache`；只要依赖可用，PostgreSQL 会在创建时自动安装它。

```sql
CREATE EXTENSION pg_adm;

SELECT * FROM adm.size_tables_biggest LIMIT 20;
SELECT * FROM adm.index_size_usage WHERE idx_scan = 0;
SELECT * FROM adm.table_bloat ORDER BY wastedbytes DESC LIMIT 20;
```

多项报告依赖优化器统计信息与累计统计信息，因此其结果只是估算或观测，不能证明某个对象可以安全删除。

### 重要对象

- 容量视图包括 `size_databases`、`size_schemas`、`size_tables` 与 `size_relations_biggest`。
- 膨胀和统计视图包括 `table_bloat`、`btree_bloat`、`no_stats` 与 `buffers_breakdown`。
- 索引诊断包括 `index_summary`、`index_duplicates`、`index_size_usage` 与 `unindexed_foreign_keys`。
- `dependency_tree(...)` 与 `dependency`：检查对象依赖。
- `extension_objects`：列出已安装扩展拥有的对象。
- `explain_count_estimate(text)`：根据执行计划估算查询行数。
- `clone_schema(...)` 与 `grant_on_tables(...)`：修改数据库状态和权限。

该扩展不可重定位，并始终使用 `adm`。应按对象分别授权：某些视图会暴露集群元数据，而模式克隆、批量授权等辅助函数会创建对象或改变权限。膨胀、重复索引和未使用索引报告都需要管理员判断，不能直接接入自动删除流程。
