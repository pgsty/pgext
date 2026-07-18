## Usage

Sources:

- [Median README at the reviewed commit](https://github.com/za-arthur/pg_median/blob/0a43525499153317020d8cd71f8c3258c2f34fc5/README.md)
- [median.control at the reviewed commit](https://github.com/za-arthur/pg_median/blob/0a43525499153317020d8cd71f8c3258c2f34fc5/median.control)
- [Aggregate definition at the reviewed commit](https://github.com/za-arthur/pg_median/blob/0a43525499153317020d8cd71f8c3258c2f34fc5/median--1.0.sql)
- [Median implementation at the reviewed commit](https://github.com/za-arthur/pg_median/blob/0a43525499153317020d8cd71f8c3258c2f34fc5/median.c)

`median` defines a parallel-safe aggregate named `median(anyelement)`. It ignores null inputs, sorts the remaining values with the type's comparison support, and returns the middle value. For an even number of values, it computes the mean of the two middle values.

### Basic Example

```sql
CREATE EXTENSION median;

SELECT median(value)
FROM (VALUES
  (1.0::numeric),
  (9.0::numeric),
  (3.0::numeric),
  (7.0::numeric)
) AS sample(value);
```

The example returns the midpoint of the two central numeric values.

### Caveats

- The input type must provide comparison support. Even-cardinality input additionally needs compatible `+` and `/` operators because the implementation calculates a mean.
- The aggregate retains and sorts all non-null values in memory, so it is not a constant-memory percentile implementation.
- Control and catalog both identify version `1.0`. Upstream publishes no current PostgreSQL compatibility matrix; test the extension against the exact server major version and input types you use.
