## 用法

来源：

- [官方架构与快速入门](https://github.com/mayflower/pg_turboquant/blob/4d27dc6d1c5de7fd3dc4a37fbec89ece597edd63/README.md)
- [索引选项参考](https://github.com/mayflower/pg_turboquant/blob/4d27dc6d1c5de7fd3dc4a37fbec89ece597edd63/docs/reference/index-options.md)
- [兼容性参考](https://github.com/mayflower/pg_turboquant/blob/4d27dc6d1c5de7fd3dc4a37fbec89ece597edd63/docs/reference/compatibility.md)

`pg_turboquant` 为 pgvector `vector` 和 `halfvec` 列添加 `turboquant` 近似最近邻索引访问方法。它的紧凑快速路径面向已归一化余弦或内积检索，精确重排则在 SQL 中显式完成。

### 核心流程

```sql
CREATE EXTENSION vector;
CREATE EXTENSION pg_turboquant;

CREATE INDEX docs_embedding_tq_idx
ON docs USING turboquant (embedding tq_cosine_ops)
WITH (
    bits = 4,
    lists = 0,
    transform = 'hadamard',
    normalized = true
);

SELECT *
FROM tq_rerank_candidates(
    'docs'::regclass,
    'id',
    'embedding',
    '[1,0,0,0]'::vector(4),
    'cosine',
    50,
    10
);
```

将 `lists = 0` 用于平坦近似扫描，或设为正数以启用 IVF 路由。文档中的生产快速路径使用 `bits = 4`、`transform = 'hadamard'`、归一化向量和自动 lanes；选择参数前，应在代表性向量上基准测试召回率、延迟、索引大小、构建时间、WAL 和并发写入。

### SQL 与可观测界面

- `tq_approx_candidates(...)` 返回近似候选项，`tq_rerank_candidates(...)` 实施文档化的 SQL 层精确重排流程。
- `tq_recommended_query_knobs(...)` 提供查询设置建议。
- `tq_index_metadata(...)`、`tq_index_heap_stats(...)` 和 `tq_last_scan_stats()` 暴露索引、堆以及当前后端最近一次扫描的诊断信息。
- `tq_maintain_index(...)` 执行扩展的轻量级 delta-tier 合并或压缩工作。

### 兼容性与维护

已复核的契约面向 PostgreSQL 16 和 17，以及 pgvector `0.8.1`。L2 和未归一化扫描使用解码向量兼容评分，而不是忠实的量化快速路径。当前格式是 `v2`；旧索引需要 `REINDEX` 或重新创建。精确重排仍位于访问方法之外，MVCC 可见性仍由 PostgreSQL 执行器处理，而实现使用通用 WAL。用于生产之前，应固定版本，并验证升级、崩溃恢复、副本、仅索引扫描、元数据过滤和维护建议。
