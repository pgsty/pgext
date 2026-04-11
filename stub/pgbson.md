
## Usage

> Syntax:
>
> ```sql
> CREATE EXTENSION pgbson;
> SELECT bson_get_datetime(bson_column, 'msg.header.event.ts') FROM my_table;
> SELECT (bson_column->'msg'->'header'->'event'->>'ts')::timestamp FROM my_table;
> ```
>
> Source: [README](https://github.com/buzzm/postgresbson)

`pgbson` adds a BSON data type to PostgreSQL together with functions and operators for creating, inspecting, and querying BSON documents. The upstream README positions it as a binary, richly typed alternative to JSON/JSONB with round-trip fidelity and first-class support for datetimes, numeric subtypes, and raw bytes.

## Why BSON

The README highlights several BSON advantages over JSON:

- datetimes are first-class values
- numeric types remain distinct (`int32`, `int64`, `float`, `decimal`)
- raw byte arrays are first-class
- round-tripping preserves exact binary representation
- native SDK support exists across many languages

## Access Patterns

The extension exposes two styles of access:

### Dotpath Accessors

These are the high-performance typed accessors documented upstream:

```sql
SELECT bson_get_datetime(bson_column, 'msg.header.event.ts') FROM my_table;
SELECT bson_get_bson(bson_column, 'msg.header.event') FROM my_table;
```

The README argues these are more memory-efficient than repeated arrow dereferences because they walk the BSON structure directly and materialize only the terminal value.

### Arrow Operators

It also supports JSON-like operators:

```sql
SELECT (bson_column->'msg'->'header'->'event'->>'ts')::timestamp
FROM my_table;
```

## JSON Interop

The BSON type can be cast to JSON using Extended JSON (EJSON) so type fidelity is preserved. This allows BSON values to be fed into JSON/JSONB functions and operators when needed:

```sql
SELECT (bson_get_bson(bson_column, 'msg.header.event')::jsonb) ?& ARRAY['id','type']
FROM my_table;
```

## Notes

The README includes examples of end-to-end BSON round-tripping across Java, Kafka, Python, and PostgreSQL, emphasizing that the stored BSON payload can be retrieved byte-for-byte unchanged when cast back to `bytea`.
