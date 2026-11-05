## Usage

Sources:

- [Upstream source 1](https://github.com/antflydb/antfly/blob/616117cc78ba29dcd6c0212abd78ac076201ee3f/rs/crates/postgres/README.md)
- [Upstream source 2](https://github.com/antflydb/antfly/blob/616117cc78ba29dcd6c0212abd78ac076201ee3f/rs/crates/postgres/antfly_postgres.control)

`antfly_postgres` 0.1.0 connects PostgreSQL search indexes and SQL functions to an external Antfly service.

### Core Workflow

The index sends existing rows during creation and synchronizes inserts. `antfly_search` returns identifiers, scores and JSONB data; `antfly_sync_trigger` provides a separate row-change path. Query builders `antfly_postgres.search`, `antfly_postgres.semantic` and `antfly_postgres.hybrid` produce structured search requests.

```sql
CREATE EXTENSION antfly_postgres;
SELECT antfly_status('http://localhost:8080/db/v1/');
CREATE TABLE antfly_docs (content text);
CREATE INDEX ON antfly_docs USING antfly (content)
  WITH (url = 'http://localhost:8080/db/v1/', collection = 'docs');
SELECT * FROM antfly_docs WHERE content @@@ 'database search';
```

### Operational Boundaries

Requires PostgreSQL 13–18, superuser installation and a reachable Antfly API. Include the /db/v1/ prefix in the configured URL; the bare default does not reach that API. Search sends data outside PostgreSQL, so configure network access and service credentials deliberately. Do not assume PostgreSQL rollback also rolls back remote writes. This is the renamed monorepo successor to pgaf.
