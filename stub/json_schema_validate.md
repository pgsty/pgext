## Usage

Sources:

- [Official README at the reviewed commit](https://github.com/adunstan/json_schema_validate/blob/f771374c64cd014fb5ba9373832e2109e6fd4e35/README.md)
- [Version 0.1.0 SQL API](https://github.com/adunstan/json_schema_validate/blob/f771374c64cd014fb5ba9373832e2109e6fd4e35/json_schema_validate--0.1.0.sql)
- [Validation implementation](https://github.com/adunstan/json_schema_validate/blob/f771374c64cd014fb5ba9373832e2109e6fd4e35/json_schema_validate.c)
- [Extension control file](https://github.com/adunstan/json_schema_validate/blob/f771374c64cd014fb5ba9373832e2109e6fd4e35/json_schema_validate.control)

`json_schema_validate` validates PostgreSQL `json` and `jsonb` values against a substantial, but incomplete, JSON Schema implementation. Use `jsonschema_is_valid` for a Boolean result, `jsonschema_validate` for structured errors, and `jsonschema_compile` when the same schema is reused many times. Version 0.1.0 requires PostgreSQL 14 or later.

### Core Workflow

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
    jsonschema_is_valid(
      data,
      jsonschema_compile('{"type":"object","required":["id"]}'::jsonb)
    )
  )
);
```

For valid input, the reviewed implementation of `jsonschema_validate` returns an empty array `[]`; for invalid input it returns an array of objects containing a path and message. Some source comments incorrectly say valid input returns SQL NULL, so rely on the implementation and regression-test the packaged build.

### Main Objects

`jsonschema_is_valid` and `jsonschema_validate` have overloads for `json`, `jsonb`, and the compiled `jsonschema_compiled` type. `jsonschema_compile(jsonb)` parses a schema once and caches regular expressions for repeated validation. The functions are declared immutable, strict, and parallel-safe, making constant schemas suitable for CHECK constraints.

### Schema Scope and Caveats

The implementation covers common type, object, array, string, numeric, composition, conditional, format, and local-reference keywords. It does not implement tuple-style `prefixItems`, dependency keywords, unevaluated-property or unevaluated-item keywords, external references, or dynamic references. Only local JSON Pointer references are supported, and unknown formats are ignored.

Unsupported or ignored keywords can make a document pass when a full standards implementation would reject it. Pin and test the exact schema vocabulary before using the extension as an integrity or security boundary, and keep resource limits in mind for deeply nested schemas, large arrays, complex regular expressions, or many accumulated validation errors.
