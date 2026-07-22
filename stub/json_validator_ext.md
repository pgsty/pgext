## Usage

Sources:

- [Official README](https://github.com/DblBee/json_validator_ext/blob/81c2935930072a1c03f0626a4239c2cb180396d9/README.md)
- [Rust implementation](https://github.com/DblBee/json_validator_ext/blob/81c2935930072a1c03f0626a4239c2cb180396d9/src/lib.rs)
- [Extension control file](https://github.com/DblBee/json_validator_ext/blob/81c2935930072a1c03f0626a4239c2cb180396d9/json_validator_ext.control)
- [Rust package manifest](https://github.com/DblBee/json_validator_ext/blob/81c2935930072a1c03f0626a4239c2cb180396d9/Cargo.toml)

`json_validator_ext` validates PostgreSQL `jsonb` values against JSON Schema. It supplies a compact operator for constraints, an equivalent function, detailed error reporting, and a per-backend LRU cache for compiled schemas.

### Core Workflow

Install the extension, then put the JSON instance on the left and its schema on the right of `@@`:

```sql
CREATE EXTENSION json_validator_ext;

CREATE TABLE support_ticket (
    id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    details jsonb NOT NULL,
    CONSTRAINT details_valid CHECK (
        details @@ '{
          "required": ["email"],
          "properties": {"email": {"type": "string", "format": "email"}}
        }'::jsonb
    )
);

SELECT *
FROM unnest(explain_json_schema_errors(
    '{"type":"integer"}'::jsonb,
    '"not-an-integer"'::jsonb
));
```

### Objects and Configuration

- `jsonb @@ jsonb` returns whether the left-hand instance satisfies the right-hand schema.
- `validate_json_schema(jsonb, jsonb)` takes the schema first and the instance second.
- `explain_json_schema_errors(jsonb, jsonb)` returns a text array with instance paths and validation messages.
- `json_validator.cache_size` controls compiled schemas cached per backend. It defaults to 100, accepts 1 through 1000, and is a superuser-settable runtime parameter.

Invalid schema definitions make the boolean APIs return false; use the diagnostic function to distinguish an invalid schema from an invalid instance. The control file marks the extension superuser-only, untrusted, and non-relocatable. The canonical extension and package name is `json_validator_ext`, although the README heading calls the project `pg_json_schema`; version `0.0.0` is an early API, so pin and regression-test the reviewed build and JSON Schema behavior before relying on it for integrity constraints.
