
## Usage

> Syntax:
>
> ```sql
> CREATE EXTENSION pgproto;
> INSERT INTO pb_schemas (name, data) VALUES ('MySchema', '\x...');
> CREATE TABLE items (id serial PRIMARY KEY, data protobuf);
> SELECT data #> '{Outer, inner, id}'::text[] FROM items;
> ```
>
> Source: [README](https://github.com/Apaezmx/pgproto)

`pgproto` adds native Protocol Buffers support to PostgreSQL. It provides a `protobuf` type, runtime schema registration, nested field extraction, update helpers, and indexing support for schema-aware access to protobuf payloads.

## Setup

Enable the extension:

```sql
CREATE EXTENSION pgproto;
```

Register protobuf schemas by loading `FileDescriptorSet` blobs:

```sql
INSERT INTO pb_schemas (name, data) VALUES ('MySchema', '\x...');
```

Create a table using the custom `protobuf` type:

```sql
CREATE TABLE items (
    id SERIAL PRIMARY KEY,
    data protobuf
);
```

## Querying

The README highlights nested field extraction with PostgreSQL-style operators:

```sql
SELECT data #> '{Outer, inner, id}'::text[] FROM items;
SELECT data #> '{Outer, tags, mykey}'::text[] FROM items;
```

It also mentions custom operators such as `->` and `#>` for schema-aware navigation.

## Modification Functions

`pgproto` includes pure functions that return a new protobuf value:

- `pb_set(...)`
- `pb_insert(...)`
- `pb_delete(...)`

Because they return modified values rather than mutating in place, they are normally used in `UPDATE` statements:

```sql
UPDATE items SET data = pb_set(data, ARRAY['Outer', 'a'], '42');
UPDATE items SET data = pb_insert(data, ARRAY['Outer', 'scores', '0'], '100');
UPDATE items SET data = pb_delete(data, ARRAY['Outer', 'a']);
```

The `||` operator merges two protobuf messages of the same type.

## Indexing

The README documents B-tree expression indexes on extracted fields:

```sql
CREATE INDEX idx_pb_id ON items ((data #> '{Outer, inner, id}'::text[]));
```

The project also advertises GIN support for retrieval workflows.

## Notes

The upstream README positions `pgproto` as more storage-efficient than JSONB for protobuf-native payloads and highlights protobuf schema evolution, enums, `oneof`, and map/repeated field access as supported use cases.
