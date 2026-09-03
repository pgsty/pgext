## 用法

来源：

- [EDB Postgres Tuner 官方文档](https://www.enterprisedb.com/docs/pg_extensions/pg_tuner/)
- [官方配置指南](https://www.enterprisedb.com/docs/pg_extensions/pg_tuner/configuring/)
- [官方使用指南](https://www.enterprisedb.com/docs/pg_extensions/pg_tuner/using/)

`edb_pg_tuner` 观察主机资源与 workload 行为，再推荐或按需应用 PostgreSQL 配置变更。

### 启用

安装 EDB Postgres Tuner 1.3.2，预加载、重启并创建扩展：

```ini
shared_preload_libraries = 'edb_pg_tuner'
edb_pg_tuner.autotune = false
```

```sql
CREATE EXTENSION edb_pg_tuner;
```

在按照 workload 与内存预算审查建议前，应保持 `edb_pg_tuner.autotune` 禁用。

### 审查建议

`edb_pg_tuner_recommendations` 返回当前建议或 SQL command。PostgreSQL 14+ 还提供 query-level spill statistic。

```sql
SELECT * FROM edb_pg_tuner_recommendations();
SELECT edb_pg_tuner_recommendations('sql');

SELECT *
FROM edb_pg_tuner_query_stats()
WHERE sort_spill > 0 OR hash_spill > 0;
```

`edb_pg_tuner_global_stats()` 报告 buffer-level tracking total。

### 运维边界

自动调优可能修改需要重启的 GUC，也可能根据历史 spill 提高 `work_mem`。`work_mem` 按 plan node 与并发 query 分配，因此观察到的 spill 倍数不是集群内存保证。应限制 `edb_pg_tuner.work_mem_pool` 与 `edb_pg_tuner.max_wal_size_limit`，审查生成 SQL，并通过与手工 PostgreSQL 配置相同的变更控制流程发布。

