## Usage

Sources:

- [Official README](https://github.com/furstenheim/jsonb_deep_sum/blob/3f1b4be67e3bc74b7cc4635fc285dc3c602ee420/README.md)
- [Extension SQL API](https://github.com/furstenheim/jsonb_deep_sum/blob/3f1b4be67e3bc74b7cc4635fc285dc3c602ee420/jsonb_deep_sum--0.0.2.sql)
- [Official regression tests](https://github.com/furstenheim/jsonb_deep_sum/blob/3f1b4be67e3bc74b7cc4635fc285dc3c602ee420/sql/jsonb_deep_sum_test.sql)

`jsonb_deep_sum` 0.0.2 recursively adds numeric leaves in JSON objects. It supplies a two-value function and an aggregate, making it useful for summing sparse, nested metric maps whose keys differ between rows.

### Core Workflow

```sql
CREATE EXTENSION jsonb_deep_sum;

SELECT jsonb_deep_add(
  '{"requests":{"ok":4},"bytes":10}'::jsonb,
  '{"requests":{"ok":3,"failed":1},"bytes":5}'::jsonb
);

CREATE TABLE metric_batch (metrics jsonb);
INSERT INTO metric_batch VALUES
  ('{"a":1,"nested":{"x":2}}'),
  ('{"a":4,"nested":{"x":3,"y":8}}'),
  (NULL);

SELECT jsonb_deep_sum(metrics) FROM metric_batch;
```

`jsonb_deep_add(jsonb, jsonb) RETURNS jsonb` merges objects recursively and adds numeric values at matching paths. Keys present on only one side are retained. `jsonb_deep_sum(jsonb)` uses that function as its transition step with `{}` as the initial state; SQL NULL input rows are skipped.

### Data Contract and Limitations

Every JSON value traversed must be either an object or a number. Strings, booleans, JSON nulls, and arrays raise errors, and incompatible shapes at the same key are not a coercion mechanism. Validate documents before aggregation and keep units consistent for every repeated path.

The implementation directly walks PostgreSQL's internal `jsonb` representation in C. Upstream examples tested PostgreSQL 12-era builds but do not publish a current major-version matrix; rebuild and run the regression suite against the exact server. Aggregation materializes a complete result object in backend memory, so high key cardinality or deeply nested documents can be expensive. For schema-stable metrics, typed columns generally provide clearer validation, indexing, and overflow handling.
