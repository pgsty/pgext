## Usage

Sources:

- [Upstream README at the reviewed commit](https://github.com/mpartel/postgres-protobuf/blob/dd1abe9f1d915656eb149eee1f84490137ec7d1d/README.md)
- [Version 0.2 SQL objects](https://github.com/mpartel/postgres-protobuf/blob/dd1abe9f1d915656eb149eee1f84490137ec7d1d/postgres_protobuf--0.2.sql)
- [C++ implementation](https://github.com/mpartel/postgres-protobuf/blob/dd1abe9f1d915656eb149eee1f84490137ec7d1d/postgres_protobuf.cpp)

`postgres_protobuf` stores Protocol Buffer schemas and queries protobuf values held as `bytea`. It can select one field, arrays, or row sets and convert messages to and from their JSON representation.

```sql
CREATE EXTENSION postgres_protobuf;
INSERT INTO protobuf_file_descriptor_sets (name, file_descriptor_set)
VALUES ('default', $1::bytea);
SELECT protobuf_query(
  'example.User:profile.email',
  protobuf_payload
)
FROM event_store;
```

Generate the descriptor set with `protoc`, including imports, and insert it in a committed transaction before querying. Related functions include `protobuf_query_array`, `protobuf_query_multi`, `protobuf_to_json_text`, and `protobuf_from_json_text`.

Upstream warns against untrusted queries or unvalidated protobuf bytes. Large or recursive map-heavy messages can consume substantial C++ heap memory outside ordinary PostgreSQL accounting; every query scans the message, schemas are cached only for one transaction, and schema-dependent query functions cannot safely serve as index expressions. Proto3 missing/default fields yield no result, deprecated groups are unsupported, and the reviewed implementation was AMD64-only. Set input limits, validate and reserialize outside the database, restrict descriptor changes, and benchmark memory and latency.
