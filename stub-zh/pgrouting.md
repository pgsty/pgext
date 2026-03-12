

## 用法

> [pgRouting - 基于 PostgreSQL 的路径规划](https://github.com/pgRouting/pgrouting)

pgRouting 扩展 PostGIS/PostgreSQL 地理空间数据库，提供地理空间路径规划和其他网络分析功能。

该库包含以下特性：

- 全对最短路径（Floyd-Warshall、Johnson）
- A* 算法（含双向变体）
- Dijkstra 算法（代价、代价矩阵、行驶距离、K 条最短路径、经由路由、最近点）
- 双向 Dijkstra
- 旅行商问题（TSP）
- 网络流（最大流、Boykov-Kolmogorov、Edmonds-Karp、预流推进）
- 生成树（Kruskal、Prim 及其 BFS/DFS/行驶距离变体）
- 图组件（连通分量、强连通分量、双连通分量、关节点、桥）
- 转弯限制最短路径（TRSP）
- WithPoints 路由（边上任意位置）
- 图压缩与实用函数

### 快速开始

启用扩展（需要 PostGIS）：

```sql
CREATE EXTENSION pgrouting CASCADE;
```

### 图的表示

pgRouting 使用返回边数据的 SQL 查询来表示图。标准边查询格式：

```sql
SELECT id, source, target, cost, reverse_cost FROM edges;
```

| 列 | 类型 | 说明 |
|----|------|------|
| `id` | ANY-INTEGER | 边标识符 |
| `source` | ANY-INTEGER | 起始顶点标识符 |
| `target` | ANY-INTEGER | 终止顶点标识符 |
| `cost` | ANY-NUMERICAL | 权重（source 到 target）；负值表示排除该边 |
| `reverse_cost` | ANY-NUMERICAL | 权重（target 到 source）；默认 -1（不存在） |

### 无几何体的简单示例

创建一个图并查找最短路径：

```sql
CREATE TABLE wiki (
  id SERIAL,
  source INTEGER,
  target INTEGER,
  cost INTEGER
);

INSERT INTO wiki (source, target, cost) VALUES
  (1, 2, 7),  (1, 3, 9), (1, 6, 14),
  (2, 3, 10), (2, 4, 15),
  (3, 6, 2),  (3, 4, 11),
  (4, 5, 6),
  (5, 6, 9);

SELECT * FROM pgr_dijkstra(
  'SELECT id, source, target, cost FROM wiki',
  1, 5, false);
```

--------

## 函数族

### Dijkstra - 最短路径

核心路由函数。支持一对一、一对多、多对一、多对多及组合签名。

```sql
pgr_dijkstra(Edges SQL, start_vid,  end_vid,  [directed])
pgr_dijkstra(Edges SQL, start_vid,  end_vids, [directed])
pgr_dijkstra(Edges SQL, start_vids, end_vid,  [directed])
pgr_dijkstra(Edges SQL, start_vids, end_vids, [directed])
pgr_dijkstra(Edges SQL, Combinations SQL,     [directed])
```

返回：`(seq, path_seq, start_vid, end_vid, node, edge, cost, agg_cost)`

**一对一：**

```sql
SELECT * FROM pgr_dijkstra(
  'SELECT id, source, target, cost, reverse_cost FROM edges',
  6, 10, true);
```

**一对多：**

```sql
SELECT * FROM pgr_dijkstra(
  'SELECT id, source, target, cost, reverse_cost FROM edges',
  6, ARRAY[10, 17]);
```

**多对多（无向）：**

```sql
SELECT * FROM pgr_dijkstra(
  'SELECT id, source, target, cost, reverse_cost FROM edges',
  ARRAY[6, 1], ARRAY[10, 17],
  directed => false);
```

**组合：**

```sql
SELECT * FROM pgr_dijkstra(
  'SELECT id, source, target, cost, reverse_cost FROM edges',
  'SELECT source, target FROM combinations',
  false);
```

#### Dijkstra 代价

仅返回聚合代价，不含路径详情：

```sql
pgr_dijkstraCost(Edges SQL, start_vid, end_vid, [directed])
```

返回：`(start_vid, end_vid, agg_cost)`

#### Dijkstra 代价矩阵

为一组顶点生成代价矩阵：

```sql
pgr_dijkstraCostMatrix(Edges SQL, vids, [directed])
```

#### Dijkstra 经由

按有序顶点序列规划路径：

```sql
pgr_dijkstraVia(Edges SQL, via_vertices, [directed, strict, U_turn_on_edge])
```

#### Dijkstra 最近点

查找距离一组目标最近的顶点：

```sql
pgr_dijkstraNear(Edges SQL, start_vid, end_vids, [directed])
```

### A* - 最短路径

使用 A* 启发式算法。需要边查询中包含额外的坐标列（`x1`、`y1`、`x2`、`y2`）。

```sql
pgr_aStar(Edges SQL, start_vid, end_vid, [directed, heuristic, factor, epsilon])
```

| 选项 | 类型 | 默认值 | 说明 |
|------|------|--------|------|
| `directed` | BOOLEAN | `true` | 图方向 |
| `heuristic` | INTEGER | `5` | 距离启发式（0-5） |
| `factor` | FLOAT | `1` | 单位换算因子 |
| `epsilon` | FLOAT | `1` | 近似因子 |

```sql
SELECT * FROM pgr_aStar(
  'SELECT id, source, target, cost, reverse_cost, x1, y1, x2, y2 FROM edges',
  6, 12,
  directed => true, heuristic => 2);
```

另有：`pgr_aStarCost`、`pgr_aStarCostMatrix`

### 双向算法

双向变体从两端同时搜索：

- `pgr_bdDijkstra`、`pgr_bdDijkstraCost`、`pgr_bdDijkstraCostMatrix`
- `pgr_bdAstar`、`pgr_bdAstarCost`、`pgr_bdAstarCostMatrix`

```sql
SELECT * FROM pgr_bdDijkstra(
  'SELECT id, source, target, cost, reverse_cost FROM edges',
  6, 10);
```

### K 条最短路径（Yen 算法）

查找两个顶点之间的 K 条最短路径：

```sql
pgr_KSP(Edges SQL, start_vid, end_vid, K, [directed, heap_paths])
```

返回：`(seq, path_id, path_seq, start_vid, end_vid, node, edge, cost, agg_cost)`

```sql
SELECT * FROM pgr_KSP(
  'SELECT id, source, target, cost, reverse_cost FROM edges',
  6, 17, 2);
```

### 行驶距离

查找给定距离内可达的所有顶点：

```sql
pgr_drivingDistance(Edges SQL, start_vid, distance, [directed])
pgr_drivingDistance(Edges SQL, start_vids, distance, [directed, equicost])
```

返回：`(seq, depth, start_vid, pred, node, edge, cost, agg_cost)`

```sql
SELECT * FROM pgr_drivingDistance(
  'SELECT id, source, target, cost, reverse_cost FROM edges',
  11, 3.0);
```

### 旅行商问题

**基于矩阵的 TSP：**

```sql
pgr_TSP(Matrix SQL, [start_id, end_id])
```

返回：`(seq, node, cost, agg_cost)`

```sql
SELECT * FROM pgr_TSP(
  $$SELECT * FROM pgr_dijkstraCostMatrix(
    'SELECT id, source, target, cost, reverse_cost FROM edges',
    ARRAY[1, 3, 5, 6, 7, 8, 9, 10, 11, 15, 16, 17],
    directed => false)$$,
  start_id => 1);
```

**欧几里得 TSP**（直接使用坐标）：

```sql
pgr_TSPeuclidean(Coordinates SQL, [start_id, end_id])
```

### 网络流

计算最大流及相关属性：

```sql
-- 最大流
pgr_maxFlow(Edges SQL, source, sink)

-- 特定算法
pgr_boykovKolmogorov(Edges SQL, source, sink)
pgr_edmondsKarp(Edges SQL, source, sink)
pgr_pushRelabel(Edges SQL, source, sink)

-- 边不相交路径
pgr_edgeDisjointPaths(Edges SQL, source, sink)

-- 最大基数匹配
pgr_maxCardinalityMatch(Edges SQL, [directed])
```

网络流的边 SQL 使用 `capacity` 和 `reverse_capacity` 替代 `cost`/`reverse_cost`。

### 生成树

**Kruskal 算法：**

```sql
pgr_kruskal(Edges SQL)         -- 最小生成树
pgr_kruskalBFS(Edges SQL, root_vid, [max_depth])
pgr_kruskalDFS(Edges SQL, root_vid, [max_depth])
pgr_kruskalDD(Edges SQL, root_vid, distance)
```

**Prim 算法：**

```sql
pgr_prim(Edges SQL)            -- 最小生成树
pgr_primBFS(Edges SQL, root_vid, [max_depth])
pgr_primDFS(Edges SQL, root_vid, [max_depth])
pgr_primDD(Edges SQL, root_vid, distance)
```

### 图组件

```sql
-- 连通分量（无向）
pgr_connectedComponents(Edges SQL)

-- 强连通分量（有向）
pgr_strongComponents(Edges SQL)

-- 双连通分量
pgr_biconnectedComponents(Edges SQL)

-- 关节点（割点）
pgr_articulationPoints(Edges SQL)

-- 桥（割边）
pgr_bridges(Edges SQL)
```

### 转弯限制最短路径（TRSP）

带禁止路径限制的路由：

```sql
pgr_trsp(Edges SQL, Restrictions SQL, start_vid, end_vid, [directed])
pgr_trspVia(Edges SQL, Restrictions SQL, via_vertices, [directed, strict, U_turn_on_edge])
pgr_trsp_withPoints(Edges SQL, Restrictions SQL, Points SQL, start_pid, end_pid, [options])
```

限制条件 SQL 格式：

| 列 | 类型 | 说明 |
|----|------|------|
| `path` | ARRAY[ANY-INTEGER] | 禁止的边 ID 序列 |
| `cost` | ANY-NUMERICAL | 禁止路径的代价 |

### WithPoints - 任意位置路由

在边上任意点（不仅是顶点）之间路由：

```sql
pgr_withPoints(Edges SQL, Points SQL, start_pid, end_pid, [options])
pgr_withPointsCost(Edges SQL, Points SQL, start_pid, end_pid, [options])
pgr_withPointsCostMatrix(Edges SQL, Points SQL, pids, [options])
pgr_withPointsKSP(Edges SQL, Points SQL, start_pid, end_pid, K, [options])
pgr_withPointsDD(Edges SQL, Points SQL, start_pid, distance, [options])
```

点 SQL 格式：

| 列 | 类型 | 默认值 | 说明 |
|----|------|--------|------|
| `pid` | ANY-INTEGER | | 点标识符 |
| `edge_id` | ANY-INTEGER | | 最近的边 |
| `fraction` | ANY-NUMERICAL | | 在边上的位置（0-1） |
| `side` | CHAR | `b` | `r`(右侧)、`l`(左侧)、`b`(两侧) |

### 图压缩

通过压缩顶点简化图：

```sql
pgr_contraction(Edges SQL, contraction_order, [options])
```

### 实用函数

```sql
-- 从边数据中提取顶点
pgr_extractVertices(Edges SQL)

-- 查找点附近的边
pgr_findCloseEdges(Edges SQL, point, tolerance, [options])

-- 分离交叉几何体
pgr_separateCrossing(Edges SQL)

-- 分离相切几何体
pgr_separateTouching(Edges SQL)

-- 版本信息
pgr_version()
pgr_full_version()
```

--------

## 使用几何体

### 构建路由拓扑

从空间边中提取顶点并构建拓扑：

```sql
-- 从边几何体中提取顶点
SELECT * INTO vertices
FROM pgr_extractVertices('SELECT id, geom FROM edges ORDER BY id');

-- 设置起始顶点
UPDATE edges AS e
SET source = v.id, x1 = x, y1 = y
FROM vertices AS v
WHERE ST_StartPoint(e.geom) = v.geom;

-- 设置终止顶点
UPDATE edges AS e
SET target = v.id, x2 = x, y2 = y
FROM vertices AS v
WHERE ST_EndPoint(e.geom) = v.geom;
```

### 基于几何体长度设置代价

```sql
UPDATE edges SET
  cost = sign(cost) * ST_Length(geom),
  reverse_cost = sign(reverse_cost) * ST_Length(geom);
```

### 获取路径几何体

将路由结果与边几何体结合：

```sql
SELECT seq, node, edge, cost, agg_cost, geom
FROM pgr_dijkstra(
  'SELECT id, source, target, cost, reverse_cost FROM edges',
  6, 10
) AS r
LEFT JOIN edges AS e ON r.edge = e.id;
```

--------

## 性能优化

将查询限制在感兴趣的区域内，减少处理的边数：

```sql
SELECT * FROM pgr_dijkstra($$
  SELECT id, source, target, cost, reverse_cost
  FROM edges
  WHERE geom && (
    SELECT ST_Buffer(ST_Union(geom), 1)
    FROM edges WHERE source IN (7) OR target IN (8))$$,
  7, 8);
```

--------

## 全对最短路径

用于计算所有顶点对之间的距离：

```sql
-- Floyd-Warshall（不需要边 ID）
pgr_floydWarshall(Edges SQL, [directed])

-- Johnson（不需要边 ID）
pgr_johnson(Edges SQL, [directed])
```

返回：`(start_vid, end_vid, agg_cost)`

```sql
SELECT * FROM pgr_floydWarshall(
  'SELECT id, source, target, cost, reverse_cost FROM edges');
```
