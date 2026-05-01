## Usage

Sources: [README](https://github.com/Apaezmx/pgproto/blob/v0.5.0/README.md), [release 0.5.0](https://github.com/Apaezmx/pgproto/releases/tag/v0.5.0), [PGXN 0.5.0](https://pgxn.org/dist/pgproto/0.5.0/), [SQL definitions](https://github.com/Apaezmx/pgproto/blob/v0.5.0/sql/pgproto--1.0.sql), [Makefile](https://github.com/Apaezmx/pgproto/blob/v0.5.0/Makefile), [pgproto.control](https://github.com/Apaezmx/pgproto/blob/v0.5.0/pgproto.control)

`pgproto` stores Protocol Buffers `proto3` payloads in PostgreSQL as a native `protobuf` type, with schema-aware extraction, update helpers, containment/index support, and text/integer path operators. The upstream package version is `0.5.0`; the extension SQL/control default version remains `1.0`.

The current upstream source is a C/PGXS extension: the official `Makefile` sets `MODULE_big = pgproto`, builds C objects from `src/*.o`, and includes `$(PGXS)`. The README describes the implementation as pure C with no external Protobuf library dependency.

```sql
CREATE EXTENSION pgproto;
```

### Schema Registry and Storage

`pgproto` needs runtime protobuf descriptors before name/path-based extraction can interpret a binary payload. Register a serialized `FileDescriptorSet` in `pb_schemas`, or call the SQL registration helper when that fits your workflow:

```sql
INSERT INTO pb_schemas (name, data)
VALUES ('MySchema', '\x...');

SELECT pb_register_schema('MySchema', '\x...');
```

Store serialized protobuf bytes in a `protobuf` column:

```sql
CREATE TABLE items (
  id serial PRIMARY KEY,
  data protobuf
);

INSERT INTO items (data) VALUES ('\x0a02082a');
```

The 0.5.0 SQL also installs a convenience cast from `protobuf` to `bytea`, so byte-oriented functions such as `length(data::bytea)` can be used when needed.

### Querying

Use the path operators for nested, repeated, and map fields:

```sql
-- Integer accessor: returns int4
SELECT data #> '{Outer, inner, id}'::text[] FROM items;

-- Text accessor: returns text
SELECT data #>> '{Outer, tags, mykey}'::text[] FROM items;

-- Array index lookup
SELECT data #> '{Outer, scores, 0}'::text[] FROM items;
```

Other user-facing extraction helpers and operators defined by the extension include:

- `pb_get_int32(protobuf, int4)` for tag-based `int4` extraction.
- `pb_get_int32_by_name(protobuf, text, text)` and `pb_get_int32_by_name_dot(protobuf, text)` for name-based integer extraction.
- `->` as shorthand for dot-path integer lookup through `pb_get_int32_by_name_dot`.
- `pb_get_int32_by_path(protobuf, text[])` behind `#>`.
- `pb_get_text_by_path(protobuf, text[])` behind `#>>`.
- `pb_to_json(protobuf, text)` for text JSON conversion when a message name is supplied.

### Updates and Merge

`pb_set`, `pb_insert`, and `pb_delete` are pure functions: they return a new `protobuf` value, so persist changes with `UPDATE ... SET`. Upstream 0.5.0 documents automatic compaction for these mutations to remove stale tags.

```sql
UPDATE items
SET data = pb_set(data, ARRAY['Outer', 'a'], '42');

UPDATE items
SET data = pb_insert(data, ARRAY['Outer', 'scores', '0'], '100');

UPDATE items
SET data = pb_insert(data, ARRAY['Outer', 'tags', 'key1'], 'value1');

UPDATE items
SET data = pb_delete(data, ARRAY['Outer', 'a']);
```

Merge two protobuf values with the `||` operator, which calls `pb_merge`:

```sql
UPDATE items
SET data = data || other.data
FROM other
WHERE items.id = other.id;
```

### Indexing and Containment

Use ordinary expression indexes on extracted fields:

```sql
CREATE INDEX idx_items_pb_id
ON items ((data #> '{Outer, inner, id}'::text[]));

SELECT *
FROM items
WHERE (data #> '{Outer, inner, id}'::text[]) = 42;
```

The SQL definitions also expose protobuf containment with `@>` and a default `protobuf_gin_ops` operator class for GIN indexes:

```sql
CREATE INDEX idx_items_data_gin
ON items USING gin (data protobuf_gin_ops);

SELECT * FROM items WHERE data @> '\x0a02082a'::protobuf;
```

### Schema Evolution

The README frames schema evolution as a normal use case: added fields read as `NULL` from older messages, deprecated or unknown fields are skipped during traversal, enums are read as standard varints, and unset `oneof` fields return `NULL`.

### Caveats

- Runtime schemas are required for schema-aware path navigation; without registered descriptors, the extension cannot resolve message field names.
- `#>` returns `int4` and `#>>` returns `text`; choose the operator/function that matches the expected field type.
- Mutator helpers do not update rows in place; the returned value must be assigned back to the column.
- The README benchmark numbers are upstream project benchmarks, not independent performance guarantees.
