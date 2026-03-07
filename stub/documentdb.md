
## Usage

Add to shared_preload_libraries first:

```ini
shared_preload_libraries = 'pg_documentdb_core, pg_stat_statements, auto_explain'
```

Example, create extension and perform DDL & CRUD

```sql
-- CASCADE will install documentdb_core, pg_cron, vector, etc.
CREATE EXTENSION IF NOT EXISTS documentdb CASCADE;
```

Currently, DocumentDB can be used with FerretDB 2.0+ as a MongoDB-compatible backend.
