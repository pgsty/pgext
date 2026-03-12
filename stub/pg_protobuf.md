


## Usage

> [pg_protobuf: Protocol Buffers support for PostgreSQL](https://github.com/afiskon/pg_protobuf)

Provides functions to decode Protocol Buffer binary data directly in SQL without schema definitions.

### Functions

- `protobuf_decode(bytea) RETURNS cstring` -- Decode protobuf binary into a human-readable format
- `protobuf_get_int(bytea, int) RETURNS bigint` -- Extract an integer field by field number
- `protobuf_get_string(bytea, int) RETURNS text` -- Extract a string field by field number
- `protobuf_get_bytes(bytea, int) RETURNS bytea` -- Extract raw bytes by field number
- `protobuf_get_bool(bytea, int) RETURNS boolean` -- Extract a boolean field by field number
- `protobuf_get_float(bytea, int) RETURNS real` -- Extract a float field by field number
- `protobuf_get_double(bytea, int) RETURNS double precision` -- Extract a double field by field number
- `protobuf_get_sfixed32(bytea, int) RETURNS int` -- Extract a signed fixed 32-bit field
- `protobuf_get_sfixed64(bytea, int) RETURNS bigint` -- Extract a signed fixed 64-bit field
- `protobuf_get_int_multi(bytea, int) RETURNS bigint[]` -- Extract repeated integer fields
- `protobuf_get_string_multi(bytea, int) RETURNS text[]` -- Extract repeated string fields
- `protobuf_get_bytes_multi(bytea, int) RETURNS bytea[]` -- Extract repeated bytes fields
- `protobuf_get_bool_multi(bytea, int) RETURNS boolean[]` -- Extract repeated boolean fields
- `protobuf_get_float_multi(bytea, int) RETURNS real[]` -- Extract repeated float fields
- `protobuf_get_double_multi(bytea, int) RETURNS double precision[]` -- Extract repeated double fields
- `protobuf_get_sfixed32_multi(bytea, int) RETURNS int[]` -- Extract repeated sfixed32 fields
- `protobuf_get_sfixed64_multi(bytea, int) RETURNS bigint[]` -- Extract repeated sfixed64 fields

### Example

```sql
CREATE EXTENSION pg_protobuf;

-- Create a table with protobuf data
CREATE TABLE heroes (x bytea);

-- Define accessor functions for specific fields
CREATE FUNCTION hero_name(x bytea) RETURNS text AS $$
BEGIN
  RETURN protobuf_get_string(x, 512);
END $$ LANGUAGE plpgsql IMMUTABLE;

CREATE FUNCTION hero_hp(x bytea) RETURNS bigint AS $$
BEGIN
  RETURN protobuf_get_int(x, 2);
END $$ LANGUAGE plpgsql IMMUTABLE;

-- Create an index on a protobuf field
CREATE INDEX hero_name_idx ON heroes USING btree(hero_name(x));

-- Query by protobuf field
SELECT hero_name(x) FROM heroes ORDER BY hero_name(x) LIMIT 10;
```

### Limitations

- No modification of Protobuf data
- Enums readable via `protobuf_get_int`
- Unsigned types not directly supported (no unsigned integers in PostgreSQL)
- `[packed=true]` not supported by `*_multi` procedures (use `protobuf_get_bytes*` instead)
