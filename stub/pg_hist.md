## Usage

Sources:

- [Official README](https://github.com/segasai/pg_hist/blob/3e67428bcc5017c8fda9a8635ca6a076417f30d7/README.md)
- [Extension SQL](https://github.com/segasai/pg_hist/blob/3e67428bcc5017c8fda9a8635ca6a076417f30d7/scripts/pg_hist--1.0.0.sql)
- [Histogram implementation](https://github.com/segasai/pg_hist/blob/3e67428bcc5017c8fda9a8635ca6a076417f30d7/pg_hist.c)

pg_hist computes one- to ten-dimensional histograms from the result of a supplied SQL query. It returns only nonempty bin coordinates and counts, with weighted variants returning floating-point totals. It is useful for ad hoc numeric summaries when the query source is trusted.

### Core Workflow

Provide equal-length arrays for bin counts, lower bounds, and upper bounds. The source query must return one numeric column per dimension:

```sql
CREATE EXTENSION pg_hist;

SELECT *
FROM pg_hist_2d(
  'SELECT x, y FROM measurement',
  ARRAY[20, 10],
  ARRAY[0.0, -5.0],
  ARRAY[100.0, 5.0]
);
```

For a weighted histogram, append the weight as the query's last column and call the matching `_w` function.

### Important Objects

- `pg_hist_1d`, `pg_hist_2d`, and `pg_hist_3d` return fixed-shape, nonempty integer-count bins.
- `pg_hist` supports arbitrary dimensions but requires a caller-supplied record definition.
- `pg_hist_1d_w`, `pg_hist_2d_w`, and `pg_hist_3d_w` return weighted sums.
- `pg_hist_w` is the arbitrary-dimensional weighted form.

### Limits and Safety Notes

The query text is prepared and executed inside the calling backend; never construct it from untrusted input. The implementation allocates a dense accumulator before omitting empty bins, permits at most 10 dimensions and 100 million total bins, and reads rows in batches of 1,000. Values outside the supplied ranges or null inputs do not contribute. Large bin products can consume substantial backend memory, so validate array sizes and run expensive histograms away from latency-sensitive sessions.
