


## Usage

> [pg_jsonschema: PostgreSQL extension providing JSON Schema validation](https://github.com/supabase/pg_jsonschema)

`pg_jsonschema` adds JSON Schema validation functions to PostgreSQL, enabling schema enforcement on JSON/JSONB columns via check constraints.

### Functions

| Function | Description |
|----------|-------------|
| `json_matches_schema(schema json, instance json)` | Validate a JSON instance against a schema, returns boolean |
| `jsonb_matches_schema(schema json, instance jsonb)` | Validate a JSONB instance against a schema, returns boolean |
| `jsonschema_is_valid(schema json)` | Check whether a JSON schema itself is valid |
| `jsonschema_validation_errors(schema json, instance json)` | Return an array of validation error messages |

### Table Constraints

Use check constraints to enforce document structure:

```sql
CREATE TABLE customer (
    id serial PRIMARY KEY,
    metadata json,
    CHECK (
        json_matches_schema(
            '{
                "type": "object",
                "properties": {
                    "tags": {
                        "type": "array",
                        "items": {
                            "type": "string",
                            "maxLength": 16
                        }
                    }
                }
            }',
            metadata
        )
    )
);

-- Valid insert (passes check constraint)
INSERT INTO customer(metadata) VALUES ('{"tags": ["vip", "darkmode-ui"]}');

-- Invalid insert (rejected by check constraint)
INSERT INTO customer(metadata) VALUES ('{"tags": [1, 3]}');
-- ERROR: new row violates check constraint
```

### Error Inspection

Retrieve detailed validation errors:

```sql
SELECT jsonschema_validation_errors('{"maxLength": 4}', '"123456789"');
-- Returns: {"\"123456789\" is longer than 4 characters"}
```

### Schema Validation

Verify that a schema is well-formed before using it:

```sql
SELECT jsonschema_is_valid('{
    "type": "object",
    "properties": {
        "name": {"type": "string"},
        "age":  {"type": "integer", "minimum": 0}
    },
    "required": ["name"]
}');
-- Returns: true
```
