## 用法

来源：

- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/README.md)
- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/LICENSE)
- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_graph/pg_graph.control)
- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_graph/pg_graph--1.0.sql)
- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_graph/meson.build)
- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_graph/pg_graph.c)

`pg_graph` `1.0` 提供普通关系表上的属性图遍历与算法。本条目限定于 AlohaDB 21.1.0 的 PostgreSQL 18 分支，是预览模块；不表示可直接安装在标准 PostgreSQL 上。

### 工作流程

```sql
CREATE EXTENSION pg_graph;
SELECT graph_create('demo');
```

安装需要超级用户；运行函数还需相应对象权限。

### 运行边界

graph_add_vertex_label 与 graph_add_edge_label 绑定已有表，并可安装外键、检查约束。graph_bfs、graph_pagerank 和 graph_sample_neighbors 提供遍历、排序与采样，时间列可启用 as_of 选择。图读取遵循表权限和行级安全。算法会将边加载到后端内存，pg_graph.max_edges 用于限定规模。启用约束前需验证已有数据，并评估大图计算成本。

此模块对应上游旧名 `alohadb_graph`，应遵循 AlohaDB 的集群及对象改名流程，不能将其他项目的同名扩展当作替代品。升级前应演练依赖对象、权限和备份恢复。
