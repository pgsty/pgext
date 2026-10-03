## 用法

来源：

- [Official gv_index.control](https://github.com/opengauss-mirror/openGauss-server/blob/9c01442e7d6811e5d3ab4c621d9153b6ce1af678/contrib/gv_index/gv_index.control)
- [Official gv_index--1.0.sql](https://github.com/opengauss-mirror/openGauss-server/blob/9c01442e7d6811e5d3ab4c621d9153b6ce1af678/contrib/gv_index/gv_index--1.0.sql)
- [Official graph_test.sql](https://github.com/opengauss-mirror/openGauss-server/blob/9c01442e7d6811e5d3ab4c621d9153b6ce1af678/contrib/gv_index/sql/graph_test.sql)
- [Official CMakeLists.txt](https://github.com/opengauss-mirror/openGauss-server/blob/9c01442e7d6811e5d3ab4c621d9153b6ce1af678/contrib/gv_index/CMakeLists.txt)

`gv_index` 1.0 为 openGauss 增加 `gv_graph` 向量访问方法。它使用 openGauss 的向量类型、操作符及 GaussVector 库，不代表支持原生 PostgreSQL 或 pgvector。

### 创建索引

```sql
CREATE EXTENSION gv_index;
SET enable_indexscan_optimization = on;
CREATE TABLE graph_items (id int, repr vector(128)) WITH (storage_type=ustore);
CREATE INDEX graph_items_idx ON graph_items USING gv_graph (repr vector_l2_ops)
  WITH (graph_degree=48, quantization_type=lvq, subgraph_count=2, num_parallels=32);
SET gv_graph_nprobes = 256;
```

### 查询与维护

将维度匹配的向量载入来源表，使用内核距离操作符按与查询向量的 L2 距离排序。SQL 脚本定义 `vector_l2_ops` 与 `vector_cosine_ops`；后者使用 openGauss 的内积操作符和范数函数。上游示例还演示 INSERT、DELETE 与 VACUUM。`gv_graph_nprobes` 控制搜索探测，构建选项包括图的度、量化方式、子图数与并行度。这些参数属于该内核模块，不属于名称相近的 pgvector 访问方法。

### 运行边界

构建要求 `VECTOR_HOME` 指向 GaussVector 头文件与 `libannlite.so`；目录缺失时会跳过此模块。示例使用 ustore 表与索引扫描优化。安装需要管理员权限。控制文件标记可重定位但非可信；库加载时完成初始化，无需共享预加载。此 openGauss 模块没有公布 PostgreSQL 大版本支持列表。
