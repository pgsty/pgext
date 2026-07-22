## 用法

来源：

- [官方 README](https://github.com/akorotkov/vgram/blob/babae3775c9c21bcf66f55c434642af3eeb0ce0d/README.md)
- [官方扩展 SQL](https://github.com/akorotkov/vgram/blob/babae3775c9c21bcf66f55c434642af3eeb0ce0d/vgram--1.0.sql)
- [官方 GIN 实现](https://github.com/akorotkov/vgram/blob/babae3775c9c21bcf66f55c434642af3eeb0ce0d/vgram_gin.c)

`vgram` 使用由可变长度 q-gram 构建的 GIN 索引加速 `LIKE` 和 `ILIKE` 子串检索。与固定三元组不同，它从目标语料学习常见短 gram，并索引更长或更少见的 gram，因此可能针对特定数据分布获得更好的选择性。

### 构建语料专用索引

版本 `1.0` 要求 PostgreSQL 13 或更高版本。在 psql 中收集常见 gram、保存结果，并把它传给 `vgram_gin_ops` 选项：

```sql
CREATE EXTENSION vgram;

SELECT qgram_stat(title, 2, 4, 0.05) AS vgrams
FROM documents
\gset

CREATE INDEX documents_title_vgram_idx
ON documents USING gin (
  title vgram_gin_ops (minq=2, maxq=4, vgrams=:'vgrams')
);

SELECT *
FROM documents
WHERE title ILIKE '%indexing%';
```

`minq` 和 `maxq` 限制提取 gram 的长度。传给 `qgram_stat()` 的阈值把语料频率达到或超过该值的 gram 标记为过于常见，不纳入索引。可用 `get_vgrams()` 检查某个候选值和已学习 gram 数组的提取结果。

### 对象与高级统计

- `qgram_stat(text, int, int, float8)` 聚合常见 q-gram。
- `get_vgrams(text, int, int, text[])` 展示从文本中选出的 gram。
- `vgram_gin_ops` 为普通 `text` 的 `LIKE` 与 `ILIKE` 建立索引。
- `vgram_text` 与文本兼容，但会安装自定义频率统计与选择性估算。
- `vgram_gin_ops2` 是 `vgram_text` 的 GIN 运算符类。

### 维护边界

学习得到的选项数组会嵌入索引定义，而且可能非常大。它只反映构建时样本；语言或分布发生重大变化后，需要重新收集统计、比较召回与计划，并通常执行 `REINDEX` 或创建替代索引。GIN 候选仍会复查，但过于常见或太短的模式仍可能退化成大范围扫描。应评估构建耗时、索引大小、写放大、待处理列表行为、VACUUM、排序规则、多字节文本、大小写折叠和规划器估算。还要保留可复现的训练查询与参数，确保副本和重建环境使用相同的运算符类选项。
