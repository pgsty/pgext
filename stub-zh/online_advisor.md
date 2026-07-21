## 用法

来源：

- [官方1.0版本README](https://github.com/knizhnik/online_advisor/blob/1.0/README.md)
- [扩展控制文件](https://github.com/knizhnik/online_advisor/blob/1.0/online_advisor.control)
- [1.0版本SQL对象](https://github.com/knizhnik/online_advisor/blob/1.0/online_advisor--1.0.sql)
- [示例预加载配置](https://github.com/knizhnik/online_advisor/blob/1.0/online_advisor.conf)

`online_advisor` 观察PostgreSQL执行计划和工作负载时间，然后推荐索引、扩展统计信息或准备语句。它仅报告候选方案；从不自动创建索引或统计对象。

### 核心流程

预加载库并重启PostgreSQL：

```conf
shared_preload_libraries = 'online_advisor'
```

在每个需要观察工作负载的数据库中创建并激活扩展：

```sql
CREATE EXTENSION online_advisor;

-- Calling an extension function activates collection in this database.
SELECT get_executor_stats();
```

运行代表性的工作负载后，检查推荐方案：

```sql
SELECT * FROM proposed_indexes;
SELECT * FROM proposed_statistics;
SELECT * FROM get_executor_stats();

-- Keep separate index candidates instead of combining compatible clauses.
SELECT * FROM propose_indexes(combine => false);
```

在应用生成的 `create_index` 或 `create_statistics` 语句之前，请审查每个建议。在创建索引或统计对象之后运行 `ANALYZE`，以便规划器可以使用当前统计数据。

### 对象和设置

- `proposed_indexes`: `propose_indexes(combine, reset)` 的视图，并带有过滤体积、调用次数、耗时以及候选的 `CREATE INDEX` 语句。
- `proposed_statistics`: `propose_statistics(combine, reset)` 的视图，并带有误估数、调用次数、耗时以及候选的 `CREATE STATISTICS` 语句。
- `get_executor_stats(reset)`: 返回聚合规划和执行时间、查询计数及规划开销比率。
- `online_advisor.filtered_threshold`: 考虑为索引建议的过滤行数最小值，默认值为 `1000`。
- `online_advisor.misestimation_threshold`: 实际到估计行数比值，用于统计；默认值为 `10`。
- `online_advisor.min_rows`: 误估分析中返回的最小行数，默认值为 `1000`。
- `online_advisor.max_index_proposals` 和 `online_advisor.max_stat_proposals`: 建议容量；在激活扩展之前设置它们。
- `online_advisor.do_instrumentation`, `online_advisor.log_duration`, 和 `online_advisor.prepare_threshold`: 控制收集和准备语句建议。

### 注意事项

- 仪器化会增加工作负载开销，请在目标系统上进行测量，并在不需要时禁用数据收集。
- 索引启发式算法不考虑复合索引、连接索引或仅用于避免排序的索引的操作符顺序问题。
- 扩展不会估计建议索引的好处。在构建昂贵的索引之前，使用计划审查或假设性索引工具进行评估。
- 建议是数据库本地的，并依赖于激活或重置以来观察到的工作负载。
