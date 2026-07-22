## Usage

Sources:

- [Official README](https://github.com/keithf4/pg_doc_store/blob/cd2c8c4bc7235bc6443ee0aa9dbec19f1f945333/README.md)
- [Extension control file](https://github.com/keithf4/pg_doc_store/blob/cd2c8c4bc7235bc6443ee0aa9dbec19f1f945333/pg_doc_store.control)
- [Official PGXN distribution](https://pgxn.org/dist/pg_doc_store/)

`pg_doc_store` implements a small document-store API in PL/pgSQL on top of ACID PostgreSQL tables. It creates a collection table for `jsonb` documents, assigns UUIDs, maintains a full-text-search vector, and exposes save, containment search, and text search functions.

### Core Workflow

The extension requires `pgcrypto` for UUID generation. A collection name must be schema-qualified.

```sql
CREATE EXTENSION pgcrypto;
CREATE EXTENSION pg_doc_store;

SELECT * FROM create_document('public.customers');

SELECT save_document(
    'public.customers',
    '{"name":"Ada","city":"London"}'::jsonb
);
```

`save_document` creates the collection when needed. A supplied `id` updates or inserts that UUID; without one, the function generates an ID and adds it to the returned document.

```sql
SELECT *
FROM find_document(
    'public.customers',
    '{"city":"London"}'::jsonb,
    'name',
    'ASC'
);

SELECT *
FROM search_document('public.customers', 'Ada');
```

### User-Facing Objects

- `create_document(p_tablename text)`: creates a collection and returns its schema and table name.
- `save_document(p_tablename text, p_doc_string jsonb)`: inserts or upserts a document and returns its stored JSON.
- `find_document(p_tablename text, p_criteria jsonb, p_orderbykey text DEFAULT 'id', p_orderby text DEFAULT 'ASC')`: returns documents containing the criteria, ordered by a document key.
- `search_document(p_tablename text, p_query text)`: returns documents ranked by full-text relevance.

Each collection contains `id uuid`, `body jsonb`, `search tsvector`, `created_at`, and `updated_at`, with a primary key, GIN indexes for JSON containment and text search, and a trigger that refreshes the search vector.

### Operational Notes

Upstream requires PostgreSQL 9.5 or later and `pgcrypto`; the control file declares that dependency and fixes the extension's schema placement. Collection tables are ordinary database objects created dynamically by the API, so include them explicitly in backup, privilege, schema-migration, and retention plans. Calls that name a collection should use trusted, schema-qualified identifiers. No shared preload is required. The upstream version `0.2.1` is an old, compact PL/pgSQL implementation; test JSON key ordering, full-text configuration, concurrent upserts, and upgrade behavior on the target PostgreSQL release.
