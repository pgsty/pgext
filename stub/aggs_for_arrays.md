

## Usage

> [aggs_for_arrays: aggregate-like functions for single arrays (column-based)](https://github.com/pjungwir/aggs_for_arrays)

Provides functions that compute statistics on a single array input. Supports `SMALLINT`, `INTEGER`, `BIGINT`, `REAL`, and `DOUBLE PRECISION`.

```sql
CREATE EXTENSION aggs_for_arrays;
```

### Functions

| Function | Description |
|---|---|
| `array_to_hist(values T[], start T, width T, count INT)` | Compute histogram bucket counts |
| `array_to_hist_2d(x[], y[], ...)` | 2-D histogram |
| `array_to_mean(values T[])` | Mean of array elements |
| `array_to_median(values T[])` | Median (unsorted input OK) |
| `sorted_array_to_median(values T[])` | Median (pre-sorted input) |
| `array_to_mode(values T[])` | Mode of array elements |
| `sorted_array_to_mode(values T[])` | Mode (pre-sorted input) |
| `array_to_percentile(values T[], pct FLOAT)` | Percentile (0 to 1) |
| `sorted_array_to_percentile(values T[], pct FLOAT)` | Percentile (pre-sorted input) |
| `array_to_percentiles(values T[], pcts FLOAT[])` | Multiple percentiles |
| `sorted_array_to_percentiles(values T[], pcts FLOAT[])` | Multiple percentiles (pre-sorted) |
| `array_to_max(values T[])` | Maximum element |
| `array_to_min(values T[])` | Minimum element |
| `array_to_min_max(values T[])` | `{min, max}` tuple |
| `array_to_skewness(values T[])` | Skewness |
| `array_to_kurtosis(values T[])` | Kurtosis |

### Examples

```sql
-- Histogram with 10 buckets starting at 0, width 10
SELECT array_to_hist(ARRAY[5,15,25,35,45], 0, 10, 5);
-- {1,1,1,1,1}

-- Mean and median
SELECT array_to_mean(ARRAY[1,2,3,4,5]);   -- 3.0
SELECT array_to_median(ARRAY[1,2,3,4,5]); -- 3.0

-- Percentiles
SELECT array_to_percentile(ARRAY[1,2,3,4,5,6,7,8,9,10], 0.95);

-- Multiple percentiles at once
SELECT array_to_percentiles(ARRAY[1,2,3,4,5], ARRAY[0.25, 0.5, 0.75]);
```
