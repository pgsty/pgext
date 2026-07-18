## Usage

Sources:

- [Upstream README](https://github.com/gavinwahl/postgres-json-schema/blob/e25c758abb6395a5457bd1124da2983a48dc2aac/README.md)
- [Extension control file](https://github.com/gavinwahl/postgres-json-schema/blob/e25c758abb6395a5457bd1124da2983a48dc2aac/postgres-json-schema.control)

`postgres-json-schema` version `0.1.1` provides a pure PL/pgSQL function for validating PostgreSQL `jsonb` values against JSON Schema draft 4. A common use is enforcing document shape with a check constraint.

The extension name contains hyphens and must be quoted:

```sql
CREATE EXTENSION "postgres-json-schema";

CREATE TABLE events (
    id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    payload jsonb NOT NULL,
    CONSTRAINT payload_is_object
        CHECK (validate_json_schema('{"type":"object"}', payload))
);

INSERT INTO events (payload) VALUES ('{"kind":"created"}');
```

The validator covers JSON Schema draft 4 except remote HTTP references. Keep referenced schemas local rather than relying on network retrieval. The extension installs no shared library and requires no preload setting.
