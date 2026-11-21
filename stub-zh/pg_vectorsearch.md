## 用法

来源：

- [README.md](https://github.com/timescale/pg_vectorsearch/blob/v0.2.0/README.md)
- [sql/pg_vectorsearch.sql](https://github.com/timescale/pg_vectorsearch/blob/v0.2.0/sql/pg_vectorsearch.sql)
- [docs/maintenance.md](https://github.com/timescale/pg_vectorsearch/blob/v0.2.0/docs/maintenance.md)
- [docs/tuning.md](https://github.com/timescale/pg_vectorsearch/blob/v0.2.0/docs/tuning.md)
- [CHANGELOG.md](https://github.com/timescale/pg_vectorsearch/blob/v0.2.0/CHANGELOG.md)
- [VERSION](https://github.com/timescale/pg_vectorsearch/blob/v0.2.0/VERSION)
- [meson.build](https://github.com/timescale/pg_vectorsearch/blob/v0.2.0/meson.build)
- [pg_vectorsearch.control.in](https://github.com/timescale/pg_vectorsearch/blob/v0.2.0/pg_vectorsearch.control.in)

`pg_vectorsearch` 0.2.0 是使用 PRISM 索引进行近似最近邻检索的 alpha 扩展，支持 PostgreSQL 18 和 19、原生 `vec32` / `vec16` 类型以及可选的 pgvector 互操作。当前面向测试，行为与维护要求仍在演进。

### 核心用法

```sql
CREATE EXTENSION pg_vectorsearch;
CREATE TABLE vectorsearch_items (
    id serial PRIMARY KEY,
    embedding vec32(3) STORAGE PLAIN
);
INSERT INTO vectorsearch_items (embedding)
SELECT ARRAY[i::real, (i % 7)::real, (i % 5)::real]::vec32
FROM generate_series(1, 1000) AS i;
CREATE INDEX vectorsearch_items_prism ON vectorsearch_items
USING prism (embedding vec32_l2_ops);
SET prism.nprobe = 40;
SELECT id FROM vectorsearch_items
ORDER BY embedding <-> '[3,1,2]' LIMIT 5;
```

### 运行边界

安装需要超级用户，无需预加载，扩展不可重定位。`prism` 访问方法提供 `vec32_l2_ops`、`vec32_ip_ops`、`vec32_cosine_ops` 及对应的 vec16 操作符类；`<->`、`<#>` 和 `<=>` 分别表示 L2、负内积和余弦距离。`prism.nprobe` 用于权衡召回与查询工作量；零表示自动调节，不代表精确召回。

应先装载代表性数据，再建立索引。L2 与余弦索引最多支持 8192 维：1968 维以内任意维度，之后要求是 2 的幂乘以不超过八的奇因子。内积索引对 vec32 仍限制为 1968 维，对 vec16 为 3584 维。不支持表达式索引与非日志索引。应将 `prism.leaf_refine_threshold` 保持为零。过滤后可能不足 LIMIT 行，受 `work_mem` 限制的候选缓冲区也可能截断大量或无上限的结果，应与精确查询基准比较。

可用 `prism_index_settings` 检查配置，并按维护指南处理持续更新。`prism_rebalance` 仅支持关闭 `centroid_fastscan` 的单层 RaBitQ 质心页索引，不支持默认多层格式。DELETE 留下墓碑，退役页仅能通过 REINDEX 回收。早期候选发布版不支持原位升级。

### 升级到 0.2.0

安装匹配文件，打开尚未加载旧库的新会话，更新 SQL 扩展并重建全部 PRISM 索引：

```sql
ALTER EXTENSION pg_vectorsearch UPDATE TO '0.2.0';
REINDEX INDEX vectorsearch_items_prism;
```

0x07 读取端拒绝旧 0x06 索引格式，两版动态库不能共存于同一会话。`vec32_norm` 改名为 `l2_norm`；`vec32_dims` 的 vec16 重载改名为 `vec16_dims`，`vector_dims` 提供兼容 pgvector 的接口。直接调用者须更新函数名，依赖视图仍保留函数 OID。

`posting_reference_format` 默认使用 float16。参考向量元素超出 ±65504 时，构建或分裂会报错，必要时选择 float32。vec16 类型现在拒绝超范围输入。已有过滤、缓冲区大小及维护限制继续适用。
