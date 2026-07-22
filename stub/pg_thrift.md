## Usage

Sources:

- [Official extension control file](https://github.com/ycui1984/pg_thrift/blob/1de265a661ab3c61aa593d7e99d3a313024170fc/pg_thrift.control)
- [Official README](https://github.com/ycui1984/pg_thrift/blob/1de265a661ab3c61aa593d7e99d3a313024170fc/README.md)
- [Official extension SQL](https://github.com/ycui1984/pg_thrift/blob/1de265a661ab3c61aa593d7e99d3a313024170fc/pg_thrift--1.0.sql)

`pg_thrift` 1.0 reads selected fields from Apache Thrift Binary Protocol or Compact Protocol values stored as `bytea`. It is useful for inspecting known Thrift payloads and building expression indexes on scalar fields. It does not load IDL schemas, generate Thrift messages, or provide a Thrift RPC client or server.

### Core Workflow

Install the extension, store an encoded struct as `bytea`, then call a protocol-specific getter with the numeric Thrift field ID. Container getters return encoded elements as `bytea[]`; feed those elements to the matching `parse_thrift_binary_*` or `parse_thrift_compact_*` function.

```sql
CREATE EXTENSION pg_thrift;

CREATE TABLE thrift_events (
    payload bytea NOT NULL
);

CREATE INDEX thrift_events_name_idx
    ON thrift_events (thrift_binary_get_string(payload, 1));

SELECT thrift_binary_get_int32(payload, 2) AS event_id,
       thrift_binary_get_string(payload, 1) AS event_name
FROM thrift_events
WHERE thrift_binary_get_string(payload, 1) = 'created';
```

The extension also defines the `thrift_binary` type. Its text input is a small JSON object such as `{"type":"int16","value":60}`; this convenience type supports the Binary Protocol only.

### Object Index

- Binary struct getters: `thrift_binary_get_bool`, `thrift_binary_get_byte`, `thrift_binary_get_double`, `thrift_binary_get_int16`, `thrift_binary_get_int32`, `thrift_binary_get_int64`, `thrift_binary_get_string`, and struct/list/set/map getters.
- Compact struct getters: the corresponding `thrift_compact_get_*` family.
- Scalar/container parsers: `parse_thrift_binary_*` and `parse_thrift_compact_*` families for encoded slices returned by container getters.
- Convenience conversion: `jsonb_to_thrift_binary(jsonb)` returns Binary Protocol bytes; `get_thrift_binary_type` and `get_thrift_binary_value` inspect the custom type.

Field IDs and wire types must match the application’s IDL. The extension has no schema registry and cannot verify that a field number has the intended business meaning.

### Compatibility and Caveats

The upstream code dates from 2018 and the catalog scope is PostgreSQL 10–11. Pin and test the exact source on newer servers before adoption. The C parser operates directly on supplied bytes; validate protocol, message size, nesting, and malformed-input behavior before accepting untrusted payloads.

All SQL functions are marked immutable, which permits expression indexes, but a changed IDL or wire encoding can make an existing indexed expression semantically obsolete. Rebuild affected indexes after such a change. `pg_thrift` is a decoding/query helper, not a replacement for application-level Thrift compatibility tests.
