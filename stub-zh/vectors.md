## 用法

来源：

- [官方 0.4.0 版本 README](https://github.com/tensorchord/pgvecto.rs/blob/2b290b34e8ba69104ea2f800fa53328c6ed6c236/README.md)
- [官方 0.4.0 版本扩展 SQL](https://github.com/tensorchord/pgvecto.rs/blob/2b290b34e8ba69104ea2f800fa53328c6ed6c236/sql/install/vectors--0.4.0.sql)
- [从 PGVecto.rs 迁移到 VectorChord 的官方指南](https://docs.vectorchord.ai/vectorchord/admin/migration.html)

`vectors` 是 PGVecto.rs 0.4.0 的规范扩展名。这是一个基于 Rust/pgrx 的向量相似度扩展，提供稠密、半精度、稀疏与二进制向量类型，以及自己的 `vectors` 索引访问方法。上游目前建议迁移到 VectorChord；`vectors` 主要适合维护现有 PGVecto.rs 部署。

### 设置与检索

必须先把库加入 `shared_preload_libraries` 并重启 PostgreSQL，才能创建扩展。对象安装在 `vectors` 模式中。

```sql
CREATE EXTENSION vectors;

CREATE TABLE items (
    id bigserial PRIMARY KEY,
    embedding vectors.vector(3) NOT NULL
);

INSERT INTO items(embedding)
VALUES ('[1,2,3]'), ('[4,5,6]');

SELECT id
FROM items
ORDER BY embedding <-> '[3,2,1]'::vectors.vector
LIMIT 5;

CREATE INDEX items_embedding_idx
ON items USING vectors (embedding vectors.vector_l2_ops)
WITH (options = '[indexing.hnsw]');
```

对于稠密 `vector`，`<->` 是平方欧氏距离，`<#>` 是负点积，`<=>` 是余弦距离。应按查询操作符选择匹配的 `vector_l2_ops`、`vector_dot_ops` 或 `vector_cos_ops` 运算符类。其他类型包括 `vecf16`、`svector` 与 `bvector`，各有对应运算符类；二进制向量还支持面向汉明/杰卡德距离的操作。

检索 GUC 包括 `vectors.hnsw_ef_search`、`vectors.ivf_nprobe`、`vectors.search_mode` 以及量化重排/快速扫描设置。应使用代表性数据调节召回率与延迟，不能把默认值或发布的基准结果当作保证。

### 生命周期与迁移

PGVecto.rs 索引存储独立于 PostgreSQL 堆存储管理，上游概览警告索引 WAL 支持尚不完整。应针对精确的 0.4.0 构建演练崩溃恢复、复制、备份恢复、升级、`REINDEX` 与删除扩展。索引列必须具有固定维度。

官方 VectorChord 迁移路径会在不同模式安装 `vchord`/`vector`，记录并删除 PGVecto.rs 索引，通过文档规定的类型转换迁移列，再重建索引。类型转换会取得 `ACCESS EXCLUSIVE` 锁，因此会产生停机。`svector` 还需要转换下标，因为 PGVecto.rs 稀疏下标从零开始，而 pgvector 从一开始。在所有依赖列和索引完成迁移并验证前，应保留旧扩展。
