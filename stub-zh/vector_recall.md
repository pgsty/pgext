## 用法

来源：

- [vector_recall.control](https://github.com/AlibabaIncubator/gpdb-faiss-vector/blob/f247fd9653361316c2e9fefed4b767d7811c75d5/vector_recall.control)
- [README.md](https://github.com/AlibabaIncubator/gpdb-faiss-vector/blob/f247fd9653361316c2e9fefed4b767d7811c75d5/README.md)
- [vector_recall--0.1.0.sql](https://github.com/AlibabaIncubator/gpdb-faiss-vector/blob/f247fd9653361316c2e9fefed4b767d7811c75d5/vector_recall--0.1.0.sql)
- [makefile](https://github.com/AlibabaIncubator/gpdb-faiss-vector/blob/f247fd9653361316c2e9fefed4b767d7811c75d5/makefile)

`vector_recall` 将 Faiss 集成到 Greenplum，以 BYTEA 保存序列化索引，并通过 SQL 创建、训练、添加向量及执行 top-k 或范围查询。这里收录的是 Greenplum 扩展，不据此声明原生 PostgreSQL 兼容性。

### 核心用法

```sql
CREATE EXTENSION vector_recall;
SELECT faiss_index_create(3, 'Flat', 1);
```

### 运行边界

共享库与 Faiss 运行时必须匹配 Greenplum 环境。安装需要特权，控制文件没有预加载要求。训练与插入返回新的序列化索引，调用方需保存返回值。查询可以按调用方提供的键缓存反序列化索引，该键应随索引修订变化。扩展还提供缓存重置、清理、用量查询与 top-k 合并函数。应限制服务器资源并使用可信的序列化索引输入；上游基准不能代表实际负载保证。
