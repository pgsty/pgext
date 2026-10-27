## Usage

Sources:

- [Extension versions](https://docs.databricks.com/aws/en/oltp/projects/extensions)
- [Enable Lakebase Search](https://docs.databricks.com/aws/en/oltp/projects/lakebase-search)
- [Extension guide](https://docs.databricks.com/aws/en/oltp/projects/lakebase-vector)

`lakebase_vector` provides the `lakebase_ann` index method for pgvector-compatible approximate vector search on managed Lakebase Postgres. The provider lists version 1.1.1 for PostgreSQL 16–18; this is not a downloadable extension package.

### Enable and Query

Enable Lakebase Search in project settings first. This operation is irreversible and restarts all project computes, disconnecting clients. Install per database with an authorized role:

```sql
CREATE EXTENSION IF NOT EXISTS lakebase_vector CASCADE;
CREATE TABLE items (id bigint, embedding vector(3));
INSERT INTO items VALUES (1, '[0.1,0.2,0.3]');
CREATE INDEX items_ann ON items USING lakebase_ann (embedding vector_l2_ops);
SELECT * FROM items ORDER BY embedding <-> '[0.2,0.1,0.3]' LIMIT 5;
```

`CASCADE` installs the `vector` dependency. Generate embeddings outside the database and keep vector dimensions consistent. ANN results approximate nearest neighbors.

### Maintenance

`ALTER EXTENSION lakebase_vector UPDATE` updates SQL objects. Index storage has a separate version: existing indexes retain their format until rebuilt with `REINDEX INDEX CONCURRENTLY`. Check the actual compute's `pg_available_extensions` before upgrading; provider compute restarts can be needed to expose new versions. Source language, library/preload settings and public license are not established by these guides.
