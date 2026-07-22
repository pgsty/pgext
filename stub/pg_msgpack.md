## Usage

Sources:

- [Official extension SQL](https://github.com/choplin/pg_msgpack/blob/7cf18817053f4a42df80c34d816003505de8b1c1/pg_msgpack--0.0.1.sql)
- [Official conversion implementation](https://github.com/choplin/pg_msgpack/blob/7cf18817053f4a42df80c34d816003505de8b1c1/pg_msgpack.c)
- [Official access-operator implementation](https://github.com/choplin/pg_msgpack/blob/7cf18817053f4a42df80c34d816003505de8b1c1/pg_msgpack_op.c)

`pg_msgpack` adds a `msgpack` value type for legacy MessagePack data, with JSON-style text I/O, casts, and field/element access. It can help inspect compact payloads already stored in PostgreSQL, but its 2013 implementation predates modern MessagePack and PostgreSQL `jsonb` facilities.

### Core Workflow

```sql
CREATE EXTENSION pg_msgpack;

SELECT '{"name":"Ada","scores":[10,20]}'::msgpack -> 'name';
SELECT ('{"name":"Ada","scores":[10,20]}'::msgpack -> 'scores') -> 1;
SELECT '{"a":1}'::json::msgpack::bytea;
```

Array positions are zero-based. A missing key, incompatible container type, or failed unpack operation normally yields SQL NULL rather than an error.

### SQL Surface

- The `msgpack` type accepts and emits JSON-like text.
- Casts between `msgpack` and `json` use `WITH INOUT` conversion.
- Assignment casts between `msgpack` and `bytea` use `WITHOUT FUNCTION`, exposing the stored binary representation without validation or transformation.
- The `->` object operator accepts a text key, and the `->` array operator accepts an integer position; both return another `msgpack` value.

The extension does not provide a `jsonb` cast, mutation API, or `GIN`/B-tree operator class. Use native JSON types when indexing, current JSON semantics, or broad operator support matters more than MessagePack wire compatibility.

### Safety and Compatibility

Never pass a negative array position to the integer access operator. The reviewed C code checks only the upper bound and then applies the signed position as a pointer offset; a negative value can access invalid memory and may destabilize the backend. Validate all indexes before the query reaches this extension.

The legacy formatter and raw-string handling do not provide a complete modern MessagePack-to-JSON round trip for every value. Treat casts and textual output as inspection conveniences, not as a canonical interchange encoding.

The reviewed upstream repository has a single 2013-era revision and uses old PostgreSQL JSON-parser and messagepack-c interfaces. Build and fuzz-test it against the exact server and library versions before deployment; isolate untrusted payloads from production backends.
