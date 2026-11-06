## 用法

来源：

- [README](https://github.com/cgkanchi/pg_lion/blob/v0.3.0/README.md)
- [参考手册](https://github.com/cgkanchi/pg_lion/blob/v0.3.0/docs/REFERENCE.md)
- [扩展控制文件](https://github.com/cgkanchi/pg_lion/blob/v0.3.0/pg_lion_vector.control)
- [扩展 SQL](https://github.com/cgkanchi/pg_lion/blob/v0.3.0/pg_lion_vector--0.3.0.sql)

`pg_lion_vector` 0.3.0 为 pgvector 的 `vector` 与 `halfvec` 值添加 `lionvec` 近邻索引访问方法。它使用父扩展 `pg_lion` 的动态库，依赖 pgvector 0.7 或更新版本。

### 核心流程

先创建父扩展与 `vector`。独立的 `lion` 索引可在按距离排序之前过滤候选行。

```sql
CREATE EXTENSION vector;
CREATE EXTENSION pg_lion;
CREATE EXTENSION pg_lion_vector;
CREATE TABLE lion_items (id integer, category integer, embedding vector(3));
INSERT INTO lion_items VALUES (1, 3, '[1,0,0]'), (2, 3, '[0,1,0]'), (3, 4, '[0,0,1]');
CREATE INDEX lion_embedding ON lion_items USING lionvec (embedding vector_cosine_ops);
CREATE INDEX lion_category ON lion_items USING lion (category);
ALTER FUNCTION cosine_distance(vector, vector) COST 100;
VACUUM ANALYZE lion_items;
SELECT id FROM lion_items
WHERE category = 3
ORDER BY embedding <=> '[1,0,0]'::vector
LIMIT 10;
```

上游建议将所选运算符类对应距离函数的开销设为 `COST 100`，便于规划器比较索引扫描与排序成本。`pg_dump` 不会保留该设置，更新 `vector` 时也可能将其重置。

### 运算符类与调优

| 类型 | L2（`<->`） | 负内积（`<#>`） | 余弦（`<=>`） |
| --- | --- | --- | --- |
| `vector` | `vector_l2_ops`（默认） | `vector_ip_ops` | `vector_cosine_ops` |
| `halfvec` | `halfvec_l2_ops`（默认） | `halfvec_ip_ops` | `halfvec_cosine_ops` |

`partitions` 控制索引构建时的聚类数量；`maintenance_work_mem` 必须容纳足够的样本。`pg_lion.vector_margin` 和 `pg_lion.vector_patience` 控制近似搜索。受支持的 lion 过滤条件产生较小候选集时，可以采用精确评分。`pg_lion.vector_postfilter` 控制 lion 索引无法回答的条件；修改前应阅读参考手册。

### 运行边界

需要 PostgreSQL 16 及以上版本，并由超级用户安装。只支持单个 `vector` 或 `halfvec` 列，最多 8128 维；通常采用近似搜索。存在过滤条件不意味着所有查询都会成为精确近邻搜索。频繁更新和删除可能需要定期 `REINDEX`。父扩展的自定义 WAL 索引仍需维持预加载，父扩展与伴随扩展文件应保持相同版本。
