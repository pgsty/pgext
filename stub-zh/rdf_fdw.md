
## 用法

> 来源：[README](https://github.com/jimjonesbr/rdf_fdw/blob/main/README.md), [v2.4 release](https://github.com/jimjonesbr/rdf_fdw/releases/tag/v2.4)

`rdf_fdw` 是面向 RDF triplestore 的 foreign data wrapper，通过 SPARQL endpoint 暴露数据。它允许 PostgreSQL 用 SQL 查询 RDF 数据，支持 SQL 子句下推，提供用于 RDF term 的 `rdfnode` 类型，并包含 SPARQL 1.1 函数支持。

### 注册 SPARQL Endpoint

使用 `CREATE SERVER` 注册一个 SPARQL endpoint：

```sql
CREATE SERVER dbpedia
FOREIGN DATA WRAPPER rdf_fdw
OPTIONS (endpoint 'https://dbpedia.org/sparql');
```

README 记录的 server options 包括：

- `endpoint`（必需）
- `batch_size`
- `enable_pushdown`
- `format`
- `http_proxy`
- `connect_timeout`

`v2.4` 把代理凭据从 `SERVER` options 挪到了 `USER MAPPING`，以降低凭据暴露风险。

### 用户映射与外部表

```sql
CREATE USER MAPPING FOR postgres
SERVER dbpedia
OPTIONS (user 'admin', password 'secret');
```

`rdf_fdw` 通过声明外部表来嵌入 SPARQL 查询，并把结果变量映射到 PostgreSQL 列。README 也特别强调了通过自定义 `rdfnode` 类型对 RDF node 的原生处理。

```sql
CREATE FOREIGN TABLE film (
  film_id text OPTIONS (variable '?film', nodetype 'iri'),
  name text OPTIONS (variable '?name', nodetype 'literal', literal_type 'xsd:string')
)
SERVER dbpedia
OPTIONS (sparql 'SELECT ?film ?name WHERE { ?film dbp:name ?name }');
```

### 下推、DML 与辅助函数

上游文档明确点名支持这些下推：

- `WHERE`
- `LIMIT`
- `ORDER BY`
- `DISTINCT`

它也记录了数据修改支持：

- `INSERT`
- `UPDATE`
- `DELETE`

SPARQL UPDATE 的批处理粒度由 `batch_size` 控制。

README 列出的工具函数包括：

- `rdf_fdw_version()`
- `rdf_fdw_settings()`
- `rdf_fdw_clone_table()`

它还说明扩展覆盖了更广泛的 SPARQL 函数，包括 aggregates、string functions、numeric functions、date/time functions、hash functions 和 custom functions。

### 注意事项

当前 README 警告：检索到的 RDF 数据会先完整载入内存，再转换为 PostgreSQL 表示，因此大结果集需要足够的 PostgreSQL 内存。文档声明的最低支持版本是 PostgreSQL 9.5+。
