## Usage

Sources:

- [Official source](https://github.com/vladeParica/jsonvd/blob/905cbc44deb3e8413412544b6611a720bef755df/README.md)
- [Official source](https://github.com/vladeParica/jsonvd/blob/905cbc44deb3e8413412544b6611a720bef755df/jsonvd.control)
- [Official source](https://github.com/vladeParica/jsonvd/blob/905cbc44deb3e8413412544b6611a720bef755df/sql/jsonvd--1.0.0.sql)
- [Official source](https://github.com/vladeParica/jsonvd/blob/905cbc44deb3e8413412544b6611a720bef755df/docs/02-design/api-contracts.md)

`jsonvd` 1.0.0 is a SQL/PLpgSQL JSONB utility extension for PostgreSQL 14–18. It searches nested values, compares and merges documents, and infers types. It is trusted, uses the fixed `jsonvd` schema and needs no shared library, preload or restart. Installation still requires permission to create an extension in the database.

### Core Workflow

```sql
CREATE EXTENSION jsonvd;
SELECT jsonvd.find_path('{"items":[{"id":1},{"id":2}]}'::jsonb, '2'::jsonb);
SELECT * FROM jsonvd.diff('{"a":1,"b":[1,2]}'::jsonb,
                          '{"a":2,"b":[1,3]}'::jsonb);
SELECT jsonvd.merge_deep('{"db":{"host":"localhost","port":5432}}'::jsonb,
                        '{"db":{"port":6543}}'::jsonb);
SELECT jsonvd.merge_patch('{"a":1,"b":2}'::jsonb, '{"a":null}'::jsonb);
```

Paths are `text[]`, compatible with PostgreSQL JSON operators and update functions. Matching uses JSONB equality. `find_path` returns the first matching path or SQL NULL; `find_paths` returns every match. `contains_deep` tests membership and `flatten` returns path, value, type and depth for each node.

### Transformation and Inference

`diff` returns add/remove/replace rows and compares arrays by position. `merge_deep` recursively combines objects but lets the right side win other conflicts; SQL NULL sides are ignored. `merge_patch` follows JSON Merge Patch: JSON null removes an object key. SQL NULL and JSON null are distinct inputs.

`set_path` creates intermediate objects, `delete_key_deep` removes a key recursively, and `push` adds or replaces an object field. `parse_value` parses valid JSON text and otherwise returns a JSON string. `infer_schema` and `recordset_definition` infer a column description from sample objects; this is a description of the sample, not a constraint on future data.

### Boundaries

Functions are declared immutable and parallel-safe. Strict functions return SQL NULL for a NULL argument; merge and diff have their documented non-strict behavior. Invalid paths or argument shapes can raise SQL errors. Review null, scalar and array behavior before substituting these helpers for an existing application merge policy.
