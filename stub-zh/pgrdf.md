


## 用法

> 来源：[pgRDF upstream README](https://github.com/styk-tv/pgRDF/blob/v0.6.4/README.md)、[pgRDF user guide](https://github.com/styk-tv/pgRDF/tree/v0.6.4/guide)、[v0.6.4 release](https://github.com/styk-tv/pgRDF/releases/tag/v0.6.4)。

`pgRDF` 在 PostgreSQL 内存储 RDF 数据，并提供可从 SQL 调用的辅助函数，用于加载 Turtle/TriG/N-Quads、执行 SPARQL 查询与更新、管理命名图、做 SHACL 校验，以及执行 RDFS/OWL 2 RL 物化推理。

```sql
CREATE EXTENSION pgrdf;
SELECT pgrdf.version();
```

### 预加载与 PostgreSQL 版本注意事项

上游文档说明支持 PostgreSQL 14-17；由于 pgRDF 仍固定在 `pgrx` 0.16，PostgreSQL 18 支持暂缓。

`pgrdf` 必须在 PostgreSQL 启动前加入 `shared_preload_libraries`。如果没有预加载，上游文档说明共享内存字典和计划缓存原子量不会初始化，第一次调用 pgRDF 时可能失败。

```conf
shared_preload_libraries = 'pgrdf'
```

修改该设置后重启 PostgreSQL，然后验证：

```sql
SHOW shared_preload_libraries;

SELECT pgrdf.parse_turtle(
  '@prefix ex: <http://example.org/> . ex:t a ex:T .',
  1::bigint,
  'http://example.org/'
);
```

### 加载 RDF

内联 Turtle 载荷使用 `parse_turtle`，服务器端文件使用 `load_turtle`。图 ID 是 `bigint` 值；命名图辅助函数负责在 ID 和 IRI 之间建立映射。版本 0.6.x 通过 `load_turtle(..., bulk_load => true)` 增加 parallel bulk loader 路径。

```sql
SELECT pgrdf.add_graph(100::bigint, 'http://example.org/graph/main');

SELECT pgrdf.parse_turtle(
  '@prefix ex: <http://example.org/> .
   ex:alice ex:knows ex:bob .
   ex:alice ex:name "Alice" .',
  100::bigint,
  'http://example.org/graph/main'
);

SELECT pgrdf.load_turtle('/srv/rdf/foaf.ttl', 100::bigint);
SELECT pgrdf.load_turtle('/srv/rdf/bulk.ttl', 100::bigint, bulk_load => true);
SELECT pgrdf.count_quads(100::bigint);
```

上游文档列出的相关导入和图管理函数包括 `parse_trig`、`parse_nquads`、`add_graph`、`drop`、`clear`、`copy`、`move_graph`、`graph_id` 和 `graph_iri`。

### 使用 SPARQL 查询

`pgrdf.sparql(text)` 将 SPARQL 结果作为 SQL 行返回。上游 v0.5 接口包含 `SELECT` 和 `ASK`、过滤、排序、限制、`OPTIONAL`、`UNION`、`MINUS`、聚合、`VALUES`、`BIND`、`CONSTRUCT`、`DESCRIBE`、命名图 `GRAPH` 子句以及属性路径。

```sql
SELECT *
FROM pgrdf.sparql(
  'PREFIX ex: <http://example.org/>
   SELECT ?person ?name
   WHERE {
     ?person ex:name ?name .
     FILTER(REGEX(?name, "^A", "i"))
   }
   ORDER BY ?name
   LIMIT 20'
);
```

命名图查询可以绑定图 IRI：

```sql
SELECT *
FROM pgrdf.sparql(
  'PREFIX ex: <http://example.org/>
   SELECT ?g (COUNT(*) AS ?n)
   WHERE { GRAPH ?g { ?s ex:name ?name } }
   GROUP BY ?g
   ORDER BY ?g'
);
```

### 更新图

上游 v0.5 接口包含 `INSERT DATA`、`DELETE DATA`、`INSERT/DELETE WHERE`、`DELETE+INSERT WHERE` 等 SPARQL Update 形式，以及图生命周期语句。

```sql
SELECT pgrdf.sparql(
  'PREFIX ex: <http://example.org/>
   INSERT DATA {
     GRAPH <http://example.org/graph/main> {
       ex:bob ex:name "Bob" .
     }
   }'
);
```

### 推理与校验

使用 `pgrdf.materialize(graph_id, profile)` 为 `rdfs` 或 `owl-rl` profile 写入推导出的三元组。物化操作设计为可重复执行；上游文档说明，在写入新的闭包前会先删除之前的推导行。

```sql
SELECT pgrdf.parse_turtle(
  '@prefix ex:   <http://example.com/> .
   @prefix rdf:  <http://www.w3.org/1999/02/22-rdf-syntax-ns#> .
   @prefix rdfs: <http://www.w3.org/2000/01/rdf-schema#> .
   ex:Engineer rdfs:subClassOf ex:Person .
   ex:Person   rdfs:subClassOf ex:Agent .
   ex:alice    rdf:type        ex:Engineer .',
  100::bigint
);

SELECT pgrdf.materialize(100::bigint, 'owl-rl');

SELECT *
FROM pgrdf.sparql(
  'PREFIX rdf: <http://www.w3.org/1999/02/22-rdf-syntax-ns#>
   PREFIX ex:  <http://example.com/>
   SELECT ?class WHERE { ex:alice rdf:type ?class }'
);
```

使用 `pgrdf.validate(data, shapes, mode)` 做 SHACL 校验；上游文档说明其输出为 JSONB `sh:ValidationReport`，并原生支持 SHACL Core。上游还说明 SHACL-SPARQL 约束执行受其 RDF 库依赖控制，因此应将 `mode => 'sparql'` 视为高级接口，并以实际安装版本验证。

### 运维辅助函数

上游文档列出的常用内省和缓存管理函数包括：

| 函数 | 用途 |
|----------|-----|
| `pgrdf.stats()` | 查看运行时计数器和缓存状态 |
| `pgrdf.shmem_reset()` | 重置共享内存字典/缓存状态 |
| `pgrdf.plan_cache_clear()` | 清除预备 SPARQL 计划缓存 |
| `pgrdf.sparql_parse(text)` | 查看解析后的 SPARQL 而不执行 |

`pgrdf.path_max_depth` 设置用于限制属性路径展开深度。

### 版本说明

`pgrdf` 0.6.4 改进 deferred-index bulk-load 路径：对于高于 `pgrdf.bulk_defer_index_min` 的 fresh bulk loads，`load_turtle(..., bulk_load => true)` 也会推迟 dictionary `unique_term` constraint，然后在同一事务中重建并验证。由于 pgRDF 仍 pin 在 `pgrx` 0.16，PostgreSQL 18 仍由上游暂缓。
