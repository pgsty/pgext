## Usage

Sources:

- [Official extension control file](https://github.com/dotvezz/pg_smolid/blob/b9840b6764545c4e59353506cfa4bcf2cfb2bf4f/pg_smolid.control)
- [Official Rust package manifest](https://github.com/dotvezz/pg_smolid/blob/b9840b6764545c4e59353506cfa4bcf2cfb2bf4f/Cargo.toml)
- [Official Rust implementation](https://github.com/dotvezz/pg_smolid/blob/b9840b6764545c4e59353506cfa4bcf2cfb2bf4f/src/lib.rs)

`pg_smolid` 0.1.0 adds the `smolid` data type and generators backed by the Rust `smolid` crate. Values have a compact text representation, compare by their underlying unsigned 64-bit value, and support B-tree and hash indexes.

### Core Workflow

```sql
CREATE EXTENSION pg_smolid;

CREATE TABLE events (
  id smolid PRIMARY KEY DEFAULT smolid(),
  payload jsonb NOT NULL
);

INSERT INTO events(payload) VALUES ('{"kind":"signup"}');
SELECT id, smolid_to_text(id), smolid_version(id) FROM events;
```

Use `smolid()` for a new value or `smolid_with_type(integer)` to request an encoded type value. `smolid_type(smolid)` returns that optional type; `smolid_version(smolid)` returns the format version. `smolid_to_bigint(smolid)` and `smolid_to_text(smolid)` expose the numeric and textual forms.

### Type Behavior and Caveats

The extension defines comparison operators, default B-tree and hash operator classes, and implicit casts from `smolid` to `bigint` and `text`. Ordering follows the encoded integer, not locale text order. The control file marks installation superuser-only and untrusted. Pin the extension and Rust crate versions across nodes, and validate serialization plus sort behavior before using the type in durable keys or logical replication.
