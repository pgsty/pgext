## 用法

来源：

- [0.5.0 版本官方 README](https://github.com/aaronsb/knowledge-graph-system/blob/b6ebe037b5d6d89921fa62e09bc8c0a54394aa18/graph-accel/README.md)
- [官方扩展控制文件](https://github.com/aaronsb/knowledge-graph-system/blob/b6ebe037b5d6d89921fa62e09bc8c0a54394aa18/graph-accel/ext/graph_accel.control)
- [官方加载实现](https://github.com/aaronsb/knowledge-graph-system/blob/b6ebe037b5d6d89921fa62e09bc8c0a54394aa18/graph-accel/ext/src/load.rs)

`graph_accel` 是 Apache AGE 图数据的只读内存遍历加速器。适合需要反复执行邻域、最短路径、度数或诱导子图查询，同时不希望每次都重新执行 Cypher 的场景。缓存属于单个 PostgreSQL 后端，不会在会话间共享。

### 核心流程

先安装 Apache AGE，再由超级用户创建 `graph_accel`，选择 AGE 图并将邻接数据加载到当前后端：

```sql
CREATE EXTENSION age;
CREATE EXTENSION graph_accel;

SET graph_accel.source_graph = 'my_graph';
SET graph_accel.node_id_property = 'concept_id';

SELECT * FROM graph_accel_load('my_graph');
SELECT * FROM graph_accel_neighborhood('concept-42', 3, 'both');
SELECT * FROM graph_accel_path('concept-42', 'concept-99', 10, 'out');
SELECT * FROM graph_accel_degree(20);
SELECT * FROM graph_accel_status();
```

`graph_accel_load()` 返回节点数、边数与加载时间。`graph_accel_neighborhood()` 返回每个可达节点的标签、应用标识、距离与路径边元数据。`graph_accel_path()` 在跳数、方向和可选置信度条件内查找无权最短路径。`graph_accel_subgraph()` 返回种子标识周围的诱导子图。

### 配置与缓存失效

- `graph_accel.source_graph` 选择 AGE 图，也可直接传给 `graph_accel_load()`。
- `graph_accel.node_id_property` 选择作为应用标识暴露的节点属性。
- `graph_accel.node_labels` 与 `graph_accel.edge_types` 限制加载的标签和边类型；`*` 表示全部。
- `graph_accel.max_memory_mb` 限制每个后端的缓存，而非整个集群的内存。
- `graph_accel.auto_reload` 与 `graph_accel.reload_debounce_sec` 控制基于代次的重新加载行为。

AGE 写入绕过本扩展使用的普通 PostgreSQL 触发器。图数据变化后，写入方必须显式推进代次：

```sql
SELECT graph_accel_invalidate('my_graph');
```

该函数会发送通知，但其他会话只有在下一次检查并发现代次不匹配时才重新加载。若没有协作式失效，它们可能继续返回过期拓扑。

### 运维说明

已复核 README 覆盖 PostgreSQL 13-18，但其预构建产物绑定文档所述镜像的精确 PostgreSQL/Apache AGE ABI；其他组合应重新构建。内存消耗会随加载同一图的后端数量成倍增加。该扩展不能替代 AGE 作为事实来源，也不提供加权最短路径或写操作。
