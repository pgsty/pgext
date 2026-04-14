## 用法
> 来源: [主页](https://onesparse.github.io/OneSparse/)、[入门文档](https://onesparse.github.io/OneSparse/docs.html)、[Matrix](https://onesparse.github.io/OneSparse/test_matrix_header.html)、[Vector](https://onesparse.github.io/OneSparse/test_vector_header.html) 和 [Algorithms](https://onesparse.github.io/OneSparse/test_examples_header.html)。

OneSparse 是一个将 SuiteSparse:GraphBLAS 绑定到 PostgreSQL 的扩展，并以新类型、函数和运算符的形式提供稀疏线性代数与图算法能力。
文档将 `matrix` 视为核心类型，`vector` 和 `scalar` 都建立在同一模型之上。

### 核心设置

```sql
CREATE EXTENSION onesparse;
SET search_path TO public,onesparse;

SELECT 'int32'::matrix;
SELECT 'int32'::vector;
SELECT 'int32:42'::scalar;
```

入门文档指出，OneSparse 将 API 放在 `onesparse` 模式中，matrix 和 vector 页面也展示了同样的 `search_path` 交互方式。

### 矩阵与向量

Matrix 页面展示了常见的矩阵操作，例如构造、打印、绘制、调整大小、类型转换和聚合。
Vector 页面展示了对应的向量 API，包括 `nvals`、`size`、`set_element`、`get_element`、`eadd`、`emult`、`reduce_scalar`、`choose` 和 `apply`。

```sql
SELECT print('int32(4:4)'::matrix);
SELECT draw('int32(4:4)[1:2:1 2:3:2 3:1:3]'::matrix);
SELECT eadd('int32[0:1 1:2 2:3]'::vector, 'int32[0:1 1:2 2:3]'::vector, 'plus_int32');
SELECT reduce_scalar('int32[0:1 1:2 2:3]'::vector, 'plus_monoid_int32');
```

### 图算法

入门文档使用 Matrix Market 文件和随机图来演示图算法。
它重点介绍了以下算法：

- `bfs(graph, 1)`，用于层级和父节点 BFS
- `sssp(cast_to(graph, 'int32'), 1::bigint, 1)`，用于单源最短路径
- `pagerank(graph)`，用于按链接结构对顶点排序
- `triangle_centrality(graph)`，用于基于三角形的中心性分析
- `betweenness(graph, ARRAY[...])` 和 `square_clustering(graph)`，用于更多图分析

文档中的代表性示例：

```sql
SELECT draw(triu(graph), (SELECT level FROM bfs(graph, 1)), false, false, true, 0.5)
FROM karate;
```

同一指南还展示了使用 `mmread(...)` 加载图，以及使用 `draw(...)` 进行图可视化。

### 范围

这套文档覆盖面很广。本 stub 只提炼了核心接口，以及在入门、矩阵、向量和算法页面中反复出现的主要示例，而不重复完整的 GraphBLAS 能力清单。
