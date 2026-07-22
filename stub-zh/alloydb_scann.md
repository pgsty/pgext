## 用法

来源：

- [创建 ScaNN 索引](https://cloud.google.com/alloydb/docs/ai/create-scann-index)
- [AlloyDB ScaNN 索引参考](https://cloud.google.com/alloydb/docs/reference/ai/scann-index-reference)
- [AlloyDB 支持的扩展](https://cloud.google.com/alloydb/docs/reference/extensions)

`alloydb_scann` 是 Google AlloyDB 为 `vector` 列提供的托管 ScaNN 访问方法，用于近似最近邻搜索。当向量查询延迟、索引构建时间或内存占用比精确最近邻保证更重要时，可以使用它。它只在 AlloyDB 与 AlloyDB Omni 中提供，并不是可移植的社区 PostgreSQL 扩展。

### 核心流程

安装托管扩展；必要时，`CASCADE` 会自动安装其依赖的 `vector` 扩展：

```sql
CREATE EXTENSION IF NOT EXISTS alloydb_scann CASCADE;
```

创建自动调优的余弦距离索引，并执行使用相同距离度量的最近邻查询：

```sql
CREATE INDEX items_embedding_scann
ON items USING scann (embedding cosine);

SELECT id, embedding <=> $1::vector AS distance
FROM items
ORDER BY embedding <=> $1::vector
LIMIT 20;
```

索引中的距离函数与查询操作符必须表示同一种度量。ScaNN 支持 `l2`、`dot_product` 与 `cosine`。简单的 `USING scann` 定义默认启用自动调优和自动维护。

### 索引模式与选项

- `mode='AUTO'` 让 AlloyDB 选择并维护树结构。`optimization='SEARCH_OPTIMIZED'` 偏向召回率与延迟；`BALANCED` 降低构建成本。
- `auto_maintenance='ON'` 是自动模式默认值，会随着数据变化调整索引。
- `mode='MANUAL'` 暴露 `num_leaves`、`quantizer` 与 `max_num_levels` 等设置；只有在代表性查询上测量召回率与延迟后再使用。
- 量化器包括 `SQ8` 与 `FLAT`；其他量化器和更深层树可能属于预览特性，并要求单独的实例参数。

### 依赖与注意事项

源表必须包含已经存储的 `vector` 嵌入。空表、近空表与分区表存在特殊创建约束；少于 10,000 行的表若使用自动调优索引，需要延迟创建。近似结果必须由应用侧测试召回率，索引构建与维护也会消耗实例资源。

版本可用性、支持的 PostgreSQL 大版本、预览参数、权限、升级与维护行为均由 AlloyDB 服务控制。应以目标实例的厂商文档为准，不要为它套用社区扩展版本。
