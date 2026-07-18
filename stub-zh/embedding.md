## 用法

来源：

- [上游 README 与弃用声明](https://github.com/neondatabase/pg_embedding/blob/5d48508aeaeb86b9e9c630468ada7a124b905795/README.md)
- [扩展 control 文件](https://github.com/neondatabase/pg_embedding/blob/5d48508aeaeb86b9e9c630468ada7a124b905795/embedding.control)
- [0.3.6 版 SQL](https://github.com/neondatabase/pg_embedding/blob/5d48508aeaeb86b9e9c630468ada7a124b905795/embedding--0.3.6.sql)

`embedding` `0.3.6` 版以 `real[]` 保存向量，并添加用于近似最近邻搜索的 HNSW 访问方法。运算符包括欧氏距离 `<->`、余弦距离 `<=>` 和曼哈顿距离 `<~>`。

### 示例

```sql
CREATE EXTENSION embedding;
CREATE TABLE documents (id integer PRIMARY KEY, embedding real[]);
INSERT INTO documents VALUES (1, ARRAY[1, 2, 3]), (2, ARRAY[4, 5, 6]);
SELECT id FROM documents ORDER BY embedding <-> ARRAY[3, 3, 3] LIMIT 1;
CREATE INDEX ON documents USING hnsw (embedding)
  WITH (dims=3, m=3, efconstruction=5, efsearch=5);
```

上游仓库已归档，并明确声明 `pg_embedding` 已弃用，强烈建议迁移到 `pgvector`。应保持向量维度一致，使用有代表性的数据验证 HNSW 召回率并调节构建与查询参数。它只适合在迁移规划期间维持现有部署，不应作为新系统的选择。
