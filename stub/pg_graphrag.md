## Usage

Sources:

- [Official source](https://github.com/Riku-KANO/pg-graph/blob/16986a73d3b45f02ff1744efc0efcc688517da9f/README.md)
- [Official source](https://github.com/Riku-KANO/pg-graph/blob/16986a73d3b45f02ff1744efc0efcc688517da9f/pg_graphrag.control)
- [Official source](https://github.com/Riku-KANO/pg-graph/blob/16986a73d3b45f02ff1744efc0efcc688517da9f/Cargo.toml)
- [Official source](https://github.com/Riku-KANO/pg-graph/blob/16986a73d3b45f02ff1744efc0efcc688517da9f/src/sql/schema.sql)
- [Official source](https://github.com/Riku-KANO/pg-graph/blob/16986a73d3b45f02ff1744efc0efcc688517da9f/examples/quickstart.sql)

`pg_graphrag` 0.1.0 stores knowledge graphs with source evidence, temporal relationships and permission-aware retrieval. The SQL schema requires PostgreSQL 18 features, including `uuidv7` and `WITHOUT OVERLAPS`, even though its Cargo manifest contains older-major build flags. This preview's repository does not establish a license.

### Core Workflow

Install as a superuser; `btree_gist` is required. Objects are in `graphrag`. The database does not call an LLM: applications supply extraction results and community summaries. The lightweight path below builds a graph without an LLM.

```sql
CREATE EXTENSION pg_graphrag CASCADE;
SELECT graphrag.create_graph('demo');
SELECT graphrag.add_document('demo', content => 'Alice works at Example Labs.',
                             external_id => 'doc1', title => 'Team');
SELECT graphrag.lazy_index('demo', id)
FROM graphrag.document WHERE external_id = 'doc1';
SELECT graphrag.build_context('demo', query => 'Alice', max_tokens => 600);
```

`ingest_extraction` stores structured extraction for a source chunk. `local_search` combines graph, dense and lexical retrieval; `relational_paths` returns explanation paths with evidence. `build_communities`, `pending_report` and `set_report` support application-generated summaries. Temporal searches can supply `as_of`; removing a document removes its derived evidence through the supported API.

### Permissions

Document/chunk ACLs contain PostgreSQL role names; empty ACLs are public. Call `graphrag.enable_row_security()` to protect direct table access, then use `graphrag.grant_reader` or `graphrag.grant_writer` deliberately. Run retrieval as the actual caller, for example with a transaction-local role, rather than as the owner or superuser that bypasses row security. The extension filters derived entities and reports by visible source evidence; ungrounded curated attributes may be public.

### Limits

No preload or restart is required. Dense retrieval scans candidates and does not require `pgvector`; it is not a claim of indexed vector search. LLM output remains untrusted application data, and summaries can become stale. Verify access rules and the supported refresh/deletion workflow with a representative corpus before deployment.
