## Usage

Sources:

- [Extension versions](https://docs.databricks.com/aws/en/oltp/projects/extensions)
- [Enable Lakebase Search](https://docs.databricks.com/aws/en/oltp/projects/lakebase-search)
- [Extension guide](https://docs.databricks.com/aws/en/oltp/projects/lakebase-text)

`lakebase_text` supplies BM25 ranking over standard PostgreSQL text-search vectors through `lakebase_bm25`. The managed Lakebase catalog lists SQL version 0.1.3 on PostgreSQL 16–18.

### Enable and Search

Enable Lakebase Search in project settings; this irreversible operation restarts every project compute and disconnects clients. Install with an authorized database role, insert text, then build the index:

```sql
CREATE EXTENSION IF NOT EXISTS lakebase_text;
CREATE TABLE documents (
  id bigint, passage text,
  terms tsvector GENERATED ALWAYS AS (to_tsvector('english', passage)) STORED
);
INSERT INTO documents (id, passage) VALUES (1, 'Postgres search');
CREATE INDEX documents_bm25 ON documents USING lakebase_bm25 (terms);
SELECT id FROM documents
ORDER BY terms <@> to_bm25query(to_tsvector('english', 'search'), 'documents_bm25')
LIMIT 5;
```

The score is negative; ascending order ranks the best matches first. Index scans may omit zero-score rows. Use matching text-search configurations for documents and queries.

### Maintenance

Build on populated tables and keep autovacuum effective; `VACUUM` refreshes statistics after bulk inserts. `ALTER EXTENSION lakebase_text UPDATE` changes SQL objects; index-format upgrades separately require `REINDEX INDEX CONCURRENTLY`. Check provider availability before updating. Public source, license, implementation language and preload settings are not documented here; this guide describes the provider-managed feature.
