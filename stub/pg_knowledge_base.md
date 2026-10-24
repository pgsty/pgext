## Usage

Sources:

- [Official source](https://github.com/igor-suhorukov/pg_knowledge_base/blob/a25f8dbec782699981e57b864afe33ab074d56a5/README.md)
- [Official source](https://github.com/igor-suhorukov/pg_knowledge_base/blob/a25f8dbec782699981e57b864afe33ab074d56a5/pg_knowledge_base.control)
- [Official source](https://github.com/igor-suhorukov/pg_knowledge_base/blob/a25f8dbec782699981e57b864afe33ab074d56a5/sql/pg_knowledge_base--0.1.sql)

`pg_knowledge_base` 0.1 combines transactional RDF tables, the Oxigraph SPARQL engine and relational-to-RDF mappings. This preview targets stock PostgreSQL 19 and was tested upstream on 19rc1. Its SQL functions work without preloading; the optional network endpoint has additional requirements.

### Core Workflow

Install as a superuser. Objects live in the fixed `kb` schema.

```sql
CREATE EXTENSION pg_knowledge_base;
SELECT kb.load('<urn:alice> <urn:knows> <urn:bob> .');
SELECT * FROM kb.sparql('SELECT ?s ?o WHERE { ?s <urn:knows> ?o }')
  AS t(s text, o text);
SELECT kb.sparql_json('SELECT ?s WHERE { ?s <urn:knows> <urn:bob> }');
```

`kb.load` ingests RDF text; `kb.load_file` reads a server-side file subject to its privilege checks. `kb.sparql` returns records with an explicit SQL column definition, while `kb.sparql_json` and `kb.sparql_text` return serialized results. `kb.sparql_update` changes the RDF store within the caller's transaction. Rollback, WAL, backup and replication use PostgreSQL's table machinery.

### Mappings and Network Access

`kb.add_mapping` registers R2RML mappings and `kb.direct_mapping` creates direct mappings. Mapped tables are read under the querying role's permissions and row security. SPARQL updates affect the RDF store, not mapped relational tables. Remote `SERVICE` evaluation is unsupported.

The HTTP endpoint is disabled by default. Enabling it requires adding `pg_knowledge_base` to `shared_preload_libraries`, setting `kb.endpoint_listen_addresses` and restarting PostgreSQL. Preserve existing preload entries. Its default port is 7878; database-specific query, update and graph-store paths authenticate through PostgreSQL roles and `pg_hba.conf`. Use TLS when sending credentials and review connection/worker limits. Optional tracing also needs preload and its documented tracing extension.

### Operational Limits

Grant access to the `kb` schema and underlying objects deliberately. Bulk loading locks the quad table and rebuilds its indexes, so plan for blocked concurrent work. The Rust engine has its own memory limit, separate from PostgreSQL `work_mem`. Review supported RDF/result formats and resource settings before exposing arbitrary queries.
