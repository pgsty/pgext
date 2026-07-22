## 用法

来源：

- [AlloyDB 列式引擎概览](https://docs.cloud.google.com/alloydb/docs/columnar-engine/about?hl=en)
- [配置列式引擎](https://docs.cloud.google.com/alloydb/docs/columnar-engine/configure)
- [手动管理列式内容](https://docs.cloud.google.com/alloydb/docs/columnar-engine/manage-content-manually)
- [AlloyDB 支持的扩展](https://docs.cloud.google.com/alloydb/docs/reference/extensions)

`google_columnar_engine` 是 AlloyDB 扩展，提供 Google Cloud 的内存列存储和向量化规划/执行能力，用于分析型扫描、连接与聚合。它只能在 AlloyDB 内运行，没有独立的社区软件包或上游版本。

### 启用

把实例数据库标志 `google_columnar_engine.enabled=on`。该操作会重启实例。实例恢复后，`alloydbsuperuser` 成员可以启用 SQL 扩展：

```sql
CREATE EXTENSION IF NOT EXISTS google_columnar_engine;

SELECT google_columnar_engine_add(
    relation => 'sales.orders',
    columns  => 'customer_id,created_at,total'
);

SELECT * FROM g_columnar_relations;

EXPLAIN (ANALYZE, COLUMNAR_ENGINE)
SELECT customer_id, sum(total)
FROM sales.orders
GROUP BY customer_id;
```

启用引擎时默认开启自动列式化，它可以根据已观察到的工作负载选择有用列。`google_columnar_engine_add` 显式管理内容，`google_columnar_engine_drop` 将内容移出列存储，`google_columnar_engine_refresh('sales.orders')` 重建失效的列块。

### 持久性与内存

手动添加的关系不会在实例重启后自动恢复。需要持久选择时，应使用 `google_columnar_engine.relations` 数据库标志；反过来，如果只在 SQL 中删除由该标志管理的内容，却不更新标志，下次重启时内容还会重新加入。引擎默认预留实例内存的 30%。修改其内存大小标志同样会重启实例，因此应规划存储大小与标志变更窗口。

### 查询与表边界

只有当一次扫描所需的全部列都已装入，并且使用受支持的内置类型时，该扫描才能使用列式引擎。叶子分区可逐个装入，但不支持非叶分区表和外部表。服务商指南指出，即使列已装入，很小的表（大约少于 5,000 行）和索引查找也可能继续使用行存储。

更新会使列块失效。应在 `g_columnar_relations` 中比较 `invalid_block_count` 与 `total_block_count`，并刷新大量变更的表；频繁更新的 OLTP 表可能并不适合。使用 `EXPLAIN (COLUMNAR_ENGINE)` 确认真正的执行路径，不要从扩展已安装推断查询已经加速。服务商标志、支持类型和资格规则会随 AlloyDB 引擎版本变化，应以当前实例文档为准。
