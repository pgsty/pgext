## Usage

Sources:

- [Official extension SQL](https://github.com/niquola/jsonb_extra/blob/329cb98a84bc9daf929783500bc6beedef4f9b59/jsonb_extra--1.0.sql)
- [Official regression SQL](https://github.com/niquola/jsonb_extra/blob/329cb98a84bc9daf929783500bc6beedef4f9b59/sql/jsonb_extra.sql)
- [Official implementation](https://github.com/niquola/jsonb_extra/blob/329cb98a84bc9daf929783500bc6beedef4f9b59/jsonb_extra.c)

`jsonb_extra` 1.0 is a small C extension that extracts every JSON value reached by a key path, converts JSON scalars to text, and replaces an existing nested value. It predates much of PostgreSQL's current JSON-path surface and should be evaluated as experimental code.

### Core Workflow

```sql
CREATE EXTENSION jsonb_extra;

SELECT jsonb_extract(
  '{"a":{"items":[{"x":5},{"x":6},{"other":7}]}}'::jsonb,
  ARRAY['a', 'items', 'x']
);

SELECT jsonb_extract_text(
  '{"a":{"items":[{"x":5},{"x":6}]}}'::jsonb,
  ARRAY['a', 'items', 'x']
);

SELECT jsonb_as_text('true'::jsonb);

SELECT jsonb_update(
  '{"a":{"b":{"c":1,"d":7}}}'::jsonb,
  ARRAY['a', 'b', 'c'],
  '5'::jsonb
);
```

`jsonb_extract` walks object keys and fans out across arrays encountered on the path, returning a `jsonb[]`. `jsonb_extract_text` performs the same traversal and returns a `text[]`. Both return null when the path produces no values.

### Important Objects

- `jsonb_extract(jsonb, text[])` returns all matching values as JSONB elements.
- `jsonb_extract_text(jsonb, text[])` returns all matching values converted to text.
- `jsonb_as_text(jsonb)` converts a top-level scalar to text; JSON null becomes an empty string, and containers use JSON text representation.
- `jsonb_update(jsonb, text[], jsonb)` replaces a value at an existing object path and returns the rebuilt document.

### Operational Notes

All four functions are strict and immutable. The implementation explicitly marks path creation and array updates as unfinished, so `jsonb_update` should only be used for existing object paths. The source also contains unresolved memory-safety comments around JSONB value handling. Test malformed, missing, nested-array, and large-document cases on the exact PostgreSQL build, and prefer PostgreSQL's maintained built-in JSONB functions when they cover the same operation.
