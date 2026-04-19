## 用法

> 来源：[homepage](https://onesparse.com/), [release v1.0.0](https://github.com/OneSparse/OneSparse/releases/tag/v1.0.0), [control file at v1.0.0](https://raw.githubusercontent.com/OneSparse/OneSparse/v1.0.0/onesparse.control), [intro docs](https://onesparse.com/docs.html), [matrix docs](https://onesparse.com/test_matrix_header.html), [vector docs](https://onesparse.com/test_vector_header.html), [algorithm examples](https://onesparse.com/test_examples_header.html)

OneSparse 是一个把 SuiteSparse:GraphBLAS 绑定进 Postgres 的 PostgreSQL 扩展，它把稀疏线性代数和图算法暴露为新的类型、函数与运算符。文档将 `matrix` 视为核心类型，`vector` 和 `scalar` 建立在同一模型之上。`v1.0.0` release 已存在，但该标签下的 extension control file 仍声明 SQL `default_version = '0.1.0'`。

### 核心设置

```sql
CREATE EXTENSION onesparse;
SET search_path TO public,onesparse;

SELECT 'int32'::matrix;
SELECT 'int32'::vector;
SELECT 'int32:42'::scalar;
```

文档站点按 `matrix`、`vector` 和 `scalar` 组织 API，交互示例主要依赖类型转换和构造器。

### Matrix 与 Vector

`matrix` 页面展示了构造、`print()`、`draw()`、赋值、提取、`cast_to()`、调整大小和聚合等常见操作。`vector` 页面记录了对应的向量 API，包括 `nvals()`, `size()`, `eadd()`, `emult()`, `reduce_scalar()`, `choose()` 和 `apply()`。

```sql
SELECT print('int32(4:4)'::matrix);
SELECT draw('int32(4:4)[1:2:1 2:3:2 3:1:3]'::matrix);
SELECT eadd('int32[0:1 1:2 2:3]'::vector, 'int32[0:1 1:2 2:3]'::vector, 'plus_int32');
SELECT reduce_scalar('int32[0:1 1:2 2:3]'::vector, 'plus_monoid_int32');
```

### 图算法

示例页使用 Matrix Market 输入与 `draw(...)` 图形可视化。文档列出的图算法包括：

- `bfs(graph, 1)`，用于 level 和 parent BFS
- `sssp(cast_to(graph, 'int32'), 1::bigint, 1)`，用于单源最短路径
- `pagerank(graph)`，用于按链接结构给顶点排序
- `triangle_centrality(graph)`，用于基于三角形的中心性
- `betweenness(graph, ARRAY[...])` 和 `square_clustering(graph)`，用于补充图分析

文档中的代表性示例：

```sql
SELECT draw(triu(graph), (SELECT level FROM bfs(graph, 1)), false, false, true, 0.5)
FROM karate;
```

同一份指南还展示了使用 `mmread('/home/postgres/onesparse/demo/karate.mtx')` 加载图。
