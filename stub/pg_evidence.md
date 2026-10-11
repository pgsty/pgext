## Usage

Sources:

- [README.md](https://github.com/amaljithkuttamath/pg-evidence/blob/f3a2ce82ce93f901458d7ac10fe4a582a82fcc2f/README.md)
- [docs/api.md](https://github.com/amaljithkuttamath/pg-evidence/blob/f3a2ce82ce93f901458d7ac10fe4a582a82fcc2f/docs/api.md)
- [docs/operations.md](https://github.com/amaljithkuttamath/pg-evidence/blob/f3a2ce82ce93f901458d7ac10fe4a582a82fcc2f/docs/operations.md)
- [src/lib.rs](https://github.com/amaljithkuttamath/pg-evidence/blob/f3a2ce82ce93f901458d7ac10fe4a582a82fcc2f/src/lib.rs)
- [Cargo.toml](https://github.com/amaljithkuttamath/pg-evidence/blob/f3a2ce82ce93f901458d7ac10fe4a582a82fcc2f/Cargo.toml)
- [pg_evidence.control](https://github.com/amaljithkuttamath/pg-evidence/blob/f3a2ce82ce93f901458d7ac10fe4a582a82fcc2f/pg_evidence.control)

`pg_evidence` 0.1.0 is a development preview for storing document versions, retrieving bounded excerpts and resolving citations to exact UTF-8 byte spans. PostgreSQL 18 and a UTF8 database are required. Embeddings are supplied by the client; `vector` is needed only for embedding-enabled collections.

### Core Workflow

```sql
CREATE EXTENSION pg_evidence;
SELECT evidence.init_collection('docs', '{}'::jsonb);
SET statement_timeout = '5s';
SELECT evidence.query('docs', '{
  "nodes":[{"id":"hits","op":"literal","text":"retrieval","limit":5}],
  "output":"hits","excerpt_bytes":512
}'::jsonb);
```

### Operational Boundaries

`evidence.stage_version()` stages source bytes, their digest and spans; `evidence.publish_version()` makes a version current. `evidence.query()` supports literal, regex, lexical and optional vector retrieval; `evidence.resolve()` verifies retained citation bytes. `evidence.annotate()` manages tags and relationships, and `evidence.purge()` logically removes retained source content. A nonzero `statement_timeout` is mandatory for queries: output budgets do not bound database work. Installation requires a superuser; functions use invoker privileges and application roles need schema and corpus grants. Trusted corpus writers can bypass API invariants through direct SQL. Corpus tables survive dropping the extension. Purge does not erase WAL, backups or other copies. License selection and native x86-64 validation remain pending upstream; no production-readiness claim is made.
