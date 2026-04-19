
## Usage

Sources: [README](https://github.com/Apaezmx/pgproto/blob/main/README.md), [release 0.3.3](https://github.com/Apaezmx/pgproto/releases/tag/v0.3.3), [pgproto.control](https://github.com/Apaezmx/pgproto/blob/main/pgproto.control)

`pgproto` adds a `protobuf` type for storing binary Protocol Buffers with schema-aware extraction and update helpers. The latest upstream release is `0.3.3`, while the extension control file advertises SQL default version `1.0`.

```sql
CREATE EXTENSION pgproto;
```

### Basic Workflow

Register a `FileDescriptorSet` so the extension can interpret message layouts:

```sql
INSERT INTO pb_schemas (name, data) VALUES ('MySchema', '\x...');
```

Store protobuf payloads in a `protobuf` column:

```sql
CREATE TABLE items (
  id serial PRIMARY KEY,
  data protobuf
);
```

### Querying

Use PostgreSQL-style path operators for nested fields:

```sql
SELECT data #> '{Outer, inner, id}'::text[] FROM items;
SELECT data #> '{Outer, tags, mykey}'::text[] FROM items;
```

The README highlights `->` and `#>` as the primary navigation operators for nested, repeated, and map fields.

### Updates and Merge

The write helpers are pure functions that return a new protobuf value:

- `pb_set(...)`
- `pb_insert(...)`
- `pb_delete(...)`
- `||` to merge two messages of the same type

```sql
UPDATE items SET data = pb_set(data, ARRAY['Outer', 'a'], '42');
UPDATE items SET data = pb_insert(data, ARRAY['Outer', 'scores', '0'], '100');
UPDATE items SET data = pb_delete(data, ARRAY['Outer', 'a']);
UPDATE items SET data = data || other_data;
```

### Indexing and Evolution

Expression indexes work on extracted fields:

```sql
CREATE INDEX idx_pb_id ON items ((data #> '{Outer, inner, id}'::text[]));
```

The README also documents schema evolution as a first-class use case: adding fields is backward-compatible, deprecated fields remain readable if present in older payloads, and re-registering schemas with `ON CONFLICT` is the expected update path.

### Caveats

- `pgproto` relies on registered runtime schemas; without the descriptor set, path-based extraction cannot interpret the payload.
- The update helpers do not mutate in place, so they need to be used in `UPDATE ... SET data = ...`.
