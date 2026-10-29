## 用法

来源：

- [README.md](https://github.com/kiransabne04/pg_stat_index_queries/blob/6a079ffc709515574271e55f4c71b528f60e9f2d/README.md)
- [INSTALL.md](https://github.com/kiransabne04/pg_stat_index_queries/blob/6a079ffc709515574271e55f4c71b528f60e9f2d/docs/INSTALL.md)
- [GUIDE.md](https://github.com/kiransabne04/pg_stat_index_queries/blob/6a079ffc709515574271e55f4c71b528f60e9f2d/docs/GUIDE.md)
- [pg_stat_index_queries.control](https://github.com/kiransabne04/pg_stat_index_queries/blob/6a079ffc709515574271e55f4c71b528f60e9f2d/pg_stat_index_queries.control)
- [pg_stat_index_queries--0.10.sql](https://github.com/kiransabne04/pg_stat_index_queries/blob/6a079ffc709515574271e55f4c71b528f60e9f2d/pg_stat_index_queries--0.10.sql)
- [ci.yml](https://github.com/kiransabne04/pg_stat_index_queries/blob/6a079ffc709515574271e55f4c71b528f60e9f2d/.github/workflows/ci.yml)

`pg_stat_index_queries` 记录哪些查询规划并执行了索引扫描，用于评估索引使用和删除风险。本次源码快照的 SQL 版本为 0.10，上游对原生 C 扩展测试 PostgreSQL 15–18。已观测负载能证明索引被使用，却不能证明未观测到的索引可以安全删除。

### 核心流程

将动态库加入已有预加载列表，并重启 PostgreSQL：

```ini
shared_preload_libraries = 'pg_stat_index_queries'
compute_query_id = on
```

```sql
CREATE EXTENSION pg_stat_index_queries;
SELECT indexname, index_role, queries, verdict
FROM pg_index_usage ORDER BY verdict;
SELECT queryid, calls_executed, last_seen
FROM pg_index_dependents WHERE indexname = 'orders_customer_idx';
SELECT * FROM pg_stat_index_queries_info;
```

除主库外，还应在所有副本上收集具有代表性的负载。原生钩子分别统计规划使用和执行使用；位图扫描会计入已执行，因为 PostgreSQL 在启动阶段就构建其扫描状态。

### 重要对象

- `pg_index_usage` 综合展示使用记录、索引用途、阻碍和提示；`pg_index_role` 说明约束、外键、副本标识等用途。
- `pg_index_column_usage` 与 `pg_index_sole_path` 帮助判断列使用情况，以及哪些已观测查询依赖某个唯一索引路径。
- `pg_index_query_stats` 可用 `pg_stat_statements` 补充索引与查询对的信息。语句时间和缓冲区统计会在查询使用的各个索引上重复展示，并非逐索引执行成本。
- `pg_index_table_stats`、`pg_index_column_stats` 与 `pg_index_hints` 使用目录和统计信息，仍受正常的采样及可见性限制。
- `pg_stat_index_queries_reset()` 清除统计，默认已从 PUBLIC 撤销执行权限。

### 配置与保留

`pg_stat_index_queries.max` 默认记录 10000 个查询与索引对，修改需要重启。`pg_stat_index_queries.track` 接受 top、all 或 none。`pg_stat_index_queries.save` 默认开启，`pg_stat_index_queries.save_interval` 默认 300 秒，`pg_stat_index_queries.flush_interval` 默认 1 秒。崩溃可能丢失最近一次保存后的计数，其他会话要等本地计数刷新后才能看到。

通过 `pg_stat_index_queries.sample_rate` 采样后，计数成为估计值，并可能遗漏罕见使用者。淘汰条目及此前重置也会限制结论。查询统计独立重置，因此显示的百分比可能超过 100。删除索引前必须审查约束、月末任务、副本和替代计划，扩展不会自动给出安全的 DROP 决策。

### 托管 PostgreSQL 替代方式

RDS/Aurora 通常不允许安装自定义 C 动态库。独立的 `pg_index_deps` SQL 版本读取已捕获计划，有自己的安装方式和覆盖边界，不能提供原生版本的已执行计划计数。
