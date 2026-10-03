## 用法

来源：

- [PGXN rdf_fdw 3.0.0](https://pgxn.org/dist/rdf_fdw/3.0.0/)
- [rdf_fdw 3.0 README](https://github.com/jimjonesbr/rdf_fdw/blob/v3.0/README.md)
- [rdf_fdw 3.0 变更日志](https://github.com/jimjonesbr/rdf_fdw/blob/v3.0/CHANGELOG.md)
- [rdf_fdw 3.0 控制文件](https://github.com/jimjonesbr/rdf_fdw/blob/v3.0/rdf_fdw.control)

`rdf_fdw` 是通过 SPARQL endpoint 查询 RDF triplestore 的 PostgreSQL foreign data wrapper。它把 SPARQL 结果变量暴露为外部表列，支持常见 SQL 子句下推，提供用于 RDF term 的原生 `rdfnode` 类型，实现了多种 SPARQL 1.1 辅助函数，并可通过可写外部表执行 SPARQL `INSERT`、`UPDATE` 和 `DELETE`。

发行包 3.0.0 对应 SQL 扩展版本 3.0。本次修复 RDF 比较、算术与下推正确性。`GROUP BY`、`DISTINCT`、`UNION` 和唯一约束按 RDF 项的书写形式区分；`sparql.uri()` 现在返回 `rdfnode`，赋值到文本列时需要显式转换。部分无法安全下推的条件改为本地执行。

### 创建扩展

```sql
CREATE EXTENSION IF NOT EXISTS rdf_fdw;

SELECT rdf_fdw_version();
SELECT * FROM rdf_fdw_settings;
```

从 2.x 升级前，先保存并删除 `rdfnode` 列上的索引，以及对其排序、分组或去重的依赖视图、物化视图或 SQL 函数体函数。升级后重建这些对象；仅执行 `REINDEX` 不够。新安装与已做好准备的升级应分别选择适用命令：

```sql
CREATE EXTENSION rdf_fdw WITH VERSION '3.0';
ALTER EXTENSION rdf_fdw UPDATE TO '3.0';
```

### 注册 SPARQL Endpoint

```sql
CREATE SERVER dbpedia
FOREIGN DATA WRAPPER rdf_fdw
OPTIONS (
  endpoint          'https://dbpedia.org/sparql',
  enable_pushdown   'true',
  request_timeout   '60',
  max_response_size '104857600',
  readonly          'true'
);
```

常用 server options：

- `endpoint`：SPARQL endpoint URL，必需。
- `batch_size`：每批 SPARQL UPDATE 的行数。
- `enable_pushdown`：启用 SQL 到 SPARQL 的下推。
- `format`：期望的 SPARQL 结果 MIME 类型。
- `http_proxy`：代理 URL；代理凭据应放在 `USER MAPPING`。
- `connect_timeout`：连接超时。
- `request_timeout`：完整 HTTP 请求超时。
- `max_response_size`：最大响应体字节数；`0` 表示不限制。
- `readonly`：在请求到达 endpoint 前阻止 `INSERT`、`UPDATE` 和 `DELETE`。
- `request_max_redirect`：`0` 拒绝重定向，正数启用并限制次数；不再接受 `-1`。`request_redirect` 已弃用。

连接公共或不可信 endpoint 时建议设置 `max_response_size`，因为 `rdf_fdw` 会先把取回的 RDF 数据载入内存，再转换为 PostgreSQL 表示。

### 用户映射

```sql
CREATE USER MAPPING FOR postgres
SERVER dbpedia
OPTIONS (
  user 'sparql_user',
  password 'secret'
);
```

v2.6.0 增加了 Bearer token 认证：

```sql
CREATE USER MAPPING FOR postgres
SERVER dbpedia
OPTIONS (
  token 'eyJhbGciOi...'
);
```

代理凭据也应放在 `USER MAPPING`：

```sql
CREATE USER MAPPING FOR app_user
SERVER dbpedia
OPTIONS (
  proxy_user 'proxy-user',
  proxy_password 'proxy-secret'
);
```

### 使用 rdfnode 列的外部表

外部表列应声明为 `rdfnode`，以保留 RDF term、IRI、blank node、语言标签和 XSD datatype。

```sql
CREATE FOREIGN TABLE dbpedia_films (
  film rdfnode OPTIONS (variable '?film'),
  name rdfnode OPTIONS (variable '?name'),
  year rdfnode OPTIONS (variable '?year')
)
SERVER dbpedia
OPTIONS (
  sparql $$
    SELECT ?film ?name ?year
    WHERE {
      ?film a dbo:Film ;
            rdfs:label ?name ;
            dbo:releaseDate ?year .
      FILTER (lang(?name) = 'en')
    }
  $$
);
```

v2.6.0 中，RDF 值使用原生 PostgreSQL 列类型的方式已被弃用。已有 native-typed 表仍可工作，但会发出 warning，并且会丢失 RDF term 细节。

### 查询与下推

```sql
SELECT film, sparql.lex(name) AS title
FROM dbpedia_films
WHERE name = '"The Matrix"@en'::rdfnode
ORDER BY year
LIMIT 10;

EXPLAIN (VERBOSE, COSTS OFF)
SELECT *
FROM dbpedia_films
WHERE film = '<http://dbpedia.org/resource/The_Matrix>'::rdfnode;
```

`rdf_fdw` 可下推 `WHERE`、`LIMIT`、`ORDER BY`、`DISTINCT` 以及受支持的比较和函数。使用 `EXPLAIN VERBOSE` 查看生成的远端 SPARQL。

### Prefix 管理

`rdf_fdw` 在 `sparql` schema 下提供 catalog 表和辅助函数，用于复用 SPARQL prefix：

```sql
SELECT sparql.add_context('default', 'Default SPARQL prefix context');
SELECT sparql.add_prefix('default', 'rdf',  'http://www.w3.org/1999/02/22-rdf-syntax-ns#');
SELECT sparql.add_prefix('default', 'rdfs', 'http://www.w3.org/2000/01/rdf-schema#');
SELECT sparql.add_prefix('default', 'xsd',  'http://www.w3.org/2001/XMLSchema#');
```

### 数据修改

可写外部表会将 PostgreSQL `INSERT`、`UPDATE` 和 `DELETE` 转换为 SPARQL UPDATE 请求。示例假设管理员已创建 rdf_write 服务，连接到自己管理的可写端点；公共 DBpedia 服务应保持只读。更新模式中的每个变量都必须映射到非空的 `rdfnode` 列，且必须设置 `sparql_update_pattern`。

```sql
CREATE FOREIGN TABLE writable_triples (
  subject rdfnode OPTIONS (variable '?s'),
  predicate rdfnode OPTIONS (variable '?p'),
  object rdfnode OPTIONS (variable '?o')
)
SERVER rdf_write
OPTIONS (
  sparql 'SELECT ?s ?p ?o WHERE { ?s ?p ?o }',
  sparql_update_pattern '?s ?p ?o .',
  readonly 'false'
);
INSERT INTO writable_triples(subject, predicate, object) VALUES (
  '<http://example.org/film/1>'::rdfnode,
  '<http://www.w3.org/2000/01/rdf-schema#label>'::rdfnode,
  '"Example Film"@en'::rdfnode
);
```

使用 `readonly = true` 禁用 server 或 table 的写入。表上显式设置的 `readonly = false` 会覆盖 server 的 `readonly = true`；如果必须禁止写入，还应通过远端 endpoint 的凭据和访问控制强制限制。

每次远端写入都会立即由三元组存储提交；PostgreSQL ROLLBACK 无法撤销，应设计应用层补偿，不能假定跨系统原子性。

### 克隆外部表

```sql
CALL rdf_fdw_clone_table(
  foreign_table := 'dbpedia_films',
  target_table  := 'dbpedia_films_local',
  fetch_size    := 1000,
  create_table  := true
);
```

`rdf_fdw_clone_table()` 会把外部表数据分批复制成本地表。v2.5 修复了克隆 RDF term 时的多项往返问题。

### SPARQL 函数

`sparql` schema 实现了许多 SPARQL 1.1 函数和聚合，包括：

- `sparql.sum`、`sparql.avg`、`sparql.min`、`sparql.max`、`sparql.group_concat`、`sparql.sample` 等聚合
- `sparql.isiri`、`sparql.isblank`、`sparql.isliteral`、`sparql.datatype`、`sparql.iri`、`sparql.strdt`、`sparql.strlang` 等 RDF term 函数
- `sparql.strlen`、`sparql.substr`、`sparql.ucase`、`sparql.lcase`、`sparql.contains`、`sparql.replace` 等字符串函数
- 数值、日期时间、哈希和其他便利函数

### 注意事项

- 上游最低基线为 PostgreSQL 9.5 或更高版本。
- 取回的 RDF 数据会在转换前累积到内存中。应设置 `max_response_size`、使用 `LIMIT`，并限制远端结果集规模。
- 优先使用 `rdfnode` 列。RDF term 使用 PostgreSQL 原生类型已被弃用，并会丢失 IRI、语言和 datatype 信息。
- 密钥应放在 `USER MAPPING`；不要把代理凭据或 endpoint token 写进 `SERVER` options。
- 公共 SPARQL endpoint 可能慢或有速率限制。需要时使用 `connect_timeout`、`request_timeout`、重试和本地物化。
- 3.0 中扩展不可重定位，且会创建固定的 `sparql` 模式。修改已有安装前，应核对 3.0 的迁移要求。
