## Usage

Sources:

- [Upstream README](https://github.com/neurondb/pg_pageindex/blob/893c95a4846787e65b2276f3016a1852b42e49ff/README.md)
- [Version 1.0 SQL API](https://github.com/neurondb/pg_pageindex/blob/893c95a4846787e65b2276f3016a1852b42e49ff/sql/pg_pageindex--1.0.sql)
- [API reference](https://github.com/neurondb/pg_pageindex/blob/893c95a4846787e65b2276f3016a1852b42e49ff/docs/api.md)
- [Upstream security guidance](https://github.com/neurondb/pg_pageindex/blob/893c95a4846787e65b2276f3016a1852b42e49ff/docs/security.md)

pg_pageindex 1.0 runs the PageIndex document-tree pipeline inside a PostgreSQL backend. It reads server-side PDF or Markdown files, returns JSON documents, and can call language models for summaries and tree search.

### Harden installation first

The install SQL grants function execution to public by default. Revoke it in the same transaction, then grant only the functions a dedicated trusted role needs:

```sql
BEGIN;
CREATE EXTENSION pg_pageindex;
REVOKE EXECUTE ON ALL FUNCTIONS IN SCHEMA pageindex FROM PUBLIC;
COMMIT;
SELECT pageindex.version();
```

After explicit grants, a trusted role can build from an approved server path:

```sql
SELECT pageindex.build_markdown(
  '/srv/pageindex/manual.md',
  '{"add_node_text":false}'::jsonb
);
```

### Caveats

- Paths are resolved on the database server and read with the PostgreSQL operating-system account. Never accept client-controlled paths; upstream equates that trust boundary with dangerous server-side program execution.
- PDF parsing and the Go, MuPDF, and PageIndex bridge run inside the calling backend. Malformed or very large documents can consume memory and CPU, raise errors, or crash the backend.
- Summary, description, and search operations can make blocking outbound requests to configured OpenAI-compatible or Anthropic endpoints. Document text, structure, prompts, and queries may leave the database host.
- Keep API keys outside SQL and database tables, restrict outbound networking, set statement and network timeouts, and review provider retention policy.
- The extension returns JSON rather than managing indexed rows. It does not create a PostgreSQL index access method or make arbitrary table queries faster.
