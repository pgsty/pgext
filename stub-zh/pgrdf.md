## 用法

来源：

- [pgRDF 0.6.20 README](https://github.com/styk-tv/pgRDF/blob/v0.6.20/README.md)
- [pgRDF 0.6.20 用户指南](https://github.com/styk-tv/pgRDF/tree/v0.6.20/guide)
- [pgRDF 0.6.20 更新日志](https://github.com/styk-tv/pgRDF/blob/v0.6.20/CHANGELOG.md)
- [pgRDF 0.6.20 发行版](https://github.com/styk-tv/pgRDF/releases/tag/v0.6.20)

`pgRDF` 将 RDF 数据存储在 PostgreSQL 中，并提供了用于加载 Turtle/TriG/N-Quads、SPARQL 查询/更新、命名图、SHACL 验证和 RDFS/OWL 2 RL 物化视图的 SQL 可调用辅助函数。

```sql
CREATE EXTENSION pgrdf;
SELECT pgrdf.version();
```

### 预加载与 PostgreSQL 版本注意事项

pgRDF 0.6.20 支持 PostgreSQL 14-18，并从 `pgrx` 0.16.1 升级到 0.19.1。上游描述 0.6.20 是一个构建/运行时迁移，没有模式或查询表面的更改；PostgreSQL 19 仍然是跟踪后续版本。

在 PostgreSQL 启动前，`pgrdf` 必须存在于 `shared_preload_libraries` 中。如果没有预加载，上游文档指出共享内存字典和计划缓存原子操作未初始化，首次调用 pgRDF 可能会失败。

```conf
shared_preload_libraries = 'pgrdf'
```

更改此设置后，请重启 PostgreSQL，并验证：

```sql
SHOW shared_preload_libraries;

SELECT pgrdf.parse_turtle(
  '@prefix ex: <http://example.org/> . ex:t a ex:T .',
  1::bigint,
  'http://example.org/'
);
```

### 加载 RDF 数据

使用 `parse_turtle` 用于内联 Turtle 载荷，使用 `load_turtle` 用于服务器端文件。图 ID 是 `bigint` 值；命名图辅助函数将 ID 映射到 IRI。版本 0.6.x 添加了通过 `load_turtle(..., bulk_load => true)` 的并行批量加载路径。

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

上游文档中相关的数据导入和图管理功能包括 `parse_trig`、`parse_nquads`、`add_graph`、`drop`、`clear`、`copy`、`move_graph`、`graph_id` 和 `graph_iri`。

### 切分图片段

0.6.x 系列增加了 `carve_graph` 重载，用于在不解码和重新编码共享字典的情况下将谓词定义的切片或有界邻域复制到另一个图：

```sql
SELECT pgrdf.carve_graph(
  100::bigint,
  'http://example.org/type'::text,
  200::bigint
);

SELECT pgrdf.carve_graph(
  100::bigint,
  ARRAY['http://example.org/alice', 'http://example.org/bob']::text[],
  201::bigint,
  2
);
```

有界邻域形式使用 `max_hops` 作为图距离边界。当必须阻止被截断的属性路径遍历时，请将 `pgrdf.on_path_truncation` 设置为 `warn` 或 `error`。

### 使用 SPARQL 查询

`pgrdf.sparql(text)` 返回 SPARQL 结果作为 SQL 行。上游 v0.5 表面包括 `SELECT` 和 `ASK`、过滤器、排序、限制、`OPTIONAL`、`UNION`、`MINUS`、聚合、`VALUES`、`BIND`、`CONSTRUCT`、`DESCRIBE`、命名图的 `GRAPH` 子句以及属性路径。

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

上游 v0.5 表面包括 SPARQL Update 形式，如 `INSERT DATA`、`DELETE DATA`、`INSERT/DELETE WHERE`、`DELETE+INSERT WHERE` 和图生命周期语句。

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

### 推理与验证

使用 `pgrdf.materialize(graph_id, profile)` 将 `rdfs` 或 `owl-rl` 语义写入推断三元组。物化旨在可重复；上游文档指出在写入新的闭包之前会删除之前的推断行。

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

使用 `pgrdf.validate(data, shapes, mode)` 进行 SHACL 验证；上游文档 JSONB `sh:ValidationReport` 输出和原生 SHACL Core 支持。SHACL-SPARQL 约束执行由其 RDF 库依赖项控制，因此将 `mode => 'sparql'` 视为高级表面以验证安装的确切版本。

### 运营辅助函数

上游文档中包含的有用内省和缓存管理辅助函数包括：

| 函数 | 用途 |
|----------|-----|
| `pgrdf.stats()` | 检查运行时计数器和缓存状态 |
| `pgrdf.shmem_reset()` | 重置共享内存字典/缓存状态 |
| `pgrdf.plan_cache_clear()` | 清除 SPARQL 计划缓存 |
| `pgrdf.sparql_parse(text)` | 检查解析的 SPARQL 而不执行 |

`pgrdf.path_max_depth` 设置保护属性路径扩展深度，而 `pgrdf.on_path_truncation = count | warn | error` 控制调用者如何得知已达到限制。

### 版本说明

从 0.6.4 到 0.6.20 的发行版在大规模 RDF 导入和查询正确性方面有了实质性改进：流式/窗口批量导入、分阶段多后端加载器、安全重复加载到填充字典、图切片辅助函数、字典包含在 `pg_dump` 中、SPARQL 表达式/聚合添加以及失败关闭路径截断处理。0.6.20 发行版本身仅将构建/运行时层更改为 `pgrx` 0.19.1，并增加了对 PostgreSQL 18 的支持。

对于非常大的新鲜 N-Triples 导入，上游文档建议使用 `pgrdf.load_turtle_staged_run` 作为可恢复的、阶段导向的路径。它分别提交解析、字典、解析和索引阶段，并且操作上不同于事务性的 `load_turtle()` 调用；在用于生产规模导入之前，请验证分阶段表空间、磁盘空闲空间以及恢复程序。
