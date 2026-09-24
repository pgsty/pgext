## Usage

Sources:

- [README](https://github.com/TeamSpringbird/stannum/blob/e163585cb9d6b6f78b8de067a9c4aa33ea238063/README.md)
- [Control file](https://github.com/TeamSpringbird/stannum/blob/e163585cb9d6b6f78b8de067a9c4aa33ea238063/postgres/stannum.control)
- [Cargo.toml](https://github.com/TeamSpringbird/stannum/blob/e163585cb9d6b6f78b8de067a9c4aa33ea238063/Cargo.toml)
- [postgres/sql/stannum--0.1.0.sql](https://github.com/TeamSpringbird/stannum/blob/e163585cb9d6b6f78b8de067a9c4aa33ea238063/postgres/sql/stannum--0.1.0.sql)

`stannum` is an experimental text-search extension derived from Lead, with persistent indexes, Boolean and phrase queries, and BM25 scoring. The reviewed 0.1.0 source has its own extension identity and storage format.

### Core Workflow

```sql
CREATE EXTENSION stannum;
CREATE TABLE documents (id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY, body text);
INSERT INTO documents (body) VALUES
  ('PostgreSQL supports full text search'), ('Search with exact phrase matching');
CREATE INDEX documents_search ON documents USING stannum (body);
ANALYZE documents;
SELECT id, stannum.full_score(ctid) AS score
FROM documents WHERE body ==> 'search' ORDER BY score DESC LIMIT 10;
SELECT * FROM stannum.segment_info('documents_search');
SELECT * FROM stannum.verify_index('documents_search', heap_check => true);
```

### Operation and Compatibility

The `==>` operator accepts TINQL. `stannum.highlight` marks matching text; `stannum.segment_info` inspects index segments and `stannum.verify_index` checks their consistency, optionally against the heap. Index maintenance follows PostgreSQL writes and vacuum.

PostgreSQL 17 and 18 are build targets, with most current lifecycle evidence on 18. A superuser installs the extension in its fixed `stannum` schema. Primary-server use loads on demand. For standby index reads, add the library to `shared_preload_libraries` on both primary and standbys and restart those servers.

### Migration Boundary

Do not install `tin` and `stannum` in the same database: their global `==>` operator conflicts. There is no supported in-place migration from the earlier extension names or storage layouts. Treat this as experimental software and validate recovery and maintenance on the intended workload.
