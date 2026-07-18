## Usage

Sources:

- [Upstream README](https://github.com/adunstan/json_schema_validate/blob/f771374c64cd014fb5ba9373832e2109e6fd4e35/README.md)
- [Extension SQL](https://github.com/adunstan/json_schema_validate/blob/f771374c64cd014fb5ba9373832e2109e6fd4e35/json_schema_validate--0.1.0.sql)
- [Extension control file](https://github.com/adunstan/json_schema_validate/blob/f771374c64cd014fb5ba9373832e2109e6fd4e35/json_schema_validate.control)

`json_schema_validate` validates `json` or `jsonb` data against JSON Schema. `jsonschema_is_valid` returns a boolean, while `jsonschema_validate` returns an empty array for valid data or structured error objects for invalid data:

```sql
CREATE EXTENSION json_schema_validate;

SELECT jsonschema_is_valid(
  '{"id":1,"name":"Ada"}'::jsonb,
  '{"type":"object","required":["id","name"]}'::jsonb
);

SELECT jsonschema_validate(
  '{"id":"wrong"}'::jsonb,
  '{"properties":{"id":{"type":"integer"}}}'::jsonb
);

CREATE TABLE validated_events (
  data jsonb CHECK (
    jsonschema_is_valid(data, '{"type":"object","required":["id"]}'::jsonschema_compiled)
  )
);
```

### Supported scope

Version `0.1.0` requires PostgreSQL 14 or later. It supports many Draft 7-and-later keywords, local JSON Pointer references, and compiled schemas, but does not implement tuple-style `prefixItems`, dependency keywords, unevaluated-item/property keywords, external references, or dynamic-reference identifiers. Unknown formats are ignored. Review the upstream keyword list before using the extension as a security or data-integrity boundary.
