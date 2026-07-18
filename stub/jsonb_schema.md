## Usage

Sources:

- [jsonb_schema README at the reviewed commit](https://github.com/postgrespro/jsonb_schema/blob/541028e86696c2bfc9c057245ed47563d92befa0/README.md)
- [jsonb_schema 1.0 install SQL at the reviewed commit](https://github.com/postgrespro/jsonb_schema/blob/541028e86696c2bfc9c057245ed47563d92befa0/jsonb_schema--1.0.sql)
- [jsonb_schema regression SQL at the reviewed commit](https://github.com/postgrespro/jsonb_schema/blob/541028e86696c2bfc9c057245ed47563d92befa0/sql/test.sql)

`jsonb_schema` version 1.0 separates repeated JSONB structure from values. `jsonb_pack` stores a deduplicated binary schema in `jsonb_schemes` and returns a two-element JSONB array containing the schema ID and encoded data; `jsonb_unpack` reconstructs the original value.

```sql
CREATE EXTENSION jsonb_schema;

CREATE TABLE packed_documents (
  id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  payload jsonb NOT NULL
);

INSERT INTO packed_documents (payload)
VALUES (jsonb_pack('{"kind":"event","value":42}'::jsonb));

SELECT id, jsonb_unpack(payload)
FROM packed_documents;

SELECT count(*) FROM jsonb_schemes;
```

### Caveats

- Packed values are an extension-specific array representation, not transparently queryable as the original object. Applications must unpack them before normal JSONB access.
- `jsonb_pack` performs an insert when a schema is new. Treat it as a write operation even though the install SQL labels the function parallel safe.
- The functions refer to `jsonb_schemes` without schema qualification. If the extension is installed outside the active search path, calls can fail or resolve an unintended relation; pin a controlled search path and test the chosen extension schema.
- Preserve and back up `jsonb_schemes` together with every packed column. A packed value cannot be reconstructed when its referenced schema row is missing.
