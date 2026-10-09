## 用法

来源：

- [README.md](https://github.com/knizhnik/hnsw/blob/929e7f5a128a151ae01836d50c889a28fc532201/README.md)
- [hnsw--0.1.0.sql](https://github.com/knizhnik/hnsw/blob/929e7f5a128a151ae01836d50c889a28fc532201/hnsw--0.1.0.sql)
- [Makefile](https://github.com/knizhnik/hnsw/blob/929e7f5a128a151ae01836d50c889a28fc532201/Makefile)
- [hnsw.control](https://github.com/knizhnik/hnsw/blob/929e7f5a128a151ae01836d50c889a28fc532201/hnsw.control)

`hnsw` 是较早的 PostgreSQL ANN 原型，使用 real 数组和按需构建的内存索引，与 pgvector 中同名的索引访问方法是不同实现。

### 核心用法

```sql
CREATE EXTENSION hnsw;
CREATE TABLE embeddings (id integer PRIMARY KEY, payload real[]);
CREATE INDEX ON embeddings USING hnsw(payload) WITH (maxelements=100000, dims=3, m=32);
SELECT id FROM embeddings ORDER BY payload <-> ARRAY[1,2,3]::real[] LIMIT 10;
```

### 运行边界

由超级用户创建扩展。SQL 定义 `l2_distance(real[], real[])`、`<->` 运算符、hnsw 访问方法和 real 数组运算符类。用 `maxelements` 限制索引容量，`dims` 必须匹配向量维度；`ef` 控制考虑的邻居数量。查询使用距离排序和行数限制。

上游源码最后更新于 2023 年，没有声明当前 PostgreSQL 大版本支持矩阵。未声明预加载要求，对象可迁移。内存设计需要明确规划容量并验证重启和恢复行为，不能套用 pgvector 的持久性或运维约束。避免在同一数据库中安装访问方法名称冲突的其他扩展。
