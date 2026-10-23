## 用法

来源：

- [官方资料](https://github.com/A-Jimider/pg_dann/blob/351737cd0d9561a290e1c06fa687e1caeea0ddb6/README.md)
- [官方资料](https://github.com/A-Jimider/pg_dann/blob/351737cd0d9561a290e1c06fa687e1caeea0ddb6/pg_dann.control)
- [官方资料](https://github.com/A-Jimider/pg_dann/blob/351737cd0d9561a290e1c06fa687e1caeea0ddb6/sql/pg_dann--0.1.0.sql)

`pg_dann` 0.1.0 是实验性的磁盘优先近似向量索引，采用 IVF 列表以及 SQ8 或 RaBitQ 量化，依赖 `vector`。本次核对的源码在 PostgreSQL 16 上进行过验证，本文不据此确认更广的兼容范围。

### 核心流程

先安装依赖，再以超级用户安装本扩展；无需预加载。

```sql
CREATE EXTENSION IF NOT EXISTS vector;
CREATE EXTENSION pg_dann;
CREATE TABLE items (id bigint, embedding vector(3));
INSERT INTO items VALUES (1, '[1,2,3]'), (2, '[3,2,1]');
CREATE INDEX items_dann ON items USING dann (embedding dann_l2_ops)
  WITH (nlist = 1, quantizer = 'sq8');
SELECT * FROM items ORDER BY embedding <-> '[1,2,3]' LIMIT 10;
```

实际使用时应以有代表性的数据训练并构建索引。`dann_l2_ops`、`dann_ip_ops` 和 `dann_cosine_ops` 分别对应 L2、内积和余弦距离操作符；查询操作符必须与所选操作符类匹配。

### 调优与检查

`dann.probes` 控制搜索的列表数。`dann.rerank_factor` 和 `dann.enable_rerank` 控制候选重排，`dann.max_latency_ms` 设置延迟预算。索引选项包括 `nlist` 与 `quantizer`，自动量化模式按维数选择方案。`dann_scan_stats()` 和 `dann_list_stats()` 提供扫描及列表信息。

重排使用候选向量的完整精度，但无法找回未进入候选集的近邻。应针对实际数据同时测量召回率与延迟。向量中的非有限数值会被拒绝。

### 维护边界

此版本通过元数据页串行处理插入，尚未实现并行扫描或构建。清理操作会删除无效条目，但不会回收空索引页，可能需要重建索引才能释放空间。应将其视为原型，并在依赖它之前验证恢复与更新行为。
