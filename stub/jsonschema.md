Source: [jsonschema v0.1.9 README](https://github.com/theory/pg-jsonschema-boon/blob/v0.1.9/README.md), [documentation](https://github.com/theory/pg-jsonschema-boon/blob/v0.1.9/doc/jsonschema.md), [control file](https://github.com/theory/pg-jsonschema-boon/blob/v0.1.9/jsonschema.control), [SQL definition](https://github.com/theory/pg-jsonschema-boon/blob/v0.1.9/sql/jsonschema--0.1.9.sql).

## Usage

`jsonschema` validates JSON and JSONB values against JSON Schema inside PostgreSQL. It is the `theory/pg-jsonschema-boon` extension and is distinct from Supabase `pg_jsonschema`, although it provides compatibility wrappers named `json_matches_schema()` and `jsonb_matches_schema()`.

The extension supports JSON Schema draft 4, draft 6, draft 7, draft 2019-09, and draft 2020-12 through the Rust `boon` validator. It has no runtime dependency beyond PostgreSQL.

### Validate a Schema and a Document

```sql
CREATE EXTENSION IF NOT EXISTS jsonschema;

SELECT jsonschema_is_valid(
  '{
     "type": "object",
     "required": ["name", "email"],
     "properties": {
       "name":  { "type": "string" },
       "age":   { "type": "number", "minimum": 0 },
       "email": { "type": "string", "format": "email" }
     }
   }'::json
);

SELECT jsonschema_validates(
  '{"name":"Amos Burton","email":"amos@rocinante.ship"}'::json,
  '{
     "type": "object",
     "required": ["name", "email"],
     "properties": {
       "name":  { "type": "string" },
       "email": { "type": "string", "format": "email" }
     }
   }'::json
);
```

`jsonschema_is_valid(schema)` returns whether the schema itself compiles and validates against the selected draft. `jsonschema_validates(data, schema)` returns whether the JSON/JSONB value satisfies the schema.

### Check Constraints

```sql
CREATE TABLE customer_profile (
  id       bigserial PRIMARY KEY,
  profile  jsonb NOT NULL,
  CHECK (
    jsonschema_validates(
      profile,
      '{
         "type": "object",
         "required": ["email"],
         "properties": {
           "email": { "type": "string", "format": "email" },
           "tags":  {
             "type": "array",
             "items": { "type": "string", "maxLength": 16 }
           }
         }
       }'::jsonb
    )
  )
);
```

Use constraints when the database should reject malformed JSON documents at write time.

### Composed Schemas

```sql
SELECT jsonschema_validates(
  jsonb_build_object(
    'first_name', 'Naomi',
    'last_name', 'Nagata',
    'shipping_address', jsonb_build_object(
      'street_address', '1 Rocinante Way',
      'city', 'Ceres Station',
      'state', 'The Belt'
    )
  ),
  'https://example.com/schemas/customer',
  '{
     "$id": "https://example.com/schemas/address",
     "type": "object",
     "required": ["street_address", "city", "state"],
     "properties": {
       "street_address": { "type": "string" },
       "city": { "type": "string" },
       "state": { "type": "string" }
     }
   }'::jsonb,
  '{
     "$id": "https://example.com/schemas/customer",
     "type": "object",
     "required": ["first_name", "last_name", "shipping_address"],
     "properties": {
       "first_name": { "type": "string" },
       "last_name": { "type": "string" },
       "shipping_address": { "$ref": "/schemas/address" }
     }
   }'::jsonb
);
```

The `id` overloads let multiple schemas reference each other by `$id`, which is useful for componentized JSON Schema definitions.

### Compatibility Functions

```sql
SELECT json_matches_schema(
  '{"type":"string","maxLength":4}'::json,
  '"1234"'::json
);

SELECT jsonb_matches_schema(
  '{"type":"object","required":["id"]}'::json,
  '{"id":42}'::jsonb
);
```

These wrappers mirror the common `pg_jsonschema` argument order: schema first, instance second.

### Draft Selection and Caveats

```sql
SET jsonschema.default_draft = 'V2020';
SET jsonschema.default_draft = 'V7';
```

If a schema omits `$schema`, `jsonschema.default_draft` controls the default draft. Supported values are `V4`, `V6`, `V7`, `V2019`, and `V2020`.

- `jsonschema_validates(data, schema)` returns NULL if either argument is NULL.
- Invalid or uncompilable schemas can raise errors in validation calls; failed document validation returns `false` and logs details at `INFO`.
- `jsonschema_is_valid(id, VARIADIC schemas)` and `jsonschema_validates(data, id, VARIADIC schemas)` require matching `$id` values for reliable composed-schema resolution.
