## Usage

Sources: [pg_projection README](https://github.com/suissa/pg_projection), [SQL definitions](https://github.com/suissa/pg_projection/blob/main/pg_projection--1.0.sql), [control file](https://github.com/suissa/pg_projection/blob/main/pg_projection.control).

`pg_projection` provides MongoDB-style read projections for PostgreSQL `jsonb`. The 1.0 SQL file defines two functions: `pg_project(jsonb, jsonb)` for one JSON document and `pg_project_set(text, jsonb)` for a query result converted to a JSON array.

### Project One JSONB Value

Projection values are numeric flags: `1` includes a field and `0` excludes a field.

```sql
CREATE EXTENSION pg_projection;

SELECT pg_project(
  '{"_id": 7, "name": "Ada", "email": "ada@example.test", "secret": "x"}'::jsonb,
  '{"name": 1, "email": 1}'::jsonb
);
-- {"_id": 7, "name": "Ada", "email": "ada@example.test"}
```

In inclusion mode, `_id` is included by default when present. Exclude it explicitly when the caller wants only the selected fields:

```sql
SELECT pg_project(
  '{"_id": 7, "name": "Ada", "email": "ada@example.test"}'::jsonb,
  '{"_id": 0, "name": 1}'::jsonb
);
-- {"name": "Ada"}
```

### Exclude Fields

When the projection uses `0`, the function starts from the original document and removes matching top-level keys:

```sql
SELECT pg_project(
  '{"name": "Ada", "internal_id": "a-1", "secret_key": "k"}'::jsonb,
  '{"internal_id": 0, "secret_key": 0}'::jsonb
);
-- {"name": "Ada"}
```

### Project A Query Result

`pg_project_set(query_text, projection_json)` executes the supplied SQL text, converts each row with `to_jsonb(t)`, applies `pg_project`, and returns a JSON array:

```sql
SELECT pg_project_set(
  'SELECT id, username, password_hash FROM users WHERE active',
  '{"password_hash": 0}'::jsonb
);
```

Because `query_text` is dynamic SQL, pass only trusted query strings assembled by application or migration code you control. Do not concatenate untrusted user input into this argument.

### Caveats

- The SQL implementation projects top-level keys; it does not implement nested MongoDB path projection.
- Projection values are cast to integers internally, so use numeric `0` and `1` flags.
- `pg_project(jsonb, jsonb)` is declared `IMMUTABLE STRICT`; `pg_project_set(text, jsonb)` is declared `STABLE`.
