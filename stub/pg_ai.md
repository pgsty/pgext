## Usage

Sources:

- [Official upstream README](https://github.com/mahipv/pg_ai/blob/890dcd622528c00f710214ec966987ab2d04c2d4/README.md)
- [Official extension control file (pg_ai.control)](https://github.com/mahipv/pg_ai/blob/890dcd622528c00f710214ec966987ab2d04c2d4/pg_ai.control)
- [Official extension SQL (pg_ai--0.0.1.sql)](https://github.com/mahipv/pg_ai/blob/890dcd622528c00f710214ec966987ab2d04c2d4/sql/pg_ai--0.0.1.sql)

`pg_ai` — PostgreSQL extension with builtin RAG capabilities, enabling the interpretation and querying of data through both natural language and SQL functions. Use it for the corresponding vector, model, or retrieval workflow. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION pg_ai;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `pg_ai_create_vector_store(store NAME, sql_query TEXT, notes NAME = NULL)` is an extension function and returns `TEXT`.
- `pg_ai_generate_image(column_name TEXT, prompt TEXT = NULL)` is an extension function and returns `TEXT`.
- `pg_ai_help()` is an extension function and returns `TEXT`.
- `pg_ai_insight(column_name TEXT, prompt TEXT = NULL)` is an extension function and returns `TEXT`.
- `pg_ai_moderation(column_name TEXT, prompt TEXT = NULL)` is an extension function and returns `TEXT`.
- `pg_ai_query_vector_store(store NAME, nl_query TEXT, count INT = 2)` is an extension function and returns `SETOF`.
- `pg_ai_generate_image_agg` is an aggregate exposed by the extension.
- `pg_ai_insight_agg` is an aggregate exposed by the extension.
- `pg_ai_moderation_agg` is an aggregate exposed by the extension.

### Requirements and Caveats

- The reviewed control file declares default version `0.0.1`.
- The control file marks the extension as non-relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
