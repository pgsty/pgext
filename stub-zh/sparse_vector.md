## 用法

来源：

- [Pinned official README](https://github.com/MarkAntipin/pg_sparse_vector/blob/e03511c809bd0e083ff977f4ce182fb90d37a29f/README.md)
- [Pinned extension SQL](https://github.com/MarkAntipin/pg_sparse_vector/blob/e03511c809bd0e083ff977f4ce182fb90d37a29f/sparse_vector--0.1.sql)

`sparse_vector` 扩展在上游项目中以 `pg_sparse_vector` 打包，只存储 `float4[]` 向量中的非零位置，并提供点积与余弦相似度计算。它是一种紧凑的计算类型；不包含最近邻索引或向量距离运算符。

### 核心流程

由于未实现文本输入，必须从数组构造值：

```sql
CREATE EXTENSION sparse_vector;

CREATE TABLE sparse_data (
  id bigint GENERATED ALWAYS AS IDENTITY,
  embedding sparse_vector
);

INSERT INTO sparse_data (embedding)
VALUES (sparse_vector(ARRAY[1.1, 0, 5, 0, 3]::float4[]));

SELECT id,
       cosine_similarity(
         sparse_vector(ARRAY[1, 0, 4, 0, 0]::float4[]),
         embedding
       ) AS similarity
FROM sparse_data;
```

当所有已存向量和查询向量都已经归一化时，可用点积避免重复计算范数：

```sql
SELECT id,
       dot_product(
         sparse_vector_norm(ARRAY[1, 0, 4, 0, 0]::float4[]),
         embedding
       ) AS similarity
FROM sparse_data;
```

### 重要对象

- `sparse_vector(float4[])`：保留非零元素及其从零开始的索引，构造稀疏值。
- `sparse_vector_norm(float4[])`：将元素除以数组的欧几里得范数后构造值。
- `dot_product(sparse_vector, sparse_vector)`：对索引相同的已存元素求点积，返回 `float4`。
- `cosine_similarity(sparse_vector, sparse_vector)`：用点积除以两个实时计算的范数。

### 注意事项

该表示不存储原始稠密维数；末尾的零会丢失，从不同数组长度创建的向量也不会被拒绝。应在应用元数据中保证维数。输入应为一维、不含空元素的 `float4[]`；C 实现直接读取数组载荷，没有说明多维数组或空元素的处理方式。

全零向量的范数为零，因此余弦相似度可能产生非有限结果。对此类向量做归一化也没有有意义的方向。`'(0, 1)'::sparse_vector` 之类的文本转换会按设计失败；必须使用构造函数。

扩展没有比较运算符、聚合、转回数组的转换、索引运算符类或近似最近邻搜索。因此，除非应用增加其他检索策略，相似度查询会进行扫描和计算。无需预加载或重启。
