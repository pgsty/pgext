## 用法

来源：

- [PGXN 0.2.0 README](https://pgxn.org/dist/pg_recall_guard/0.2.0/README.html)
- [pg_recall_guard 控制文件](https://api.pgxn.org/src/pg_recall_guard/pg_recall_guard-0.2.0/pg_recall_guard.control)
- [pg_recall_guard 0.2.0 SQL 定义](https://api.pgxn.org/src/pg_recall_guard/pg_recall_guard-0.2.0/pg_recall_guard--0.2.0.sql)
- [PostgreSQL 许可证](https://api.pgxn.org/src/pg_recall_guard/pg_recall_guard-0.2.0/LICENSE)

`pg_recall_guard` 用精确结果衡量近似最近邻索引的召回率，并记录运维人员认可的数值。它从 PostgreSQL 系统目录发现可排序的距离操作符，而不是写死某一种向量扩展。

### 核心流程

```sql
CREATE EXTENSION pg_recall_guard CASCADE;

SELECT index_name, table_name, access_method, operator
FROM recall_guard.vector_indexes;

SELECT recall_guard.approve(
    'items_hnsw',
    p_k => 10,
    p_sample_size => 30
);

SELECT * FROM recall_guard.check();
```

`CASCADE` 会在可用时安装所需的 contrib 扩展 `tsm_system_rows`。`recall_guard.approve(...)` 测量所选索引并保存认可基线。`recall_guard.check()` 重新测量全部基线，返回基线值、当前召回率、漂移与判定。

### 对象与测量

- `recall_guard.vector_indexes` 列出通过 `pg_amop` 暴露排序操作符的索引。
- `recall_guard.measure(...)` 抽取查询向量，将索引最近邻与精确顺序扫描结果比较。
- `recall_guard.baselines` 保存认可的设置与召回率。
- `recall_guard.measurements` 保存后续观测值。
- `recall_guard.evaluate_query(...)` 是底层比较辅助函数。

实现会验证近似侧确实使用目标索引，并确认真值侧没有使用它。每个样本行自身的零距离匹配会被排除，避免这种必然命中抬高召回率。

### 成本与边界

每个样本查询都会执行精确扫描，因此大表上的测量可能很昂贵。应从较小的 `p_sample_size` 开始，把检查安排在低峰时段，并观察 I/O 与 CPU。样本来自被索引表，而非生产查询流量；它只是代理，不能替代工作负载跟踪。

PGXN 元数据声明支持 PostgreSQL 13 及以上，上游测试则覆盖 18.6 与 19 beta，PostgreSQL 13 至 17 尚未测试。0.2.0 也没有经过生产工作负载的现场验证，已发布实验只覆盖有限的向量维度与索引设置。阈值应按工作负载设定，只能在独立质量测试后认可基线。
