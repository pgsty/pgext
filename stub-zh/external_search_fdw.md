## 用法

来源：

- [AlloyDB Elasticsearch FDW 文档](https://docs.cloud.google.com/alloydb/docs/elastic-search?hl=en)
- [AlloyDB 发布说明](https://docs.cloud.google.com/alloydb/docs/release-notes)

`external_search_fdw` 是 Google 托管的 AlloyDB 扩展，可将 Elasticsearch 索引暴露为只读外部表。适合需要用 SQL、Lucene 语法、Query DSL 或 AlloyDB 混合搜索直接处理 Elasticsearch 数据，而又不希望将数据复制进 PostgreSQL 的场景。

### 核心流程

先创建 Elasticsearch 只读 API key，将其存入 Secret Manager，授予 AlloyDB 服务账户读取该 Secret 的权限，并启用出站连接，然后再配置 FDW。

```sql
CREATE EXTENSION external_search_fdw;

CREATE SERVER elasticsearch_server
FOREIGN DATA WRAPPER external_search_fdw
OPTIONS (
    server 'https://elastic.example.com:9200',
    search_provider 'elastic',
    auth_mode 'secret_manager',
    auth_method 'ApiKey',
    secret_path 'projects/123456789012/secrets/es-key/versions/1'
);

CREATE USER MAPPING FOR CURRENT_USER
SERVER elasticsearch_server;

CREATE FOREIGN TABLE es_documents (
    metadata external_search_fdw_schema.OpaqueMetadata,
    id bigint,
    body text,
    qubits integer
)
SERVER elasticsearch_server
OPTIONS (remote_table_name 'documents');

SELECT id, body
FROM es_documents
WHERE qubits < 105
ORDER BY metadata <@> 'body:"quantum computing"'
LIMIT 10;
```

当 PostgreSQL 列名与 Elasticsearch 字段名不同时，在外部表列上使用 `remote_field_name`。`metadata <@>` 排序表达式可接收 Lucene 查询文本；同一个 metadata 对象也可接收文档中定义的 Query DSL JSON 形式。

### 服务器选项与查询行为

`auth_method` 可取 `ApiKey` 或 `Basic`。可选的 `max_deadline_ms`、`pagination_num_results` 与 `pagination_context_timeout_ms` 控制请求和分页行为。AlloyDB 会尝试把 `SELECT`、`WHERE`、`ORDER BY` 与 `LIMIT` 下推到 Elasticsearch。重要查询应使用 `EXPLAIN VERBOSE` 检查；缺少限制或无法下推限制时，可能调用资源开销较高的 Elasticsearch Scroll API 分页。

### 注意事项

该集成是面向 PostgreSQL `17` 及更高版本 AlloyDB 的 Preview 功能。它只能读取而不能写入 Elasticsearch，因此数据同步仍由应用负责。并非所有 Elasticsearch 类型都受支持，`geo_point` 等专用类型明确不支持。`LIKE` 与部分其他文本条件不能下推。可用性与升级遵循 AlloyDB，而不是社区软件包。
