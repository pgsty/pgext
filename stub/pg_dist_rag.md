## Usage

Sources:

- [src/postgres/yb-extensions/pg_dist_rag/README.md](https://github.com/yugabyte/yugabyte-db/blob/771f263d34aa95952ca311bb4c63e1d7fa8e2159/src/postgres/yb-extensions/pg_dist_rag/README.md)
- [src/postgres/yb-extensions/pg_dist_rag/pg_dist_rag--0.0.1.sql](https://github.com/yugabyte/yugabyte-db/blob/771f263d34aa95952ca311bb4c63e1d7fa8e2159/src/postgres/yb-extensions/pg_dist_rag/pg_dist_rag--0.0.1.sql)
- [src/postgres/yb-extensions/pg_dist_rag/pg_dist_rag--0.0.1--0.0.2.sql](https://github.com/yugabyte/yugabyte-db/blob/771f263d34aa95952ca311bb4c63e1d7fa8e2159/src/postgres/yb-extensions/pg_dist_rag/pg_dist_rag--0.0.1--0.0.2.sql)
- [src/postgres/yb-extensions/pg_dist_rag/pg_dist_rag.control](https://github.com/yugabyte/yugabyte-db/blob/771f263d34aa95952ca311bb4c63e1d7fa8e2159/src/postgres/yb-extensions/pg_dist_rag/pg_dist_rag.control)

`pg_dist_rag` 0.0.2 manages sources, vector indexes, document-processing queues and column-embedding registrations. The reviewed distribution is part of YugabyteDB; its SQL catalog does not itself run an embedding service.

### Core Workflow

```sql
CREATE EXTENSION vector;
CREATE EXTENSION pg_dist_rag;
SELECT dist_rag.create_source(r_source_uri := 's3://example-bucket/documents/');
SELECT dist_rag.init_vector_index(
    r_index_name := 'knowledge',
    r_embedding_model_params := '{"dimensions":1536}'::jsonb
);
SELECT * FROM dist_rag.work_queue;
```

### Operational Boundaries

Install `vector` first. Objects live in `dist_rag`; the extension is not relocatable and declares no preload. No vanilla PostgreSQL major-support matrix is asserted for this YugabyteDB component.

`dist_rag.create_source` queues source work; `dist_rag.init_vector_index` creates embedding storage, and `dist_rag.build_index` queues preprocessing. External workers and configured model providers must process those queues. Monitor source, document and pipeline states before treating data as indexed.

Version 0.0.2 adds column mappings with `dist_rag.create_column_embedding_mapping` and activation through `dist_rag.init_column_embedding`; progress and pause/resume APIs manage the registration. Some management routines run with definer privileges. Limit execution and direct-table privileges to appropriate operators and workers, and protect stored provider configuration and credentials.
