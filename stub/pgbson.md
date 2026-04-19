
## Usage

Sources: [README](https://github.com/buzzm/postgresbson/blob/master/README.md), [META.json 2.0.2](https://github.com/buzzm/postgresbson/blob/master/META.json), [pgbson.control](https://github.com/buzzm/postgresbson/blob/master/pgbson.control)

`pgbson` adds a BSON data type plus BSON-aware accessors and operators. Upstream documents the package release as `2.0.2`, while the extension control file still exposes SQL default version `2.0`; this matches the packaging note that the dist version is ahead of the extension SQL version.

```sql
CREATE EXTENSION pgbson;
```

### Core Access Patterns

Typed dotpath accessors walk the BSON structure directly and are the upstream-recommended fast path:

```sql
SELECT bson_get_datetime(bson_column, 'msg.header.event.ts') FROM my_table;
SELECT bson_get_bson(bson_column, 'msg.header.event') FROM my_table;
SELECT bson_get_string(bson_column, 'data.payload.product.definition.id') FROM my_table;
```

JSON-style operators are also supported:

```sql
SELECT (bson_column->'msg'->'header'->'event'->>'ts')::timestamp
FROM my_table;
```

### Main Functions and Operators

- Typed getters such as `bson_get_string`, `bson_get_int32`, `bson_get_int64`, `bson_get_double`, `bson_get_decimal`, `bson_get_datetime`, `bson_get_binary`, and `bson_get_boolean`.
- `bson_get_bson` to return a BSON subdocument.
- `bson_get_jsonb_array` when a path resolves to an array and you want native `jsonb` array operators afterward.
- Arrow operators `->` and `->>` similar to PostgreSQL JSON types.
- Casts to `json`/`jsonb` using Extended JSON so type fidelity is preserved.

### Interop and Indexing

Cast BSON to `jsonb` when you want PostgreSQL JSON operators:

```sql
SELECT (bson_get_bson(bson_column, 'msg.header.event')::jsonb) ?& ARRAY['id', 'type']
FROM my_table;
```

Build expression indexes on extracted paths:

```sql
CREATE INDEX ON data_collection (bson_get_string(data, 'd.recordId'));
```

The README also notes BSON values can round-trip byte-for-byte through `bytea` casts.

### Caveats

- Dotpath accessors are usually faster and more memory-efficient than long `->` chains because they avoid materializing intermediate substructures.
- `bson_get_bson()` returns `NULL` for scalar endpoints because simple scalars are not BSON documents.
- Upstream explicitly calls out array handling and wrong-type accessor behavior as rough edges that still need better ergonomics.
