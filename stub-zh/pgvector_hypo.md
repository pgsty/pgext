## 用法

来源：

- [Official README.md](https://github.com/varadfromeast/pgvector_hypo/blob/154b8b9312fb9c55a5d0843f3d82c5aa5fdc3224/README.md)
- [Official pgvector_hypo.control](https://github.com/varadfromeast/pgvector_hypo/blob/154b8b9312fb9c55a5d0843f3d82c5aa5fdc3224/pgvector_hypo.control)
- [Official pgvector_hypo--0.1.0.sql](https://github.com/varadfromeast/pgvector_hypo/blob/154b8b9312fb9c55a5d0843f3d82c5aa5fdc3224/sql/pgvector_hypo--0.1.0.sql)
- [Official compatibility.md](https://github.com/varadfromeast/pgvector_hypo/blob/154b8b9312fb9c55a5d0843f3d82c5aa5fdc3224/docs/compatibility.md)

`pgvector_hypo` 0.1.0 让 PostgreSQL 16 在不实际构建索引的情况下，估计使用 pgvector 0.8.x 假设 IVFFlat 索引的计划。它预测规划器选择与成本，不预测执行速度、召回率或构建时间。

### 比较计划

```sql
CREATE EXTENSION vector;
CREATE EXTENSION pgvector_hypo;
CREATE TABLE items (id bigint, embedding vector(3));
SELECT * FROM pgvector_hypo_create_index(
  'CREATE INDEX ON items USING ivfflat (embedding vector_l2_ops) WITH (lists = 100)');
EXPLAIN SELECT id FROM items ORDER BY embedding <-> '[0,0,0]'::vector LIMIT 5;
SELECT * FROM pgvector_hypo_list_indexes;
SELECT pgvector_hypo_reset();
```

### 会话状态与范围

`pgvector_hypo_create_index()` 接受 CREATE INDEX 语句并返回合成的 OID 与名称。`pgvector_hypo_drop_index()` 删除单项注册，`pgvector_hypo_reset_index()` 是重置函数的别名。状态只属于当前会话，断开连接后消失。支持固定维度的 `vector`、`halfvec` 与 `bit`、普通列及受支持的不可变表达式，并涵盖部分索引与兼容的分区表定义。仅覆盖文档列出的 IVFFlat 操作符类；本版不支持 HNSW。

### 运行边界

先安装匹配的 vector 扩展，再由管理员安装本扩展，无需预加载。假设索引只参与顶层普通 `EXPLAIN`；普通查询、`EXPLAIN ANALYZE` 与 `EXPLAIN EXECUTE` 均不能使用它。扩展不会创建持久索引目录记录或关系文件。上游仍记录了固定长度位向量边界上的两处规划选择不一致，且受构建环境影响，因此结果仅供实验性规划分析，不能作为生产性能保证。
