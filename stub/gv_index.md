## Usage

Sources:

- [Official gv_index.control](https://github.com/opengauss-mirror/openGauss-server/blob/9c01442e7d6811e5d3ab4c621d9153b6ce1af678/contrib/gv_index/gv_index.control)
- [Official gv_index--1.0.sql](https://github.com/opengauss-mirror/openGauss-server/blob/9c01442e7d6811e5d3ab4c621d9153b6ce1af678/contrib/gv_index/gv_index--1.0.sql)
- [Official graph_test.sql](https://github.com/opengauss-mirror/openGauss-server/blob/9c01442e7d6811e5d3ab4c621d9153b6ce1af678/contrib/gv_index/sql/graph_test.sql)
- [Official CMakeLists.txt](https://github.com/opengauss-mirror/openGauss-server/blob/9c01442e7d6811e5d3ab4c621d9153b6ce1af678/contrib/gv_index/CMakeLists.txt)

`gv_index` 1.0 adds the `gv_graph` vector access method to openGauss. It uses openGauss vector types and operators plus the GaussVector library; it is not a stock PostgreSQL/pgvector compatibility claim.

### Create an Index

```sql
CREATE EXTENSION gv_index;
SET enable_indexscan_optimization = on;
CREATE TABLE graph_items (id int, repr vector(128)) WITH (storage_type=ustore);
CREATE INDEX graph_items_idx ON graph_items USING gv_graph (repr vector_l2_ops)
  WITH (graph_degree=48, quantization_type=lvq, subgraph_count=2, num_parallels=32);
SET gv_graph_nprobes = 256;
```

### Query and Maintain

Load dimension-matched vectors into the source table and order a query by L2 distance to a query vector, using the kernel distance operator. The SQL script defines `vector_l2_ops` and `vector_cosine_ops`; the latter uses the openGauss inner-product operator and norm function. The upstream example also demonstrates INSERT, DELETE and VACUUM. `gv_graph_nprobes` controls search probing, while construction options include graph degree, quantization, subgraph count and parallelism. These controls belong to this kernel module, not the similarly named pgvector access methods.

### Runtime Boundary

The build requires `VECTOR_HOME` pointing to GaussVector headers and `libannlite.so`; when the directory is absent the module is skipped. The example uses ustore tables and index-scan optimization. Install with administrator privileges. The control is relocatable but not trusted; library initialization occurs when loaded and does not require shared preload. No PostgreSQL-major support list is published for this openGauss module.
