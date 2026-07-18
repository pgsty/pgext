## Usage

Sources:

- [jsonb_deep_sum README at the reviewed commit](https://github.com/furstenheim/jsonb_deep_sum/blob/3f1b4be67e3bc74b7cc4635fc285dc3c602ee420/README.md)
- [jsonb_deep_sum.control at the reviewed commit](https://github.com/furstenheim/jsonb_deep_sum/blob/3f1b4be67e3bc74b7cc4635fc285dc3c602ee420/jsonb_deep_sum.control)
- [jsonb_deep_sum on PGXN](https://pgxn.org/dist/jsonb_deep_sum/)

`jsonb_deep_sum` recursively combines numeric fields in JSONB objects. It provides `jsonb_deep_add` for adding two objects and the `jsonb_deep_sum` aggregate for reducing a column across rows while preserving nested object structure.

### Add and Aggregate JSONB Values

```sql
CREATE EXTENSION jsonb_deep_sum;

SELECT jsonb_deep_add(
  '{"a": 1, "nested": {"b": 2}}'::jsonb,
  '{"a": 4, "nested": {"b": 3}}'::jsonb
);

CREATE TABLE measurements (data jsonb);
INSERT INTO measurements VALUES
  ('{"a": 1}'),
  ('{"a": 2, "b": 1}'),
  ('{"a": 5}');

SELECT jsonb_deep_sum(data) FROM measurements;
```

The final aggregate in this example produces an object equivalent to `{"a": 8, "b": 1}`.

### Caveats

- Version 0.0.2 supports JSON objects and numeric values. Upstream says nulls, strings, arrays, and booleans inside the JSON value raise an error.
- The algorithm traverses JSONB's internally sorted tree and performs a merge; it does not define array-merging or nonnumeric conflict semantics.
- Both the control file and PGXN distribution identify version 0.0.2, matching the catalog version.
