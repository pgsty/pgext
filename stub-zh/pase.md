## 用法

来源：

- [已复核 commit 的上游 README](https://github.com/alipay/PASE/blob/eec38006c17af8b00ab65d03132bc2ae1264e947/README.md)
- [0.0.1 版本 SQL 对象](https://github.com/alipay/PASE/blob/eec38006c17af8b00ab65d03132bc2ae1264e947/pase--0.0.1.sql)
- [HNSW 实现](https://github.com/alipay/PASE/tree/eec38006c17af8b00ab65d03132bc2ae1264e947/hnsw)

`pase` 是用于超高维向量的早期近似最近邻扩展。0.0.1 版本定义可变长 `pase` 查询类型、距离操作符，以及适用于 `float4[]` 和文本向量的 `pase_hnsw` 与 `pase_ivfflat` 索引访问方法。

```sql
CREATE EXTENSION pase;
CREATE TABLE item (id text PRIMARY KEY, vector float4[] NOT NULL);
CREATE INDEX item_hnsw ON item USING pase_hnsw (vector);
SELECT id, vector <?> pase(ARRAY[0.1, 0.2]::float4[]) AS distance
FROM item
ORDER BY distance
LIMIT 10;
```

近似索引用召回率换延迟；应在真实维度、数据分布、过滤、插入、删除、vacuum、重建和并发条件下，与精确搜索比较召回率。向量维度以及 L2/内积模式必须保持一致。

不要向不可信用户开放内置 `ivfflat_search` 和 `hnsw_search` 辅助函数。已复核 SQL 把查询向量、谓词、限制与表名插入动态 SQL，没有安全的标识符/值绑定，存在 SQL 注入风险。安装时还会修改规划器设置，部分重载明显未完成。应撤销辅助函数执行权，直接使用参数化操作符查询，并在完整审计和测试前把这个 0.0.1 代码视为仅供研究。
