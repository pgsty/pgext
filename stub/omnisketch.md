

## Usage

> [omnisketch: OmniSketch data structure for multi-dimensional stream analytics](https://github.com/tvondra/omnisketch)

Implements OmniSketch for on-line aggregation into approximate sketches and answering count queries with arbitrary predicates on multi-dimensional data.

```sql
CREATE EXTENSION omnisketch;
```

### Functions

| Function | Description |
|---|---|
| `omnisketch(epsilon, delta, record)` | Build a sketch from data with accuracy parameters |
| `omnisketch(sketch)` | Combine multiple compatible sketches |
| `omnisketch_count(sketch)` | Return total records added to the sketch |
| `omnisketch_estimate(sketch, record)` | Estimate count of records matching predicates |

### Parameters

- `epsilon` -- accuracy relative to total records, range `[0,1]` (lower = more accurate, larger sketch)
- `delta` -- accuracy, range `[0,1]`

### Examples

```sql
-- Create sample data
CREATE TABLE data (id INT, a INT, b INT);
INSERT INTO data SELECT i, mod(i,100), mod(i,100) FROM generate_series(1,1000000) s(i);

-- Pre-calculate sketches on partitions
CREATE TABLE sketches AS
SELECT mod(id,10) AS p, omnisketch(0.01, 0.01, (a, b)) AS s
FROM data GROUP BY mod(id,10);

-- Estimate count for condition (a = 10 AND b = 10)
SELECT omnisketch_estimate(omnisketch(s), (10, 10)) FROM sketches;

-- Get total record count
SELECT omnisketch_count(omnisketch(s)) FROM sketches;
```
