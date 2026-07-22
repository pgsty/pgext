## Usage

Sources:

- [Official README](https://github.com/gavinwahl/postgres-json-schema/blob/e25c758abb6395a5457bd1124da2983a48dc2aac/README.md)
- [Official extension SQL](https://github.com/gavinwahl/postgres-json-schema/blob/e25c758abb6395a5457bd1124da2983a48dc2aac/postgres-json-schema--0.1.1.sql)
- [Official control file](https://github.com/gavinwahl/postgres-json-schema/blob/e25c758abb6395a5457bd1124da2983a48dc2aac/postgres-json-schema.control)

`postgres-json-schema` provides a PL/pgSQL validator for checking JSONB values against JSON Schema draft v4. Version `0.1.1` is useful for database constraints tied to that historical draft, but it does not implement later JSON Schema drafts or remote HTTP references.

### Core Workflow

Install the extension and use `validate_json_schema` in a check constraint:

```sql
CREATE EXTENSION "postgres-json-schema";

CREATE TABLE events (
  id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  payload jsonb NOT NULL,
  CONSTRAINT events_payload_schema CHECK (
    validate_json_schema(
      '{"type":"object","required":["kind"],"properties":{"kind":{"type":"string"}}}'::jsonb,
      payload
    )
  )
);
```

Valid rows return true; invalid rows make the check constraint reject the write.

### Function and Supported Features

`validate_json_schema(schema jsonb, data jsonb, root_schema jsonb DEFAULT NULL)` returns a boolean and recursively evaluates draft-v4 keywords. The source covers types, properties, required keys, array items, numeric and length bounds, composition keywords, enums, uniqueness, dependencies, patterns, formats, and root-anchored local `$ref` paths. The third parameter is used internally to retain the root schema during recursion and normally should be omitted.

### Operational Caveats

The README explicitly excludes remote HTTP references. Treat schemas as trusted configuration: malformed or unexpected schema structures may raise SQL errors rather than simply return false. Validation runs synchronously in PL/pgSQL for every inserted or updated row covered by a constraint, so benchmark large documents and high write rates. Changing a referenced schema literal requires a new constraint definition and validation of existing rows. Pin draft v4 semantics in application contracts; applications requiring newer drafts, remote resolution, or detailed error paths should use a maintained validator outside the database or a newer extension.
