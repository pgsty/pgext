## Usage

Sources:

- [README.md](https://github.com/TheKnarf/pg-automerge/blob/77ce017eeb3c33cb47e011cabac99a6976967cf3/README.md)
- [sql/snapshots/pg_automerge--0.1.0.sql](https://github.com/TheKnarf/pg-automerge/blob/77ce017eeb3c33cb47e011cabac99a6976967cf3/sql/snapshots/pg_automerge--0.1.0.sql)
- [docs/src/pages/operations/limitations.mdx](https://github.com/TheKnarf/pg-automerge/blob/77ce017eeb3c33cb47e011cabac99a6976967cf3/docs/src/pages/operations/limitations.mdx)
- [docs/src/pages/reference/configuration.mdx](https://github.com/TheKnarf/pg-automerge/blob/77ce017eeb3c33cb47e011cabac99a6976967cf3/docs/src/pages/reference/configuration.mdx)
- [pg_automerge.control](https://github.com/TheKnarf/pg-automerge/blob/77ce017eeb3c33cb47e011cabac99a6976967cf3/pg_automerge.control)

`pg_automerge` stores Automerge documents as a native PostgreSQL 18 type. Applications create changes and own actor IDs; the extension stores, merges and queries them. It does not edit documents or create changes itself.

### Core Workflow

```sql
CREATE EXTENSION pg_automerge;
CREATE TABLE docs (id uuid PRIMARY KEY, doc automerge NOT NULL);
PREPARE store_document(uuid, bytea) AS
INSERT INTO docs (id, doc) VALUES ($1, $2)
ON CONFLICT (id) DO UPDATE SET doc = merge(docs.doc, EXCLUDED.doc);
SELECT doc->>'title' FROM docs WHERE doc @> '{"status":"open"}';
```

### Operational Boundaries

Input must be valid Automerge binary saves, not JSON text. The implicit `jsonb` cast exposes the current document; `merge()` and `merge_agg()` combine histories. Incremental changes require an existing document, so an initial insert must use a full save. Use a distinct actor ID for each writer and compare heads rather than serialized bytes when identifying histories. `pg_automerge.max_load_memory` limits estimated per-document memory, but allocations bypass PostgreSQL memory accounting and do not bound concurrent sessions. A native load or merge is not promptly interruptible. Installation requires a superuser; keep writers trusted and bound input resource costs.
