## 用法

来源：

- [当前上游架构与快速入门](https://github.com/mayflower/pg_turboquant/blob/4d27dc6d1c5de7fd3dc4a37fbec89ece597edd63/README.md)
- [索引选项参考](https://github.com/mayflower/pg_turboquant/blob/4d27dc6d1c5de7fd3dc4a37fbec89ece597edd63/docs/reference/index-options.md)
- [兼容性参考](https://github.com/mayflower/pg_turboquant/blob/4d27dc6d1c5de7fd3dc4a37fbec89ece597edd63/docs/reference/compatibility.md)

`pg_turboquant` 为 pgvector 的 `vector` 与 `halfvec` 列添加 `turboquant` 近似最近邻索引访问方法。它的紧凑四位路径面向归一化余弦或内积检索，精确重排在 SQL 层完成。

```sql
CREATE EXTENSION vector;
CREATE EXTENSION pg_turboquant;
CREATE INDEX docs_embedding_tq_idx
ON docs USING turboquant (embedding tq_cosine_ops)
WITH (bits = 4, lists = 0, transform = 'hadamard', normalized = true);
```

已复核兼容性契约面向 PostgreSQL 16/17 与 pgvector `0.8.1`。L2 和未归一化扫描走兼容回退路径，而非忠实快速路径；精确重排仍在访问方法之外。格式变化可能要求 `REINDEX`；部署前应在代表性向量上测量召回率、延迟、WAL、维护、并发写入和索引大小。
