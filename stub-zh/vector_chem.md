## 用法

来源：

- [vector_chem.control](https://github.com/leotaku/pgvector_chem/blob/d87e9c8688faca1afbf825095107452b3b0e52cb/vector_chem.control)
- [README.md](https://github.com/leotaku/pgvector_chem/blob/d87e9c8688faca1afbf825095107452b3b0e52cb/README.md)
- [sql/vector_chem.sql](https://github.com/leotaku/pgvector_chem/blob/d87e9c8688faca1afbf825095107452b3b0e52cb/sql/vector_chem.sql)
- [Makefile](https://github.com/leotaku/pgvector_chem/blob/d87e9c8688faca1afbf825095107452b3b0e52cb/Makefile)

`vector_chem` 为二值 pgvector 向量增加 Tanimoto/Jaccard 距离，提供 `<^>` 运算符与 `tanimoto_distance` 函数。

### 核心用法

```sql
CREATE EXTENSION vector;
CREATE EXTENSION vector_chem;
SELECT tanimoto_distance('[1,0,1]'::vector, '[1,0,0]'::vector);
```

### 运行边界

先安装匹配的 pgvector 库并启用 `vector`；虽然控制文件未声明依赖，SQL 仍使用该类型。安装需要超级用户权限，没有已声明的预加载要求。本扩展不支持连续值向量或近似最近邻索引，只提供标量距离计算。
