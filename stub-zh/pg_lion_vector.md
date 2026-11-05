## 用法

来源：

- [上游来源 1](https://github.com/cgkanchi/pg_lion/blob/e8c3fea6d9c9fa1da8d501966f61703ba53d17fa/README.md)
- [上游来源 2](https://github.com/cgkanchi/pg_lion/blob/e8c3fea6d9c9fa1da8d501966f61703ba53d17fa/pg_lion_vector.control)

`pg_lion_vector` 0.3.0 通过 lionvec 索引访问方法，为 pgvector 向量提供带过滤的近邻搜索。

### 核心流程

此伴随扩展依赖 `pg_lion` 及 pgvector 0.7 或更新版本，支持 `vector` 与 `halfvec`，需选择对应的距离运算符类。独立的 `lion` 过滤索引可在近邻排序前缩小候选集合。

```sql
CREATE EXTENSION vector;
CREATE EXTENSION pg_lion;
CREATE EXTENSION pg_lion_vector;
CREATE TABLE lion_items (id integer, category integer, embedding vector(3));
CREATE INDEX ON lion_items USING lionvec (embedding vector_cosine_ops);
CREATE INDEX ON lion_items USING lion (category);
SELECT id FROM lion_items
WHERE category = 3
ORDER BY embedding <=> '[1,0,0]'::vector
LIMIT 10;
```

### 运行边界

需要 PostgreSQL 16 及以上版本，并由超级用户安装。父扩展和伴随扩展文件应保持相同版本。应结合负载核对上游索引选项及召回行为，不能将索引选择视为普遍的精确搜索或性能保证。父扩展的自定义 WAL 索引仍需维持预加载。
