## Usage

Sources:

- [pg_median_utils documentation at the reviewed commit](https://github.com/greenape/pg_median_utils/blob/f457fe5eb85376fb81f5c8e0ccdab517bcc5012f/doc/pg_median_utils.md)
- [pg_median_utils.control at the reviewed commit](https://github.com/greenape/pg_median_utils/blob/f457fe5eb85376fb81f5c8e0ccdab517bcc5012f/pg_median_utils.control)
- [Installed SQL API](https://github.com/greenape/pg_median_utils/blob/f457fe5eb85376fb81f5c8e0ccdab517bcc5012f/sql/pg_median_utils--0.0.1.sql)
- [Window-function implementations](https://github.com/greenape/pg_median_utils/tree/f457fe5eb85376fb81f5c8e0ccdab517bcc5012f/src)
- [Current PGXN distribution page](https://pgxn.org/dist/pg_median_utils/)

`pg_median_utils` provides five `double precision` window functions: `median_filter`, `iterated_median_filter`, `rolling_median`, `backfilled_rolling_median`, and `rolling_median_impute`. They operate on the ordered rows of each window partition and require an odd window size.

### Rolling Median

```sql
CREATE EXTENSION pg_median_utils;

WITH samples(sample_no, value) AS (
  VALUES (1, 1.0), (2, 5.0), (3, 2.0), (4, 8.0), (5, 3.0)
)
SELECT sample_no, value,
       rolling_median(value, 3) OVER (ORDER BY sample_no) AS median_3
FROM samples
ORDER BY sample_no;
```

The first two rows return null because a complete three-row trailing window is not yet available.

### Caveats

- Catalog, control, and the actual extension install script report `0.0.1`, while documentation, tags, metadata, and PGXN label the distribution `0.0.7`. Do not use the distribution label as `ALTER EXTENSION` target unless a matching SQL upgrade path exists.
- Ordering comes from the `OVER` clause; the functions inspect partition positions rather than a caller-defined SQL window frame. Always specify deterministic ordering.
- The iterative filter repeatedly processes an entire partition until convergence, and imputation builds partition-sized arrays. Bound partition and window sizes and test backend memory and latency.
- Upstream says PostgreSQL 9.6 or later, but the C window-API code has no current compatibility matrix.
