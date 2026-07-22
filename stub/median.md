## Usage

Sources:

- [Median README at the reviewed commit](https://github.com/za-arthur/pg_median/blob/0a43525499153317020d8cd71f8c3258c2f34fc5/README.md)
- [Aggregate definition at the reviewed commit](https://github.com/za-arthur/pg_median/blob/0a43525499153317020d8cd71f8c3258c2f34fc5/median--1.0.sql)
- [Median implementation at the reviewed commit](https://github.com/za-arthur/pg_median/blob/0a43525499153317020d8cd71f8c3258c2f34fc5/median.c)

`median` defines a parallel-safe aggregate named `median(anyelement)`. It ignores null inputs, sorts every remaining value, and returns a value of the input type. It is convenient for compact analytical queries, but its type-dependent arithmetic and per-group memory use matter on large or mixed-type workloads.

### Core Workflow

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

For an odd number of non-null inputs, the aggregate returns the sorted middle value. For an even number, it adds the two middle values and divides the result by a type-constructed value of two. An empty or all-null group returns `NULL`.

The SQL interface exposes the aggregate plus internal transition, final, combine, serialize, and deserialize functions. Those support parallel aggregation; they are implementation details rather than application-level APIs.

### Type and Scale Caveats

- The input type must be sortable. Even-cardinality input additionally needs compatible `+` and `/` operators.
- Because the return type is the input type, integer inputs use integer arithmetic: the midpoint can truncate, and adding the two middle integers can overflow. Cast integers to `numeric` when an exact fractional midpoint is required.
- Each aggregate state retains all non-null values and the final function sorts them. Expect approximately linear state memory and `O(n log n)` sorting work per group; parallel workers each maintain their own state before combining.
- PostgreSQL's built-in `percentile_cont(0.5) WITHIN GROUP` is an alternative when interpolation semantics or a standard ordered-set aggregate is preferable. Confirm its return type and query plan separately.
- Upstream version `1.0` publishes no current PostgreSQL compatibility matrix. Test the extension with the exact server major version and input types used in production.
