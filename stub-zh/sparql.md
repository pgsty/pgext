

## 用法

> [sparql: PostgreSQL 的 SPARQL 查询支持](https://github.com/lacanoid/pgsparql)

从 PostgreSQL 查询 SPARQL 端点（如 DBpedia/Virtuoso）。SPARQL 查询被编译为 PostgreSQL 视图，可在 SQL 中使用。

### 获取资源的属性

```sql
SELECT * FROM sparql.get_properties('dbpedia', 'http://dbpedia.org/resource/Johann_Sebastian_Bach');
```

### 获取资源的引用

```sql
SELECT * FROM sparql.get_references('dbpedia', 'http://dbpedia.org/resource/Johann_Sebastian_Bach');
```

### 将 SPARQL 查询编译为 SQL 视图

```sql
SELECT sparql.compile_query(endpoint, identifier, sparql_query [, grouping]);
```

参数：
- `endpoint` -- 默认 SPARQL 端点名称
- `identifier` -- 创建的函数和视图的 SQL 标识符
- `sparql_query` -- 要编译的 SPARQL 查询
- `grouping` -- 可选的分组标识符数组（未分组的列会被聚合为数组）

### 示例

```sql
SELECT sparql.compile_query('dbpedia', 'ludwig_van', $$
  SELECT ?predicate, ?object
  WHERE {
    <http://dbpedia.org/resource/Ludwig_van_Beethoven> ?predicate ?object.
  }
$$, '{predicate}');

-- 现在可以通过创建的视图进行查询
SELECT * FROM ludwig_van;
```

这将创建一个函数 `ludwig_van()` 和一个视图 `ludwig_van`，用于查询 SPARQL 端点并以 SQL 表的形式返回结果。
