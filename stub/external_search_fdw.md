## Usage

Sources:

- [AlloyDB Elasticsearch FDW documentation](https://docs.cloud.google.com/alloydb/docs/elastic-search?hl=en)
- [AlloyDB release notes](https://docs.cloud.google.com/alloydb/docs/release-notes)

`external_search_fdw` is a Google-managed AlloyDB extension that exposes Elasticsearch indexes as read-only foreign tables. It is useful when SQL, Lucene syntax, Query DSL, or AlloyDB hybrid search must operate over Elasticsearch data without copying it into PostgreSQL.

### Core Workflow

Create an Elasticsearch read-only API key, store it in Secret Manager, grant the AlloyDB service account access to that secret, and enable outbound connectivity before configuring the FDW.

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

Use `remote_field_name` on a foreign-table column when its PostgreSQL name differs from the Elasticsearch field. The `metadata <@>` ordering expression accepts Lucene query text; the same metadata object can receive the documented Query DSL JSON form.

### Server Options and Query Behavior

`auth_method` accepts `ApiKey` or `Basic`. Optional `max_deadline_ms`, `pagination_num_results`, and `pagination_context_timeout_ms` options control request and pagination behavior. AlloyDB attempts to push `SELECT`, `WHERE`, `ORDER BY`, and `LIMIT` work to Elasticsearch. Inspect important queries with `EXPLAIN VERBOSE`; missing or non-pushable limits can invoke Elasticsearch Scroll API pagination.

### Caveats

This integration is a Preview feature for AlloyDB running PostgreSQL `17` or later. It reads but does not write Elasticsearch data, so synchronization remains the application's responsibility. Not every Elasticsearch type is supported, and specialized types such as `geo_point` are excluded. `LIKE` and some other text predicates are not pushed down. Availability and upgrades follow AlloyDB rather than a community package.
