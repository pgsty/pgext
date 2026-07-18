## Usage

Sources:

- [Upstream README](https://github.com/furstenheim/is_jsonb_valid/blob/3afa28d70500791a233dc17c919d6136b83cb2ef/README.md)
- [Extension SQL](https://github.com/furstenheim/is_jsonb_valid/blob/3afa28d70500791a233dc17c919d6136b83cb2ef/is_jsonb_valid--0.1.4.sql)
- [Extension control file](https://github.com/furstenheim/is_jsonb_valid/blob/3afa28d70500791a233dc17c919d6136b83cb2ef/is_jsonb_valid.control)
- [PGXN distribution](https://pgxn.org/dist/is_jsonb_valid/)

`is_jsonb_valid` validates a JSONB value against a JSON Schema and returns a boolean. Use `is_jsonb_valid(schema, data)` for draft 4 or `is_jsonb_valid_draft_v7(schema, data)` for draft 7:

```sql
CREATE EXTENSION is_jsonb_valid;

SELECT is_jsonb_valid(
  '{"type":"object","required":["name"]}'::jsonb,
  '{"name":"Ada"}'::jsonb
);

SELECT is_jsonb_valid_draft_v7(
  '{"type":"integer","minimum":0}'::jsonb,
  '7'::jsonb
);
```

### Schema support

Upstream version `0.1.4` documents PostgreSQL 9.6 and later. Remote references are unsupported, local references are limited to definitions at the schema root, and JSON Schema format validation is not implemented. Treat unsupported keywords and reference layouts as application-level validation gaps rather than assuming full draft compliance.
