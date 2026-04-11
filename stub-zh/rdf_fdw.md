
## 用法

> 语法：
>
> ```sql
> CREATE EXTENSION rdf_fdw;
> CREATE SERVER dbpedia FOREIGN DATA WRAPPER rdf_fdw
>   OPTIONS (endpoint 'https://dbpedia.org/sparql');
> ```
>
> 来源：[README](https://github.com/jimjonesbr/rdf_fdw)

`rdf_fdw` 是面向 RDF 三元组存储的外部数据包装器，通过 SPARQL endpoint 与 PostgreSQL 集成。它允许用 SQL 查询 RDF 数据，支持 SQL 子句下推，增加了用于 RDF term 的 `rdfnode` 类型，并提供 SPARQL 1.1 函数支持。

## 服务器配置

使用 `CREATE SERVER` 注册 SPARQL endpoint：

```sql
CREATE SERVER dbpedia
FOREIGN DATA WRAPPER rdf_fdw
OPTIONS (endpoint 'https://dbpedia.org/sparql');
```

README 记录的服务器参数包括：

- `endpoint`（必需）
- `batch_size`
- `enable_pushdown`
- `format`
- `http_proxy`
- `connect_timeout`

代理凭据应放在 user mapping 中。

## 外部表

`rdf_fdw` 通过声明外部表来嵌入 SPARQL 查询，并将结果变量映射到 PostgreSQL 列。README 还强调了通过自定义 `rdfnode` 类型对 RDF 节点的原生处理。

## 下推与 DML

上游文档明确指出支持以下下推：

- `WHERE`
- `LIMIT`
- `ORDER BY`
- `DISTINCT`

它也记录了数据修改能力：

- `INSERT`
- `UPDATE`
- `DELETE`

SPARQL UPDATE 的批处理粒度由 `batch_size` 控制。

## 辅助函数

README 列出了以下工具函数：

- `rdf_fdw_version()`
- `rdf_fdw_settings()`
- `rdf_fdw_clone_table()`

同时它还记录了更广泛的 SPARQL 函数支持，包括聚合、字符串函数、数值函数、日期时间函数、哈希函数和自定义函数。

## 说明

当前 README 提示：检索到的 RDF 数据在转换为 PostgreSQL 之前会先加载到内存，因此大结果集需要足够的 PostgreSQL 内存。文档还说明支持的最低 PostgreSQL 版本为 9.5+。
