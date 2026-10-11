## 用法

来源：

- [README.md](https://api.pgxn.org/src/pg_recall_guard/pg_recall_guard-0.2.9/README.md)
- [pg_recall_guard.control](https://api.pgxn.org/src/pg_recall_guard/pg_recall_guard-0.2.9/pg_recall_guard.control)
- [CHANGELOG.md](https://api.pgxn.org/src/pg_recall_guard/pg_recall_guard-0.2.9/CHANGELOG.md)
- [pg_recall_guard--0.2.8--0.2.9.sql](https://api.pgxn.org/src/pg_recall_guard/pg_recall_guard-0.2.9/pg_recall_guard--0.2.8--0.2.9.sql)

`pg_recall_guard` 0.2.9 将近似最近邻结果与精确扫描比较，并保存认可的索引基线。它从 PostgreSQL 目录发现距离运算符，依赖 `tsm_system_rows`。

### 核心用法

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

### 运行边界

`recall_guard.vector_indexes` 列出可用索引，`recall_guard.approve()` 保存测量结果与搜索设置，`recall_guard.check()` 复用这些设置，`recall_guard.measure()` 执行单次比较。等距离邻居均计为命中，因此按旧版重复向量行为批准的基线可能需要重新批准。样本来自非 NULL 向量，`k` 范围为 1–1000，`sample_size` 范围为 1–10000。部分索引、表达式索引和临时索引会被拒绝。行级安全将基线和测量结果的访问限制为有权读取的表。每个样本执行精确计算并持有关系锁，应谨慎安排。样本不能替代真实查询流量，按名称绑定的基线也可能跨越索引替换。尽管元数据声明 PG13+，上游测试仅覆盖 PostgreSQL 18.6/19 beta。这是无需预加载的纯 SQL 扩展，Apache-2.0 从 0.2.3 开始适用。

### 0.2.9 的状态值变化

0.2.9 修改了 `recall_guard.check()` 返回的状态值：`degradado` 改为 `degraded`，`critico` 改为 `critical`，`NO SE PUDO MEDIR: <error>` 改为 `COULD NOT MEASURE: <error>`；`ok` 保持不变。按旧词匹配的监控过滤条件需要更新。升级脚本用英文重新定义函数和注释，不增删公开对象。
